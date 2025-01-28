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

func (ldcuc *ListDriverContractsUseCase) ListDriverContracts(cnh *string) ([]value.DriverListContracts, error) {
	contracts, err := ldcuc.repositories.ContractRepository.FindAllByCnh(cnh)
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
		Record: contracts.Record,
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
		Kid: value.GetKidContract{
			ID:           contracts.Kid.ID,
			Name:         contracts.Kid.Name,
			Period:       contracts.Kid.Shift,
			ProfileImage: contracts.Kid.ProfileImage,
		},
		Responsible: value.GetParentContract{
			ID:           contracts.Kid.Responsible.ID,
			Name:         contracts.Kid.Responsible.Name,
			Phone:        contracts.Kid.Responsible.Phone,
			Email:        contracts.Kid.Responsible.Email,
			ProfileImage: contracts.Kid.Responsible.ProfileImage,
			Address: utils.BuildAddress(
				contracts.Kid.Responsible.Address.Street,
				contracts.Kid.Responsible.Address.Number,
				contracts.Kid.Responsible.Address.Complement,
				contracts.Kid.Responsible.Address.Zip,
			),
		},
		CreatedAt: contracts.CreatedAt,
		ExpireAt:  contracts.ExpireAt,
	}
}
