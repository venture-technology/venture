package usecase

import (
	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/domain/service/adapters"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
	"github.com/venture-technology/venture/pkg/utils"
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
	// 1. buscar contrato no s3
	// 2. adicionar contrato no retorno
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
		KidName:   contract.Kid.Name,
		Period:    contract.Kid.Shift,
		Amount:    contract.Amount,
		Record:    contract.Record,
		CreatedAt: contract.CreatedAt,
		ExpireAt:  contract.ExpireAt,
		Driver: value.GetDriverContract{
			ID:    contract.Driver.ID,
			Name:  contract.Driver.Name,
			Email: contract.Driver.Email,
			Address: utils.BuildAddress(
				contract.Driver.Address.Street,
				contract.Driver.Address.Number,
				contract.Driver.Address.Complement,
				contract.Driver.Address.Zip,
			),
			Phone:        contract.Driver.Phone,
			ProfileImage: contract.Driver.ProfileImage,
		},
		School: value.GetSchoolContract{
			ID:   contract.School.ID,
			Name: contract.School.Name,
			Address: utils.BuildAddress(
				contract.School.Address.Street,
				contract.School.Address.Number,
				contract.School.Address.Complement,
				contract.School.Address.Zip,
			),
			Phone:        contract.School.Phone,
			ProfileImage: contract.School.ProfileImage,
		},
		Responsible: value.GetParentContract{
			ID:    contract.Kid.Responsible.ID,
			Name:  contract.Kid.Responsible.Name,
			Email: contract.Kid.Responsible.Email,
			Address: utils.BuildAddress(
				contract.Kid.Responsible.Address.Street,
				contract.Kid.Responsible.Address.Number,
				contract.Kid.Responsible.Address.Complement,
				contract.Kid.Responsible.Address.Zip,
			),
			Phone:        contract.Kid.Responsible.Phone,
			ProfileImage: contract.Kid.Responsible.ProfileImage,
		},
		Invoices: invoices,
	}, nil
}
