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

func TestListKidUsecase_ListKid(t *testing.T) {
	cpf := "123"

	t.Run("if repository returns error", func(t *testing.T) {
		repository := mocks.NewKidRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewListKidsUseCase(
			&persistence.PostgresRepositories{
				KidRepository: repository,
			},
			logger,
		)

		repository.On("FindAll", mock.Anything).Return([]entity.Kid{}, errors.New("database error"))

		_, err := usecase.ListKids(&cpf)

		assert.EqualError(t, err, "database error")
	})

	t.Run("if list returns sucess", func(t *testing.T) {
		repository := mocks.NewKidRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewListKidsUseCase(
			&persistence.PostgresRepositories{
				KidRepository: repository,
			},
			logger,
		)

		repository.On("FindAll", mock.Anything).Return([]entity.Kid{}, nil)

		_, err := usecase.ListKids(&cpf)

		assert.NoError(t, err)
	})
}
