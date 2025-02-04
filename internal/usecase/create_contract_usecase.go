package usecase

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/domain/service/adapters"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/pkg/utils"
)

var createSeat = map[string]func(ccuc *CreateContractUseCase, contract *entity.Contract) error{
	"morning": func(ccuc *CreateContractUseCase, contract *entity.Contract) error {
		return ccuc.repositories.DriverRepository.Update(contract.Driver.CNH, map[string]interface{}{
			"seats_remaining": contract.Driver.Seats.Remaining - 1,
			"seats_morning":   contract.Driver.Seats.Morning - 1,
		})
	},
	"afternoon": func(ccuc *CreateContractUseCase, contract *entity.Contract) error {
		return ccuc.repositories.DriverRepository.Update(contract.Driver.CNH, map[string]interface{}{
			"seats_remaining": contract.Driver.Seats.Remaining - 1,
			"seats_afternoon": contract.Driver.Seats.Afternoon - 1,
		})
	},
	"night": func(ccuc *CreateContractUseCase, contract *entity.Contract) error {
		return ccuc.repositories.DriverRepository.Update(contract.Driver.CNH, map[string]interface{}{
			"seats_remaining": contract.Driver.Seats.Remaining - 1,
			"seats_night":     contract.Driver.Seats.Night - 1,
		})
	},
}

type CreateContractUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
	adapters     adapters.Adapters
}

func NewCreateContractUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
	adapters adapters.Adapters,
) *CreateContractUseCase {
	return &CreateContractUseCase{
		repositories: repositories,
		logger:       logger,
		adapters:     adapters,
	}
}

func (ccuc *CreateContractUseCase) CreateContract(contract *entity.Contract) error {
	var err error
	contract.Amount, err = ccuc.calcAmount(contract)
	if err != nil {
		return err
	}

	if hasAmount := contract.ValidateAmount(); !hasAmount {
		return fmt.Errorf("contract amount is invalid")
	}

	contract, err = ccuc.createStripeItems(contract)
	if err != nil {
		return err
	}

	err = ccuc.repositories.ContractRepository.Create(contract)
	if err != nil {
		return err
	}

	return createSeat[contract.Kid.Shift](ccuc, contract)
}

func (ccuc *CreateContractUseCase) calcAmount(contract *entity.Contract) (float64, error) {
	dist, err := ccuc.adapters.AddressService.GetDistance(
		buildResponsibleAddress(&contract.Kid.Responsible),
		buildSchoolAddress(&contract.School),
	)
	if err != nil {
		return 0, err
	}
	return utils.CalculateContract(*dist, float64(contract.Driver.Amount)), nil
}

func (ccuc *CreateContractUseCase) createStripeItems(contract *entity.Contract) (*entity.Contract, error) {
	prodt, err := ccuc.adapters.PaymentsService.CreateProduct(contract)
	if err != nil {
		return nil, err
	}

	contract.StripeSubscription.Product = prodt.ID
	pr, err := ccuc.adapters.PaymentsService.CreatePrice(contract)
	if err != nil {
		return nil, err
	}

	contract.StripeSubscription.Price = pr.ID
	subs, err := ccuc.adapters.PaymentsService.CreateSubscription(contract)
	if err != nil {
		return nil, err
	}

	contract.StripeSubscription.ID = subs.ID
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	contract.Record = id
	return contract, nil
}

func buildResponsibleAddress(responsible *entity.Responsible) string {
	return fmt.Sprintf(
		"%s,%s,%s",
		responsible.Address.Street,
		responsible.Address.Number,
		responsible.Address.Zip,
	)
}

func buildSchoolAddress(school *entity.School) string {
	return fmt.Sprintf(
		"%s,%s,%s",
		school.Address.Street,
		school.Address.Number,
		school.Address.Zip,
	)
}
