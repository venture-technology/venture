package usecase

import (
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type UpdateKidUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewUpdateKidUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *UpdateKidUseCase {
	return &UpdateKidUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (ucuc *UpdateKidUseCase) UpdateKid(rg string, attributes map[string]interface{}) error {
	return ucuc.repositories.KidRepository.Update(rg, attributes)
}
