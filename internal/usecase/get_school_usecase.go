package usecase

import (
	"fmt"

	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
)

type GetSchoolUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewGetSchoolUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *GetSchoolUseCase {
	return &GetSchoolUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (gsuc *GetSchoolUseCase) GetSchool(cnpj string) (value.GetSchool, error) {
	school, err := gsuc.repositories.SchoolRepository.Get(cnpj)
	if err != nil {
		return value.GetSchool{}, err
	}
	return value.GetSchool{
		ID:    school.ID,
		Name:  school.Name,
		Email: school.Email,
		Phone: school.Phone,
		Address: fmt.Sprintf(
			"%s, %s, %s",
			school.Address.Street,
			school.Address.Number,
			school.Address.ZIP,
		),
		ProfileImage: school.ProfileImage,
		CreatedAt:    school.CreatedAt,
	}, nil
}
