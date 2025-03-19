package usecase

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stripe/stripe-go/v79"
	"github.com/venture-technology/venture/internal/domain/service/adapters"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/mocks"
)

func TestCancelContractUsecase_CancelContract(t *testing.T) {
	var uuid uuid.UUID

	t.Run("if get contract return error", func(t *testing.T) {
		cr := mocks.NewContractRepository(t)
		rr := mocks.NewResponsibleRepository(t)
		ps := mocks.NewPaymentsService(t)
		logger := mocks.NewLogger(t)

		uc := NewCancelContractUseCase(
			&persistence.PostgresRepositories{
				ContractRepository:    cr,
				ResponsibleRepository: rr,
			},
			logger,
			adapters.Adapters{
				PaymentsService: ps,
			},
		)

		cr.On("Get", mock.Anything).Return(nil, errors.New("contract get repo error"))

		err := uc.CancelContract(uuid)

		assert.EqualError(t, err, "contract get repo error")
		assert.Error(t, err)
	})

	t.Run("if get responsible return error", func(t *testing.T) {
		cr := mocks.NewContractRepository(t)
		rr := mocks.NewResponsibleRepository(t)
		ps := mocks.NewPaymentsService(t)
		logger := mocks.NewLogger(t)

		uc := NewCancelContractUseCase(
			&persistence.PostgresRepositories{
				ContractRepository:    cr,
				ResponsibleRepository: rr,
			},
			logger,
			adapters.Adapters{
				PaymentsService: ps,
			},
		)

		cr.On("Get", mock.Anything).Return(&entity.Contract{}, nil)
		rr.On("Get", mock.Anything).Return(nil, errors.New("responsible get repo error"))

		err := uc.CancelContract(uuid)

		assert.EqualError(t, err, "responsible get repo error")
		assert.Error(t, err)
	})

	t.Run("if stripe list invoice return error", func(t *testing.T) {
		cr := mocks.NewContractRepository(t)
		rr := mocks.NewResponsibleRepository(t)
		ps := mocks.NewPaymentsService(t)
		logger := mocks.NewLogger(t)

		uc := NewCancelContractUseCase(
			&persistence.PostgresRepositories{
				ContractRepository:    cr,
				ResponsibleRepository: rr,
			},
			logger,
			adapters.Adapters{
				PaymentsService: ps,
			},
		)

		cr.On("Get", mock.Anything).Return(&entity.Contract{}, nil)
		rr.On("Get", mock.Anything).Return(&entity.Responsible{}, nil)
		ps.On("ListInvoices", mock.Anything).Return(map[string]entity.InvoiceInfo{}, errors.New("list invoices error"))

		err := uc.CancelContract(uuid)

		assert.EqualError(t, err, "list invoices error")
		assert.Error(t, err)
	})

	t.Run("if fine responsible return error", func(t *testing.T) {
		cr := mocks.NewContractRepository(t)
		rr := mocks.NewResponsibleRepository(t)
		ps := mocks.NewPaymentsService(t)
		logger := mocks.NewLogger(t)

		uc := NewCancelContractUseCase(
			&persistence.PostgresRepositories{
				ContractRepository:    cr,
				ResponsibleRepository: rr,
			},
			logger,
			adapters.Adapters{
				PaymentsService: ps,
			},
		)

		cr.On("Get", mock.Anything).Return(&entity.Contract{}, nil)
		rr.On("Get", mock.Anything).Return(&entity.Responsible{}, nil)
		ps.On("ListInvoices", mock.Anything).Return(map[string]entity.InvoiceInfo{}, nil)
		ps.On("CalculateRemainingValueSubscription", mock.Anything, mock.Anything).Return(float64(0))
		ps.On("FineResponsible", mock.Anything, mock.Anything, mock.Anything).Return(&stripe.PaymentIntent{}, errors.New("fine responsible error"))

		err := uc.CancelContract(uuid)

		assert.EqualError(t, err, "fine responsible error")
		assert.Error(t, err)
	})

	t.Run("if cancel contract return error", func(t *testing.T) {
		cr := mocks.NewContractRepository(t)
		rr := mocks.NewResponsibleRepository(t)
		ps := mocks.NewPaymentsService(t)
		logger := mocks.NewLogger(t)

		uc := NewCancelContractUseCase(
			&persistence.PostgresRepositories{
				ContractRepository:    cr,
				ResponsibleRepository: rr,
			},
			logger,
			adapters.Adapters{
				PaymentsService: ps,
			},
		)

		cr.On("Get", mock.Anything).Return(&entity.Contract{}, nil)
		rr.On("Get", mock.Anything).Return(&entity.Responsible{}, nil)
		ps.On("ListInvoices", mock.Anything).Return(map[string]entity.InvoiceInfo{}, nil)
		ps.On("CalculateRemainingValueSubscription", mock.Anything, mock.Anything).Return(float64(0))
		ps.On("FineResponsible", mock.Anything, mock.Anything, mock.Anything).Return(&stripe.PaymentIntent{}, nil)
		cr.On("Cancel", mock.Anything).Return(errors.New("cancel contract repo error"))

		err := uc.CancelContract(uuid)

		assert.EqualError(t, err, "cancel contract repo error")
		assert.Error(t, err)
	})

	t.Run("when usecase return success", func(t *testing.T) {
		cr := mocks.NewContractRepository(t)
		rr := mocks.NewResponsibleRepository(t)
		ps := mocks.NewPaymentsService(t)
		logger := mocks.NewLogger(t)

		uc := NewCancelContractUseCase(
			&persistence.PostgresRepositories{
				ContractRepository:    cr,
				ResponsibleRepository: rr,
			},
			logger,
			adapters.Adapters{
				PaymentsService: ps,
			},
		)

		cr.On("Get", mock.Anything).Return(&entity.Contract{}, nil)
		rr.On("Get", mock.Anything).Return(&entity.Responsible{}, nil)
		ps.On("ListInvoices", mock.Anything).Return(map[string]entity.InvoiceInfo{}, nil)
		ps.On("CalculateRemainingValueSubscription", mock.Anything, mock.Anything).Return(float64(0))
		ps.On("FineResponsible", mock.Anything, mock.Anything, mock.Anything).Return(&stripe.PaymentIntent{}, nil)
		cr.On("Cancel", mock.Anything).Return(nil)

		err := uc.CancelContract(uuid)

		assert.NoError(t, err)
		assert.Nil(t, err)
	})
}
