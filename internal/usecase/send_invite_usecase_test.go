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

func TestSendInviteUsecase_Accept_Invite(t *testing.T) {
	invite := entity.Invite{}

	t.Run("if Invite repository returns error", func(t *testing.T) {
		logger := mocks.NewLogger(t)
		repository := mocks.NewInviteRepository(t)

		usecase := NewSendInviteUseCase(
			&persistence.PostgresRepositories{
				InviteRepository: repository,
			},
			logger,
		)
		repository.On("Create", mock.Anything).Return(errors.New("database error"))

		err := usecase.SendInvite(&invite)

		assert.EqualError(t, err, "database error")
	})

	t.Run("when invite repository returns success", func(t *testing.T) {
		logger := mocks.NewLogger(t)
		repository := mocks.NewInviteRepository(t)

		usecase := NewSendInviteUseCase(
			&persistence.PostgresRepositories{
				InviteRepository: repository,
			},
			logger,
		)
		repository.On("Create", mock.Anything).Return(nil)

		err := usecase.SendInvite(&invite)

		assert.Nil(t, err)
	})
}
