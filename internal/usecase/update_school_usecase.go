package usecase

import (
	"fmt"

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

	if _, exists := attributes["password"]; exists {
		ok, errors := utils.ValidatePassword(attributes["password"].(string))
		if !ok {
			return fmt.Errorf(errors)
		}

		hash, err := utils.MakeHash(attributes["password"].(string))
		if err != nil {
			return fmt.Errorf("error hashing password: %w", err)
		}

		attributes["password"] = hash
	}

	return usuc.repositories.SchoolRepository.Update(cnpj, attributes)
}
