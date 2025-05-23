package workers

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"
	"github.com/venture-technology/venture/internal/domain/service/adapters"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/value"
	"github.com/venture-technology/venture/pkg/utils"
)

type WorkerQueue struct {
	ch         chan *value.CreateContractParams
	logger     contracts.Logger
	adapters   adapters.Adapters
	bucket     contracts.S3Iface
	converters contracts.Converters
	email      contracts.WorkerEmail
}

func NewWorkerCreateContract(
	buffer int,
	logger contracts.Logger,
	bucket contracts.S3Iface,
	adapters adapters.Adapters,
	converters contracts.Converters,
	email contracts.WorkerEmail,
) contracts.WorkerCreateContract {
	queue := &WorkerQueue{
		ch:         make(chan *value.CreateContractParams, buffer),
		bucket:     bucket,
		logger:     logger,
		adapters:   adapters,
		converters: converters,
		email:      email,
	}

	go queue.worker()
	return queue
}

func (w *WorkerQueue) Enqueue(payload *value.CreateContractParams) error {
	w.ch <- payload
	return nil
}

func (w *WorkerQueue) worker() {
	for payload := range w.ch {
		if payload == nil {
			w.logger.Infof("payload is nil")
			w.notify(fmt.Errorf("payload nil"), payload)
			continue
		}

		w.logger.Infof("worker - calculating amount")
		payload, err := w.calcuateAmount(payload)
		if err != nil {
			w.logger.Infof(fmt.Sprintf("error calculating amount: %v", err))
			w.notify(err, payload)
			continue
		}

		w.logger.Infof("worker - parsing contract")
		file, err := w.parseContract(payload)
		if err != nil {
			w.logger.Infof(fmt.Sprintf("error parsing contract: %v", err))
			w.notify(err, payload)
			continue
		}

		w.logger.Infof("worker - saving file in bucket")
		payload.FileURL, err = w.bucket.Save(
			value.GetBucketContract(),
			"contracts",
			payload.UUID,
			file,
			value.PDF,
		)
		if err != nil {
			w.logger.Infof(fmt.Sprintf("error saving file: %v", err))
			w.notify(err, payload)
			continue
		}

		w.logger.Infof("worker - sending contract")
		resp, err := w.adapters.AgreementService.SignatureRequest(*payload)
		if err != nil {
			w.logger.Infof(fmt.Sprintf("error creating signature request: %v", err))
			w.notify(err, payload)
			continue
		}
		w.logger.Infof(
			fmt.Sprintf(
				"Driver: %s, Responsible: %s, Kid: %s, School: %s",
				resp.Metadata.Keys.DriverCNH,
				resp.Metadata.Keys.ResponsibleCPF,
				resp.Metadata.Keys.KidRG,
				resp.Metadata.Keys.SchoolCNPJ,
			),
		)
		w.notify(err, payload)
	}
}

func (w *WorkerQueue) calcuateAmount(
	params *value.CreateContractParams,
) (*value.CreateContractParams, error) {
	w.logger.Infof(fmt.Sprintf("%s, %s", params.ResponsibleAddr, params.SchoolAddr))
	distance, err := w.adapters.AddressService.GetDistance(
		params.ResponsibleAddr,
		params.SchoolAddr,
	)
	if err != nil {
		return nil, err
	}

	amount := utils.CalculateContract(*distance, params.DriverAmount)
	params.AmountCents = amount
	params.AmountAnualCents = amount * 12

	return params, nil
}

func (w *WorkerQueue) parseContract(
	params *value.CreateContractParams,
) ([]byte, error) {
	html, err := w.getHtmlParsed()
	if err != nil {
		return nil, err
	}

	pdf, err := w.converters.ConvertHTMLtoPDF(html, *params)
	if err != nil {
		return nil, err
	}

	return pdf, nil
}

func (w *WorkerQueue) notify(err error, payload *value.CreateContractParams) {
	if err != nil {
		w.email.Enqueue(&entity.Email{
			Recipient: payload.ResponsibleEmail,
			Subject:   "Tivemos problema com a geração do seu contrato. Sentimos muito!",
			Body: fmt.Sprintf(
				"Verifique se todas suas informações internas estão corretas ou entre em contrato com a nossa equipe: %s",
				viper.GetString("AWS_SES_EMAIL_FROM"),
			),
		})
	}
	w.email.Enqueue(&entity.Email{
		Recipient: payload.ResponsibleEmail,
		Subject:   "Seu contrato foi gerado com sucesso!",
		Body: fmt.Sprintf(
			"No link abaixo, está seu contrato na Dropbox, por favor assine para que seu filho começe a usufruir do transporte. \n %s",
			payload.FileURL,
		),
	})
}

func getPath() (string, error) {
	path, err := utils.GetAbsPath()
	if err != nil {
		return "", err
	}

	return filepath.Join(path, filePath()), nil
}

func filePath() string {
	return "../../internal/domain/service/agreements/template/agreement.html"
}

func (w *WorkerQueue) getHtmlParsed() ([]byte, error) {
	path, err := getPath()
	if err != nil {
		return nil, err
	}

	content, err := w.adapters.AgreementService.BuildContract(path)
	if err != nil {
		return nil, err
	}

	return content, nil
}
