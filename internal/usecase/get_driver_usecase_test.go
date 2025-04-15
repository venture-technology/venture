package usecase

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/mocks"
)

func TestGetDriverUsecase_GetDriver(t *testing.T) {

	cnh := "123"
	driver := entity.Driver{
		Name:  "motorista",
		Email: "motorista@gmail.com",
	}

	t.Run("if repository returns error", func(t *testing.T) {
		repository := mocks.NewDriverRepository(t)
		logger := mocks.NewLogger(t)
		s3iface := mocks.NewS3Iface(t)

		usecase := NewGetDriverUseCase(
			&persistence.PostgresRepositories{
				DriverRepository: repository,
			},
			logger,
			s3iface,
		)

		repository.On("Get", mock.Anything).Return(&driver, errors.New("database error"))

		_, err := usecase.GetDriver(cnh)

		assert.EqualError(t, err, "database error")
	})

	t.Run("ir list returns success", func(t *testing.T) {
		repository := mocks.NewDriverRepository(t)
		logger := mocks.NewLogger(t)
		s3iface := mocks.NewS3Iface(t)

		usecase := NewGetDriverUseCase(
			&persistence.PostgresRepositories{
				DriverRepository: repository,
			},
			logger,
			s3iface,
		)

		repository.On("Get", mock.Anything).Return(&driver, nil)
		images := []string{"image1.jpg", "image2.jpg"}
		s3iface.On("List", mock.Anything, fmt.Sprintf("%s/gallery", cnh)).Return(images, nil)

		response, err := usecase.GetDriver(cnh)

		assert.Nil(t, err)
		assert.EqualValues(t, driver, entity.Driver{
			Name:  response.Name,
			Email: response.Email,
		})
	})
}
