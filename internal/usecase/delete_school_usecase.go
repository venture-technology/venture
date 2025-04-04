package usecase

import (
	"fmt"

	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type DeleteSchoolUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewDeleteSchoolUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *DeleteSchoolUseCase {
	return &DeleteSchoolUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (dsuc *DeleteSchoolUseCase) DeleteSchool(cnpj string) error {
	SchoolHasContract, err := dsuc.repositories.ContractRepository.SchoolHasEnableContract(cnpj)
	if err != nil {
		return err
	}

	if SchoolHasContract {
		return fmt.Errorf("impossivel deletar escola possuindo contrato ativo")
	}

	return dsuc.repositories.SchoolRepository.Delete(cnpj)
}
