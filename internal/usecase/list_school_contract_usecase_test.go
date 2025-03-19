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

func TestListSchoolContractsUsecase_ListSchoolContracts(t *testing.T) {
	t.Run("get contract repository return error", func(t *testing.T) {
		contractRepository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)

		contractRepository.On("GetBySchool", mock.Anything).Return([]entity.EnableContract{}, errors.New("database error"))

		uc := NewListSchoolContractUseCase(
			&persistence.PostgresRepositories{
				ContractRepository: contractRepository,
			},
			logger,
		)

		_, err := uc.ListSchoolContract("123456789")

		assert.EqualError(t, err, "database error")
		assert.Error(t, err)
	})

	t.Run("list school contracts return success", func(t *testing.T) {
		contractRepository := mocks.NewContractRepository(t)
		logger := mocks.NewLogger(t)

		contractRepository.On("GetBySchool", mock.Anything).Return([]entity.EnableContract{}, nil)

		uc := NewListSchoolContractUseCase(
			&persistence.PostgresRepositories{
				ContractRepository: contractRepository,
			},
			logger,
		)

		_, err := uc.ListSchoolContract("123456789")

		assert.NoError(t, err)
		assert.Nil(t, err)
	})
}
