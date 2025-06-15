package usecase

import (
	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
)

type GetContractUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewGetContractUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *GetContractUseCase {
	return &GetContractUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (gcuc *GetContractUseCase) GetContract(uuid uuid.UUID) (value.GetContractOutput, error) {
	contract, err := gcuc.repositories.ContractRepository.GetByUUID(uuid)
	if err != nil {
		return value.GetContractOutput{}, err
	}
	return value.GetContractOutput{
		ID:                   contract.ID,
		UUID:                 contract.UUID,
		Status:               contract.Status,
		SigningURL:           contract.SigningURL,
		StripeSubscriptionID: contract.StripeSubscriptionID,
		CreatedAt:            contract.CreatedAt,
		ExpiredAt:            contract.ExpireAt,
		Driver: value.GetDriverContract{
			ID:           contract.Driver.ID,
			Name:         contract.Driver.Name,
			Email:        contract.Driver.Email,
			Phone:        contract.Driver.Phone,
			ProfileImage: contract.Driver.ProfileImage,
		},
		Responsible: value.GetParentContract{
			ID:           contract.Responsible.ID,
			Name:         contract.Responsible.Name,
			Email:        contract.Responsible.Email,
			Address:      contract.Responsible.Address.GetFullAddress(),
			Phone:        contract.Responsible.Phone,
			ProfileImage: contract.Responsible.ProfileImage,
		},
		Kid: value.GetKidContract{
			ID:           contract.Kid.ID,
			Name:         contract.Kid.Name,
			Period:       contract.Kid.Shift,
			ProfileImage: contract.Kid.ProfileImage,
		},
		School: value.GetSchoolContract{
			ID:           contract.School.ID,
			Name:         contract.School.Name,
			Address:      contract.School.Address.GetFullAddress(),
			Phone:        contract.School.Phone,
			ProfileImage: contract.School.ProfileImage,
		},
	}, nil
}
