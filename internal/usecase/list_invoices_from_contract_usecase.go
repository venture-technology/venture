package usecase

import (
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type ListInvoicesFromContractUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewListInvoicesFromContractUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *ListInvoicesFromContractUseCase {
	return &ListInvoicesFromContractUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (lifcuc *ListInvoicesFromContractUseCase) ListInvoicesFromContract() {

}
