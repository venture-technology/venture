package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/venture-technology/venture/internal/domain/service/adapters"
	"github.com/venture-technology/venture/internal/domain/service/agreements"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
	"github.com/venture-technology/venture/mocks"
)

func TestCreateContractUsecase_CreateContract(t *testing.T) {
	var params value.CreateContractRequestParams

	t.Run("repo temp contract get by everyone return error", func(t *testing.T) {
		tcr := mocks.NewTempContractRepository(t)
		dr := mocks.NewDriverRepository(t)
		sc := mocks.NewSchoolRepository(t)
		kr := mocks.NewKidRepository(t)
		rr := mocks.NewResponsibleRepository(t)
		das := mocks.NewAddressService(t)
		s3 := mocks.NewS3Iface(t)
		dbs := mocks.NewAgreementService(t)
		logger := mocks.NewLogger(t)
		converter := mocks.NewConverters(t)

		tcr.On("GetByEveryone", mock.Anything).Return(false, errors.New("database error"))
		logger.On("Infof", mock.Anything, mock.Anything).Return(logger)

		uc := NewCreateContractUseCase(
			&persistence.PostgresRepositories{
				TempContractRepository: tcr,
				DriverRepository:       dr,
				SchoolRepository:       sc,
				KidRepository:          kr,
				ResponsibleRepository:  rr,
			},
			logger,
			adapters.Adapters{
				AddressService:   das,
				AgreementService: dbs,
			},
			s3,
			converter,
		)

		_, err := uc.CreateContract(&params)

		assert.EqualError(t, err, "database error")
		assert.Error(t, err)
	})

	t.Run("repo temp contract get by everyone return error", func(t *testing.T) {
		tcr := mocks.NewTempContractRepository(t)
		dr := mocks.NewDriverRepository(t)
		sc := mocks.NewSchoolRepository(t)
		kr := mocks.NewKidRepository(t)
		rr := mocks.NewResponsibleRepository(t)
		das := mocks.NewAddressService(t)
		s3 := mocks.NewS3Iface(t)
		dbs := mocks.NewAgreementService(t)
		logger := mocks.NewLogger(t)
		converter := mocks.NewConverters(t)

		tcr.On("GetByEveryone", mock.Anything).Return(true, nil)
		logger.On("Infof", mock.Anything, mock.Anything).Return(logger)

		uc := NewCreateContractUseCase(
			&persistence.PostgresRepositories{
				TempContractRepository: tcr,
				DriverRepository:       dr,
				SchoolRepository:       sc,
				KidRepository:          kr,
				ResponsibleRepository:  rr,
			},
			logger,
			adapters.Adapters{
				AddressService:   das,
				AgreementService: dbs,
			},
			s3,
			converter,
		)

		_, err := uc.CreateContract(&params)

		assert.EqualError(t, err, "temp contract already exists")
		assert.Error(t, err)
	})

	t.Run("get agreement html file return error", func(t *testing.T) {
		tcr := mocks.NewTempContractRepository(t)
		dr := mocks.NewDriverRepository(t)
		sc := mocks.NewSchoolRepository(t)
		kr := mocks.NewKidRepository(t)
		rr := mocks.NewResponsibleRepository(t)
		das := mocks.NewAddressService(t)
		s3 := mocks.NewS3Iface(t)
		dbs := mocks.NewAgreementService(t)
		logger := mocks.NewLogger(t)
		converter := mocks.NewConverters(t)

		tcr.On("GetByEveryone", mock.Anything).Return(false, nil)
		logger.On("Infof", mock.Anything, mock.Anything).Return(logger)
		dbs.On("GetAgreementHtml", mock.Anything).Return([]byte{}, errors.New("agreement service error"))

		uc := NewCreateContractUseCase(
			&persistence.PostgresRepositories{
				TempContractRepository: tcr,
				DriverRepository:       dr,
				SchoolRepository:       sc,
				KidRepository:          kr,
				ResponsibleRepository:  rr,
			},
			logger,
			adapters.Adapters{
				AddressService:   das,
				AgreementService: dbs,
			},
			s3,
			converter,
		)

		_, err := uc.CreateContract(&params)

		assert.EqualError(t, err, "agreement service error")
		assert.Error(t, err)
	})

	t.Run("get driver repository return error", func(t *testing.T) {
		tcr := mocks.NewTempContractRepository(t)
		dr := mocks.NewDriverRepository(t)
		sc := mocks.NewSchoolRepository(t)
		kr := mocks.NewKidRepository(t)
		rr := mocks.NewResponsibleRepository(t)
		das := mocks.NewAddressService(t)
		s3 := mocks.NewS3Iface(t)
		dbs := mocks.NewAgreementService(t)
		logger := mocks.NewLogger(t)
		converter := mocks.NewConverters(t)

		tcr.On("GetByEveryone", mock.Anything).Return(false, nil)
		dbs.On("GetAgreementHtml", mock.Anything).Return([]byte{}, nil)
		dr.On("Get", mock.Anything).Return(&entity.Driver{}, errors.New("get driver repository error"))
		logger.On("Infof", mock.Anything, mock.Anything).Return(logger)

		uc := NewCreateContractUseCase(
			&persistence.PostgresRepositories{
				TempContractRepository: tcr,
				DriverRepository:       dr,
				SchoolRepository:       sc,
				KidRepository:          kr,
				ResponsibleRepository:  rr,
			},
			logger,
			adapters.Adapters{
				AddressService:   das,
				AgreementService: dbs,
			},
			s3,
			converter,
		)

		_, err := uc.CreateContract(&params)

		assert.EqualError(t, err, "get driver repository error")
		assert.Error(t, err)
	})

	t.Run("get school repository return error", func(t *testing.T) {
		tcr := mocks.NewTempContractRepository(t)
		dr := mocks.NewDriverRepository(t)
		sc := mocks.NewSchoolRepository(t)
		kr := mocks.NewKidRepository(t)
		rr := mocks.NewResponsibleRepository(t)
		das := mocks.NewAddressService(t)
		s3 := mocks.NewS3Iface(t)
		dbs := mocks.NewAgreementService(t)
		logger := mocks.NewLogger(t)
		converter := mocks.NewConverters(t)

		tcr.On("GetByEveryone", mock.Anything).Return(false, nil)
		dbs.On("GetAgreementHtml", mock.Anything).Return([]byte{}, nil)
		dr.On("Get", mock.Anything).Return(&entity.Driver{}, nil)
		sc.On("Get", mock.Anything).Return(&entity.School{}, errors.New("get school repository error"))
		logger.On("Infof", mock.Anything, mock.Anything).Return(logger)

		uc := NewCreateContractUseCase(
			&persistence.PostgresRepositories{
				TempContractRepository: tcr,
				DriverRepository:       dr,
				SchoolRepository:       sc,
				KidRepository:          kr,
				ResponsibleRepository:  rr,
			},
			logger,
			adapters.Adapters{
				AddressService:   das,
				AgreementService: dbs,
			},
			s3,
			converter,
		)

		_, err := uc.CreateContract(&params)

		assert.EqualError(t, err, "get school repository error")
		assert.Error(t, err)
	})

	t.Run("get kid repository return error", func(t *testing.T) {
		tcr := mocks.NewTempContractRepository(t)
		dr := mocks.NewDriverRepository(t)
		sc := mocks.NewSchoolRepository(t)
		kr := mocks.NewKidRepository(t)
		rr := mocks.NewResponsibleRepository(t)
		das := mocks.NewAddressService(t)
		s3 := mocks.NewS3Iface(t)
		dbs := mocks.NewAgreementService(t)
		logger := mocks.NewLogger(t)
		converter := mocks.NewConverters(t)

		tcr.On("GetByEveryone", mock.Anything).Return(false, nil)
		dbs.On("GetAgreementHtml", mock.Anything).Return([]byte{}, nil)
		dr.On("Get", mock.Anything).Return(&entity.Driver{}, nil)
		sc.On("Get", mock.Anything).Return(&entity.School{}, nil)
		kr.On("Get", mock.Anything).Return(&entity.Kid{}, errors.New("get kid repository error"))
		logger.On("Infof", mock.Anything, mock.Anything).Return(logger)

		uc := NewCreateContractUseCase(
			&persistence.PostgresRepositories{
				TempContractRepository: tcr,
				DriverRepository:       dr,
				SchoolRepository:       sc,
				KidRepository:          kr,
				ResponsibleRepository:  rr,
			},
			logger,
			adapters.Adapters{
				AddressService:   das,
				AgreementService: dbs,
			},
			s3,
			converter,
		)

		_, err := uc.CreateContract(&params)

		assert.EqualError(t, err, "get kid repository error")
		assert.Error(t, err)
	})

	t.Run("get responsible repository return error", func(t *testing.T) {
		tcr := mocks.NewTempContractRepository(t)
		dr := mocks.NewDriverRepository(t)
		sc := mocks.NewSchoolRepository(t)
		kr := mocks.NewKidRepository(t)
		rr := mocks.NewResponsibleRepository(t)
		das := mocks.NewAddressService(t)
		s3 := mocks.NewS3Iface(t)
		dbs := mocks.NewAgreementService(t)
		logger := mocks.NewLogger(t)
		converter := mocks.NewConverters(t)

		tcr.On("GetByEveryone", mock.Anything).Return(false, nil)
		dbs.On("GetAgreementHtml", mock.Anything).Return([]byte{}, nil)
		dr.On("Get", mock.Anything).Return(&entity.Driver{}, nil)
		sc.On("Get", mock.Anything).Return(&entity.School{}, nil)
		kr.On("Get", mock.Anything).Return(&entity.Kid{}, nil)
		rr.On("Get", mock.Anything).Return(&entity.Responsible{}, errors.New("get responsible repository error"))
		logger.On("Infof", mock.Anything, mock.Anything).Return(logger)

		uc := NewCreateContractUseCase(
			&persistence.PostgresRepositories{
				TempContractRepository: tcr,
				DriverRepository:       dr,
				SchoolRepository:       sc,
				KidRepository:          kr,
				ResponsibleRepository:  rr,
			},
			logger,
			adapters.Adapters{
				AddressService:   das,
				AgreementService: dbs,
			},
			s3,
			converter,
		)

		_, err := uc.CreateContract(&params)

		assert.EqualError(t, err, "get responsible repository error")
		assert.Error(t, err)
	})

	t.Run("get distance address service return error", func(t *testing.T) {
		tcr := mocks.NewTempContractRepository(t)
		dr := mocks.NewDriverRepository(t)
		sc := mocks.NewSchoolRepository(t)
		kr := mocks.NewKidRepository(t)
		rr := mocks.NewResponsibleRepository(t)
		das := mocks.NewAddressService(t)
		s3 := mocks.NewS3Iface(t)
		dbs := mocks.NewAgreementService(t)
		logger := mocks.NewLogger(t)
		converter := mocks.NewConverters(t)

		var value float64

		tcr.On("GetByEveryone", mock.Anything).Return(false, nil)
		dbs.On("GetAgreementHtml", mock.Anything).Return([]byte{}, nil)
		dr.On("Get", mock.Anything).Return(&entity.Driver{}, nil)
		sc.On("Get", mock.Anything).Return(&entity.School{}, nil)
		kr.On("Get", mock.Anything).Return(&entity.Kid{}, nil)
		rr.On("Get", mock.Anything).Return(&entity.Responsible{}, nil)
		das.On("GetDistance", mock.Anything, mock.Anything).Return(&value, errors.New("get distance error"))
		logger.On("Infof", mock.Anything, mock.Anything).Return(logger)

		uc := NewCreateContractUseCase(
			&persistence.PostgresRepositories{
				TempContractRepository: tcr,
				DriverRepository:       dr,
				SchoolRepository:       sc,
				KidRepository:          kr,
				ResponsibleRepository:  rr,
			},
			logger,
			adapters.Adapters{
				AddressService:   das,
				AgreementService: dbs,
			},
			s3,
			converter,
		)

		_, err := uc.CreateContract(&params)

		assert.EqualError(t, err, "get distance error")
		assert.Error(t, err)
	})

	t.Run("convert pdf to html error", func(t *testing.T) {
		tcr := mocks.NewTempContractRepository(t)
		dr := mocks.NewDriverRepository(t)
		sc := mocks.NewSchoolRepository(t)
		kr := mocks.NewKidRepository(t)
		rr := mocks.NewResponsibleRepository(t)
		das := mocks.NewAddressService(t)
		s3 := mocks.NewS3Iface(t)
		dbs := mocks.NewAgreementService(t)
		logger := mocks.NewLogger(t)
		converter := mocks.NewConverters(t)

		var value float64

		tcr.On("GetByEveryone", mock.Anything).Return(false, nil)
		dbs.On("GetAgreementHtml", mock.Anything).Return([]byte{}, nil)
		dr.On("Get", mock.Anything).Return(&entity.Driver{}, nil)
		sc.On("Get", mock.Anything).Return(&entity.School{}, nil)
		kr.On("Get", mock.Anything).Return(&entity.Kid{}, nil)
		rr.On("Get", mock.Anything).Return(&entity.Responsible{}, nil)
		das.On("GetDistance", mock.Anything, mock.Anything).Return(&value, nil)
		converter.On("ConvertPDFtoHTML", mock.Anything, mock.Anything).Return([]byte{}, errors.New("pdf to html error"))
		logger.On("Infof", mock.Anything, mock.Anything).Return(logger)

		uc := NewCreateContractUseCase(
			&persistence.PostgresRepositories{
				TempContractRepository: tcr,
				DriverRepository:       dr,
				SchoolRepository:       sc,
				KidRepository:          kr,
				ResponsibleRepository:  rr,
			},
			logger,
			adapters.Adapters{
				AddressService:   das,
				AgreementService: dbs,
			},
			s3,
			converter,
		)

		_, err := uc.CreateContract(&params)

		assert.EqualError(t, err, "pdf to html error")
		assert.Error(t, err)
	})

	t.Run("save s3 with type return error", func(t *testing.T) {
		tcr := mocks.NewTempContractRepository(t)
		dr := mocks.NewDriverRepository(t)
		sc := mocks.NewSchoolRepository(t)
		kr := mocks.NewKidRepository(t)
		rr := mocks.NewResponsibleRepository(t)
		das := mocks.NewAddressService(t)
		s3 := mocks.NewS3Iface(t)
		dbs := mocks.NewAgreementService(t)
		logger := mocks.NewLogger(t)
		converter := mocks.NewConverters(t)

		var value float64

		tcr.On("GetByEveryone", mock.Anything).Return(false, nil)
		dbs.On("GetAgreementHtml", mock.Anything).Return([]byte{}, nil)
		dr.On("Get", mock.Anything).Return(&entity.Driver{}, nil)
		sc.On("Get", mock.Anything).Return(&entity.School{}, nil)
		kr.On("Get", mock.Anything).Return(&entity.Kid{}, nil)
		rr.On("Get", mock.Anything).Return(&entity.Responsible{}, nil)
		das.On("GetDistance", mock.Anything, mock.Anything).Return(&value, nil)
		converter.On("ConvertPDFtoHTML", mock.Anything, mock.Anything).Return([]byte{}, nil)
		s3.On("PDF").Return("pdf")
		s3.On("SaveWithType", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return("", errors.New("save s3 error"))

		uc := NewCreateContractUseCase(
			&persistence.PostgresRepositories{
				TempContractRepository: tcr,
				DriverRepository:       dr,
				SchoolRepository:       sc,
				KidRepository:          kr,
				ResponsibleRepository:  rr,
			},
			logger,
			adapters.Adapters{
				AddressService:   das,
				AgreementService: dbs,
			},
			s3,
			converter,
		)

		_, err := uc.CreateContract(&params)

		assert.EqualError(t, err, "save s3 error")
		assert.Error(t, err)
	})

	t.Run("signature request send to dropbox return error", func(t *testing.T) {
		tcr := mocks.NewTempContractRepository(t)
		dr := mocks.NewDriverRepository(t)
		sc := mocks.NewSchoolRepository(t)
		kr := mocks.NewKidRepository(t)
		rr := mocks.NewResponsibleRepository(t)
		das := mocks.NewAddressService(t)
		s3 := mocks.NewS3Iface(t)
		dbs := mocks.NewAgreementService(t)
		logger := mocks.NewLogger(t)
		converter := mocks.NewConverters(t)

		var value float64

		tcr.On("GetByEveryone", mock.Anything).Return(false, nil)
		dbs.On("GetAgreementHtml", mock.Anything).Return([]byte{}, nil)
		dr.On("Get", mock.Anything).Return(&entity.Driver{}, nil)
		sc.On("Get", mock.Anything).Return(&entity.School{}, nil)
		kr.On("Get", mock.Anything).Return(&entity.Kid{}, nil)
		rr.On("Get", mock.Anything).Return(&entity.Responsible{}, nil)
		das.On("GetDistance", mock.Anything, mock.Anything).Return(&value, nil)
		converter.On("ConvertPDFtoHTML", mock.Anything, mock.Anything).Return([]byte{}, nil)
		s3.On("PDF").Return("pdf")
		s3.On("SaveWithType", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return("", nil)
		dbs.On("SignatureRequest", mock.Anything).Return(agreements.ContractRequest{}, errors.New("signature request error"))
		logger.On("Infof", mock.Anything, mock.Anything).Return(logger)

		uc := NewCreateContractUseCase(
			&persistence.PostgresRepositories{
				TempContractRepository: tcr,
				DriverRepository:       dr,
				SchoolRepository:       sc,
				KidRepository:          kr,
				ResponsibleRepository:  rr,
			},
			logger,
			adapters.Adapters{
				AddressService:   das,
				AgreementService: dbs,
			},
			s3,
			converter,
		)

		_, err := uc.CreateContract(&params)

		assert.EqualError(t, err, "signature request error")
		assert.Error(t, err)
	})

	t.Run("create contract usecase return success", func(t *testing.T) {
		tcr := mocks.NewTempContractRepository(t)
		dr := mocks.NewDriverRepository(t)
		sc := mocks.NewSchoolRepository(t)
		kr := mocks.NewKidRepository(t)
		rr := mocks.NewResponsibleRepository(t)
		das := mocks.NewAddressService(t)
		s3 := mocks.NewS3Iface(t)
		dbs := mocks.NewAgreementService(t)
		logger := mocks.NewLogger(t)
		converter := mocks.NewConverters(t)

		var value float64

		tcr.On("GetByEveryone", mock.Anything).Return(false, nil)
		dbs.On("GetAgreementHtml", mock.Anything).Return([]byte{}, nil)
		dr.On("Get", mock.Anything).Return(&entity.Driver{}, nil)
		sc.On("Get", mock.Anything).Return(&entity.School{}, nil)
		kr.On("Get", mock.Anything).Return(&entity.Kid{}, nil)
		rr.On("Get", mock.Anything).Return(&entity.Responsible{}, nil)
		das.On("GetDistance", mock.Anything, mock.Anything).Return(&value, nil)
		converter.On("ConvertPDFtoHTML", mock.Anything, mock.Anything).Return([]byte{}, nil)
		s3.On("PDF").Return("pdf")
		s3.On("SaveWithType", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return("", nil)
		dbs.On("SignatureRequest", mock.Anything).Return(agreements.ContractRequest{}, nil)

		uc := NewCreateContractUseCase(
			&persistence.PostgresRepositories{
				TempContractRepository: tcr,
				DriverRepository:       dr,
				SchoolRepository:       sc,
				KidRepository:          kr,
				ResponsibleRepository:  rr,
			},
			logger,
			adapters.Adapters{
				AddressService:   das,
				AgreementService: dbs,
			},
			s3,
			converter,
		)

		_, err := uc.CreateContract(&params)

		assert.NoError(t, err)
		assert.Nil(t, err)
	})
}
