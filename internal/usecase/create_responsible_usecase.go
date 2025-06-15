package usecase

import (
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type CreateResponsibleUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewCreateResponsibleUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *CreateResponsibleUseCase {
	return &CreateResponsibleUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (cruc *CreateResponsibleUseCase) CreateResponsible(responsible *entity.Responsible) error {
	err := cruc.repositories.ResponsibleRepository.Create(responsible)
	if err != nil {
		return err
	}

	return err
}
