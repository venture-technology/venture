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

func TestListDriverInviteUsecase_ListDriverInvites(t *testing.T) {
	cnh := "87987896000122"

	t.Run("if repository returns error", func(t *testing.T) {
		repository := mocks.NewInviteRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewListDriverInvitesUseCase(
			&persistence.PostgresRepositories{
				InviteRepository: repository,
			},
			logger,
		)
		repository.On("GetByDriver", mock.Anything).Return([]entity.School{}, errors.New("database error"))

		_, err := usecase.ListDriverInvites(cnh)

		assert.EqualError(t, err, "database error")
	})

	t.Run("if list returns sucess", func(t *testing.T) {
		repository := mocks.NewInviteRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewListDriverInvitesUseCase(
			&persistence.PostgresRepositories{
				InviteRepository: repository,
			},
			logger,
		)
		repository.On("GetByDriver", mock.Anything).Return([]entity.School{}, nil)
		logger.On("Infof", mock.Anything, mock.Anything).Return()

		_, err := usecase.ListDriverInvites(cnh)

		assert.NoError(t, err)
	})
}
