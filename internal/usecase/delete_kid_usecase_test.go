package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/mocks"
)

func TestDeleteKidUsecase_DeleteKid(t *testing.T) {

	t.Run("if delete kid on repository returns error", func(t *testing.T) {
		repository := mocks.NewKidRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewDeleteKidUseCase(
			&persistence.PostgresRepositories{
				KidRepository: repository,
			},
			logger,
		)
		rg := "1234"
		repository.On("Delete", &rg).Return(errors.New("kid repository delete error")).Once()

		err := usecase.DeleteKid(&rg)

		assert.EqualError(t, err, "kid repository delete error")
	})

	t.Run("when delete kid return success", func(t *testing.T) {
		repository := mocks.NewKidRepository(t)
		logger := mocks.NewLogger(t)

		usecase := NewDeleteKidUseCase(
			&persistence.PostgresRepositories{
				KidRepository: repository,
			},
			logger,
		)
		rg := "1234"
		repository.On("Delete", &rg).Return(nil).Once()
		err := usecase.DeleteKid(&rg)

		assert.Nil(t, err)
	})
}
