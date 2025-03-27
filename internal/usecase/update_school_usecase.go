package usecase

import (
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
	"github.com/venture-technology/venture/pkg/utils"
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
	err := utils.ValidateUpdate(attributes, value.SchoollAllowedKeys)
	if err != nil {
		return err
	}

	return usuc.repositories.SchoolRepository.Update(cnpj, attributes)
}
