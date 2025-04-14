package usecase

import (
	"github.com/venture-technology/venture/internal/domain/service/adapters"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type WebhookPaymentEventsUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
	adapters     adapters.Adapters
}

func NewWebhookPaymentEventsUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
	adapters adapters.Adapters,
) *WebhookPaymentEventsUseCase {
	return &WebhookPaymentEventsUseCase{
		repositories: repositories,
		logger:       logger,
		adapters:     adapters,
	}
}

