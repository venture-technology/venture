package usecase

import (
	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type ExpireContractUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewExpireContractUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *ExpireContractUseCase {
	return &ExpireContractUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (ecuc *ExpireContractUseCase) ExpireContract(uuid uuid.UUID) error {
	return ecuc.repositories.ContractRepository.Expired(uuid)
}
