package usecase

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/mocks"
)

func TestCancelTempContractUsecase_CancelContract(t *testing.T) {
	var uuid uuid.UUID

	t.Run("when repository return error", func(t *testing.T) {
		tempContractRepository := mocks.NewTempContractRepository(t)
		logger := mocks.NewLogger(t)

		uc := NewCancelTempContractUsecase(
			&persistence.PostgresRepositories{
				TempContractRepository: tempContractRepository,
			},
			logger,
		)

		tempContractRepository.On("Cancel", mock.Anything).Return(errors.New("database error"))

		err := uc.CancelTempContract(uuid)

		assert.EqualError(t, err, "database error")
		assert.Error(t, err)
	})

	t.Run("when usecase return success", func(t *testing.T) {
		tempContractRepository := mocks.NewTempContractRepository(t)
		logger := mocks.NewLogger(t)

		uc := NewCancelTempContractUsecase(
			&persistence.PostgresRepositories{
				TempContractRepository: tempContractRepository,
			},
			logger,
		)

		tempContractRepository.On("Cancel", mock.Anything).Return(nil)

		err := uc.CancelTempContract(uuid)

		assert.Nil(t, err)
		assert.NoError(t, err)
	})
}
