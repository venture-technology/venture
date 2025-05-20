package setup

import (
	"log"
	"os"

	"github.com/venture-technology/venture/config"
	"github.com/venture-technology/venture/internal/domain/service/addresses"
	"github.com/venture-technology/venture/internal/domain/service/agreements"
	"github.com/venture-technology/venture/internal/domain/service/converters"
	"github.com/venture-technology/venture/internal/domain/service/decorator"
	"github.com/venture-technology/venture/internal/domain/service/payments"
	"github.com/venture-technology/venture/internal/infra"
	"github.com/venture-technology/venture/internal/infra/bucket"
	"github.com/venture-technology/venture/internal/infra/cache"
	"github.com/venture-technology/venture/internal/infra/database"
	"github.com/venture-technology/venture/internal/infra/email"
	"github.com/venture-technology/venture/internal/infra/logger"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/internal/infra/queue"
	"github.com/venture-technology/venture/internal/infra/workers"
)

const (
	ServiceName = "venture"
)

type Setup struct {
	app          *infra.Application
	repositories *persistence.PostgresRepositories
}

func NewSetup() Setup {
	err := config.LoadServerEnvironmentVars(ServiceName, os.Getenv(config.ServerEnvironment))
	if err != nil {
		log.Fatal(err)
	}

	return Setup{
		app:          &infra.Application{},
		repositories: &persistence.PostgresRepositories{},
	}
}

func (s Setup) Postgres() {
	s.app.Postgres, _ = database.NewPGGORMImpl()
}

func (s Setup) Repositories() {
	s.app.Repositories.KidRepository = persistence.KidRepositoryImpl{Postgres: s.app.Postgres}
	s.app.Repositories.ContractRepository = persistence.ContractRepositoryImpl{Postgres: s.app.Postgres}
	s.app.Repositories.DriverRepository = persistence.DriverRepositoryImpl{Postgres: s.app.Postgres}
	s.app.Repositories.InviteRepository = persistence.InviteRepositoryImpl{Postgres: s.app.Postgres}
	s.app.Repositories.PartnerRepository = persistence.PartnerRepositoryImpl{Postgres: s.app.Postgres}
	s.app.Repositories.ResponsibleRepository = persistence.ResponsibleRepositoryImpl{Postgres: s.app.Postgres}
	s.app.Repositories.SchoolRepository = persistence.SchoolRepositoryImpl{Postgres: s.app.Postgres}
	s.app.Repositories.TempContractRepository = persistence.TempContractRepositoryImpl{Postgres: s.app.Postgres}
}

// Cache need started before SQL Database.
//
// Because, how we used cache like decorator with Repository, Repository cant receive a null instance of cache
func (s Setup) Cache() {
	s.app.Cache = cache.NewCacheImpl()
}

func (s Setup) Bucket() {
	s.app.Bucket = bucket.NewS3Impl()
}

func (s Setup) Email() {
	s.app.Email = email.NewSesImpl()
}

func (s Setup) Logger(taskname string) {
	s.app.Logger, _ = logger.New(taskname)
}

func (s Setup) Finish() {
	infra.App = *s.app
}

func (s Setup) Adapters() {
	s.app.Adapters.AddressService = decorator.AddressDecorator{
		AddressAdapter: addresses.NewGoogleAdapter(),
		Cache:          s.app.Cache,
	}
	s.app.Adapters.PaymentsService = payments.NewStripeAdapter()
	s.app.Adapters.AgreementService = agreements.NewAgreementService(s.app.Logger, &s.app.Repositories)
}

func (s Setup) Converters() {
	s.app.Converters = converters.NewConverter()
}

func (s Setup) Queue() {
	s.app.Queue = queue.NewSQSQueue()
}

// workers must be started after logger, database and adapters.
func (s Setup) WorkerCreateContract() {
	s.app.WorkerCreateContract = workers.NewWorkerCreateLabel(
		100,
		s.app.Logger,
		s.app.Bucket,
		s.app.Adapters,
		s.app.Converters,
		&s.app.Repositories,
	)
}

func (s Setup) WorkerAcceptContract() {
}

func (s Setup) WorkerSendEmail() {
}
