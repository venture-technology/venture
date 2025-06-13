package signatures

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/spf13/viper"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
	"github.com/venture-technology/venture/pkg/realtime"
)

var (
	exp = realtime.Now().Add(2 * 24 * time.Hour).Unix()
)

const (
	urlBase      = "https://api.hellosign.com/v3"
	urlSignature = urlBase + "/signature_request/send"
)

type Signature interface {
	// Create a new file contract
	Create(params value.CreateContractParams) (*ContractParams, error)
}

type signature struct {
	accessToken string

	logger       contracts.Logger
	repositories *persistence.PostgresRepositories
}

func NewSignature(
	accessToken string,
	logger contracts.Logger,
	repositories *persistence.PostgresRepositories,
) *signature {
	return &signature{
		accessToken:  accessToken,
		logger:       logger,
		repositories: repositories,
	}
}

func (s *signature) Create(params value.CreateContractParams) (*ContractParams, error) {
	contract := s.resource(params)
	payload, err := json.Marshal(contract)
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequest(
		http.MethodPost, urlSignature,
		bytes.NewBuffer(payload),
	)
	if err != nil {
		return nil, err
	}

	r = s.httpCfg(r)
	return s.createResource(r, contract)
}

func (s *signature) resource(contract value.CreateContractParams) ContractParams {
	return ContractParams{
		Title:   "Contrato de Prestação de Serviço",
		Subject: "Assinatura Anual - Venture",
		Message: "Por favor, reveja o contrato para assinatura e utilização dos serviços prestados pelo motorista",
		Signers: []Signer{
			{
				EmailAddress: contract.DriverEmail,
				Name:         contract.DriverName,
			},
			{
				EmailAddress: contract.ResponsibleEmail,
				Name:         contract.ResponsibleName,
			},
		},
		CCEmailAddresses: []string{viper.GetString("DROPBOX_CC_EMAIL")},
		FileUrls:         []string{contract.FileURL},
		Metadata: Metadata{
			CustomID: contract.UUID,
			Keys: struct {
				UUID               string `json:"uuid"`
				DriverCNH          string `json:"driver_cnh"`
				ResponsibleCPF     string `json:"responsible_cpf"`
				KidRG              string `json:"kid_rg"`
				SchoolCNPJ         string `json:"school_cnpj"`
				AmountContract     int64  `json:"amount_contract"`
				AnualContractValue int64  `json:"anual_contract_value"`
			}{
				UUID:               contract.UUID,
				DriverCNH:          contract.DriverCNH,
				ResponsibleCPF:     contract.ResponsibleCPF,
				KidRG:              contract.KidRG,
				SchoolCNPJ:         contract.SchoolCNPJ,
				AmountContract:     contract.AmountCents,
				AnualContractValue: contract.AmountCents * 12,
			},
		},
		SigningOptions: struct {
			Draw        bool   `json:"draw"`
			Type        bool   `json:"type"`
			Upload      bool   `json:"upload"`
			Phone       bool   `json:"phone"`
			DefaultType string `json:"default_type"`
		}{
			Draw:        true,
			Type:        true,
			Upload:      true,
			Phone:       false,
			DefaultType: "draw",
		},
		FieldOptions: struct {
			DateFormat string `json:"date_format"`
		}{
			DateFormat: "DD - MM - YYYY",
		},
		ExpiresAt: exp,
		TestMode:  true,
	}
}

func (s *signature) httpCfg(r *http.Request) *http.Request {
	auth := base64.StdEncoding.EncodeToString(
		[]byte(s.accessToken),
	)
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", fmt.Sprintf("Basic %s", auth))
	return r
}

func (s *signature) createResource(r *http.Request, resource ContractParams) (*ContractParams, error) {
	c := &http.Client{}
	resp, err := c.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf(
			"unexpected status code from signature platform: %d, response: %s",
			resp.StatusCode, string(b),
		)
	}

	var sr SignatureResponse
	err = json.Unmarshal(b, &sr)
	if err != nil {
		return nil, err
	}

	err = s.createPrev(sr)
	if err != nil {
		return nil, err
	}

	return &resource, nil
}

func (s *signature) createPrev(sr SignatureResponse) error {
	prev := &entity.TempContract{
		SigningURL:         sr.SignatureRequest.SigningURL,
		SignatureRequestID: sr.SignatureRequest.SignatureRequestID,
		CreatedAt:          sr.SignatureRequest.CreatedAt,
		ExpiredAt:          exp,
		Status:             TemporaryContractPending,
		DriverCNH:          sr.SignatureRequest.Metadata.Keys.DriverCNH,
		ResponsibleCPF:     sr.SignatureRequest.Metadata.Keys.ResponsibleCPF,
		KidRG:              sr.SignatureRequest.Metadata.Keys.KidRG,
		SchoolCNPJ:         sr.SignatureRequest.Metadata.Keys.SchoolCNPJ,
		UUID:               sr.SignatureRequest.Metadata.Keys.UUID,
	}

	return s.repositories.TempContractRepository.Create(prev)
}
