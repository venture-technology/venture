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

func TestGetTempContractsDriverUseCase_GetDriverTempContracts(t *testing.T) {
	cnh := "123"

	t.Run("if repository returns error", func(t *testing.T) {
		repository := mocks.NewTempContractRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewGetTempContractsDriverUseCase(
			&persistence.PostgresRepositories{
				TempContractRepository: repository,
			},
			logger,
		)

		repository.On("GetByDriver", mock.Anything).Return([]entity.TempContract{}, errors.New("database error"))

		_, err := usecase.GetDriverTempContracts(cnh)

		assert.EqualError(t, err, "database error")
	})

	t.Run("ir list returns success", func(t *testing.T) {
		repository := mocks.NewTempContractRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewGetTempContractsDriverUseCase(
			&persistence.PostgresRepositories{
				TempContractRepository: repository,
			},
			logger,
		)

		repository.On("GetByDriver", mock.Anything).Return([]entity.TempContract{}, nil)

		_, err := usecase.GetDriverTempContracts(cnh)
		assert.Nil(t, err)
	})
}
