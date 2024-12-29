package usecase

import (
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type ListSchoolContractUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewListSchoolContractUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *ListSchoolContractUseCase {
	return &ListSchoolContractUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (lscuc *ListSchoolContractUseCase) ListSchoolContract() {

}
