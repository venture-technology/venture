package usecase

import (
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
)

type GetKidUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewGetKidUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *GetKidUseCase {
	return &GetKidUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (gcuc *GetKidUseCase) GetKid(rg *string) (value.GetKid, error) {
	kid, err := gcuc.repositories.KidRepository.Get(rg)
	if err != nil {
		return value.GetKid{}, err
	}
	return value.GetKid{
		ID:              kid.ID,
		Name:            kid.Name,
		RG:              kid.RG,
		ResponsibleName: kid.Responsible.Name,
		Address:         kid.Responsible.Address.GetFullAddress(),
		Period:          kid.Shift,
		ProfileImage:    kid.ProfileImage,
	}, nil
}
