package usecase

import (
	"github.com/venture-technology/venture/internal/domain/service/adapters"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
	"github.com/venture-technology/venture/pkg/utils"
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

func (lruc *ListResponsibleContractsUseCase) ListResponsibleContracts(cpf string) ([]value.ResponsibleListContracts, error) {
	contracts, err := lruc.repositories.ContractRepository.GetByResponsible(cpf)
	if err != nil {
		return []value.ResponsibleListContracts{}, err
	}
	response := []value.ResponsibleListContracts{}
	for _, contract := range contracts {
		response = append(response, buildResponsibleListContracts(&contract))
	}
	return response, nil
}

func buildResponsibleListContracts(contracts *entity.EnableContract) value.ResponsibleListContracts {
	return value.ResponsibleListContracts{
		ID:        contracts.ID,
		UUID:      contracts.UUID,
		Status:    contracts.Status,
		KidName:   contracts.Kid.Name,
		Period:    contracts.Kid.Shift,
		Amount:    contracts.Amount,
		CreatedAt: contracts.CreatedAt,
		ExpireAt:  contracts.ExpireAt,
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
		Driver: value.GetDriverContract{
			ID:           contracts.Driver.ID,
			Name:         contracts.Driver.Name,
			Email:        contracts.Driver.Email,
			Phone:        contracts.Driver.Phone,
			ProfileImage: contracts.Driver.ProfileImage,
		},
	}
}
