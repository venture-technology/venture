package usecase

import (
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
	"github.com/venture-technology/venture/pkg/utils"
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
		KidName:   contracts.Kid.Name,
		Period:    contracts.Kid.Shift,
		CreatedAt: contracts.CreatedAt,
		ExpireAt:  contracts.ExpireAt,
		Driver: value.GetDriverContract{
			ID:    contracts.Driver.ID,
			Name:  contracts.Driver.Name,
			Email: contracts.Driver.Email,
			Address: utils.BuildAddress(
				contracts.Driver.Address.Street,
				contracts.Driver.Address.Number,
				contracts.Driver.Address.Complement,
				contracts.Driver.Address.Zip,
			),
			Phone:        contracts.Driver.Phone,
			ProfileImage: contracts.Driver.ProfileImage,
		},
		Responsible: value.GetParentContract{
			ID:    contracts.Kid.Responsible.ID,
			Name:  contracts.Kid.Responsible.Name,
			Email: contracts.Kid.Responsible.Email,
			Phone: contracts.Kid.Responsible.Phone,
		},
	}
}
