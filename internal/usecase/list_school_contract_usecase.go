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

func (lscuc *ListSchoolContractUseCase) ListSchoolContract(cnpj string) ([]value.SchoolListContracts, error) {
	contracts, err := lscuc.repositories.ContractRepository.GetBySchool(cnpj)
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
		Status:    contracts.Status,
		KidName:   contracts.Kid.Name,
		Period:    contracts.Kid.Shift,
		Amount:    contracts.Amount,
		UUID:      contracts.UUID,
		CreatedAt: contracts.CreatedAt,
		ExpireAt:  contracts.ExpireAt,
		Driver: value.GetDriverContract{
			ID:           contracts.Driver.ID,
			Name:         contracts.Driver.Name,
			Email:        contracts.Driver.Email,
			Phone:        contracts.Driver.Phone,
			ProfileImage: contracts.Driver.ProfileImage,
		},
		Responsible: value.GetParentContract{
			ID:           contracts.Responsible.ID,
			Name:         contracts.Responsible.Name,
			Email:        contracts.Responsible.Email,
			Phone:        contracts.Responsible.Phone,
			ProfileImage: contracts.Responsible.ProfileImage,
			Address: utils.BuildAddress(
				contracts.Responsible.Address.Street,
				contracts.Responsible.Address.Number,
				contracts.Responsible.Address.Complement,
				contracts.Responsible.Address.Zip,
			),
		},
	}
}
