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

func TestListSchoolInviteUsecase_ListSchoolInvites(t *testing.T) {
	cnpj := "87987896000122"

	t.Run("if repository returns error", func(t *testing.T) {
		repository := mocks.NewInviteRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewListSchoolInvitesUseCase(
			&persistence.PostgresRepositories{
				InviteRepository: repository,
			},
			logger,
		)
		repository.On("GetBySchool", mock.Anything).Return([]entity.Driver{}, errors.New("database error"))

		_, err := usecase.ListSchoolInvites(cnpj)

		assert.EqualError(t, err, "database error")
	})

	t.Run("if list returns sucess", func(t *testing.T) {
		repository := mocks.NewInviteRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewListSchoolInvitesUseCase(
			&persistence.PostgresRepositories{
				InviteRepository: repository,
			},
			logger,
		)
		repository.On("GetBySchool", mock.Anything).Return([]entity.Driver{}, nil)

		_, err := usecase.ListSchoolInvites(cnpj)

		assert.NoError(t, err)
	})
}
