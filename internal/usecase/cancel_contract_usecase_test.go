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

func TestCancelContractUsecase_CancelContract(t *testing.T) {
	var uuid uuid.UUID

	t.Run("get contract repository return error", func(t *testing.T) {
		repository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)
		ps := mocks.NewPaymentsService(t)

		repository.On("GetByUUID", mock.Anything).Return(nil, errors.New("database error"))

		uc := NewCancelContractUseCase(
			&persistence.PostgresRepositories{
				ContractRepository: repository,
			},
			logger,
			adapters.Adapters{
				PaymentsService: ps,
			},
		)

		err := uc.CancelContract(uuid)

		assert.EqualError(t, err, "database error")
		assert.Error(t, err)
	})

	t.Run("list invoices return error", func(t *testing.T) {
		repository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)
		ps := mocks.NewPaymentsService(t)

		repository.On("GetByUUID", mock.Anything).Return(&entity.EnableContract{}, nil)
		ps.On("ListInvoices", mock.Anything).Return(nil, errors.New("list invoices error"))

		uc := NewCancelContractUseCase(
			&persistence.PostgresRepositories{
				ContractRepository: repository,
			},
			logger,
			adapters.Adapters{
				PaymentsService: ps,
			},
		)

		err := uc.CancelContract(uuid)

		assert.EqualError(t, err, "list invoices error")
		assert.Error(t, err)
	})

	t.Run("fine responsible return error", func(t *testing.T) {
		repository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)
		ps := mocks.NewPaymentsService(t)

		repository.On("GetByUUID", mock.Anything).Return(&entity.EnableContract{}, nil)
		ps.On("ListInvoices", mock.Anything).Return(map[string]entity.InvoiceInfo{}, nil)
		ps.On("CalculateRemainingValueSubscription", mock.Anything, mock.Anything).Return(float64(0))
		ps.On("FineResponsible", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("fine responsible error"))

		uc := NewCancelContractUseCase(
			&persistence.PostgresRepositories{
				ContractRepository: repository,
			},
			logger,
			adapters.Adapters{
				PaymentsService: ps,
			},
		)

		err := uc.CancelContract(uuid)

		assert.EqualError(t, err, "fine responsible error")
		assert.Error(t, err)
	})

	t.Run("cancel contract repository return error", func(t *testing.T) {
		repository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)
		ps := mocks.NewPaymentsService(t)

		repository.On("GetByUUID", mock.Anything).Return(&entity.EnableContract{}, nil)
		ps.On("ListInvoices", mock.Anything).Return(map[string]entity.InvoiceInfo{}, nil)
		ps.On("CalculateRemainingValueSubscription", mock.Anything, mock.Anything).Return(float64(0))
		ps.On("FineResponsible", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
		repository.On("Cancel", mock.Anything).Return(errors.New("cancel contract error"))

		uc := NewCancelContractUseCase(
			&persistence.PostgresRepositories{
				ContractRepository: repository,
			},
			logger,
			adapters.Adapters{
				PaymentsService: ps,
			},
		)

		err := uc.CancelContract(uuid)

		assert.EqualError(t, err, "cancel contract error")
		assert.Error(t, err)
	})

	t.Run("cancel contract usecase return success", func(t *testing.T) {
		repository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)
		ps := mocks.NewPaymentsService(t)

		repository.On("GetByUUID", mock.Anything).Return(&entity.EnableContract{}, nil)
		ps.On("ListInvoices", mock.Anything).Return(map[string]entity.InvoiceInfo{}, nil)
		ps.On("CalculateRemainingValueSubscription", mock.Anything, mock.Anything).Return(float64(0))
		ps.On("FineResponsible", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
		repository.On("Cancel", mock.Anything).Return(nil)

		uc := NewCancelContractUseCase(
			&persistence.PostgresRepositories{
				ContractRepository: repository,
			},
			logger,
			adapters.Adapters{
				PaymentsService: ps,
			},
		)

		err := uc.CancelContract(uuid)

		assert.NoError(t, err)
		assert.Nil(t, err)
	})
}
