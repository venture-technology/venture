package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/venture-technology/venture/config"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/mocks"
	"github.com/venture-technology/venture/pkg/utils"
)

func TestResponsibleLoginUsecase_LoginResponsible(t *testing.T) {
	pwd := "password"
	password, err := utils.MakeHash(pwd)
	if err != nil {
		t.Errorf("make hash func error %s", err.Error())
	}

	t.Run("when repository not found user", func(t *testing.T) {
		rr := mocks.NewResponsibleRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewResponsibleLoginUsecase(
			&persistence.PostgresRepositories{
				ResponsibleRepository: rr,
			},
			logger,
			config.Config{
				Server: config.Server{
					Secret: "teste",
				},
			},
		)

		rr.On("GetByEmail", mock.Anything).Return(&entity.Responsible{}, errors.New("user not found"))

		_, err := usecase.LoginResponsible("email", "password")

		assert.Error(t, err)
	})

	t.Run("when validate hash return comparing error", func(t *testing.T) {
		rr := mocks.NewResponsibleRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewResponsibleLoginUsecase(
			&persistence.PostgresRepositories{
				ResponsibleRepository: rr,
			},
			logger,
			config.Config{
				Server: config.Server{
					Secret: "teste",
				},
			},
		)

		rr.On("GetByEmail", mock.Anything).Return(&entity.Responsible{
			Password: password,
		}, nil)

		_, err := usecase.LoginResponsible("email", "passwords")

		assert.Error(t, err)
	})

	t.Run("when login return success", func(t *testing.T) {
		rr := mocks.NewResponsibleRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewResponsibleLoginUsecase(
			&persistence.PostgresRepositories{
				ResponsibleRepository: rr,
			},
			logger,
			config.Config{
				Server: config.Server{
					Secret: "teste",
				},
			},
		)

		rr.On("GetByEmail", mock.Anything).Return(&entity.Responsible{
			Password: password,
		}, nil)

		_, err := usecase.LoginResponsible("email", "password")

		assert.Nil(t, err)
	})
}
