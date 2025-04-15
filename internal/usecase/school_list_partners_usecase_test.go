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

func TestListSchoolPartnersUsecase_ListDriver(t *testing.T) {
	cpnj := "12323"

	t.Run("if partner repository returns error", func(t *testing.T) {
		repository := mocks.NewPartnerRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewSchoolListPartnersUseCase(
			&persistence.PostgresRepositories{
				PartnerRepository: repository,
			},
			logger,
		)
		repository.On("GetBySchool", mock.Anything).Return([]entity.Partner{}, errors.New("database error"))

		_, err := usecase.SchoolListPartners(cpnj)

		assert.EqualError(t, err, "database error")
	})

	t.Run("if list returns sucess", func(t *testing.T) {
		repository := mocks.NewPartnerRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewSchoolListPartnersUseCase(
			&persistence.PostgresRepositories{
				PartnerRepository: repository,
			},
			logger,
		)
		repository.On("GetBySchool", mock.Anything).Return([]entity.Partner{}, nil)

		_, err := usecase.SchoolListPartners(cpnj)

		assert.NoError(t, err)
	})
}
