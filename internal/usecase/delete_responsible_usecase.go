package usecase

import (
	"fmt"

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
	responsible, err := druc.repositories.ResponsibleRepository.Get(cpf)
	if err != nil {
		return err
	}

	responsibleHasContract, err := druc.repositories.ContractRepository.ResponsibleHasEnableContract(cpf)
	if err != nil {
		return err
	}

	if responsibleHasContract {
		return fmt.Errorf("is not possible to delete responsible with active contract")
	}

	return druc.repositories.ResponsibleRepository.Delete(responsible.CPF)
}
