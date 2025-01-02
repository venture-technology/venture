package usecase

import (
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type UpdateChildUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewUpdateChildUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *UpdateChildUseCase {
	return &UpdateChildUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (ucuc *UpdateChildUseCase) UpdateChild(child *entity.Child, attributes map[string]interface{}) error {
	return ucuc.repositories.ChildRepository.Update(child, attributes)
}
