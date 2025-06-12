package infra

import (
	"github.com/venture-technology/venture/internal/domain/service/address"
	"github.com/venture-technology/venture/internal/domain/service/payments"
	"github.com/venture-technology/venture/internal/domain/service/signatures"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type Application struct {
	Repositories         persistence.PostgresRepositories
	Postgres             contracts.PostgresIface
	Cache                contracts.Cacher
	Bucket               contracts.S3Iface
	Email                contracts.SESIface
	Logger               contracts.Logger
	Address              address.Address
	Payments             payments.Payments
	Signature            signatures.Signature
	Converters           contracts.Converters
	Queue                contracts.Queue
	WorkerCreateContract contracts.WorkerCreateContract
	WorkerEmail          contracts.WorkerEmail
}

var App Application
