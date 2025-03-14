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

func TestGetTempContractsResponsibleUseCase_GetResponsibleTempContracts(t *testing.T) {
	cpf := "123"

	t.Run("if repository returns error", func(t *testing.T) {
		repository := mocks.NewTempContractRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewGetTempContractsResponsibleUseCase(
			&persistence.PostgresRepositories{
				TempContractRepository: repository,
			},
			logger,
		)

		repository.On("FindAllByResponsible", mock.Anything).Return([]entity.TempContract{}, errors.New("database error"))

		_, err := usecase.GetResponsibleTempContracts(cpf)

		assert.EqualError(t, err, "database error")
	})

	t.Run("ir list returns success", func(t *testing.T) {
		repository := mocks.NewTempContractRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewGetTempContractsResponsibleUseCase(
			&persistence.PostgresRepositories{
				TempContractRepository: repository,
			},
			logger,
		)

		repository.On("FindAllByResponsible", mock.Anything).Return([]entity.TempContract{}, nil)

		_, err := usecase.GetResponsibleTempContracts(cpf)

		assert.Nil(t, err)
	})
}
