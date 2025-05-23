package agreements

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
	"github.com/venture-technology/venture/pkg/realtime"
)

type AgreementService struct {
	logger       contracts.Logger
	repositories *persistence.PostgresRepositories
}

func NewAgreementService(
	logger contracts.Logger,
	repositories *persistence.PostgresRepositories,
) *AgreementService {
	return &AgreementService{
		logger:       logger,
		repositories: repositories,
	}
}

func (as *AgreementService) SignatureRequest(contract value.CreateContractParams) (ContractRequest, error) {
	url := viper.GetString("DROPBOX_API_URL")
	apiKey := viper.GetString("DROPBOX_SECRET_KEY")

	signatureContract := as.MappingContractInfo(contract)
	payload, err := json.Marshal(signatureContract)
	if err != nil {
		as.logger.Infof(fmt.Sprintf("error marshalling payload: %v", err))
		return ContractRequest{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		as.logger.Infof(fmt.Sprintf("error create request: %v", err))
		return ContractRequest{}, err
	}

	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:", apiKey)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", auth))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		as.logger.Infof(fmt.Sprintf("error to do payload request: %v", err))
		return ContractRequest{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		as.logger.Infof(fmt.Sprintf("error to check payload: %v", err))
		return ContractRequest{}, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return ContractRequest{}, fmt.Errorf(
			"unexpected status code from signature platform: %d, response: %s",
			resp.StatusCode, string(body),
		)
	}

	var signatureResponse SignatureResponse
	err = json.Unmarshal(body, &signatureResponse)
	if err != nil {
		as.logger.Infof(fmt.Sprintf("error unmarshalling response: %v", err))
		return ContractRequest{}, err
	}

	tempContract := as.buildTemporaryContract(signatureResponse)
	err = as.repositories.TempContractRepository.Create(tempContract)
	if err != nil {
		as.logger.Infof(fmt.Sprintf("error creating temporary contract: %v", err))
		return ContractRequest{}, err
	}

	return signatureContract, nil
}

func (as *AgreementService) BuildContract(path string) ([]byte, error) {
	htmlFile, err := os.ReadFile(path)
	if err != nil {
		as.logger.Infof(fmt.Sprintf("error reading html file: %v", err))
		return nil, err
	}
	return htmlFile, nil
}

// return json to send dropbox
func (as *AgreementService) MappingContractInfo(contract value.CreateContractParams) ContractRequest {
	return ContractRequest{
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
		ExpiresAt: as.GetExpireTime(),
		TestMode:  true,
	}
}

func (as *AgreementService) GetExpireTime() int64 {
	return realtime.Now().Add(2 * 24 * time.Hour).Unix()
}

func (as *AgreementService) buildTemporaryContract(signatureResponse SignatureResponse) *entity.TempContract {
	return &entity.TempContract{
		SigningURL:         signatureResponse.SignatureRequest.SigningURL,
		SignatureRequestID: signatureResponse.SignatureRequest.SignatureRequestID,
		CreatedAt:          signatureResponse.SignatureRequest.CreatedAt,
		ExpiredAt:          as.GetExpireTime(),
		Status:             TemporaryContractPending,
		DriverCNH:          signatureResponse.SignatureRequest.Metadata.Keys.DriverCNH,
		ResponsibleCPF:     signatureResponse.SignatureRequest.Metadata.Keys.ResponsibleCPF,
		KidRG:              signatureResponse.SignatureRequest.Metadata.Keys.KidRG,
		SchoolCNPJ:         signatureResponse.SignatureRequest.Metadata.Keys.SchoolCNPJ,
		UUID:               signatureResponse.SignatureRequest.Metadata.Keys.UUID,
	}
}

func (as *AgreementService) HandleCallbackVerification() (any, error) {
	return true, nil
}

func (as *AgreementService) SignatureRequestAllSigned(httpContext *gin.Context) (ASRASOutput, error) {
	var requestParams SignatureRequestAllSigned
	bodyBytes, err := io.ReadAll(httpContext.Request.Body)
	if err != nil {
		as.logger.Errorf(err.Error())
	}

	httpContext.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	if err := httpContext.BindJSON(&requestParams); err != nil {
		return ASRASOutput{}, fmt.Errorf("invalid body")
	}

	return ASRASOutput{
		Contract: entity.Contract{
			UUID:           requestParams.SignatureRequest.Metadata.Keys.UUID,
			Status:         value.ContractCurrently,
			SigningURL:     requestParams.SignatureRequest.SigningURL,
			DriverCNH:      requestParams.SignatureRequest.Metadata.Keys.DriverCNH,
			SchoolCNPJ:     requestParams.SignatureRequest.Metadata.Keys.SchoolCNPJ,
			KidRG:          requestParams.SignatureRequest.Metadata.Keys.KidRG,
			ResponsibleCPF: requestParams.SignatureRequest.Metadata.Keys.ResponsibleCPF,
			CreatedAt:      requestParams.SignatureRequest.CreatedAt,
			ExpireAt:       realtime.Now().Add(365 * 24 * time.Hour).Unix(),
			Amount:         requestParams.SignatureRequest.Metadata.Keys.AmountContract,
			AnualAmount:    requestParams.SignatureRequest.Metadata.Keys.AnualContractValue,
		},
		Signatures: requestParams.SignatureRequest.Signatures,
	}, nil
}
