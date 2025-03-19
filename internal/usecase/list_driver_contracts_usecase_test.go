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
	t.Run("get contract repository return error", func(t *testing.T) {
		repository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)

		repository.On("GetByDriver", mock.Anything).Return([]entity.EnableContract{}, errors.New("database error"))

		uc := NewListDriverContractsUseCase(
			&persistence.PostgresRepositories{
				ContractRepository: repository,
			},
			logger,
		)

		_, err := uc.ListDriverContracts("123456789")

		assert.EqualError(t, err, "database error")
		assert.Error(t, err)
	})

	t.Run("list driver contracts return success", func(t *testing.T) {
		repository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)

		repository.On("GetByDriver", mock.Anything).Return([]entity.EnableContract{}, nil)

		uc := NewListDriverContractsUseCase(
			&persistence.PostgresRepositories{
				ContractRepository: repository,
			},
			logger,
		)

		_, err := uc.ListDriverContracts("123456789")

		assert.NoError(t, err)
		assert.Nil(t, err)
	})
}
