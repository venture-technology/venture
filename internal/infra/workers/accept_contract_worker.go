package workers

type workerQueueAcceptContract struct {
	ch chan *string
}

func NewWorkerAcceptContract(
	buffer int,
) *workerQueueAcceptContract {
	queue := &workerQueueAcceptContract{
		ch: make(chan *string, buffer),
	}

	go queue.worker()
	return queue
}

func (w *workerQueueAcceptContract) Enqueue(payload *string) error {
	w.ch <- payload
	return nil
}

func (w *workerQueueAcceptContract) worker() {
	for payload := range w.ch {
		if payload == nil {
			continue
		}
	}
}
