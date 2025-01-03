package setup

import (
	"log"

	"github.com/venture-technology/venture/config"
	"github.com/venture-technology/venture/internal/infra"
	"github.com/venture-technology/venture/internal/infra/bucket"
	"github.com/venture-technology/venture/internal/infra/database"
	"github.com/venture-technology/venture/internal/infra/email"
	"github.com/venture-technology/venture/internal/infra/logger"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type Setup struct {
	app          *infra.Application
	repositories *persistence.PostgresRepositories
}

func NewSetup() Setup {
	Config, err := config.Load("config/config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	return Setup{
		app: &infra.Application{
			Config: *Config,
		},
		repositories: &persistence.PostgresRepositories{},
	}
}

func (s Setup) Postgres() {
	s.app.Postgres, _ = database.NewPGGORMImpl(s.app.Config)
}

func (s Setup) Repositories() {
	s.repositories.ChildRepository = persistence.ChildRepositoryImpl{Postgres: s.app.Postgres}
	s.repositories.ContractRepository = persistence.ContractRepositoryImpl{Postgres: s.app.Postgres}
	s.repositories.DriverRepository = persistence.DriverRepositoryImpl{Postgres: s.app.Postgres}
	s.repositories.InviteRepository = persistence.InviteRepositoryImpl{Postgres: s.app.Postgres}
	s.repositories.PartnerRepository = persistence.PartnerRepositoryImpl{Postgres: s.app.Postgres}
	s.repositories.ResponsibleRepository = persistence.ResponsibleRepositoryImpl{Postgres: s.app.Postgres}
	s.repositories.SchoolRepository = persistence.SchoolRepositoryImpl{Postgres: s.app.Postgres}
}

func (s Setup) Redis() {
	s.app.Redis = database.NewRedisImpl(s.app.Config)
}

func (s Setup) RedisRepositories() {
	s.app.RedisRepositories.AdminRepository = persistence.AdminRepositoryImpl{Redis: s.app.Redis}
}

func (s Setup) Bucket() {
	s.app.Bucket = bucket.NewS3Impl(s.app.Config)
}

func (s Setup) Email() {
	s.app.Email = email.NewSesImpl(s.app.Config)
}

func (s Setup) Logger(taskname string) {
	s.app.Logger, _ = logger.New(taskname)
}
