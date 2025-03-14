package usecase

import (
	"github.com/venture-technology/venture/internal/domain/service/adapters"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type CreateResponsibleUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
	adapters     adapters.Adapters
}

func NewCreateResponsibleUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
	adapters adapters.Adapters,
) *CreateResponsibleUseCase {
	return &CreateResponsibleUseCase{
		repositories: repositories,
		logger:       logger,
		adapters:     adapters,
	}
}

func (cruc *CreateResponsibleUseCase) CreateResponsible(responsible *entity.Responsible) error {
	var err error
	responsible.CustomerId, err = cruc.adapters.PaymentsService.CreateCustomer(responsible)
	if err != nil {
		return err
	}

	paymentMethod, err := cruc.adapters.PaymentsService.CreatePaymentMethod(responsible.CardToken)
	if err != nil {
		return err
	}

	responsible.PaymentMethodId = paymentMethod.ID

	// this err is ignored, because we can be attach later if fail
	_, err = cruc.adapters.PaymentsService.AttachCardToResponsible(responsible.CustomerId, responsible.PaymentMethodId)

	return cruc.repositories.ResponsibleRepository.Create(responsible)
}
