package infra

import (
	"github.com/venture-technology/venture/config"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type Application struct {
	Repositories      persistence.PostgresRepositories
	RedisRepositories persistence.RedisRepositories
	Postgres          contracts.PostgresIface
	Redis             contracts.RedisIface
	Bucket            contracts.S3Iface
	Email             contracts.SESIface
	Logger            contracts.Logger
	Config            config.Config
}

var App Application
