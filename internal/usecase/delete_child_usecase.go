package usecase

import (
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type DeleteChildUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewDeleteChildUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *DeleteChildUseCase {
	return &DeleteChildUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (dcuc *DeleteChildUseCase) DeleteChild(rg *string) error {
	return dcuc.repositories.ChildRepository.Delete(rg)
}
