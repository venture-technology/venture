package workers

import (
	"fmt"

	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
)

type emailWorker struct {
	ch     chan *entity.Email
	ses    contracts.SESIface
	logger contracts.Logger
}

func NewWorkerEmail(
	buffer int,
	ses contracts.SESIface,
	logger contracts.Logger,
) *emailWorker {
	queue := &emailWorker{
		ch:     make(chan *entity.Email, buffer),
		ses:    ses,
		logger: logger,
	}

	go queue.worker()
	return queue
}

func (w *emailWorker) Enqueue(payload *entity.Email) error {
	w.ch <- payload
	return nil
}

func (w *emailWorker) worker() {
	for payload := range w.ch {
		if payload == nil {
			continue
		}
		w.logger.Infof(
			fmt.Sprintf(
				"sending a new email to: %s with subject: %s", payload.Recipient, payload.Subject,
			),
		)
		w.ses.SendEmail(payload)
	}
}
