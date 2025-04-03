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

func TestDeleteDriverUsecase_DeleteDriver(t *testing.T) {
	driver := entity.Driver{
		CNH:      "37886632129",
		Schedule: "afternoon",
		Car: entity.Car{
			Capacity: 60,
		},
	}

	t.Run("if contract repository returns error", func(t *testing.T) {
		repository := mocks.NewDriverRepository(t)
		logger := mocks.NewLogger(t)
		cr := mocks.NewContractRepository(t)

		usecase := NewDeleteDriverUseCase(
			&persistence.PostgresRepositories{
				ContractRepository: cr,
				DriverRepository:   repository,
			},
			logger,
		)
		cr.On("DriverHasEnableContract", mock.Anything).Return(true, nil)
		err := usecase.DeleteDriver(driver.CNH)

		assert.EqualError(t, err, "impossivel deletar motorista possuindo contrata ativo")
	})

	t.Run("if delete Driver on repository returns error", func(t *testing.T) {
		repository := mocks.NewDriverRepository(t)
		logger := mocks.NewLogger(t)
		cr := mocks.NewContractRepository(t)

		usecase := NewDeleteDriverUseCase(
			&persistence.PostgresRepositories{
				ContractRepository: cr,
				DriverRepository:   repository,
			},
			logger,
		)
		cr.On("DriverHasEnableContract", mock.Anything).Return(false, nil)
		repository.On("Delete", mock.Anything).Return(errors.New("Driver repository delete error"))

		err := usecase.DeleteDriver(driver.CNH)

		assert.EqualError(t, err, "Driver repository delete error")
	})

	t.Run("when delete Driver return success", func(t *testing.T) {
		repository := mocks.NewDriverRepository(t)
		logger := mocks.NewLogger(t)
		cr := mocks.NewContractRepository(t)

		usecase := NewDeleteDriverUseCase(
			&persistence.PostgresRepositories{
				ContractRepository: cr,
				DriverRepository:   repository,
			},
			logger,
		)
		cr.On("DriverHasEnableContract", mock.Anything).Return(false, nil)
		repository.On("Delete", mock.Anything).Return(nil)

		err := usecase.DeleteDriver(driver.CNH)

		assert.Nil(t, err)
	})
}
