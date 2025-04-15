package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/mocks"
)

func TestUpdateSchoolUsecase_UpdateSchool(t *testing.T) {
	t.Run("if someone try send unknown field to update", func(t *testing.T) {
		repository := mocks.NewSchoolRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewUpdateSchoolUseCase(
			&persistence.PostgresRepositories{
				SchoolRepository: repository,
			},
			logger,
		)

		err := usecase.UpdateSchool("123", map[string]interface{}{
			"cpnj": "123",
		})

		assert.EqualError(t, err, "chaves não permitidas: cpnj")
		assert.Error(t, err)
	})

	t.Run("when proxy to update returns error", func(t *testing.T) {
		SchoolRepository := mocks.NewSchoolRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewUpdateSchoolUseCase(
			&persistence.PostgresRepositories{
				SchoolRepository: SchoolRepository,
			},
			logger,
		)
		SchoolRepository.On("Update", mock.Anything, mock.Anything).Return(errors.New("database error"))

		err := usecase.UpdateSchool("123", map[string]interface{}{
			"name":     "123",
			"email":    "amostradinho@gmail.com",
			"password": "17",
		})

		assert.EqualError(t, err, "database error")
	})

	t.Run("when proxy to update give success ", func(t *testing.T) {
		SchoolRepository := mocks.NewSchoolRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewUpdateSchoolUseCase(
			&persistence.PostgresRepositories{
				SchoolRepository: SchoolRepository,
			},
			logger,
		)
		SchoolRepository.On("Update", mock.Anything, mock.Anything).Return(nil)

		err := usecase.UpdateSchool("123", map[string]interface{}{
			"email": "amostradinho@gmail.com",
		})

		assert.Nil(t, err)
	})
}
