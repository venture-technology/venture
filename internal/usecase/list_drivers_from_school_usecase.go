package usecase

import (
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type ListDriversFromSchoolUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewListDriversFromSchoolUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *ListDriversFromSchoolUseCase {
	return &ListDriversFromSchoolUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (ldfsuc *ListDriversFromSchoolUseCase) ListDriversFromSchool() {

}
