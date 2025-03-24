package usecase

import (
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
	"github.com/venture-technology/venture/pkg/utils"
)

type UpdateDriverUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewUpdateDriverUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *UpdateDriverUseCase {
	return &UpdateDriverUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (uduc *UpdateDriverUseCase) UpdateDriver(cnh string, attributes map[string]interface{}) error {
	err := utils.ValidateUpdate(attributes, value.DriverAllowedKeys)
	if err != nil {
		return err
	}

	return uduc.repositories.DriverRepository.Update(cnh, attributes)
}
