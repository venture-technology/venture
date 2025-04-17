package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/mocks"
)

func TestUpdateDriverUsecase_UpdateDriver(t *testing.T) {
	// atualizar com uma chave nao permitida
	t.Run("when driver try change unknown key", func(t *testing.T) {
		repository := mocks.NewDriverRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewUpdateDriverUseCase(
			&persistence.PostgresRepositories{
				DriverRepository: repository,
			},
			logger,
		)

		err := usecase.UpdateDriver("123", map[string]interface{}{
			"teste": "123",
		})

		assert.EqualError(t, err, "chaves não permitidas: teste")
	})

	// tentar trocar endereço sem passar todas chaves necessarias
	t.Run("when there's change address with all address fields", func(t *testing.T) {
		repository := mocks.NewDriverRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewUpdateDriverUseCase(
			&persistence.PostgresRepositories{
				DriverRepository: repository,
			},
			logger,
		)

		err := usecase.UpdateDriver("123", map[string]interface{}{
			"street": "123",
		})

		assert.EqualError(t, err, "os seguintes campos são obrigatórios: number, complement, zip")
	})

	// trocar estado quando há contrato disponível
	t.Run("when there's driver get error", func(t *testing.T) {
		driverRepository := mocks.NewDriverRepository(t)
		contractRepository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewUpdateDriverUseCase(
			&persistence.PostgresRepositories{
				DriverRepository:   driverRepository,
				ContractRepository: contractRepository,
			},
			logger,
		)

		driverRepository.On("Get", mock.Anything).Return(&entity.Driver{}, errors.New("get driver error"))

		err := usecase.UpdateDriver("123", map[string]interface{}{
			"state": "BA",
			"phone": "123",
		})

		assert.EqualError(t, err, "get driver error")
	})

	t.Run("when there's error to get driver enable contract", func(t *testing.T) {
		driverRepository := mocks.NewDriverRepository(t)
		contractRepository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewUpdateDriverUseCase(
			&persistence.PostgresRepositories{
				DriverRepository:   driverRepository,
				ContractRepository: contractRepository,
			},
			logger,
		)

		driverRepository.On("Get", mock.Anything).Return(&entity.Driver{}, nil)
		contractRepository.On("DriverHasEnableContract", mock.Anything).Return(false, errors.New("get driver enable contract error"))

		err := usecase.UpdateDriver("123", map[string]interface{}{
			"state": "BA",
			"phone": "123",
		})

		assert.EqualError(t, err, "get driver enable contract error")
	})

	// trocar estado quando há contrato disponível
	t.Run("when there's driver try change state with contract enable", func(t *testing.T) {
		driverRepository := mocks.NewDriverRepository(t)
		contractRepository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewUpdateDriverUseCase(
			&persistence.PostgresRepositories{
				DriverRepository:   driverRepository,
				ContractRepository: contractRepository,
			},
			logger,
		)

		driverRepository.On("Get", mock.Anything).Return(&entity.Driver{}, nil)
		contractRepository.On("DriverHasEnableContract", mock.Anything).Return(true, nil)

		err := usecase.UpdateDriver("123", map[string]interface{}{
			"state": "BA",
			"phone": "123",
		})

		assert.EqualError(t, err, "impossible change provincy when has enable contract")
	})

	t.Run("when there's driver update returns error", func(t *testing.T) {
		driverRepository := mocks.NewDriverRepository(t)
		contractRepository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewUpdateDriverUseCase(
			&persistence.PostgresRepositories{
				DriverRepository:   driverRepository,
				ContractRepository: contractRepository,
			},
			logger,
		)

		driverRepository.On("Get", mock.Anything).Return(&entity.Driver{}, nil)
		contractRepository.On("DriverHasEnableContract", mock.Anything).Return(false, nil)
		driverRepository.On("Update", mock.Anything, mock.Anything).Return(errors.New("update driver error"))

		err := usecase.UpdateDriver("123", map[string]interface{}{
			"state": "BA",
			"phone": "123",
		})

		assert.EqualError(t, err, "update driver error")
	})

	t.Run("when there's driver update returns error", func(t *testing.T) {
		driverRepository := mocks.NewDriverRepository(t)
		contractRepository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewUpdateDriverUseCase(
			&persistence.PostgresRepositories{
				DriverRepository:   driverRepository,
				ContractRepository: contractRepository,
			},
			logger,
		)

		driverRepository.On("Get", mock.Anything).Return(&entity.Driver{}, nil)
		contractRepository.On("DriverHasEnableContract", mock.Anything).Return(false, nil)
		driverRepository.On("Update", mock.Anything, mock.Anything).Return(nil)

		err := usecase.UpdateDriver("123", map[string]interface{}{
			"state": "BA",
			"phone": "123",
		})

		assert.NoError(t, err)
	})
}
