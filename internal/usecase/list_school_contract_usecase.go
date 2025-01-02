package usecase

import (
	"fmt"

	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
)

type ListSchoolContractUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewListSchoolContractUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *ListSchoolContractUseCase {
	return &ListSchoolContractUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (lscuc *ListSchoolContractUseCase) ListSchoolContract(cnpj *string) ([]value.SchoolListContracts, error) {
	contracts, err := lscuc.repositories.ContractRepository.FindAllByCnpj(cnpj)
	if err != nil {
		return nil, err
	}
	response := []value.SchoolListContracts{}
	for _, contract := range contracts {
		response = append(response, buildSchoolListContracts(&contract))
	}
	return response, nil
}

func buildSchoolListContracts(contracts *entity.Contract) value.SchoolListContracts {
	return value.SchoolListContracts{
		ID:        contracts.ID,
		Record:    contracts.Record,
		Status:    contracts.Status,
		Amount:    contracts.Amount,
		ChildName: contracts.Child.Name,
		Period:    contracts.Child.Shift,
		CreatedAt: contracts.CreatedAt,
		ExpireAt:  contracts.ExpireAt,
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
		Responsible: value.GetParentContract{
			ID:    contracts.Child.Responsible.ID,
			Name:  contracts.Child.Responsible.Name,
			Email: contracts.Child.Responsible.Email,
			Phone: contracts.Child.Responsible.Phone,
		},
	}
}
