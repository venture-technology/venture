package usecase

import (
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
	"github.com/venture-technology/venture/pkg/utils"
)

type ListDriverContractsUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewListDriverContractsUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *ListDriverContractsUseCase {
	return &ListDriverContractsUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (ldcuc *ListDriverContractsUseCase) ListDriverContracts(cnh string) ([]value.DriverListContracts, error) {
	contracts, err := ldcuc.repositories.ContractRepository.GetByDriver(cnh)
	if err != nil {
		return []value.DriverListContracts{}, err
	}
	response := []value.DriverListContracts{}
	for _, contract := range contracts {
		response = append(response, buildDriverListContracts(&contract))
	}
	return response, nil
}

func buildDriverListContracts(contracts *entity.Contract) value.DriverListContracts {
	return value.DriverListContracts{
		ID:     contracts.ID,
		UUID:   contracts.UUID,
		Status: contracts.Status,
		Amount: contracts.Amount,
		School: value.GetSchoolContract{
			ID:   contracts.School.ID,
			Name: contracts.School.Name,
			Address: utils.BuildAddress(
				contracts.School.Address.Street,
				contracts.School.Address.Number,
				contracts.School.Address.Complement,
				contracts.School.Address.Zip,
			),
			Phone:        contracts.School.Phone,
			ProfileImage: contracts.School.ProfileImage,
		},
		Responsible: value.GetParentContract{
			ID:    contracts.Responsible.ID,
			Name:  contracts.Responsible.Name,
			Email: contracts.Responsible.Email,
			Address: utils.BuildAddress(
				contracts.Responsible.Address.Street,
				contracts.Responsible.Address.Number,
				contracts.Responsible.Address.Complement,
				contracts.Responsible.Address.Zip,
			),
			Phone:        contracts.Responsible.Phone,
			ProfileImage: contracts.Responsible.ProfileImage,
		},
		Kid: value.GetKidContract{
			ID:           contracts.Kid.ID,
			Name:         contracts.Kid.Name,
			Period:       contracts.Kid.Shift,
			ProfileImage: contracts.Kid.ProfileImage,
		},
		CreatedAt: contracts.CreatedAt,
		ExpireAt:  contracts.UpdatedAt,
	}
}
