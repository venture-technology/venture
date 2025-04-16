package usecase

import (
	"fmt"

	"github.com/venture-technology/venture/internal/entity"
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

	fields := []string{"street", "number", "complement", "zip"}
	err = utils.ValidateRequiredGroup(attributes, fields)
	if err != nil {
		return err
	}

	driver, err := uduc.repositories.DriverRepository.Get(cnh)
	if err != nil {
		return err
	}

	err = uduc.validateProvincy(driver, attributes)
	if err != nil {
		return err
	}

	return uduc.repositories.DriverRepository.Update(cnh, attributes)
}

func (uduc *UpdateDriverUseCase) validateProvincy(driver *entity.Driver, attrs map[string]interface{}) error {
	_, exists := attrs["state"]
	if exists {
		_, ok := attrs["state"].(string)
		if !ok {
			return fmt.Errorf("state type error")
		}

		hasContract, err := uduc.repositories.ContractRepository.DriverHasEnableContract(driver.CNH)
		if err != nil {
			return err
		}

		if hasContract {
			return fmt.Errorf("impossible change provincy when has enable contract")
		}
	}

	return nil
}
