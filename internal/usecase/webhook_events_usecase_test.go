package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/venture-technology/venture/internal/domain/service/adapters"
	"github.com/venture-technology/venture/internal/domain/service/agreements"
	"github.com/venture-technology/venture/internal/infra/persistence"
	"github.com/venture-technology/venture/mocks"
)

func TestWebhookEventsUsecase_Execute(t *testing.T) {
	t.Run("if event not exists", func(t *testing.T) {
		logger := mocks.NewLogger(t)
		agreementService := mocks.NewAgreementService(t)

		event := agreements.EventWrapper{
			Event: agreements.Event{
				EventType: "not found",
			},
		}

		usecase := NewWebhookEventsUseCase(
			&persistence.PostgresRepositories{},
			logger,
			adapters.Adapters{
				AgreementService: agreementService,
			},
		)

		resp, err := usecase.Execute(event)

		assert.EqualError(t, err, "event type not used")
		assert.Nil(t, resp)
	})

	t.Run("receive callback test webhook and returns fail", func(t *testing.T) {
		logger := mocks.NewLogger(t)
		agreementService := mocks.NewAgreementService(t)

		event := agreements.EventWrapper{
			Event: agreements.Event{
				EventType: "callback_test",
			},
		}

		usecase := NewWebhookEventsUseCase(
			&persistence.PostgresRepositories{},
			logger,
			adapters.Adapters{
				AgreementService: agreementService,
			},
		)

		agreementService.On("HandleCallbackVerification").Return(false, errors.New("handle callback verification error"))
		logger.On("Errorf", mock.Anything, mock.Anything)

		_, err := usecase.Execute(event)

		assert.EqualError(t, err, "handle callback verification error")
	})

	t.Run("receive callback test webhook and returns sucess", func(t *testing.T) {
		logger := mocks.NewLogger(t)
		agreementService := mocks.NewAgreementService(t)

		event := agreements.EventWrapper{
			Event: agreements.Event{
				EventType: "callback_test",
			},
		}

		usecase := NewWebhookEventsUseCase(
			&persistence.PostgresRepositories{},
			logger,
			adapters.Adapters{
				AgreementService: agreementService,
			},
		)

		agreementService.On("HandleCallbackVerification").Return(true, nil)
		logger.On("Infof", mock.Anything, mock.Anything)

		_, err := usecase.Execute(event)

		assert.Nil(t, err)
	})
}
