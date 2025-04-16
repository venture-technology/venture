package main

import (
	"github.com/venture-technology/venture/internal/infra"
	"github.com/venture-technology/venture/internal/setup"
	"github.com/venture-technology/venture/internal/usecase"
)

// This file is used to expire temporary contracts in the database.
// It's running everyday, twice a day.
func main() {
	setup := setup.NewSetup()
	setup.Logger("venture-task-expire-temp-contract")
	setup.Cache()
	setup.Postgres()
	setup.Repositories()
	setup.Bucket()
	setup.Email()
	setup.Adapters()
	setup.Converters()

	setup.Finish()

	err := handler()
	if err != nil {
		infra.App.Logger.Errorf(err.Error())
		return
	}

	infra.App.Logger.Infof("Temporary contracts expired successfully")
	infra.App.Logger.Infof("Task finished")
}

func handler() error {
	usecase := usecase.NewExpireTemporaryContractUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	if err := usecase.ExpireTemporaryContracts(); err != nil {
		return err
	}
	return nil
}
