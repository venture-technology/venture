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

func TestGetResponsibleUsecase_GetResponsible(t *testing.T) {
	cpf := "123"

	responsible := entity.Responsible{}

	t.Run("if repository returns error", func(t *testing.T) {
		repository := mocks.NewResponsibleRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewGetResponsibleUseCase(
			&persistence.PostgresRepositories{
				ResponsibleRepository: repository,
			},
			logger,
		)

		repository.On("Get", mock.Anything).Return(&responsible, errors.New("database error"))

		_, err := usecase.GetResponsible(cpf)

		assert.EqualError(t, err, "database error")
	})

	t.Run("ir list returns success", func(t *testing.T) {
		repository := mocks.NewResponsibleRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewGetResponsibleUseCase(
			&persistence.PostgresRepositories{
				ResponsibleRepository: repository,
			},
			logger,
		)

		repository.On("Get", mock.Anything).Return(&responsible, nil)

		_, err := usecase.GetResponsible(cpf)

		assert.Nil(t, err)
	})
}
