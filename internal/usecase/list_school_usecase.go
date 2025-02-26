package usecase

import (
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
	"github.com/venture-technology/venture/pkg/utils"
)

type ListSchoolUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewListSchoolUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *ListSchoolUseCase {
	return &ListSchoolUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (lsuc *ListSchoolUseCase) ListSchool() ([]value.ListSchool, error) {
	schools, err := lsuc.repositories.SchoolRepository.FindAll()
	if err != nil {
		return nil, err
	}
	response := []value.ListSchool{}
	for _, school := range schools {
		response = append(response, buildListSchool(school))
	}
	return response, nil
}

func buildListSchool(school entity.School) value.ListSchool {
	return value.ListSchool{
		ID:           school.ID,
		Name:         school.Name,
		Email:        school.Email,
		Phone:        school.Phone,
		ProfileImage: school.ProfileImage,
		CreatedAt:    school.CreatedAt,
		Address: utils.BuildAddress(
			school.Address.Street,
			school.Address.Number,
			school.Address.Complement,
			school.Address.Zip,
		),
		City:   school.City,
		States: school.States,
	}
}
