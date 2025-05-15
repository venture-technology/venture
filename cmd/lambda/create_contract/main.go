package main

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/venture-technology/venture/internal/infra"
	"github.com/venture-technology/venture/internal/setup"
	"github.com/venture-technology/venture/internal/usecase"
)

func main() {
	setup := setup.NewSetup()
	setup.Logger("venture-lambda-create-contract")
	setup.Postgres()
	setup.Repositories()
	setup.Bucket()
	setup.Queue()
	setup.Adapters()
	setup.Converters()

	setup.Finish()

	sqs := infra.App.Queue

	msgs, err := sqs.PullMessages(viper.GetString("CREATE_LABEL_CONTRACT_QUEUE"))
	if err != nil {
		infra.App.Logger.Errorf(fmt.Sprintf("error pulling messages from queue: %v", err.Error()))
		return
	}

	for _, msg := range msgs {
		infra.App.Logger.Infof(fmt.Sprintf("message: %s", msg.Body))
		err := handler(msg.Body)
		if err == nil {
			sqs.DeleteMessage(viper.GetString("CREATE_CONTRACT_QUEUE"), msg.ReceiptHandle)
		}
	}
}

func handler(msg string) error {
	uc := usecase.NewCreateContractUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
		infra.App.Adapters,
		infra.App.Bucket,
		infra.App.Queue,
		infra.App.Converters,
	)

	err := uc.Execute(msg)
	if err != nil {
		return err
	}

	return nil
}
