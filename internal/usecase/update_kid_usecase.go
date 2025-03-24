package usecase

import (
	"fmt"

	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
	"github.com/venture-technology/venture/pkg/utils"
)

type UpdateKidUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewUpdateKidUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *UpdateKidUseCase {
	return &UpdateKidUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (ucuc *UpdateKidUseCase) UpdateKid(rg string, attributes map[string]interface{}) error {
	err := utils.ValidateUpdate(attributes, value.KidAllowedKeys)
	if err != nil {
		return err
	}

	_, exists := attributes["shift"]
	if exists {
		shift, ok := attributes["shift"].(string)
		if !ok {
			return fmt.Errorf("shift invalido")
		}

		err = ucuc.ValidateShift(shift)
		if err != nil {
			return err
		}

		// checando se esse kid já tem contrato
		kidHasContract, err := ucuc.repositories.ContractRepository.KidHasContract(rg)
		if err != nil {
			return err
		}

		if kidHasContract {
			return fmt.Errorf("impossivel trocar horario possuindo contrato, contate o atendimento")
		}

	}

	return ucuc.repositories.KidRepository.Update(rg, attributes)
}

func (ucuc *UpdateKidUseCase) ValidateShift(period string) error {
	_, exists := value.Shifts[period]
	if !exists {
		return fmt.Errorf("shift inexistente")
	}
	return nil
}
