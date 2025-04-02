package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/mocks"
)

func TestUpdateResponsibleUsecase_UpdateResponsible(t *testing.T) {
	t.Run("if someone try send unknown field to update", func(t *testing.T) {
		repository := mocks.NewResponsibleRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewUpdateResponsibleUseCase(
			&persistence.PostgresRepositories{
				ResponsibleRepository: repository,
			},
			logger,
		)

		err := usecase.UpdateResponsible("123", map[string]interface{}{
			"cpf": "123",
		})

		assert.EqualError(t, err, "chaves não permitidas: cpf")
		assert.Error(t, err)
	})

	t.Run("when responsible send unknown address to update", func(t *testing.T) {
		responsibleRepository := mocks.NewResponsibleRepository(t)
		contractRepository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewUpdateResponsibleUseCase(
			&persistence.PostgresRepositories{
				ResponsibleRepository: responsibleRepository,
				ContractRepository:    contractRepository,
			},
			logger,
		)

		err := usecase.UpdateResponsible("123", map[string]interface{}{
			"street": "para lanches",
		})

		assert.EqualError(t, err, "os seguintes campos são obrigatórios: number, complement, zip")
		assert.Error(t, err)
	})

	t.Run("when responsible has enable contract and try change his address", func(t *testing.T) {
		responsibleRepository := mocks.NewResponsibleRepository(t)
		contractRepository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewUpdateResponsibleUseCase(
			&persistence.PostgresRepositories{
				ResponsibleRepository: responsibleRepository,
				ContractRepository:    contractRepository,
			},
			logger,
		)

		contractRepository.On("ResponsibleHasEnableContract", mock.Anything).Return(true, nil)

		err := usecase.UpdateResponsible("123", map[string]interface{}{
			"street":     "Rua Veredas",
			"number":     "17",
			"complement": "Bloco B",
			"zip":        "02344328",
		})

		assert.EqualError(t, err, "impossivel trocar endereco possuindo contrato ativo, contate o atendimento")
		assert.Error(t, err)
	})

	t.Run("when proxy to update returns error", func(t *testing.T) {
		responsibleRepository := mocks.NewResponsibleRepository(t)
		contractRepository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewUpdateResponsibleUseCase(
			&persistence.PostgresRepositories{
				ResponsibleRepository: responsibleRepository,
				ContractRepository:    contractRepository,
			},
			logger,
		)

		contractRepository.On("ResponsibleHasEnableContract", mock.Anything).Return(false, nil)
		responsibleRepository.On("Update", mock.Anything, mock.Anything).Return(errors.New("database error"))

		err := usecase.UpdateResponsible("123", map[string]interface{}{
			"phone":      "123",
			"street":     "Rua Veredas",
			"number":     "17",
			"complement": "Bloco B",
			"zip":        "02344328",
		})

		assert.EqualError(t, err, "database error")
	})

	t.Run("when proxy to update give success without address", func(t *testing.T) {
		responsibleRepository := mocks.NewResponsibleRepository(t)
		contractRepository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewUpdateResponsibleUseCase(
			&persistence.PostgresRepositories{
				ResponsibleRepository: responsibleRepository,
				ContractRepository:    contractRepository,
			},
			logger,
		)

		responsibleRepository.On("Update", mock.Anything, mock.Anything).Return(nil)

		err := usecase.UpdateResponsible("123", map[string]interface{}{
			"phone": "123",
		})

		assert.Nil(t, err)
	})
}
