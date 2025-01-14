package usecase

import (
	"fmt"

	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
)

type GetChildUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewGetChildUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *GetChildUseCase {
	return &GetChildUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (gcuc *GetChildUseCase) GetChild(rg *string) (value.GetChild, error) {
	child, err := gcuc.repositories.ChildRepository.Get(rg)
	if err != nil {
		return value.GetChild{}, err
	}
	return value.GetChild{
		ID:              child.ID,
		Name:            child.Name,
		RG:              child.RG,
		ResponsibleName: child.Responsible.Name,
		Address: fmt.Sprintf(
			"%s, %s, %s",
			child.Responsible.Address.Street,
			child.Responsible.Address.Number,
			child.Responsible.Address.ZIP,
		),
		Period:       child.Shift,
		ProfileImage: child.ProfileImage,
	}, nil
}
