package usecase

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/mocks"
)

func TestGetSchoolUseCase_GetSchool(t *testing.T) {
	t.Run("when repository returns error", func(t *testing.T) {
		schoolRepository := mocks.NewSchoolRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewGetSchoolUseCase(
			&persistence.PostgresRepositories{
				SchoolRepository: schoolRepository,
			},
			logger,
		)

		schoolRepository.On("Get", mock.Anything).Return(&entity.School{}, fmt.Errorf("database error"))

		_, err := usecase.GetSchool("123")

		assert.EqualError(t, err, "database error")
	})

	t.Run("when usecase returns success", func(t *testing.T) {
		schoolRepository := mocks.NewSchoolRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewGetSchoolUseCase(
			&persistence.PostgresRepositories{
				SchoolRepository: schoolRepository,
			},
			logger,
		)

		school := entity.School{
			Name:  "Escola",
			Email: "email@escola.com",
		}

		schoolRepository.On("Get", mock.Anything).Return(&school, nil)

		response, err := usecase.GetSchool("123")

		assert.Nil(t, err)
		assert.EqualValues(t, school, entity.School{
			Name:  response.Name,
			Email: response.Email,
		})
	})
}
