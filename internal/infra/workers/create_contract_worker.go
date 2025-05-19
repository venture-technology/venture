package workers

import (
	"fmt"
	"path/filepath"

	"github.com/venture-technology/venture/internal/domain/service/adapters"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/value"
	"github.com/venture-technology/venture/pkg/utils"
)

type workerQueue struct {
	ch           chan *value.CreateContractParams
	logger       contracts.Logger
	adapters     adapters.Adapters
	bucket       contracts.S3Iface
	converters   contracts.Converters
	repositories *persistence.PostgresRepositories
}

func NewWorkerCreateLabel(
	buffer int,
	logger contracts.Logger,
	bucket contracts.S3Iface,
	adapters adapters.Adapters,
	converters contracts.Converters,
	repositories *persistence.PostgresRepositories,
) contracts.Workers {
	queue := &workerQueue{
		ch:           make(chan *value.CreateContractParams, buffer),
		bucket:       bucket,
		logger:       logger,
		adapters:     adapters,
		converters:   converters,
		repositories: repositories,
	}

	go queue.worker()
	return queue
}

func (w *workerQueue) Enqueue(payload *value.CreateContractParams) error {
	w.ch <- payload
	return nil
}

func (w *workerQueue) worker() {
	for payload := range w.ch {
		if payload == nil {
			w.logger.Infof("payload is nil")
			continue
		}

		w.logger.Infof("worker - calculating amount")
		payload, err := w.calcuateAmount(payload)
		if err != nil {
			w.logger.Infof(fmt.Sprintf("error calculating amount: %v", err))
			continue
		}

		w.logger.Infof("worker - parsing contract")
		file, err := w.parseContract(payload)
		if err != nil {
			w.logger.Infof(fmt.Sprintf("error parsing contract: %v", err))
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
			continue
		}

		w.logger.Infof("worker - sending contract")
		resp, err := w.adapters.AgreementService.SignatureRequest(*payload)
		if err != nil {
			w.logger.Infof(fmt.Sprintf("error creating signature request: %v", err))
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
	}
}

func (ccuc *workerQueue) calcuateAmount(
	params *value.CreateContractParams,
) (*value.CreateContractParams, error) {
	distance, err := ccuc.adapters.AddressService.GetDistance(
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

func (ccuc *workerQueue) parseContract(
	params *value.CreateContractParams,
) ([]byte, error) {
	html, err := ccuc.getHtmlParsed()
	if err != nil {
		return nil, err
	}

	pdf, err := ccuc.converters.ConvertHTMLtoPDF(html, *params)
	if err != nil {
		return nil, err
	}

	return pdf, nil
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

func (ccuc *workerQueue) getHtmlParsed() ([]byte, error) {
	path, err := getPath()
	if err != nil {
		return nil, err
	}

	content, err := ccuc.adapters.AgreementService.BuildContract(path)
	if err != nil {
		return nil, err
	}

	return content, nil
}
