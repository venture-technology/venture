package usecase

import (
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
	return dsuc.repositories.SchoolRepository.Delete(cnpj)
}
