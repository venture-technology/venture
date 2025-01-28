package usecase

import (
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
)

type ListKidsUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewListKidsUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *ListKidsUseCase {
	return &ListKidsUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (lcuc *ListKidsUseCase) ListKids(cpf *string) ([]value.ListKid, error) {
	kids, err := lcuc.repositories.KidRepository.FindAll(cpf)
	if err != nil {
		return []value.ListKid{}, err
	}
	response := []value.ListKid{}
	for _, kid := range kids {
		response = append(response, buildListKid(&kid))
	}
	return response, nil
}

func buildListKid(kid *entity.Kid) value.ListKid {
	return value.ListKid{
		ID:           kid.ID,
		Name:         kid.Name,
		Period:       kid.Shift,
		ProfileImage: kid.ProfileImage,
	}
}
