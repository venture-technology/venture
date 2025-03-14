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

func TestGetContractUsecase_GetContract(t *testing.T) {
	uuid := uuid.New()
	contract := entity.Contract{}

	t.Run("if repository returns error", func(t *testing.T) {
		repository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)
		payments := mocks.NewPaymentsService(t)

		usecase := NewGetContractUseCase(
			&persistence.PostgresRepositories{
				ContractRepository: repository,
			},
			logger,
			adapters.Adapters{
				PaymentsService: payments,
			},
		)

		repository.On("Get", mock.Anything).Return(&contract, errors.New("database error"))

		_, err := usecase.GetContract(uuid)

		assert.EqualError(t, err, "database error")
	})

	t.Run("when list invoices returns error", func(t *testing.T) {
		repository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)
		payments := mocks.NewPaymentsService(t)

		usecase := NewGetContractUseCase(
			&persistence.PostgresRepositories{
				ContractRepository: repository,
			},
			logger,
			adapters.Adapters{
				PaymentsService: payments,
			},
		)

		repository.On("Get", mock.Anything).Return(&contract, nil)
		payments.On("ListInvoices", mock.Anything).Return(map[string]entity.InvoiceInfo{}, errors.New("invoice list error"))

		_, err := usecase.GetContract(uuid)

		assert.EqualError(t, err, "invoice list error")
	})

	t.Run("if list return success", func(t *testing.T) {
		repository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)
		payments := mocks.NewPaymentsService(t)

		usecase := NewGetContractUseCase(
			&persistence.PostgresRepositories{
				ContractRepository: repository,
			},
			logger,
			adapters.Adapters{
				PaymentsService: payments,
			},
		)

		repository.On("Get", mock.Anything).Return(&contract, nil)
		payments.On("ListInvoices", mock.Anything).Return(map[string]entity.InvoiceInfo{}, nil)

		_, err := usecase.GetContract(uuid)

		assert.Nil(t, err)
	})
}
