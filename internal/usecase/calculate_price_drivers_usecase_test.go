package usecase

import (
	"errors"
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/venture-technology/venture/internal/domain/service/adapters"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/mocks"
)

func TestCalculatePrice_ResponsibleRepoError(t *testing.T) {
	responsible := &entity.Responsible{
		CPF: "12345678900",
		Address: entity.Address{
			Street:     "Rua 1",
			Number:     "100",
			Complement: "Apto 101",
			Zip:        "12345-678",
		},
	}
	school := &entity.School{
		CNPJ: "98765432000100",
		Address: entity.Address{
			Street:     "Rua 2",
			Number:     "200",
			Complement: "Bloco B",
			Zip:        "23456-789",
		},
	}

	responsibleAddress := fmt.Sprintf("%s, %s, %s, %s",
		responsible.Address.Street,
		responsible.Address.Number,
		responsible.Address.Complement,
		responsible.Address.Zip,
	)
	schoolAddress := fmt.Sprintf("%s, %s, %s, %s",
		school.Address.Street,
		school.Address.Number,
		school.Address.Complement,
		school.Address.Zip,
	)
	t.Run("if responsible repository returns error", func(t *testing.T) {
		logger := mocks.NewLogger(t)
		rr := mocks.NewResponsibleRepository(t)
		cr := mocks.NewContractRepository(t)
		pr := mocks.NewPartnerRepository(t)
		as := mocks.NewAddressService(t)
		sr := mocks.NewSchoolRepository(t)

		usecase := NewCalculatePriceDriversUseCase(
			&persistence.PostgresRepositories{
				PartnerRepository:     pr,
				ContractRepository:    cr,
				ResponsibleRepository: rr,
				SchoolRepository:      sr,
			},
			logger,
			adapters.Adapters{
				AddressService: as,
			},
		)

		// Configurar o mock para erro no repositório de responsáveis
		rr.On("Get", responsible.CPF).Return(responsible, errors.New("responsible get error"))
		logger.On("Infof", mock.Anything, mock.Anything).Return(logger)

		// Executar o método
		_, err := usecase.CalculatePrice(responsible.CPF, school.CNPJ)

		assert.EqualError(t, err, "responsible get error")
	})

	t.Run("if school repository returns error", func(t *testing.T) {
		logger := mocks.NewLogger(t)
		rr := mocks.NewResponsibleRepository(t)
		cr := mocks.NewContractRepository(t)
		pr := mocks.NewPartnerRepository(t)
		as := mocks.NewAddressService(t)
		sr := mocks.NewSchoolRepository(t)

		usecase := NewCalculatePriceDriversUseCase(
			&persistence.PostgresRepositories{
				PartnerRepository:     pr,
				ContractRepository:    cr,
				ResponsibleRepository: rr,
				SchoolRepository:      sr,
			},
			logger,
			adapters.Adapters{
				AddressService: as,
			},
		)

		// Configurar o mock para erro no repositório de responsáveis
		rr.On("Get", responsible.CPF).Return(responsible, nil)
		sr.On("Get", school.CNPJ).Return(school, errors.New("school get error"))
		logger.On("Infof", mock.Anything, mock.Anything).Return(logger)

		// Executar o método
		_, err := usecase.CalculatePrice(responsible.CPF, school.CNPJ)

		assert.EqualError(t, err, "school get error")
	})

	t.Run("if list Driver returns error", func(t *testing.T) {
		logger := mocks.NewLogger(t)
		rr := mocks.NewResponsibleRepository(t)
		cr := mocks.NewContractRepository(t)
		pr := mocks.NewPartnerRepository(t)
		as := mocks.NewAddressService(t)
		sr := mocks.NewSchoolRepository(t)

		usecase := NewCalculatePriceDriversUseCase(
			&persistence.PostgresRepositories{
				PartnerRepository:     pr,
				ContractRepository:    cr,
				ResponsibleRepository: rr,
				SchoolRepository:      sr,
			},
			logger,
			adapters.Adapters{
				AddressService: as,
			},
		)
		expectedDistance := 10.5
		rr.On("Get", responsible.CPF).Return(responsible, nil)
		sr.On("Get", school.CNPJ).Return(school, nil)
		as.On("GetDistance", responsibleAddress, schoolAddress).Return(&expectedDistance, nil)
		pr.On("GetBySchool", school.CNPJ).Return([]entity.Partner{}, errors.New("database error"))
		logger.On("Infof", mock.Anything, mock.Anything).Return(logger)

		_, err := usecase.CalculatePrice(responsible.CPF, school.CNPJ)

		assert.EqualError(t, err, "database error")

	})

	t.Run("if calculate price drivers usecase returns sucess", func(t *testing.T) {
		logger := mocks.NewLogger(t)
		rr := mocks.NewResponsibleRepository(t)
		cr := mocks.NewContractRepository(t)
		pr := mocks.NewPartnerRepository(t)
		as := mocks.NewAddressService(t)
		sr := mocks.NewSchoolRepository(t)

		usecase := NewCalculatePriceDriversUseCase(
			&persistence.PostgresRepositories{
				PartnerRepository:     pr,
				ContractRepository:    cr,
				ResponsibleRepository: rr,
				SchoolRepository:      sr,
			},
			logger,
			adapters.Adapters{
				AddressService: as,
			},
		)
		expectedDistance := 10.5
		// Configurar o mock para erro no repositório de responsáveis
		rr.On("Get", responsible.CPF).Return(responsible, nil)
		sr.On("Get", school.CNPJ).Return(school, nil)
		as.On("GetDistance", responsibleAddress, schoolAddress).Return(&expectedDistance, nil)
		pr.On("GetBySchool", school.CNPJ).Return([]entity.Partner{}, nil)
		// Executar o método
		result, err := usecase.CalculatePrice(responsible.CPF, school.CNPJ)

		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
}
