package usecase

import (
	"fmt"

	"github.com/venture-technology/venture/internal/domain/service/adapters"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type CreateResponsibleUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
	adapters     adapters.Adapters
	email        contracts.WorkerEmail
}

func NewCreateResponsibleUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
	adapters adapters.Adapters,
	email contracts.WorkerEmail,
) *CreateResponsibleUseCase {
	return &CreateResponsibleUseCase{
		repositories: repositories,
		logger:       logger,
		adapters:     adapters,
		email:        email,
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

	err = cruc.repositories.ResponsibleRepository.Create(responsible)
	if err != nil {
		return err
	}

	err = cruc.email.Enqueue(&entity.Email{
		Recipient: responsible.Email,
		Subject:   fmt.Sprintf("Olá %s, sua conta foi criada no Venture! Estamos felizes em contar com você.", responsible.Name),
		Body:      "Seja bem-vindo a nossa plataforma, registre seu filho, encontre um motorista, economize tempo e dinheiro.",
	})

	return err
}
