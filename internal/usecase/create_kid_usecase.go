package usecase

import (
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type CreateKidUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewCreateKidUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *CreateKidUseCase {
	return &CreateKidUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (ccuc *CreateKidUseCase) CreateKid(kid *entity.Kid) error {
	return ccuc.repositories.KidRepository.Create(kid)
}
