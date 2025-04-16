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

func TestExpireTemporaryContractUseCase_ExpireTemporaryContracts(t *testing.T) {
	t.Run("if TempContractRepository.GetExpiredContracts returns error", func(t *testing.T) {
		repository := mocks.NewTempContractRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewExpireTemporaryContractUseCase(
			&persistence.PostgresRepositories{
				TempContractRepository: repository,
			},
			logger,
		)

		repository.On("GetExpiredContracts").Return(nil, errors.New("failed to get expired contracts"))
		logger.On("Errorf", mock.Anything, mock.Anything).Return()

		err := usecase.ExpireTemporaryContracts()

		assert.EqualError(t, err, "failed to get expired contracts")
	})

	t.Run("if UUID parsing fails", func(t *testing.T) {
		repository := mocks.NewTempContractRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewExpireTemporaryContractUseCase(
			&persistence.PostgresRepositories{
				TempContractRepository: repository,
			},
			logger,
		)

		expiredContracts := []entity.TempContract{
			{UUID: "invalid-uuid"},
		}

		repository.On("GetExpiredContracts").Return(expiredContracts, nil)
		logger.On("Errorf", mock.Anything, mock.Anything).Return()

		err := usecase.ExpireTemporaryContracts()

		assert.Nil(t, err)
	})

	t.Run("if TempContractRepository.Expire fails", func(t *testing.T) {
		repository := mocks.NewTempContractRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewExpireTemporaryContractUseCase(
			&persistence.PostgresRepositories{
				TempContractRepository: repository,
			},
			logger,
		)

		expiredContracts := []entity.TempContract{
			{UUID: "f47ac10b-58cc-4372-a567-0e02b2c3d479"},
		}

		repository.On("GetExpiredContracts").Return(expiredContracts, nil)
		repository.On("Expire", mock.Anything).Return(errors.New("repository error"))
		logger.On("Errorf", mock.Anything, mock.Anything).Return()

		err := usecase.ExpireTemporaryContracts()

		assert.EqualError(t, err, "repository error")
	})

	t.Run("if everything succeeds", func(t *testing.T) {
		repository := mocks.NewTempContractRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewExpireTemporaryContractUseCase(
			&persistence.PostgresRepositories{
				TempContractRepository: repository,
			},
			logger,
		)

		expiredContracts := []entity.TempContract{
			{UUID: "f47ac10b-58cc-4372-a567-0e02b2c3d479"},
		}

		repository.On("GetExpiredContracts").Return(expiredContracts, nil)
		repository.On("Expire", mock.Anything).Return(nil)

		err := usecase.ExpireTemporaryContracts()

		assert.Nil(t, err)
	})
}
