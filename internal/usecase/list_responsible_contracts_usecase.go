package usecase

import (
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type ListResponsibleContractsUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewListResponsibleContractsUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *ListResponsibleContractsUseCase {
	return &ListResponsibleContractsUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (lruc *ListResponsibleContractsUseCase) ListResponsibleContracts() {

}
