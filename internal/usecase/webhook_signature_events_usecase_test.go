package usecase

import (
	"errors"
	"testing"

	"github.com/gin-gonic/gin"
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

		usecase := NewWebhookSignatureEventsUseCase(
			&persistence.PostgresRepositories{},
			logger,
			adapters.Adapters{
				AgreementService: agreementService,
			},
		)

		resp, err := usecase.Execute(&gin.Context{}, event)

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

		usecase := NewWebhookSignatureEventsUseCase(
			&persistence.PostgresRepositories{},
			logger,
			adapters.Adapters{
				AgreementService: agreementService,
			},
		)

		agreementService.On("HandleCallbackVerification").Return(false, errors.New("handle callback verification error"))
		logger.On("Errorf", mock.Anything, mock.Anything)

		_, err := usecase.Execute(&gin.Context{}, event)

		assert.EqualError(t, err, "handle callback verification error")
	})

	t.Run("receive signature request all signed webhook and returns fail", func(t *testing.T) {
		logger := mocks.NewLogger(t)
		agreementService := mocks.NewAgreementService(t)

		event := agreements.EventWrapper{
			Event: agreements.Event{
				EventType: "signature_request_all_signed",
			},
		}

		usecase := NewWebhookSignatureEventsUseCase(
			&persistence.PostgresRepositories{},
			logger,
			adapters.Adapters{
				AgreementService: agreementService,
			},
		)

		agreementService.On("SignatureRequestAllSigned", mock.Anything).Return(agreements.ASRASOutput{}, errors.New("signature request all signed error"))
		logger.On("Errorf", mock.Anything, mock.Anything)

		_, err := usecase.Execute(&gin.Context{}, event)

		assert.EqualError(t, err, "signature request all signed error")
	})

	t.Run("receive callback test webhook and returns sucess", func(t *testing.T) {
		logger := mocks.NewLogger(t)
		agreementService := mocks.NewAgreementService(t)

		event := agreements.EventWrapper{
			Event: agreements.Event{
				EventType: "callback_test",
			},
		}

		usecase := NewWebhookSignatureEventsUseCase(
			&persistence.PostgresRepositories{},
			logger,
			adapters.Adapters{
				AgreementService: agreementService,
			},
		)

		agreementService.On("HandleCallbackVerification").Return(true, nil)
		logger.On("Infof", mock.Anything, mock.Anything)

		_, err := usecase.Execute(&gin.Context{}, event)

		assert.Nil(t, err)
	})
}
