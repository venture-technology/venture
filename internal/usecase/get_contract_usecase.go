package usecase

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/domain/service/adapters"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
)

type GetContractUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
	adapters     adapters.Adapters
}

func NewGetContractUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
	adapters adapters.Adapters,
) *GetContractUseCase {
	return &GetContractUseCase{
		repositories: repositories,
		logger:       logger,
		adapters:     adapters,
	}
}

func (gcuc *GetContractUseCase) GetContract(uuid uuid.UUID) (value.GetContract, error) {
	contract, err := gcuc.repositories.ContractRepository.Get(uuid)
	if err != nil {
		return value.GetContract{}, err
	}
	invoices, err := gcuc.adapters.PaymentsService.ListInvoices(contract.StripeSubscription.ID)
	if err != nil {
		return value.GetContract{}, err
	}
	return value.GetContract{
		ID:        contract.ID,
		Status:    contract.Status,
		ChildName: contract.Child.Name,
		Period:    contract.Child.Shift,
		Amount:    contract.Amount,
		Record:    contract.Record,
		CreatedAt: contract.CreatedAt,
		ExpireAt:  contract.ExpireAt,
		Driver: value.GetDriverContract{
			ID:    contract.Driver.ID,
			Name:  contract.Driver.Name,
			Email: contract.Driver.Email,
			Address: fmt.Sprintf(
				"%s, %s, %s",
				contract.Driver.Address.Street,
				contract.Driver.Address.Number,
				contract.Driver.Address.ZIP,
			),
			Phone:        contract.Driver.Phone,
			ProfileImage: contract.Driver.ProfileImage,
		},
		School: value.GetSchoolContract{
			ID:   contract.School.ID,
			Name: contract.School.Name,
			Address: fmt.Sprintf(
				"%s, %s, %s",
				contract.School.Address.Street,
				contract.School.Address.Number,
				contract.School.Address.ZIP,
			),
			Phone:        contract.School.Phone,
			ProfileImage: contract.School.ProfileImage,
		},
		Responsible: value.GetParentContract{
			ID:    contract.Child.Responsible.ID,
			Name:  contract.Child.Responsible.Name,
			Email: contract.Child.Responsible.Email,
			Address: fmt.Sprintf(
				"%s, %s, %s",
				contract.Child.Responsible.Address.Street,
				contract.Child.Responsible.Address.Number,
				contract.Child.Responsible.Address.ZIP,
			),
			Phone:        contract.Child.Responsible.Phone,
			ProfileImage: contract.Child.Responsible.ProfileImage,
		},
		Invoices: invoices,
	}, nil
}
