package infra

import (
	"github.com/venture-technology/venture/internal/domain/service/adapters"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type Application struct {
	Repositories persistence.PostgresRepositories
	Postgres     contracts.PostgresIface
	Cache        contracts.Cacher
	Bucket       contracts.S3Iface
	Email        contracts.SESIface
	Logger       contracts.Logger
	Adapters     adapters.Adapters
	Converters   contracts.Converters
	Queue        contracts.Queue
	Workers      contracts.Workers
}

var App Application
