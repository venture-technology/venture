package usecase

import (
	"fmt"

	"github.com/venture-technology/venture/internal/domain/service/agreements"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type WebhookEventsUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewWebhookEventsUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *WebhookEventsUseCase {
	return &WebhookEventsUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

var eventHandlers = map[string]func() (any, error){
	"callback_test": func() (any, error) {
		return agreements.HandleCallbackVerification()
	},
}

func (weuc *WebhookEventsUseCase) Execute(eventWrapper agreements.EventWrapper) (any, error) {
	eventType := eventWrapper.Event.EventType

	if handler, exists := eventHandlers[eventType]; exists {
		response, err := handler()
		if err != nil {
			weuc.logger.Errorf(fmt.Sprintf("error handling event: %s", err))
			return nil, err
		}
		weuc.logger.Infof(fmt.Sprintf("event handled: %s, %s", response, eventType))
		return response, nil
	}

	return nil, fmt.Errorf("event type not used")
}
