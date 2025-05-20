package workers

import "github.com/venture-technology/venture/internal/entity"

type emailWorker struct {
	ch chan *entity.Email
}

func NewWorkerSendEmail(
	buffer int,
) *emailWorker {
	queue := &emailWorker{
		ch: make(chan *entity.Email, buffer),
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
	}
}
