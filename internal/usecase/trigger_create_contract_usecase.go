package usecase

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/viper"
	"github.com/venture-technology/venture/internal/domain/service/adapters"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
	"github.com/venture-technology/venture/pkg/realtime"
	"github.com/venture-technology/venture/pkg/stringcommon"
	"github.com/venture-technology/venture/pkg/utils"
)

type TriggerContractUsecase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
	adapters     adapters.Adapters
	S3           contracts.S3Iface
	queue        contracts.Queue
	converter    contracts.Converters
}

func NewTriggerContractUsecase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
	adapters adapters.Adapters,
	S3 contracts.S3Iface,
	queue contracts.Queue,
	converter contracts.Converters,
) *TriggerContractUsecase {
	return &TriggerContractUsecase{
		repositories: repositories,
		logger:       logger,
		adapters:     adapters,
		S3:           S3,
		queue:        queue,
		converter:    converter,
	}
}

func (tcu *TriggerContractUsecase) TriggerExecute(requestParams *value.CreateContractRequestParams) error {
	err := tcu.validate(requestParams)
	if err != nil {
		return err
	}

	input, err := tcu.parseInfo(*requestParams)
	if err != nil {
		return err
	}

	input, err = tcu.CalcuateAmount(input)
	if err != nil {
		return err
	}

	message, err := stringcommon.RawMessage(input)
	if err != nil {
		return err
	}

	tcu.logger.Infof(message)
	err = tcu.queue.SendMessage(viper.GetString("QUEUE_CREATE_LABEL_CONTRACT"), message)
	if err != nil {
		return err
	}

	return nil
}

func (tcu *TriggerContractUsecase) validate(requestParams *value.CreateContractRequestParams) error {
	input := entity.TempContract{
		DriverCNH:      requestParams.DriverCNH,
		KidRG:          requestParams.KidRG,
		SchoolCNPJ:     requestParams.SchoolCNPJ,
		ResponsibleCPF: requestParams.ResponsibleCPF,
	}

	alreadyExists, err := tcu.repositories.TempContractRepository.HasTemporaryContract(&input)
	if err != nil {
		return err
	}

	if alreadyExists {
		return fmt.Errorf("contract already exists")
	}

	hasParent, err := tcu.repositories.KidRepository.FindByResponsible(&input.ResponsibleCPF)
	if err != nil {
		return err
	}

	if !hasParent {
		return fmt.Errorf("kid and responsible arent parents")
	}

	return nil
}

func (tcu *TriggerContractUsecase) CalcuateAmount(input *value.CreateContractParams) (*value.CreateContractParams, error) {
	distance, err := tcu.adapters.AddressService.GetDistance(input.ResponsibleAddr, input.SchoolAddr)
	if err != nil {
		return nil, err
	}

	amount := utils.CalculateContract(*distance, input.DriverAmount)

	input.AmountCents = amount
	input.AmountAnualCents = amount * 12

	return input, nil
}

func (tcu *TriggerContractUsecase) parseInfo(requestParams value.CreateContractRequestParams) (*value.CreateContractParams, error) {
	responsible, err := tcu.repositories.ResponsibleRepository.Get(requestParams.ResponsibleCPF)
	if err != nil {
		return nil, err
	}

	kid, err := tcu.repositories.KidRepository.Get(&requestParams.KidRG)
	if err != nil {
		return nil, err
	}

	driver, err := tcu.repositories.DriverRepository.Get(requestParams.DriverCNH)
	if err != nil {
		return nil, err
	}

	school, err := tcu.repositories.SchoolRepository.Get(requestParams.SchoolCNPJ)
	if err != nil {
		return nil, err
	}

	time := realtime.Now()

	return &value.CreateContractParams{
		UUID:             uuid.NewString(),
		ResponsibleCPF:   responsible.CPF,
		ResponsibleName:  responsible.Name,
		ResponsibleAddr:  responsible.Address.GetFullAddress(),
		ResponsibleEmail: responsible.Email,
		ResponsiblePhone: responsible.Phone,
		KidRG:            kid.RG,
		KidShift:         kid.Shift,
		KidName:          kid.Name,
		DriverCNH:        driver.CNH,
		DriverName:       driver.Name,
		DriverAmount:     driver.Amount,
		SchoolCNPJ:       school.CNPJ,
		SchoolName:       school.Name,
		SchoolAddr:       school.Address.GetFullAddress(),
		Time:             time,
		DateTime:         time.Format("02/01/2006"),
	}, nil
}
