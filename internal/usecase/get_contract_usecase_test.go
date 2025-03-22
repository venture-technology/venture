package usecase

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/venture-technology/venture/internal/domain/service/adapters"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/mocks"
)

func TestGetContracatUsecase_GetContract(t *testing.T) {
	var uuid uuid.UUID

	t.Run("get contract repository return error", func(t *testing.T) {
		repository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)
		ps := mocks.NewPaymentsService(t)

		repository.On("GetByUUID", mock.Anything).Return(nil, errors.New("database error"))

		uc := NewGetContractUseCase(
			&persistence.PostgresRepositories{
				ContractRepository: repository,
			},
			logger,
			adapters.Adapters{
				PaymentsService: ps,
			},
		)

		_, err := uc.GetContract(uuid)

		assert.EqualError(t, err, "database error")
		assert.Error(t, err)
	})

	t.Run("list invoices return error", func(t *testing.T) {
		repository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)
		ps := mocks.NewPaymentsService(t)

		repository.On("GetByUUID", mock.Anything).Return(&entity.Contract{}, nil)
		ps.On("ListInvoices", mock.Anything).Return(nil, errors.New("payment service error"))

		uc := NewGetContractUseCase(
			&persistence.PostgresRepositories{
				ContractRepository: repository,
			},
			logger,
			adapters.Adapters{
				PaymentsService: ps,
			},
		)

		_, err := uc.GetContract(uuid)

		assert.EqualError(t, err, "payment service error")
		assert.Error(t, err)
	})

	t.Run("get contract usecase return success", func(t *testing.T) {
		repository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)
		ps := mocks.NewPaymentsService(t)

		repository.On("GetByUUID", mock.Anything).Return(&entity.Contract{}, nil)
		ps.On("ListInvoices", mock.Anything).Return(nil, nil)

		uc := NewGetContractUseCase(
			&persistence.PostgresRepositories{
				ContractRepository: repository,
			},
			logger,
			adapters.Adapters{
				PaymentsService: ps,
			},
		)

		_, err := uc.GetContract(uuid)

		assert.NoError(t, err)
		assert.Nil(t, err)
	})
}
