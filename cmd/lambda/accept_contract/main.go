package main

import (
	"github.com/venture-technology/venture/internal/setup"
)

func main() {
	setup := setup.NewSetup()
	setup.Logger("venture-lambda-accept-contract")
	setup.Postgres()
	setup.Repositories()
	setup.Bucket()
	setup.Queue()
	setup.Adapters()
	setup.Converters()

	setup.Finish()
}
