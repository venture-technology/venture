package contracts

import (
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/value"
)

type WorkerCreateContract interface {
	Enqueue(payload *value.CreateContractParams) error
}

type WorkerAcceptContract interface {
	Enqueue(payload *string) error
}

type WorkerSendEmail interface {
	Enqueue(payload *entity.Email) error
}
