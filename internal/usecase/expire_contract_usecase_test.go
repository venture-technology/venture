package usecase

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/mocks"
)

func TestExpireContractUsecase_ExpireContract(t *testing.T) {
	uuid := uuid.New()
	contract := entity.Contract{}

	t.Run("if contract not found", func(t *testing.T) {
		driverRepository := mocks.NewDriverRepository(t)
		contractRepository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewExpireContractUseCase(
			&persistence.PostgresRepositories{
				DriverRepository:   driverRepository,
				ContractRepository: contractRepository,
			},
			logger,
		)

		contractRepository.On("Get", mock.Anything).Return(&contract, errors.New("contract not found"))

		err := usecase.ExpireContract(uuid)

		assert.EqualError(t, err, "contract not found")
	})

	t.Run("if expired contract action return error", func(t *testing.T) {
		driverRepository := mocks.NewDriverRepository(t)
		contractRepository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewExpireContractUseCase(
			&persistence.PostgresRepositories{
				DriverRepository:   driverRepository,
				ContractRepository: contractRepository,
			},
			logger,
		)

		contractRepository.On("Get", mock.Anything).Return(&contract, nil)
		contractRepository.On("Expired", mock.Anything).Return(errors.New("expire error"))

		err := usecase.ExpireContract(uuid)

		assert.EqualError(t, err, "expire error")
	})

	t.Run("if seat up/down morning repository return error", func(t *testing.T) {
		driverRepository := mocks.NewDriverRepository(t)
		contractRepository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewExpireContractUseCase(
			&persistence.PostgresRepositories{
				DriverRepository:   driverRepository,
				ContractRepository: contractRepository,
			},
			logger,
		)

		contract := entity.Contract{
			Kid: entity.Kid{
				Shift: "morning",
			},
		}

		contractRepository.On("Get", mock.Anything).Return(&contract, nil)
		contractRepository.On("Expired", mock.Anything).Return(nil)
		driverRepository.On("Update", mock.Anything, mock.Anything).Return(errors.New("update driver error"))

		err := usecase.ExpireContract(uuid)

		assert.EqualError(t, err, "update driver error")
	})

	t.Run("if seat up/down afternoon repository return error", func(t *testing.T) {
		driverRepository := mocks.NewDriverRepository(t)
		contractRepository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewExpireContractUseCase(
			&persistence.PostgresRepositories{
				DriverRepository:   driverRepository,
				ContractRepository: contractRepository,
			},
			logger,
		)

		contract := entity.Contract{
			Kid: entity.Kid{
				Shift: "afternoon",
			},
		}

		contractRepository.On("Get", mock.Anything).Return(&contract, nil)
		contractRepository.On("Expired", mock.Anything).Return(nil)
		driverRepository.On("Update", mock.Anything, mock.Anything).Return(errors.New("update driver error"))

		err := usecase.ExpireContract(uuid)

		assert.EqualError(t, err, "update driver error")
	})

	t.Run("if seat up/down night repository return error", func(t *testing.T) {
		driverRepository := mocks.NewDriverRepository(t)
		contractRepository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewExpireContractUseCase(
			&persistence.PostgresRepositories{
				DriverRepository:   driverRepository,
				ContractRepository: contractRepository,
			},
			logger,
		)

		contract := entity.Contract{
			Kid: entity.Kid{
				Shift: "night",
			},
		}

		contractRepository.On("Get", mock.Anything).Return(&contract, nil)
		contractRepository.On("Expired", mock.Anything).Return(nil)
		driverRepository.On("Update", mock.Anything, mock.Anything).Return(errors.New("update driver error"))

		err := usecase.ExpireContract(uuid)

		assert.EqualError(t, err, "update driver error")
	})

	t.Run("expired contract morning with success", func(t *testing.T) {
		driverRepository := mocks.NewDriverRepository(t)
		contractRepository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewExpireContractUseCase(
			&persistence.PostgresRepositories{
				DriverRepository:   driverRepository,
				ContractRepository: contractRepository,
			},
			logger,
		)

		contract := entity.Contract{
			Kid: entity.Kid{
				Shift: "morning",
			},
		}

		contractRepository.On("Get", mock.Anything).Return(&contract, nil)
		contractRepository.On("Expired", mock.Anything).Return(nil)
		driverRepository.On("Update", mock.Anything, mock.Anything).Return(nil)

		err := usecase.ExpireContract(uuid)

		assert.Nil(t, err)
	})

	t.Run("expired contract afternoon with success", func(t *testing.T) {
		driverRepository := mocks.NewDriverRepository(t)
		contractRepository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewExpireContractUseCase(
			&persistence.PostgresRepositories{
				DriverRepository:   driverRepository,
				ContractRepository: contractRepository,
			},
			logger,
		)

		contract := entity.Contract{
			Kid: entity.Kid{
				Shift: "afternoon",
			},
		}

		contractRepository.On("Get", mock.Anything).Return(&contract, nil)
		contractRepository.On("Expired", mock.Anything).Return(nil)
		driverRepository.On("Update", mock.Anything, mock.Anything).Return(nil)

		err := usecase.ExpireContract(uuid)

		assert.Nil(t, err)
	})

	t.Run("expired contract night with success", func(t *testing.T) {
		driverRepository := mocks.NewDriverRepository(t)
		contractRepository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewExpireContractUseCase(
			&persistence.PostgresRepositories{
				DriverRepository:   driverRepository,
				ContractRepository: contractRepository,
			},
			logger,
		)

		contract := entity.Contract{
			Kid: entity.Kid{
				Shift: "night",
			},
		}

		contractRepository.On("Get", mock.Anything).Return(&contract, nil)
		contractRepository.On("Expired", mock.Anything).Return(nil)
		driverRepository.On("Update", mock.Anything, mock.Anything).Return(nil)

		err := usecase.ExpireContract(uuid)

		assert.Nil(t, err)
	})
}
