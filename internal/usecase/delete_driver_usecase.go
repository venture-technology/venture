package usecase

import (
	"fmt"

	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type DeleteDriverUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewDeleteDriverUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *DeleteDriverUseCase {
	return &DeleteDriverUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (dduc *DeleteDriverUseCase) DeleteDriver(cnh string) error {
	DriverHasContract, err := dduc.repositories.ContractRepository.DriverHasEnableContract(cnh)
	if err != nil {
		return err
	}

	if DriverHasContract {
		return fmt.Errorf("impossivel deletar motorista possuindo contrata ativo")
	}
	return dduc.repositories.DriverRepository.Delete(cnh)
}
