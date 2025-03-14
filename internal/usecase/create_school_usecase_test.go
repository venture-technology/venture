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

func TestCreateSchoolUseCase_CreateSchool(t *testing.T) {
	t.Run("when repository returns error", func(t *testing.T) {
		schoolRepository := mocks.NewSchoolRepository(t)
		logger := mocks.NewLogger(t)

		school := entity.School{}

		usecase := NewCreateSchoolUseCase(
			&persistence.PostgresRepositories{
				SchoolRepository: schoolRepository,
			},
			logger,
		)

		schoolRepository.On("Create", mock.Anything).Return(fmt.Errorf("database error"))

		err := usecase.CreateSchool(&school)

		assert.EqualError(t, err, "database error")
	})

	t.Run("when usecase returns success", func(t *testing.T) {
		schoolRepository := mocks.NewSchoolRepository(t)
		logger := mocks.NewLogger(t)

		school := entity.School{}

		usecase := NewCreateSchoolUseCase(
			&persistence.PostgresRepositories{
				SchoolRepository: schoolRepository,
			},
			logger,
		)

		schoolRepository.On("Create", mock.Anything).Return(nil)

		err := usecase.CreateSchool(&school)

		assert.Nil(t, err)
	})
}
