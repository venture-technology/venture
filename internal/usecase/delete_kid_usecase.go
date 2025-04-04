package usecase

import (
	"fmt"

	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type DeleteKidUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewDeleteKidUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *DeleteKidUseCase {
	return &DeleteKidUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (dcuc *DeleteKidUseCase) DeleteKid(rg *string) error {
	kidHasContract, err := dcuc.repositories.ContractRepository.KidHasEnableContract(*rg)
	if err != nil {
		return err
	}

	if kidHasContract {
		return fmt.Errorf("impossivel deletar criança possuindo contrato ativo")
	}

	return dcuc.repositories.KidRepository.Delete(rg)
}
