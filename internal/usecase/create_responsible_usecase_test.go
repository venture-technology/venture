package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stripe/stripe-go/v79"
	"github.com/venture-technology/venture/internal/domain/service/adapters"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/mocks"
)

func TestCreateResponsibleUsecase_CreateResponsible(t *testing.T) {
	responsible := entity.Responsible{}
	t.Run("if create customer fails", func(t *testing.T) {
		repository := mocks.NewResponsibleRepository(t)
		payments := mocks.NewPaymentsService(t)
		logger := mocks.NewLogger(t)

		usecase := NewCreateResponsibleUseCase(
			&persistence.PostgresRepositories{
				ResponsibleRepository: repository,
			},
			logger,
			adapters.Adapters{
				PaymentsService: payments,
			},
		)

		payments.On("CreateCustomer", mock.Anything).Return("", errors.New("create customer fails"))

		err := usecase.CreateResponsible(&responsible)

		assert.EqualError(t, err, "create customer fails")
	})

	t.Run("if create payment method fails", func(t *testing.T) {
		repository := mocks.NewResponsibleRepository(t)
		payments := mocks.NewPaymentsService(t)
		logger := mocks.NewLogger(t)

		usecase := NewCreateResponsibleUseCase(
			&persistence.PostgresRepositories{
				ResponsibleRepository: repository,
			},
			logger,
			adapters.Adapters{
				PaymentsService: payments,
			},
		)

		payments.On("CreateCustomer", mock.Anything).Return("123", nil)
		payments.On("CreatePaymentMethod", mock.Anything).Return(&stripe.PaymentMethod{}, errors.New("payment method creation error"))

		err := usecase.CreateResponsible(&responsible)

		assert.EqualError(t, err, "payment method creation error")
	})

	t.Run("if create responsible on repository returns error", func(t *testing.T) {
		repository := mocks.NewResponsibleRepository(t)
		payments := mocks.NewPaymentsService(t)
		logger := mocks.NewLogger(t)

		usecase := NewCreateResponsibleUseCase(
			&persistence.PostgresRepositories{
				ResponsibleRepository: repository,
			},
			logger,
			adapters.Adapters{
				PaymentsService: payments,
			},
		)

		payments.On("CreateCustomer", mock.Anything).Return("123", nil)
		payments.On("CreatePaymentMethod", mock.Anything).Return(&stripe.PaymentMethod{}, nil)
		payments.On("AttachCardToResponsible", mock.Anything, mock.Anything).Return(&stripe.PaymentMethod{}, nil)
		repository.On("Create", mock.Anything).Return(errors.New("repository error"))

		err := usecase.CreateResponsible(&responsible)

		assert.EqualError(t, err, "repository error")
	})

	t.Run("when creation of responsible return success", func(t *testing.T) {
		repository := mocks.NewResponsibleRepository(t)
		payments := mocks.NewPaymentsService(t)
		logger := mocks.NewLogger(t)

		usecase := NewCreateResponsibleUseCase(
			&persistence.PostgresRepositories{
				ResponsibleRepository: repository,
			},
			logger,
			adapters.Adapters{
				PaymentsService: payments,
			},
		)

		payments.On("CreateCustomer", mock.Anything).Return("123", nil)
		payments.On("CreatePaymentMethod", mock.Anything).Return(&stripe.PaymentMethod{}, nil)
		payments.On("AttachCardToResponsible", mock.Anything, mock.Anything).Return(&stripe.PaymentMethod{}, nil)
		repository.On("Create", mock.Anything).Return(nil)

		err := usecase.CreateResponsible(&responsible)

		assert.Nil(t, err)
	})
}
