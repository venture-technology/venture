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

func TestListSchoolContractUsecase_ListSchoolContract(t *testing.T) {
	cnpj := "123"

	t.Run("if repository returns error", func(t *testing.T) {
		repository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewListSchoolContractUseCase(
			&persistence.PostgresRepositories{
				ContractRepository: repository,
			},
			logger,
		)

		repository.On("FindAllByCnpj", mock.Anything).Return([]entity.Contract{}, errors.New("database error"))

		_, err := usecase.ListSchoolContract(&cnpj)

		assert.EqualError(t, err, "database error")
	})

	t.Run("if list returns sucess", func(t *testing.T) {
		repository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewListSchoolContractUseCase(
			&persistence.PostgresRepositories{
				ContractRepository: repository,
			},
			logger,
		)

		repository.On("FindAllByCnpj", mock.Anything).Return([]entity.Contract{}, nil)

		_, err := usecase.ListSchoolContract(&cnpj)

		assert.NoError(t, err)
	})
}
