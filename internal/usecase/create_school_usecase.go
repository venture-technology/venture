package usecase

import (
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/pkg/utils"
)

type CreateSchoolUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewCreateSchoolUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *CreateSchoolUseCase {
	return &CreateSchoolUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (csuc *CreateSchoolUseCase) CreateSchool(school *entity.School) error {
	err := utils.ValidadeZip(school.Address.Zip)
	if err != nil {
		return err
	}

	return csuc.repositories.SchoolRepository.Create(school)
}
