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

func TestDriverLoginUsecase_LoginDriver(t *testing.T) {
	pwd := "password"
	password, err := utils.MakeHash(pwd)
	if err != nil {
		t.Errorf("make hash func error %s", err.Error())
	}

	t.Run("when repository not found user", func(t *testing.T) {
		dr := mocks.NewDriverRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewDriverLoginUsecase(
			&persistence.PostgresRepositories{
				DriverRepository: dr,
			},
			logger,
			config.Config{
				Server: config.Server{
					Secret: "teste",
				},
			},
		)

		dr.On("GetByEmail", mock.Anything).Return(&entity.Driver{}, errors.New("user not found"))

		_, err := usecase.LoginDriver("email", "password")

		assert.Error(t, err)
	})

	t.Run("when validate hash return comparing error", func(t *testing.T) {
		dr := mocks.NewDriverRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewDriverLoginUsecase(
			&persistence.PostgresRepositories{
				DriverRepository: dr,
			},
			logger,
			config.Config{
				Server: config.Server{
					Secret: "teste",
				},
			},
		)

		dr.On("GetByEmail", mock.Anything).Return(&entity.Driver{
			Password: password,
		}, nil)

		_, err := usecase.LoginDriver("email", "passwords")

		assert.Error(t, err)
	})

	t.Run("when login return success", func(t *testing.T) {
		dr := mocks.NewDriverRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewDriverLoginUsecase(
			&persistence.PostgresRepositories{
				DriverRepository: dr,
			},
			logger,
			config.Config{
				Server: config.Server{
					Secret: "teste",
				},
			},
		)

		dr.On("GetByEmail", mock.Anything).Return(&entity.Driver{
			Password: password,
		}, nil)

		_, err := usecase.LoginDriver("email", "password")

		assert.Nil(t, err)
	})
}
