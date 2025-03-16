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

func TestGetKidUsecase_GetKid(t *testing.T) {
	rg := "123"

	kid := entity.Kid{}

	t.Run("if repository returns error", func(t *testing.T) {
		repository := mocks.NewKidRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewGetKidUseCase(
			&persistence.PostgresRepositories{
				KidRepository: repository,
			},
			logger,
		)

		repository.On("Get", mock.Anything).Return(&kid, errors.New("database error"))

		_, err := usecase.GetKid(&rg)

		assert.EqualError(t, err, "database error")
	})

	t.Run("ir list returns success", func(t *testing.T) {
		repository := mocks.NewKidRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewGetKidUseCase(
			&persistence.PostgresRepositories{
				KidRepository: repository,
			},
			logger,
		)

		repository.On("Get", mock.Anything).Return(&kid, nil)

		_, err := usecase.GetKid(&rg)

		assert.Nil(t, err)
	})
}
