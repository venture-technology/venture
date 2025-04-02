package usecase

import (
	"fmt"

	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
	"github.com/venture-technology/venture/pkg/utils"
)

type UpdateResponsibleUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewUpdateResponsibleUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *UpdateResponsibleUseCase {
	return &UpdateResponsibleUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (uruc *UpdateResponsibleUseCase) UpdateResponsible(cpf string, attributes map[string]interface{}) error {
	err := utils.ValidateUpdate(attributes, value.ResponsibleAllowedKeys)
	if err != nil {
		return err
	}

	fields := []string{"street", "number", "complement", "zip"}
	err = utils.ValidateRequiredGroup(attributes, fields)
	if err != nil {
		return err
	}

	if utils.KeysExist(attributes, fields) {
		exists, err := uruc.repositories.ContractRepository.ResponsibleHasEnableContract(cpf)
		if err != nil {
			return err
		}

		if exists {
			return fmt.Errorf("impossivel trocar endereco possuindo contrato ativo, contate o atendimento")
		}
	}

	return uruc.repositories.ResponsibleRepository.Update(cpf, attributes)
}
