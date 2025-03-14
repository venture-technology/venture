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

func TestListDriverContractsUsecase_ListDriverContracts(t *testing.T) {
	cnh := "123"

	t.Run("if repository returns error", func(t *testing.T) {
		repository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewListDriverContractsUseCase(
			&persistence.PostgresRepositories{
				ContractRepository: repository,
			},
			logger,
		)

		repository.On("FindAllByCnh", mock.Anything).Return([]entity.Contract{}, errors.New("database error"))

		_, err := usecase.ListDriverContracts(&cnh)

		assert.EqualError(t, err, "database error")
	})

	t.Run("if list returns sucess", func(t *testing.T) {
		repository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewListDriverContractsUseCase(
			&persistence.PostgresRepositories{
				ContractRepository: repository,
			},
			logger,
		)

		repository.On("FindAllByCnh", mock.Anything).Return([]entity.Contract{}, nil)

		_, err := usecase.ListDriverContracts(&cnh)

		assert.NoError(t, err)
	})
}
