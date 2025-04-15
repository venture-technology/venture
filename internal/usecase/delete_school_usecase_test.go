package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/mocks"
)

func TestDeleteSchoolUsecase_DeleteSchool(t *testing.T) {
	t.Run("if contract repository returns error", func(t *testing.T) {
		repository := mocks.NewSchoolRepository(t)
		logger := mocks.NewLogger(t)
		cr := mocks.NewContractRepository(t)

		usecase := NewDeleteSchoolUseCase(
			&persistence.PostgresRepositories{
				ContractRepository: cr,
				SchoolRepository:   repository,
			},
			logger,
		)
		cnpj := "1234"
		cr.On("SchoolHasEnableContract", mock.Anything).Return(true, errors.New("impossivel deletar escola possuindo contrato ativo"))
		err := usecase.DeleteSchool(cnpj)

		assert.EqualError(t, err, "impossivel deletar escola possuindo contrato ativo")
	})

	t.Run("if delete kid on repository returns error", func(t *testing.T) {
		repository := mocks.NewSchoolRepository(t)
		logger := mocks.NewLogger(t)
		cr := mocks.NewContractRepository(t)

		usecase := NewDeleteSchoolUseCase(
			&persistence.PostgresRepositories{
				ContractRepository: cr,
				SchoolRepository:   repository,
			},
			logger,
		)
		cnpj := "1234"
		cr.On("SchoolHasEnableContract", mock.Anything).Return(false, nil)
		repository.On("Delete", cnpj).Return(errors.New("school repository delete error")).Once()

		err := usecase.DeleteSchool(cnpj)

		assert.EqualError(t, err, "school repository delete error")
	})

	t.Run("when delete kid return success", func(t *testing.T) {
		repository := mocks.NewSchoolRepository(t)
		logger := mocks.NewLogger(t)
		cr := mocks.NewContractRepository(t)

		usecase := NewDeleteSchoolUseCase(
			&persistence.PostgresRepositories{
				ContractRepository: cr,
				SchoolRepository:   repository,
			},
			logger,
		)
		cnpj := "1234"
		cr.On("SchoolHasEnableContract", mock.Anything).Return(false, nil)
		repository.On("Delete", cnpj).Return(nil).Once()
		err := usecase.DeleteSchool(cnpj)

		assert.Nil(t, err)
	})
}
