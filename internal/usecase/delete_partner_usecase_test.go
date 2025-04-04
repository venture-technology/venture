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

func TestDeletePartnerUsecase_DeletePartner(t *testing.T) {

	t.Run("if contract repository returns error", func(t *testing.T) {
		logger := mocks.NewLogger(t)
		cr := mocks.NewContractRepository(t)

		usecase := NewDeletePartnerUseCase(
			&persistence.PostgresRepositories{
				ContractRepository: cr,
			},
			logger,
		)
		partnerID := "123"
		cr.On("PartnerHasEnableContract", partnerID).Return([]entity.Contract{}, errors.New("database error"))
		err := usecase.DeletePartner(partnerID)

		assert.EqualError(t, err, "database error")
	})

	t.Run("if delete Partner on repository returns error", func(t *testing.T) {
		logger := mocks.NewLogger(t)
		cr := mocks.NewContractRepository(t)
		repository := mocks.NewPartnerRepository(t)

		usecase := NewDeletePartnerUseCase(
			&persistence.PostgresRepositories{
				ContractRepository: cr,
				PartnerRepository:  repository,
			},
			logger,
		)

		partnerID := "123"
		cr.On("PartnerHasEnableContract", partnerID).Return([]entity.Contract{}, nil)
		repository.On("Delete", mock.Anything).Return(errors.New("impossible delete partner because it has enable contract"))

		err := usecase.DeletePartner(partnerID)

		assert.EqualError(t, err, "impossible delete partner because it has enable contract")
	})

	t.Run("when delete partner return success", func(t *testing.T) {
		logger := mocks.NewLogger(t)
		cr := mocks.NewContractRepository(t)
		repository := mocks.NewPartnerRepository(t)

		usecase := NewDeletePartnerUseCase(
			&persistence.PostgresRepositories{
				ContractRepository: cr,
				PartnerRepository:  repository,
			},
			logger,
		)
		partnerID := "123"
		cr.On("PartnerHasEnableContract", partnerID).Return([]entity.Contract{}, nil)
		repository.On("Delete", partnerID).Return(nil)
		err := usecase.DeletePartner(partnerID)

		assert.Nil(t, err)
	})
}
