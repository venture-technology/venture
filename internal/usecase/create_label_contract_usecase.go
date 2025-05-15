package usecase

import (
	"encoding/json"
	"path/filepath"

	"github.com/spf13/viper"
	"github.com/venture-technology/venture/internal/domain/service/adapters"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/value"
	"github.com/venture-technology/venture/pkg/stringcommon"
	"github.com/venture-technology/venture/pkg/utils"
)

type CreateLabelContractUsecase struct {
	logger    contracts.Logger
	adapters  adapters.Adapters
	s3        contracts.S3Iface
	queue     contracts.Queue
	converter contracts.Converters
}

func NewCreateLabelContractUsecase(
	logger contracts.Logger,
	adapters adapters.Adapters,
	s3 contracts.S3Iface,
	queue contracts.Queue,
	converter contracts.Converters,
) *CreateLabelContractUsecase {
	return &CreateLabelContractUsecase{
		logger:    logger,
		adapters:  adapters,
		s3:        s3,
		queue:     queue,
		converter: converter,
	}
}

func (clc *CreateLabelContractUsecase) Execute(requestParams string) error {
	var input value.CreateContractParams
	err := json.Unmarshal([]byte(requestParams), &input)
	if err != nil {
		return err
	}

	file, err := clc.buildContract(input)
	if err != nil {
		return err
	}

	input.FileURL, err = clc.s3.Save(
		value.GetBucketContract(),
		"contracts",
		input.UUID,
		file,
		value.PDF,
	)
	if err != nil {
		return err
	}

	message, err := stringcommon.RawMessage(input)
	if err != nil {
		return err
	}

	err = clc.queue.SendMessage(viper.GetString("CREATE_CONTRACT_QUEUE"), message)
	if err != nil {
		return err
	}

	clc.logger.Infof(message)
	return nil
}

func (clc *CreateLabelContractUsecase) buildContract(createContractInput value.CreateContractParams) ([]byte, error) {
	path, err := GetHtml()
	if err != nil {
		return nil, err
	}

	html, err := clc.adapters.AgreementService.BuildContract(path)
	if err != nil {
		return nil, err
	}

	pdf, err := clc.converter.ConvertHTMLtoPDF(html, createContractInput)
	if err != nil {
		return nil, err
	}

	return pdf, nil
}

func GetHtml() (string, error) {
	path, err := utils.GetAbsPath()
	if err != nil {
		return "", err
	}

	return filepath.Join(path, PathAgreement()), nil
}

func PathAgreement() string {
	return "../../internal/domain/service/agreements/template/agreement.html"
}
