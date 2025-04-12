package usecase

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stripe/stripe-go/v79"
	"github.com/venture-technology/venture/internal/domain/service/adapters"
	"github.com/venture-technology/venture/internal/domain/service/agreements"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/mocks"
)

func TestAcceptContractUsecase_AcceptContract(t *testing.T) {
	makeASRAS := func(uuidStr string) agreements.ASRASOutput {
		return agreements.ASRASOutput{
			Contract: entity.Contract{
				UUID: uuidStr,
			},
			Signatures: []agreements.Signature{
				{SignedAt: 1},
				{SignedAt: 2},
			},
		}
	}

	t.Run("get contract already exists repository return error", func(t *testing.T) {
		rr := mocks.NewResponsibleRepository(t)
		cr := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)
		ps := mocks.NewPaymentsService(t)

		cr.On("ContractAlreadyExist", mock.Anything).Return(false, errors.New("contract already exists error"))

		uc := NewAcceptContractUseCase(
			&persistence.PostgresRepositories{
				ResponsibleRepository: rr,
				ContractRepository:    cr,
			},
			logger,
			adapters.Adapters{
				PaymentsService: ps,
			},
		)

		err := uc.AcceptContract(agreements.ASRASOutput{})

		assert.EqualError(t, err, "contract already exists error")
		assert.Error(t, err)
	})

	t.Run("if contract already exists", func(t *testing.T) {
		rr := mocks.NewResponsibleRepository(t)
		cr := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)
		ps := mocks.NewPaymentsService(t)

		cr.On("ContractAlreadyExist", mock.Anything).Return(true, nil)

		uc := NewAcceptContractUseCase(
			&persistence.PostgresRepositories{
				ResponsibleRepository: rr,
				ContractRepository:    cr,
			},
			logger,
			adapters.Adapters{
				PaymentsService: ps,
			},
		)

		err := uc.AcceptContract(agreements.ASRASOutput{})

		assert.EqualError(t, err, "contract already exists")
		assert.Error(t, err)
	})

	t.Run("get responsible repository return error", func(t *testing.T) {
		rr := mocks.NewResponsibleRepository(t)
		cr := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)
		ps := mocks.NewPaymentsService(t)

		cr.On("ContractAlreadyExist", mock.Anything).Return(false, nil)
		rr.On("Get", mock.Anything).Return(nil, errors.New("responsible get error"))

		uc := NewAcceptContractUseCase(
			&persistence.PostgresRepositories{
				ResponsibleRepository: rr,
				ContractRepository:    cr,
			},
			logger,
			adapters.Adapters{
				PaymentsService: ps,
			},
		)

		err := uc.AcceptContract(agreements.ASRASOutput{})

		assert.EqualError(t, err, "responsible get error")
		assert.Error(t, err)
	})

	t.Run("stripe create product return error", func(t *testing.T) {
		rr := mocks.NewResponsibleRepository(t)
		cr := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)
		ps := mocks.NewPaymentsService(t)

		cr.On("ContractAlreadyExist", mock.Anything).Return(false, nil)
		rr.On("Get", mock.Anything).Return(&entity.Responsible{}, nil)
		ps.On("CreateProduct", mock.Anything).Return(&stripe.Product{}, errors.New("stripe create product error"))

		uc := NewAcceptContractUseCase(
			&persistence.PostgresRepositories{
				ResponsibleRepository: rr,
				ContractRepository:    cr,
			},
			logger,
			adapters.Adapters{
				PaymentsService: ps,
			},
		)

		err := uc.AcceptContract(agreements.ASRASOutput{})

		assert.EqualError(t, err, "stripe create product error")
		assert.Error(t, err)
	})

	t.Run("stripe create price return error", func(t *testing.T) {
		rr := mocks.NewResponsibleRepository(t)
		cr := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)
		ps := mocks.NewPaymentsService(t)

		cr.On("ContractAlreadyExist", mock.Anything).Return(false, nil)
		rr.On("Get", mock.Anything).Return(&entity.Responsible{}, nil)
		ps.On("CreateProduct", mock.Anything).Return(&stripe.Product{}, nil)
		ps.On("CreatePrice", mock.Anything, mock.Anything).Return(&stripe.Price{}, errors.New("stripe create price error"))

		uc := NewAcceptContractUseCase(
			&persistence.PostgresRepositories{
				ResponsibleRepository: rr,
				ContractRepository:    cr,
			},
			logger,
			adapters.Adapters{
				PaymentsService: ps,
			},
		)

		err := uc.AcceptContract(agreements.ASRASOutput{})

		assert.EqualError(t, err, "stripe create price error")
		assert.Error(t, err)
	})

	t.Run("stripe create subscription return error", func(t *testing.T) {
		rr := mocks.NewResponsibleRepository(t)
		cr := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)
		ps := mocks.NewPaymentsService(t)

		cr.On("ContractAlreadyExist", mock.Anything).Return(false, nil)
		rr.On("Get", mock.Anything).Return(&entity.Responsible{}, nil)
		ps.On("CreateProduct", mock.Anything).Return(&stripe.Product{}, nil)
		ps.On("CreatePrice", mock.Anything, mock.Anything).Return(&stripe.Price{}, nil)
		ps.On("CreateSubscription", mock.Anything, mock.Anything).Return(&stripe.Subscription{}, errors.New("stripe create subscription error"))

		uc := NewAcceptContractUseCase(
			&persistence.PostgresRepositories{
				ResponsibleRepository: rr,
				ContractRepository:    cr,
			},
			logger,
			adapters.Adapters{
				PaymentsService: ps,
			},
		)

		err := uc.AcceptContract(agreements.ASRASOutput{})

		assert.EqualError(t, err, "stripe create subscription error")
		assert.Error(t, err)
	})

	t.Run("update temp contract repository return error", func(t *testing.T) {
		rr := mocks.NewResponsibleRepository(t)
		cr := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)
		ps := mocks.NewPaymentsService(t)
		tcr := mocks.NewTempContractRepository(t)

		uuidStr := uuid.New().String()
		parsedUUID, _ := uuid.Parse(uuidStr)

		cr.On("ContractAlreadyExist", mock.Anything).Return(false, nil)
		rr.On("Get", mock.Anything).Return(&entity.Responsible{}, nil)
		ps.On("CreateProduct", mock.Anything).Return(&stripe.Product{}, nil)
		ps.On("CreatePrice", mock.Anything, mock.Anything).Return(&stripe.Price{}, nil)
		ps.On("CreateSubscription", mock.Anything, mock.Anything).Return(&stripe.Subscription{}, nil)
		tcr.On("Update", parsedUUID, map[string]interface{}{
			"responsible_signed_at": int64(1),
			"driver_signed_at":      int64(2),
		}).Return(errors.New("update temp contract error"))

		uc := NewAcceptContractUseCase(
			&persistence.PostgresRepositories{
				ResponsibleRepository:  rr,
				ContractRepository:     cr,
				TempContractRepository: tcr,
			},
			logger,
			adapters.Adapters{
				PaymentsService: ps,
			},
		)

		err := uc.AcceptContract(makeASRAS(uuidStr))

		assert.EqualError(t, err, "update temp contract error")
		assert.Error(t, err)
	})

	t.Run("accept contract repository return error", func(t *testing.T) {
		rr := mocks.NewResponsibleRepository(t)
		cr := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)
		ps := mocks.NewPaymentsService(t)
		tcr := mocks.NewTempContractRepository(t)

		uuidStr := uuid.New().String()
		parsedUUID, _ := uuid.Parse(uuidStr)

		cr.On("ContractAlreadyExist", mock.Anything).Return(false, nil)
		rr.On("Get", mock.Anything).Return(&entity.Responsible{}, nil)
		ps.On("CreateProduct", mock.Anything).Return(&stripe.Product{}, nil)
		ps.On("CreatePrice", mock.Anything, mock.Anything).Return(&stripe.Price{}, nil)
		ps.On("CreateSubscription", mock.Anything, mock.Anything).Return(&stripe.Subscription{}, nil)
		tcr.On("Update", parsedUUID, map[string]interface{}{
			"responsible_signed_at": int64(1),
			"driver_signed_at":      int64(2),
		}).Return(nil)
		cr.On("Accept", mock.Anything).Return(errors.New("accept contract error"))

		uc := NewAcceptContractUseCase(
			&persistence.PostgresRepositories{
				ResponsibleRepository:  rr,
				ContractRepository:     cr,
				TempContractRepository: tcr,
			},
			logger,
			adapters.Adapters{
				PaymentsService: ps,
			},
		)

		err := uc.AcceptContract(agreements.ASRASOutput{
			Contract: entity.Contract{
				UUID: uuidStr,
			},
			Signatures: []agreements.Signature{
				{
					SignedAt: 1,
				},
				{
					SignedAt: 2,
				},
			},
		})

		assert.EqualError(t, err, "accept contract error")
		assert.Error(t, err)
	})

	t.Run("accept contract usecase return success", func(t *testing.T) {
		rr := mocks.NewResponsibleRepository(t)
		cr := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)
		ps := mocks.NewPaymentsService(t)
		tcr := mocks.NewTempContractRepository(t)

		uuidStr := uuid.New().String()
		parsedUUID, _ := uuid.Parse(uuidStr)

		cr.On("ContractAlreadyExist", mock.Anything).Return(false, nil)
		rr.On("Get", mock.Anything).Return(&entity.Responsible{}, nil)
		ps.On("CreateProduct", mock.Anything).Return(&stripe.Product{}, nil)
		ps.On("CreatePrice", mock.Anything, mock.Anything).Return(&stripe.Price{}, nil)
		ps.On("CreateSubscription", mock.Anything, mock.Anything).Return(&stripe.Subscription{}, nil)
		tcr.On("Update", parsedUUID, map[string]interface{}{
			"responsible_signed_at": int64(1),
			"driver_signed_at":      int64(2),
		}).Return(nil)
		cr.On("Accept", mock.Anything).Return(nil)

		uc := NewAcceptContractUseCase(
			&persistence.PostgresRepositories{
				ResponsibleRepository:  rr,
				ContractRepository:     cr,
				TempContractRepository: tcr,
			},
			logger,
			adapters.Adapters{
				PaymentsService: ps,
			},
		)

		err := uc.AcceptContract(agreements.ASRASOutput{
			Contract: entity.Contract{
				UUID: uuidStr,
			},
			Signatures: []agreements.Signature{
				{
					SignedAt: 1,
				},
				{
					SignedAt: 2,
				},
			},
		})

		assert.NoError(t, err)
		assert.Nil(t, err)
	})
}
