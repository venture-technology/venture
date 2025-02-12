package usecase

import (
	"errors"

	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/customer"
	"github.com/venture-technology/venture/internal/domain/repository"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type DeleteResponsibleUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
}

func NewDeleteResponsibleUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
) *DeleteResponsibleUseCase {
	return &DeleteResponsibleUseCase{
		repositories: repositories,
		logger:       logger,
	}
}

func (druc *DeleteResponsibleUseCase) DeleteResponsible(cpf string) error {
	return druc.repositories.ResponsibleRepository.Delete(cpf)
}

type ResponsibleUseCase struct {
	responsibleRepo repository.ResponsibleRepository
	stripeSecretKey string
}

func NewResponsibleUseCase(responsibleRepo repository.ResponsibleRepository, stripeSecretKey string) *ResponsibleUseCase {
	return &ResponsibleUseCase{
		responsibleRepo: responsibleRepo,
		stripeSecretKey: stripeSecretKey,
	}
}

// DeleteResponsible deleta o responsável e o cliente no Stripe
func (uc *ResponsibleUseCase) DeleteResponsible(responsibleID int) error {
	// Busca o CustomerID do responsável
	customerID, err := uc.responsibleRepo.GetCustomerIDByResponsibleID(responsibleID)
	if err != nil {
		return err
	}

	// Deleta o cliente no Stripe
	if customerID != "" {
		err = uc.deleteStripeCustomer(customerID)
		if err != nil {
			return errors.New("failed to delete Stripe customer: " + err.Error())
		}
	}

	// Deleta o responsável no banco de dados
	err = uc.responsibleRepo.DeleteResponsible(responsibleID)
	if err != nil {
		return errors.New("failed to delete responsible: " + err.Error())
	}

	return nil
}

// deleteStripeCustomer deleta um cliente no Stripe
func (uc *ResponsibleUseCase) deleteStripeCustomer(customerID string) error {
	stripe.Key = uc.stripeSecretKey
	if stripe.Key == "" {
		return errors.New("Stripe secret key is not set")
	}

	_, err := customer.Del(customerID, nil)
	if err != nil {
		return err
	}

	return nil
}
