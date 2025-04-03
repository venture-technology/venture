package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/mocks"
)

func TestDeleteKidUsecase_DeleteKid(t *testing.T) {

	t.Run("if contract repository returns error", func(t *testing.T) {
		repository := mocks.NewKidRepository(t)
		logger := mocks.NewLogger(t)
		cr := mocks.NewContractRepository(t)

		usecase := NewDeleteKidUseCase(
			&persistence.PostgresRepositories{
				ContractRepository: cr,
				KidRepository:      repository,
			},
			logger,
		)
		rg := "1234"
		cr.On("KidHasEnableContract", mock.Anything).Return(true, nil)
		err := usecase.DeleteKid(&rg)

		assert.EqualError(t, err, "impossivel deletar criança possuindo contrato ativo")
	})

	t.Run("if delete kid on repository returns error", func(t *testing.T) {
		repository := mocks.NewKidRepository(t)
		logger := mocks.NewLogger(t)
		cr := mocks.NewContractRepository(t)

		usecase := NewDeleteKidUseCase(
			&persistence.PostgresRepositories{
				ContractRepository: cr,
				KidRepository:      repository,
			},
			logger,
		)
		rg := "1234"
		cr.On("KidHasEnableContract", mock.Anything).Return(false, nil)
		repository.On("Delete", &rg).Return(errors.New("kid repository delete error")).Once()

		err := usecase.DeleteKid(&rg)

		assert.EqualError(t, err, "kid repository delete error")
	})

	t.Run("when delete kid return success", func(t *testing.T) {
		repository := mocks.NewKidRepository(t)
		logger := mocks.NewLogger(t)
		cr := mocks.NewContractRepository(t)

		usecase := NewDeleteKidUseCase(
			&persistence.PostgresRepositories{
				ContractRepository: cr,
				KidRepository:      repository,
			},
			logger,
		)
		rg := "1234"
		cr.On("KidHasEnableContract", mock.Anything).Return(false, nil)
		repository.On("Delete", &rg).Return(nil).Once()
		err := usecase.DeleteKid(&rg)

		assert.Nil(t, err)
	})
}
