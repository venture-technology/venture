package usecase

import (
	"fmt"

	"github.com/venture-technology/venture/internal/domain/service/adapters"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type DeleteResponsibleUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
	adapters     adapters.Adapters
}

func NewDeleteResponsibleUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
	adapters adapters.Adapters,
) *DeleteResponsibleUseCase {
	return &DeleteResponsibleUseCase{
		repositories: repositories,
		logger:       logger,
		adapters:     adapters,
	}
}

func (druc *DeleteResponsibleUseCase) DeleteResponsible(cpf string) error {
	responsible, err := druc.repositories.ResponsibleRepository.Get(cpf)
	if err != nil {
		return err
	}

	ResponsibleHasContract, err := druc.repositories.ContractRepository.ResponsibleHasEnableContract(cpf)
	if err != nil {
		return err
	}

	if ResponsibleHasContract {
		return fmt.Errorf("impossivel deletar responsavel possuindo contrato ativo")
	}

	_, err = druc.adapters.PaymentsService.DeleteStripeUser(responsible.CustomerId)
	if err != nil {
		return err
	}
	return druc.repositories.ResponsibleRepository.Delete(cpf)
}
