package usecase

import (
	"encoding/json"
	"fmt"

	"github.com/venture-technology/venture/internal/domain/service/adapters"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
)

type CreateContractUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
	adapters     adapters.Adapters
	s3           contracts.S3Iface
	queue        contracts.Queue
	converter    contracts.Converters
}

func NewCreateContractUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
	adapters adapters.Adapters,
	s3 contracts.S3Iface,
	queue contracts.Queue,
	converter contracts.Converters,
) *CreateContractUseCase {
	return &CreateContractUseCase{
		repositories: repositories,
		logger:       logger,
		adapters:     adapters,
		s3:           s3,
		queue:        queue,
		converter:    converter,
	}
}

func (ccuc *CreateContractUseCase) Execute(requestParams string) (err error) {
	defer func() {
		if err != nil {
			ccuc.logger.Infof(fmt.Sprintf("create contract uc error returned: %s", err.Error()))
		}
	}()

	var input value.CreateContractParams
	err = json.Unmarshal([]byte(requestParams), &input)
	if err != nil {
		return err
	}

	request, err := ccuc.adapters.AgreementService.SignatureRequest(input)
	if err != nil {
		return err
	}

	ccuc.logger.Infof(
		fmt.Sprintf(
			"Driver: %s, Responsible: %s, Kid: %s, School: %s",
			request.Metadata.Keys.DriverCNH,
			request.Metadata.Keys.ResponsibleCPF,
			request.Metadata.Keys.KidRG,
			request.Metadata.Keys.SchoolCNPJ,
		),
	)

	return nil
}
