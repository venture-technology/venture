package usecase

import (
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type CreateChildUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewCreateChildUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *CreateChildUseCase {
	return &CreateChildUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (ccuc *CreateChildUseCase) CreateChild(child *entity.Child) error {
	return ccuc.repositories.ChildRepository.Create(child)
}
