package usecase

import (
	"encoding/json"
	"path/filepath"

	"github.com/spf13/viper"
	"github.com/venture-technology/venture/internal/domain/service/adapters"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
	"github.com/venture-technology/venture/pkg/utils"
)

type CreateLabelContractUsecase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
	adapters     adapters.Adapters
	s3           contracts.S3Iface
	queue        contracts.Queue
	converter    contracts.Converters
}

func NewCreateLabelContractUsecase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
	adapters adapters.Adapters,
	s3 contracts.S3Iface,
	queue contracts.Queue,
	converter contracts.Converters,
) *CreateLabelContractUsecase {
	return &CreateLabelContractUsecase{
		repositories: repositories,
		logger:       logger,
		adapters:     adapters,
		s3:           s3,
		queue:        queue,
		converter:    converter,
	}
}

func (clc *CreateLabelContractUsecase) Execute(requestParams string) error {
	var input value.SQSCreateLabelContractParams
	err := json.Unmarshal([]byte(requestParams), &input)
	if err != nil {
		return err
	}

	file, err := clc.buildContract(input)
	if err != nil {
		return err
	}

	url, err := clc.s3.Save(
		value.GetBucketContract(),
		"contracts",
		input.UUID,
		file,
		value.PDF,
	)

	message, err := rawMessage(
		value.SQSCreateContractParams{
			AmountCents:      input.AmountCents,
			AmountAnualCents: input.AmountAnualCents,
			UUID:             input.UUID,
			ResponsibleID:    input.ResponsibleCPF,
			KidID:            input.KidRG,
			DriverID:         input.DriverCPF,
			SchoolID:         input.SchoolCNPJ,
			FileURL:          url,
		},
	)

	err = clc.queue.SendMessage(viper.GetString("CREATE_CONTRACT_QUEUE"), message)
	if err != nil {
		return err
	}

	return nil
}

func (clc *CreateLabelContractUsecase) buildContract(createContractInput value.SQSCreateLabelContractParams) ([]byte, error) {
	path, err := getHtml()
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

func rawMessage(input value.SQSCreateContractParams) (string, error) {
	raw, err := json.Marshal(input)
	if err != nil {
		return "", err
	}

	return string(raw), nil
}

func getHtml() (string, error) {
	path, err := utils.GetAbsPath()
	if err != nil {
		return "", err
	}

	return filepath.Join(path, pathAgreement()), nil
}

func pathAgreement() string {
	return "../domain/service/agreements/template/agreement.html"
}
