package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/venture-technology/venture/internal/domain/service/adapters"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/mocks"
)

func TestListResponsibleContractsUsecase_ListResponsibleContracts(t *testing.T) {
	cpf := "123"

	t.Run("if repository returns error", func(t *testing.T) {
		repository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewListResponsibleContractsUseCase(
			&persistence.PostgresRepositories{
				ContractRepository: repository,
			},
			logger,
			adapters.Adapters{},
		)

		repository.On("FindAllByCpf", mock.Anything).Return([]entity.Contract{}, errors.New("database error"))

		_, err := usecase.ListResponsibleContracts(&cpf)

		assert.EqualError(t, err, "database error")
	})

	t.Run("if list returns sucess", func(t *testing.T) {
		repository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewListResponsibleContractsUseCase(
			&persistence.PostgresRepositories{
				ContractRepository: repository,
			},
			logger,
			adapters.Adapters{},
		)

		repository.On("FindAllByCpf", mock.Anything).Return([]entity.Contract{}, nil)

		_, err := usecase.ListResponsibleContracts(&cpf)

		assert.NoError(t, err)
	})
}
