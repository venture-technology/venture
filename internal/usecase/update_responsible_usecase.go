package usecase

import (
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type UpdateResponsibleUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewUpdateResponsibleUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *UpdateResponsibleUseCase {
	return &UpdateResponsibleUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (uruc *UpdateResponsibleUseCase) UpdateResponsible(cpf string, attributes map[string]interface{}) error {
	return uruc.repositories.ResponsibleRepository.Update(cpf, attributes)
}
