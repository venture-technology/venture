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

func TestListSchoolUsecase_ListSchool(t *testing.T) {
	t.Run("if repository returns error", func(t *testing.T) {
		repository := mocks.NewSchoolRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewListSchoolUseCase(
			&persistence.PostgresRepositories{
				SchoolRepository: repository,
			},
			logger,
		)
		repository.On("FindAll", mock.Anything).Return([]entity.School{}, errors.New("database error"))

		_, err := usecase.ListSchool()

		assert.EqualError(t, err, "database error")
	})

	t.Run("if list returns sucess", func(t *testing.T) {
		repository := mocks.NewSchoolRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewListSchoolUseCase(
			&persistence.PostgresRepositories{
				SchoolRepository: repository,
			},
			logger,
		)
		repository.On("FindAll", mock.Anything).Return([]entity.School{}, nil)

		_, err := usecase.ListSchool()

		assert.NoError(t, err)
	})
}
