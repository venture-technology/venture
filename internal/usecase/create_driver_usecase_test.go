package usecase

import (
	"errors"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/mocks"
)

func TestCreateDriverUsecase_CreateDriver(t *testing.T) {
	driver := &entity.Driver{
		CNH:      "37886632129",
		Schedule: "afternoon",
		Car: entity.Car{
			Capacity: 60,
		},
		Password: "123Teste@",
	}

	t.Run("if s3 returns error", func(t *testing.T) {
		repository := mocks.NewDriverRepository(t)
		logger := mocks.NewLogger(t)
		s3iface := mocks.NewS3Iface(t)

		usecase := NewCreateDriverUseCase(
			&persistence.PostgresRepositories{
				DriverRepository: repository,
			},
			logger,
			s3iface,
		)
		s3iface.On("Save", mock.Anything, driver.CNH, "qrcode", mock.Anything, mock.Anything).Return("", errors.New("s3iface error"))

		err := usecase.CreateDriver(driver)

		assert.EqualError(t, err, "s3iface error")
	})

	t.Run("if driver repository returns error", func(t *testing.T) {
		repository := mocks.NewDriverRepository(t)
		logger := mocks.NewLogger(t)
		s3iface := mocks.NewS3Iface(t)

		usecase := NewCreateDriverUseCase(
			&persistence.PostgresRepositories{
				DriverRepository: repository,
			},
			logger,
			s3iface,
		)
		s3iface.On("Save", mock.Anything, driver.CNH, "qrcode", mock.Anything, mock.Anything).Return("links3", nil)
		repository.On("Create", mock.Anything).Return(errors.New("failed driver repository"))

		driver.Schedule = "afternoon"
		log.Println(driver)
		err := usecase.CreateDriver(driver)

		assert.EqualError(t, err, "failed driver repository")
	})

	t.Run("if s3 returns sucess", func(t *testing.T) {
		repository := mocks.NewDriverRepository(t)
		logger := mocks.NewLogger(t)
		s3iface := mocks.NewS3Iface(t)

		usecase := NewCreateDriverUseCase(
			&persistence.PostgresRepositories{
				DriverRepository: repository,
			},
			logger,
			s3iface,
		)
		s3iface.On("Save", mock.Anything, driver.CNH, "qrcode", mock.Anything, mock.Anything).Return("link3", nil)
		repository.On("Create", mock.Anything).Return(nil)

		driver.Schedule = "afternoon"

		err := usecase.CreateDriver(driver)

		assert.NoError(t, err)
	})
}
