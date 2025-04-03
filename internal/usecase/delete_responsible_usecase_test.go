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

func TestDeleteResponsibleUsecase_DeleteResponsible(t *testing.T) {
	responsible := entity.Responsible{}

	t.Run("if contract repository returns error", func(t *testing.T) {
		repository := mocks.NewResponsibleRepository(t)
		payments := mocks.NewPaymentsService(t)
		logger := mocks.NewLogger(t)
		cr := mocks.NewContractRepository(t)

		usecase := NewDeleteResponsibleUseCase(
			&persistence.PostgresRepositories{
				ContractRepository:    cr,
				ResponsibleRepository: repository,
			},
			logger,
			adapters.Adapters{
				PaymentsService: payments,
			},
		)
		repository.On("Get", mock.Anything).Return(&responsible, nil)
		cr.On("ResponsibleHasEnableContract", mock.Anything).Return(true, nil)
		err := usecase.DeleteResponsible("123")

		assert.EqualError(t, err, "impossivel deletar responsavel possuindo contrato ativo")
	})

	t.Run("if receive error from responsible repository get", func(t *testing.T) {
		repository := mocks.NewResponsibleRepository(t)
		payments := mocks.NewPaymentsService(t)
		logger := mocks.NewLogger(t)
		cr := mocks.NewContractRepository(t)

		usecase := NewDeleteResponsibleUseCase(
			&persistence.PostgresRepositories{
				ResponsibleRepository: repository,
				ContractRepository:    cr,
			},
			logger,
			adapters.Adapters{
				PaymentsService: payments,
			},
		)
		repository.On("Get", mock.Anything).Return(&responsible, errors.New("responsible repository get error"))

		err := usecase.DeleteResponsible("123")

		assert.EqualError(t, err, "responsible repository get error")
	})
	t.Run("if receive error from stripe delete user", func(t *testing.T) {
		repository := mocks.NewResponsibleRepository(t)
		payments := mocks.NewPaymentsService(t)
		logger := mocks.NewLogger(t)
		cr := mocks.NewContractRepository(t)

		usecase := NewDeleteResponsibleUseCase(
			&persistence.PostgresRepositories{
				ResponsibleRepository: repository,
				ContractRepository:    cr,
			},
			logger,
			adapters.Adapters{
				PaymentsService: payments,
			},
		)
		cr.On("ResponsibleHasEnableContract", mock.Anything).Return(false, nil)
		repository.On("Get", mock.Anything).Return(&responsible, nil)
		payments.On("DeleteStripeUser", mock.Anything).Return(&stripe.Customer{}, errors.New("delete stripe error from stripe"))

		err := usecase.DeleteResponsible("123")

		assert.EqualError(t, err, "delete stripe error from stripe")
	})

	t.Run("if receive error from responsible repository delete", func(t *testing.T) {
		repository := mocks.NewResponsibleRepository(t)
		payments := mocks.NewPaymentsService(t)
		logger := mocks.NewLogger(t)
		cr := mocks.NewContractRepository(t)

		usecase := NewDeleteResponsibleUseCase(
			&persistence.PostgresRepositories{
				ResponsibleRepository: repository,
				ContractRepository:    cr,
			},
			logger,
			adapters.Adapters{
				PaymentsService: payments,
			},
		)
		cr.On("ResponsibleHasEnableContract", mock.Anything).Return(false, nil)
		repository.On("Get", mock.Anything).Return(&responsible, nil)
		payments.On("DeleteStripeUser", mock.Anything).Return(&stripe.Customer{}, nil)
		repository.On("Delete", mock.Anything).Return(errors.New("responsible repository delete error"))

		err := usecase.DeleteResponsible("123")

		assert.EqualError(t, err, "responsible repository delete error")
	})

	t.Run("if delete responsible returns success", func(t *testing.T) {
		repository := mocks.NewResponsibleRepository(t)
		payments := mocks.NewPaymentsService(t)
		logger := mocks.NewLogger(t)
		cr := mocks.NewContractRepository(t)

		usecase := NewDeleteResponsibleUseCase(
			&persistence.PostgresRepositories{
				ResponsibleRepository: repository,
				ContractRepository:    cr,
			},
			logger,
			adapters.Adapters{
				PaymentsService: payments,
			},
		)
		cr.On("ResponsibleHasEnableContract", mock.Anything).Return(false, nil)
		repository.On("Get", mock.Anything).Return(&responsible, nil)
		payments.On("DeleteStripeUser", mock.Anything).Return(&stripe.Customer{}, nil)
		repository.On("Delete", mock.Anything).Return(nil)

		err := usecase.DeleteResponsible("123")

		assert.Nil(t, err)
	})
}
