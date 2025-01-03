package usecase

import (
	"fmt"

	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
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
			Address: fmt.Sprintf(
				"%s, %s, %s",
				contracts.School.Address.Street,
				contracts.School.Address.Number,
				contracts.School.Address.ZIP,
			),
			Phone:        contracts.School.Phone,
			ProfileImage: contracts.School.ProfileImage,
		},
		Child: value.GetChildContract{
			ID:           contracts.Child.ID,
			Name:         contracts.Child.Name,
			Period:       contracts.Child.Shift,
			ProfileImage: contracts.Child.ProfileImage,
		},
		Responsible: value.GetParentContract{
			ID:           contracts.Child.Responsible.ID,
			Name:         contracts.Child.Responsible.Name,
			Phone:        contracts.Child.Responsible.Phone,
			Email:        contracts.Child.Responsible.Email,
			ProfileImage: contracts.Child.Responsible.ProfileImage,
			Address: fmt.Sprintf(
				"%s, %s, %s",
				contracts.Child.Responsible.Address.Street,
				contracts.Child.Responsible.Address.Number,
				contracts.Child.Responsible.Address.ZIP,
			),
		},
		CreatedAt: contracts.CreatedAt,
		ExpireAt:  contracts.ExpireAt,
	}
}
