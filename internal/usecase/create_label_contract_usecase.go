package usecase

import (
	"encoding/json"

	"github.com/venture-technology/venture/internal/domain/service/adapters"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
	"github.com/venture-technology/venture/pkg/utils"
)

type CreateLabelContractUsecase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
	adapters     adapters.Adapters
	s3           contracts.S3Iface
	queue        contracts.Queue
	converter    contracts.Converters
}

func NewCreateLabelContractUsecase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
	adapters adapters.Adapters,
	s3 contracts.S3Iface,
	queue contracts.Queue,
	converter contracts.Converters,
) *CreateLabelContractUsecase {
	return &CreateLabelContractUsecase{
		repositories: repositories,
		logger:       logger,
		adapters:     adapters,
		s3:           s3,
		queue:        queue,
		converter:    converter,
	}
}

func (clc *CreateLabelContractUsecase) CreateLabelContract(msg string) error {
	var CreateContractInput value.CreateContractInput
	err := json.Unmarshal([]byte(msg), &CreateContractInput)
	if err != nil {
		return err
	}

	path, err := utils.GetAbsPath()
	if err != nil {
		return err
	}

	// seto contrato
	

	// converto pra pdf

	// jogo no bucket

	// jogo na fila

	return nil
}
