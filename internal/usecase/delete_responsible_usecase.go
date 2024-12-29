package usecase

import (
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type DeleteResponsibleUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewDeleteResponsibleUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *DeleteResponsibleUseCase {
	return &DeleteResponsibleUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (druc *DeleteResponsibleUseCase) DeleteResponsible(cpf string) error {
	return druc.repositories.ResponsibleRepository.Delete(cpf)
}
