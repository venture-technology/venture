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

func TestListDriverUsecase_ListDriver(t *testing.T) {
	cnpj := "87987896000122"

	t.Run("if repository returns error", func(t *testing.T) {
		repository := mocks.NewPartnerRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewListDriverFromSchoolUseCase(
			&persistence.PostgresRepositories{
				PartnerRepository: repository,
			},
			logger,
		)
		repository.On("GetBySchool", mock.Anything).Return([]entity.Partner{}, errors.New("database error"))

		_, err := usecase.ListDriverFromSchool(cnpj)

		assert.EqualError(t, err, "database error")
	})

	t.Run("if list returns sucess", func(t *testing.T) {
		repository := mocks.NewPartnerRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewListDriverFromSchoolUseCase(
			&persistence.PostgresRepositories{
				PartnerRepository: repository,
			},
			logger,
		)
		repository.On("GetBySchool", mock.Anything).Return([]entity.Partner{}, nil)

		_, err := usecase.ListDriverFromSchool(cnpj)

		assert.NoError(t, err)
	})
}
