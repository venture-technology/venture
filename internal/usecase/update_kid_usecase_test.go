package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/mocks"
)

func TestUpdateKidUsecase_UpdateKid(t *testing.T) {
	t.Run("if someone try send unknown field to update", func(t *testing.T) {
		repository := mocks.NewKidRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewUpdateKidUseCase(
			&persistence.PostgresRepositories{
				KidRepository: repository,
			},
			logger,
		)

		err := usecase.UpdateKid("123", map[string]interface{}{
			"rg": "1234567890",
		})

		assert.EqualError(t, err, "chaves não permitidas: rg")
		assert.Error(t, err)
	})

	t.Run("if someone try send unknown shift to update kid", func(t *testing.T) {
		repository := mocks.NewKidRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewUpdateKidUseCase(
			&persistence.PostgresRepositories{
				KidRepository: repository,
			},
			logger,
		)

		err := usecase.UpdateKid("123", map[string]interface{}{
			"shift": "teste",
		})

		assert.EqualError(t, err, "shift inexistente")
		assert.Error(t, err)
	})

	t.Run("if someone try send unknown type of shift to update kid", func(t *testing.T) {
		repository := mocks.NewKidRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewUpdateKidUseCase(
			&persistence.PostgresRepositories{
				KidRepository: repository,
			},
			logger,
		)

		err := usecase.UpdateKid("123", map[string]interface{}{
			"shift": 123,
		})

		assert.EqualError(t, err, "shift invalido")
		assert.Error(t, err)
	})

	t.Run("if someone try send unknown type of shift to update kid", func(t *testing.T) {
		repository := mocks.NewKidRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewUpdateKidUseCase(
			&persistence.PostgresRepositories{
				KidRepository: repository,
			},
			logger,
		)

		err := usecase.UpdateKid("123", map[string]interface{}{
			"shift": 123,
		})

		assert.EqualError(t, err, "shift invalido")
		assert.Error(t, err)
	})

	t.Run("when repository returns error", func(t *testing.T) {
		repository := mocks.NewKidRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewUpdateKidUseCase(
			&persistence.PostgresRepositories{
				KidRepository: repository,
			},
			logger,
		)

		repository.On("Update", mock.Anything, mock.Anything).Return(errors.New("database error"))

		err := usecase.UpdateKid("123", map[string]interface{}{
			"shift": "afternoon",
		})

		assert.EqualError(t, err, "database error")
	})

	t.Run("when proxy give success with shift", func(t *testing.T) {
		repository := mocks.NewKidRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewUpdateKidUseCase(
			&persistence.PostgresRepositories{
				KidRepository: repository,
			},
			logger,
		)

		repository.On("Update", mock.Anything, mock.Anything).Return(nil)

		err := usecase.UpdateKid("123", map[string]interface{}{
			"shift": "afternoon",
		})

		assert.Nil(t, err)
	})

	t.Run("when proxy give success without shift", func(t *testing.T) {
		repository := mocks.NewKidRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewUpdateKidUseCase(
			&persistence.PostgresRepositories{
				KidRepository: repository,
			},
			logger,
		)

		repository.On("Update", mock.Anything, mock.Anything).Return(nil)

		err := usecase.UpdateKid("123", map[string]interface{}{
			"attendance_permission": false,
		})

		assert.Nil(t, err)
	})
}
