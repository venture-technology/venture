package usecase

import (
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type UpdateSchoolUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewUpdateSchoolUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *UpdateSchoolUseCase {
	return &UpdateSchoolUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (usuc *UpdateSchoolUseCase) UpdateSchool(cnpj string, attributes map[string]interface{}) error {
	return usuc.repositories.SchoolRepository.Update(cnpj, attributes)
}
