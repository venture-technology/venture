package usecase

import (
	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/domain/service/adapters"
	"github.com/venture-technology/venture/internal/entity"
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

var seatsEvent = map[string]func(ccuc *CancelContractUseCase, contract *entity.Contract) error{
	"morning": func(ccuc *CancelContractUseCase, contract *entity.Contract) error {
		return ccuc.repositories.DriverRepository.Update(contract.Driver.CNH, map[string]interface{}{
			"seats_remaining": contract.Driver.Seats.Remaining + 1,
			"seats_morning":   contract.Driver.Seats.Morning + 1,
		})
	},
	"afternoon": func(ccuc *CancelContractUseCase, contract *entity.Contract) error {
		return ccuc.repositories.DriverRepository.Update(contract.Driver.CNH, map[string]interface{}{
			"seats_remaining": contract.Driver.Seats.Remaining + 1,
			"seats_afternoon": contract.Driver.Seats.Afternoon + 1,
		})
	},
	"night": func(ccuc *CancelContractUseCase, contract *entity.Contract) error {
		return ccuc.repositories.DriverRepository.Update(contract.Driver.CNH, map[string]interface{}{
			"seats_remaining": contract.Driver.Seats.Remaining + 1,
			"seats_night":     contract.Driver.Seats.Night + 1,
		})
	},
}

func (ccuc *CancelContractUseCase) CancelContract(id uuid.UUID) error {
	contract, err := ccuc.repositories.ContractRepository.Get(id)
	if err != nil {
		return err
	}

	invoices, err := ccuc.adapters.PaymentsService.ListInvoices(contract.StripeSubscription.ID)
	if err != nil {
		return err
	}

	fine := ccuc.adapters.PaymentsService.CalculateRemainingValueSubscription(invoices, contract.Amount)
	_, err = ccuc.adapters.PaymentsService.FineResponsible(contract, int64(fine))
	if err != nil {
		return nil
	}

	err = ccuc.repositories.ContractRepository.Cancel(id)
	if err != nil {
		return err
	}

	return seatsEvent[contract.Child.Shift](ccuc, contract)
}
