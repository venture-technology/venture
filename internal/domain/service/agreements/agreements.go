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
	"github.com/venture-technology/venture/config"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/pkg/realtime"
	"github.com/venture-technology/venture/pkg/utils"
)

type AgreementService struct {
	config       config.Config
	logger       contracts.Logger
	repositories *persistence.PostgresRepositories
}

func NewAgreementService(
	config config.Config,
	logger contracts.Logger,
	repositories *persistence.PostgresRepositories,
) *AgreementService {
	return &AgreementService{
		config:       config,
		logger:       logger,
		repositories: repositories,
	}
}

func (as *AgreementService) SignatureRequest(contract entity.ContractProperty) (ContractRequest, error) {
	url := as.config.Dropbox.SignatureRequestEndpoint
	apiKey := as.config.Dropbox.ApiKey

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

func (as *AgreementService) GetAgreementHtml(path string) ([]byte, error) {
	htmlFile, err := os.ReadFile(path)
	if err != nil {
		as.logger.Infof(fmt.Sprintf("error reading html file: %v", err))
		return nil, err
	}
	return htmlFile, nil
}

func (as *AgreementService) MappingContractInfo(contract entity.ContractProperty) ContractRequest {
	return ContractRequest{
		Title:   "Contrato de Prestação de Serviço",
		Subject: "Assinatura Anual - Venture",
		Message: "Por favor, reveja o contrato para assinatura e utilização dos serviços prestados pelo motorista",
		Signers: []Signer{
			{
				EmailAddress: contract.ContractParams.Driver.Email,
				Name:         contract.ContractParams.Driver.Name,
			},
			{
				EmailAddress: contract.ContractParams.Kid.Responsible.Email,
				Name:         contract.ContractParams.Kid.Responsible.Name,
			},
		},
		CCEmailAddresses: []string{as.config.Admin.AdminEmail},
		FileUrls:         []string{contract.URL},
		Metadata: Metadata{
			CustomID: contract.UUID,
			Keys: struct {
				UUID               string    `json:"uuid"`
				DriverID           string    `json:"driver_id"`
				DriverName         string    `json:"driver_name"`
				ResponsibleID      string    `json:"responsible_id"`
				ResponsibleName    string    `json:"responsible_name"`
				ResponsibleCPF     string    `json:"responsible_cpf"`
				ResponsibleEmail   string    `json:"responsible_email"`
				ResponsiblePhone   string    `json:"responsible_phone"`
				ResponsibleAddr    string    `json:"responsible_addr"`
				KidID              string    `json:"kid_id"`
				KidName            string    `json:"kid_name"`
				SchoolID           string    `json:"school_id"`
				SchoolName         string    `json:"school_name"`
				SchoolAddr         string    `json:"school_addr"`
				DateTime           string    `json:"date_time"`
				AmountContract     float64   `json:"amount_contract"`
				AnualContractValue float64   `json:"anual_contract_value"`
				Time               time.Time `json:"time"`
			}{
				UUID:             contract.UUID,
				DriverID:         contract.ContractParams.Driver.CNH,
				DriverName:       contract.ContractParams.Driver.Name,
				ResponsibleID:    contract.ContractParams.Kid.Responsible.CPF,
				ResponsibleName:  contract.ContractParams.Kid.Responsible.Name,
				ResponsibleCPF:   contract.ContractParams.Kid.Responsible.CPF,
				ResponsibleEmail: contract.ContractParams.Kid.Responsible.Email,
				ResponsiblePhone: contract.ContractParams.Kid.Responsible.Phone,
				ResponsibleAddr: utils.BuildAddress(
					contract.ContractParams.Kid.Responsible.Address.Street,
					contract.ContractParams.Kid.Responsible.Address.Number,
					contract.ContractParams.Kid.Responsible.Address.Complement,
					contract.ContractParams.Kid.Responsible.Address.Zip,
				),
				KidID:      contract.ContractParams.Kid.RG,
				KidName:    contract.ContractParams.Kid.Name,
				SchoolID:   contract.ContractParams.School.CNPJ,
				SchoolName: contract.ContractParams.School.Name,
				SchoolAddr: utils.BuildAddress(
					contract.ContractParams.School.Address.Street,
					contract.ContractParams.School.Address.Number,
					contract.ContractParams.School.Address.Complement,
					contract.ContractParams.School.Address.Zip,
				),
				DateTime:           contract.Time.Format("02/01/2006"),
				AmountContract:     contract.ContractParams.Amount,
				AnualContractValue: contract.ContractParams.Amount * 12,
				Time:               contract.Time,
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
		DriverCNH:          signatureResponse.SignatureRequest.Metadata.Keys.DriverID,
		ResponsibleCPF:     signatureResponse.SignatureRequest.Metadata.Keys.ResponsibleID,
		KidRG:              signatureResponse.SignatureRequest.Metadata.Keys.KidID,
		SchoolCNPJ:         signatureResponse.SignatureRequest.Metadata.Keys.SchoolID,
		UUID:               signatureResponse.SignatureRequest.Metadata.Keys.UUID,
	}
}

func (as *AgreementService) HandleCallbackVerification() (any, error) {
	return true, nil
}

func (as *AgreementService) SignatureRequestAllSigned(httpContext *gin.Context) (ASRASOutput, error) {
	var requestParams SignatureRequestAllSigned

	if err := httpContext.BindJSON(&requestParams); err != nil {
		return ASRASOutput{}, fmt.Errorf("invalid body")
	}

	return ASRASOutput{
		Contract: entity.Contract{
			UUID:           requestParams.SignatureRequest.Metadata.Keys.UUID,
			Status:         "currently",
			SigningURL:     requestParams.SigningURL,
			DriverCNH:      requestParams.SignatureRequest.Metadata.Keys.DriverID,
			SchoolCNPJ:     requestParams.SignatureRequest.Metadata.Keys.SchoolID,
			KidRG:          requestParams.SignatureRequest.Metadata.Keys.KidID,
			ResponsibleCPF: requestParams.SignatureRequest.Metadata.Keys.ResponsibleID,
			CreatedAt:      requestParams.CreatedAt,
			ExpireAt:       realtime.Now().Add(365 * 24 * time.Hour).Unix(),
			Amount:         requestParams.SignatureRequest.Metadata.Keys.AmountContract,
			AnualAmount:    requestParams.SignatureRequest.Metadata.Keys.AnualContractValue,
		},
		Signatures: requestParams.Signatures,
	}, nil
}
