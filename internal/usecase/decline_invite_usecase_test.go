package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/mocks"
)

func TestDeclineInviteUsecase_Accept_Invite(t *testing.T) {
	id := "123"
	t.Run("if Invite repository returns error", func(t *testing.T) {
		logger := mocks.NewLogger(t)
		repository := mocks.NewInviteRepository(t)

		usecase := NewDeclineInviteUseCase(
			&persistence.PostgresRepositories{
				InviteRepository: repository,
			},
			logger,
		)
		repository.On("Decline", mock.Anything).Return(errors.New("database error"))

		err := usecase.DeclineInvite(id)

		assert.EqualError(t, err, "database error")
	})

	t.Run("when invite repository returns success", func(t *testing.T) {
		logger := mocks.NewLogger(t)
		repository := mocks.NewInviteRepository(t)

		usecase := NewDeclineInviteUseCase(
			&persistence.PostgresRepositories{
				InviteRepository: repository,
			},
			logger,
		)
		repository.On("Decline", mock.Anything).Return(nil)

		err := usecase.DeclineInvite(id)

		assert.Nil(t, err)
	})
}
