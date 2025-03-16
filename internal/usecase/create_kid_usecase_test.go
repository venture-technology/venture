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

func TestCreateKidUsecase_CreateKid(t *testing.T) {
	kid := entity.Kid{}

	t.Run("if create kid on repository returns error", func(t *testing.T) {
		repository := mocks.NewKidRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewCreateKidUseCase(
			&persistence.PostgresRepositories{
				KidRepository: repository,
			},
			logger,
		)
		repository.On("Create", mock.Anything).Return(errors.New("repository error"))

		err := usecase.CreateKid(&kid)

		assert.EqualError(t, err, "repository error")
	})

	t.Run("when creation of kid return success", func(t *testing.T) {
		repository := mocks.NewKidRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewCreateKidUseCase(
			&persistence.PostgresRepositories{
				KidRepository: repository,
			},
			logger,
		)

		repository.On("Create", mock.Anything).Return(nil)

		err := usecase.CreateKid(&kid)

		assert.Nil(t, err)
	})
}
