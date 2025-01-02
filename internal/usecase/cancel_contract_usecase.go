package usecase

import (
	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/domain/payments"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type CancelContractUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
	payments     payments.IStripe
}

func NewCancelContractUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
	payments payments.IStripe,
) *CancelContractUseCase {
	return &CancelContractUseCase{
		repositories: repositories,
		logger:       logger,
		payments:     payments,
	}
}

func (ccuc *CancelContractUseCase) CancelContract(id uuid.UUID) error {
	contract, err := ccuc.repositories.ContractRepository.Get(id)
	if err != nil {
		return err
	}

	invoices, err := ccuc.payments.ListInvoices(contract.StripeSubscription.ID)
	if err != nil {
		return err
	}

	fine := ccuc.payments.CalculateRemainingValueSubscription(invoices, contract.Amount)
	_, err = ccuc.payments.FineResponsible(contract, int64(fine))
	if err != nil {
		return nil
	}

	err = ccuc.repositories.ContractRepository.Cancel(id)
	if err != nil {
		return err
	}

	return nil
}
