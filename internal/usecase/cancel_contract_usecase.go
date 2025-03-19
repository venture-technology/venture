package usecase

import (
	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/domain/service/adapters"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type CancelContractUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
	adapters     adapters.Adapters
}

func NewCancelContractUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
	adapters adapters.Adapters,
) *CancelContractUseCase {
	return &CancelContractUseCase{
		repositories: repositories,
		logger:       logger,
		adapters:     adapters,
	}
}

func (ccuc *CancelContractUseCase) CancelContract(uuid uuid.UUID) error {
	contract, err := ccuc.repositories.ContractRepository.Get(uuid)
	if err != nil {
		return err
	}

	responsible, err := ccuc.repositories.ResponsibleRepository.Get(contract.ResponsibleCPF)
	if err != nil {
		return err
	}

	invoices, err := ccuc.adapters.PaymentsService.ListInvoices(contract.StripeSubscriptionID)
	if err != nil {
		return err
	}

	fine := ccuc.adapters.PaymentsService.CalculateRemainingValueSubscription(invoices, contract.Amount)

	_, err = ccuc.adapters.PaymentsService.FineResponsible(responsible.CustomerId, responsible.PaymentMethodId, int64(fine))
	if err != nil {
		return err
	}

	return ccuc.repositories.ContractRepository.Cancel(uuid)
}
