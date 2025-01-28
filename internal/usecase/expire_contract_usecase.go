package usecase

import (
	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

var expireSeat = map[string]func(ecuc *ExpireContractUseCase, contract *entity.Contract) error{
	"morning": func(ecuc *ExpireContractUseCase, contract *entity.Contract) error {
		return ecuc.repositories.DriverRepository.Update(contract.Driver.CNH, map[string]interface{}{
			"seats_remaining": contract.Driver.Seats.Remaining + 1,
			"seats_morning":   contract.Driver.Seats.Morning + 1,
		})
	},
	"afternoon": func(ecuc *ExpireContractUseCase, contract *entity.Contract) error {
		return ecuc.repositories.DriverRepository.Update(contract.Driver.CNH, map[string]interface{}{
			"seats_remaining": contract.Driver.Seats.Remaining + 1,
			"seats_afternoon": contract.Driver.Seats.Afternoon + 1,
		})
	},
	"night": func(ecuc *ExpireContractUseCase, contract *entity.Contract) error {
		return ecuc.repositories.DriverRepository.Update(contract.Driver.CNH, map[string]interface{}{
			"seats_remaining": contract.Driver.Seats.Remaining + 1,
			"seats_night":     contract.Driver.Seats.Night + 1,
		})
	},
}

type ExpireContractUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewExpireContractUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *ExpireContractUseCase {
	return &ExpireContractUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (ecuc *ExpireContractUseCase) ExpireContract(uuid uuid.UUID) error {
	contract, err := ecuc.repositories.ContractRepository.Get(uuid)
	if err != nil {
		return err
	}

	err = ecuc.repositories.ContractRepository.Expired(uuid)
	if err != nil {
		return err
	}

	return expireSeat[contract.Kid.Shift](ecuc, contract)
}
