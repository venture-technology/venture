package usecase

import (
	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type CancelTempContractUsecase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewCancelTempContractUsecase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *CancelTempContractUsecase {
	return &CancelTempContractUsecase{
		repositories: repositories,
		logger:       logger,
	}
}

func (ctcuc *CancelTempContractUsecase) CancelTempContract(uuid uuid.UUID) error {
	return ctcuc.repositories.TempContractRepository.Cancel(uuid)
}
