package usecase

import (
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type DeleteDriverUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewDeleteDriverUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *DeleteDriverUseCase {
	return &DeleteDriverUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (dduc *DeleteDriverUseCase) DeleteDriver(cnh string) error {
	return dduc.repositories.DriverRepository.Delete(cnh)
}
