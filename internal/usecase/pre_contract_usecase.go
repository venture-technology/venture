package usecase

import (
	"fmt"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/domain/service/adapters"
	"github.com/venture-technology/venture/internal/domain/service/agreements"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
	"github.com/venture-technology/venture/pkg/realtime"
	"github.com/venture-technology/venture/pkg/utils"
)

type CreateContractUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
	adapters     adapters.Adapters
	S3           contracts.S3Iface
	converter    contracts.Converters
}

func NewCreateContractUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
	adapters adapters.Adapters,
	S3 contracts.S3Iface,
	converter contracts.Converters,
) *CreateContractUseCase {
	return &CreateContractUseCase{
		repositories: repositories,
		logger:       logger,
		adapters:     adapters,
		S3:           S3,
		converter:    converter,
	}
}

func (ccuc *CreateContractUseCase) CreateContract(
	requestParams *value.CreateContractRequestParams,
) (agreements.ContractRequest, error) {
	alreadyExists, err := ccuc.repositories.TempContractRepository.HasTemporaryContract(&entity.TempContract{
		DriverCNH:      requestParams.DriverCNH,
		KidRG:          requestParams.KidRG,
		SchoolCNPJ:     requestParams.SchoolCNPJ,
		ResponsibleCPF: requestParams.ResponsibleCPF,
	})
	if err != nil {
		ccuc.logger.Infof(fmt.Sprintf("error getting temp contract: %v", err.Error()))
		return agreements.ContractRequest{}, err
	}

	if alreadyExists {
		ccuc.logger.Infof(fmt.Sprintf("temp contract already exists"))
		return agreements.ContractRequest{}, fmt.Errorf("temp contract already exists")
	}

	// preciso buscar dados que vao ser enviados para fila de criacao de etiqueta contratua

	var contractProperty entity.ContractProperty
	request, err := ccuc.adapters.AgreementService.SignatureRequest(contractProperty)
	if err != nil {
		ccuc.logger.Infof(fmt.Sprintf("error sending signature request: %v", err))
		return agreements.ContractRequest{}, err
	}

	return request, nil
}

func (ccuc *CreateContractUseCase) SetContractProperty(
	requestParams value.CreateContractRequestParams,
) (entity.ContractProperty, error) {
	driver, err := ccuc.repositories.DriverRepository.Get(requestParams.DriverCNH)
	if err != nil {
		return entity.ContractProperty{}, err
	}

	school, err := ccuc.repositories.SchoolRepository.Get(requestParams.SchoolCNPJ)
	if err != nil {
		return entity.ContractProperty{}, err
	}

	kid, err := ccuc.repositories.KidRepository.Get(&requestParams.KidRG)
	if err != nil {
		return entity.ContractProperty{}, err
	}

	responsible, err := ccuc.repositories.ResponsibleRepository.Get(requestParams.ResponsibleCPF)
	if err != nil {
		return entity.ContractProperty{}, err
	}

	distance, err := ccuc.adapters.AddressService.GetDistance(
		responsible.Address.GetFullAddress(),
		school.Address.GetFullAddress(),
	)
	if err != nil {
		return entity.ContractProperty{}, err
	}

	contractValue := utils.CalculateContract(*distance, driver.Amount)
	if contractValue == 0 {
		return entity.ContractProperty{}, fmt.Errorf("invalid contract value")
	}

	return entity.ContractProperty{
		UUID: uuid.New().String(),
		ContractParams: entity.ContractParams{
			Driver:      *driver,
			School:      *school,
			Kid:         *kid,
			Responsible: *responsible,
			Amount:      contractValue,
			AnualAmount: contractValue * 12,
		},
		Time:     realtime.Now(),
		DateTime: realtime.Now().Format("02/01/2006"),
	}, nil
}

func getHtmlPaths() (string, error) {
	path, err := utils.GetAbsPath()
	if err != nil {
		return "", err
	}

	return filepath.Join(path, pathAgreement()), nil
}

func pathAgreements() string {
	return "../domain/service/agreements/template/agreement.html"
}
