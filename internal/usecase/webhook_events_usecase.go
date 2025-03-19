package usecase

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/domain/service/adapters"
	"github.com/venture-technology/venture/internal/domain/service/agreements"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type WebhookEventsUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
	adapters     adapters.Adapters
}

func NewWebhookEventsUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
	adapters adapters.Adapters,
) *WebhookEventsUseCase {
	return &WebhookEventsUseCase{
		repositories: repositories,
		logger:       logger,
		adapters:     adapters,
	}
}

func (weuc *WebhookEventsUseCase) Execute(
	httpCtx *gin.Context,
	eventWrapper agreements.EventWrapper,
) (any, error) {
	var eventHandlers = map[string]func(httpCtx *gin.Context) (any, error){
		"callback_test": func(httpCtx *gin.Context) (any, error) {
			return weuc.adapters.AgreementService.HandleCallbackVerification()
		},
		"signature_request_all_signed": func(httpCtx *gin.Context) (any, error) {
			asras, err := weuc.adapters.AgreementService.SignatureRequestAllSigned(httpCtx)
			if err != nil {
				return agreements.ASRASOutput{}, err
			}

			uc := NewAcceptContractUseCase(
				weuc.repositories,
				weuc.logger,
				weuc.adapters,
			)

			err = uc.AcceptContract(asras)
			if err != nil {
				return agreements.ASRASOutput{}, err
			}

			return asras, nil
		},
	}

	eventType := eventWrapper.Event.EventType

	if handler, exists := eventHandlers[eventType]; exists {
		response, err := handler(httpCtx)
		if err != nil {
			weuc.logger.Errorf(fmt.Sprintf("error handling event: %s", eventType))
			return nil, err
		}
		weuc.logger.Infof(fmt.Sprintf("event handled: %s, %s", response, eventType))
		return response, nil
	}

	return nil, fmt.Errorf("event type not used")
}
