package usecase

import (
	"fmt"

	"github.com/venture-technology/venture/internal/domain/service/adapters"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
)

type ListResponsibleContractsUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
	adapters     adapters.Adapters
}

func NewListResponsibleContractsUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
	adapters adapters.Adapters,
) *ListResponsibleContractsUseCase {
	return &ListResponsibleContractsUseCase{
		repositories: repositories,
		logger:       logger,
		adapters:     adapters,
	}
}

func (lruc *ListResponsibleContractsUseCase) ListResponsibleContracts(cpf *string) ([]value.ResponsibleListContracts, error) {
	contracts, err := lruc.repositories.ContractRepository.FindAllByCpf(cpf)
	if err != nil {
		return []value.ResponsibleListContracts{}, err
	}
	response := []value.ResponsibleListContracts{}
	for _, contract := range contracts {
		response = append(response, buildResponsibleListContracts(&contract))
	}
	return response, nil
}

func buildResponsibleListContracts(contracts *entity.Contract) value.ResponsibleListContracts {
	return value.ResponsibleListContracts{
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
		Driver: value.GetDriverContract{
			ID:    contracts.Driver.ID,
			Name:  contracts.Driver.Name,
			Email: contracts.Driver.Email,
			Address: fmt.Sprintf(
				"%s, %s, %s",
				contracts.Driver.Address.Street,
				contracts.Driver.Address.Number,
				contracts.Driver.Address.ZIP,
			),
			Phone:        contracts.Driver.Phone,
			ProfileImage: contracts.Driver.ProfileImage,
		},
		ChildName: contracts.Child.Name,
		Period:    contracts.Child.Shift,
		CreatedAt: contracts.CreatedAt,
		ExpireAt:  contracts.ExpireAt,
	}
}
