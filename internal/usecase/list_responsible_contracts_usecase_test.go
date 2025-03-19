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

func TestListResponsibleContractsUsecase_ListResponsibleContracts(t *testing.T) {
	t.Run("get contract repository return error", func(t *testing.T) {
		repository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)

		repository.On("GetByResponsible", mock.Anything).Return([]entity.EnableContract{}, errors.New("database error"))

		uc := NewListResponsibleContractsUseCase(
			&persistence.PostgresRepositories{
				ContractRepository: repository,
			},
			logger,
		)

		_, err := uc.ListResponsibleContracts("123456789")

		assert.EqualError(t, err, "database error")
		assert.Error(t, err)
	})

	t.Run("list responsible contracts return success", func(t *testing.T) {
		repository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)

		repository.On("GetByResponsible", mock.Anything).Return([]entity.EnableContract{}, nil)

		uc := NewListResponsibleContractsUseCase(
			&persistence.PostgresRepositories{
				ContractRepository: repository,
			},
			logger,
		)

		_, err := uc.ListResponsibleContracts("123456789")

		assert.NoError(t, err)
		assert.Nil(t, err)
	})
}
