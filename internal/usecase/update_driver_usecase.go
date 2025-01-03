package usecase

import (
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type UpdateDriverUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewUpdateDriverUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *UpdateDriverUseCase {
	return &UpdateDriverUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (uduc *UpdateDriverUseCase) UpdateDriver(cnh string, attributes map[string]interface{}) error {
	return uduc.repositories.DriverRepository.Update(cnh, attributes)
}
