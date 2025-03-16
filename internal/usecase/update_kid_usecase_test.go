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
			"phone": "123",
		})

		assert.EqualError(t, err, "database error")
	})

	t.Run("when proxy give success", func(t *testing.T) {
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
			"phone": "123",
		})

		assert.Nil(t, err)
	})
}
