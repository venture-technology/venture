package usecase

import (
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type ListSchoolUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewListSchoolUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *ListSchoolUseCase {
	return &ListSchoolUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (lsuc *ListSchoolUseCase) ListSchool() {

}
