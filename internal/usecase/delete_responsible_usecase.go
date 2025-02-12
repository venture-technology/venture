package usecase

import (
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/customer"
	"github.com/venture-technology/venture/config"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type DeleteResponsibleUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
	config       config.Config
}

func NewDeleteResponsibleUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
	config config.Config,
) *DeleteResponsibleUseCase {
	return &DeleteResponsibleUseCase{
		repositories: repositories,
		logger:       logger,
		config:       config,
	}
}

func (druc *DeleteResponsibleUseCase) DeleteResponsible(cpf string) error {
	responsible, err := druc.repositories.ResponsibleRepository.Get(cpf)
	if err != nil {
		return err
	}
	_, err = druc.deleteStripeUser(responsible.CustomerId)
	if err != nil {
		return err
	}
	return druc.repositories.ResponsibleRepository.Delete(cpf)
}

func (druc *DeleteResponsibleUseCase) deleteStripeUser(customerId string) (*stripe.Customer, error) {
	stripe.Key = druc.config.StripeEnv.SecretKey

	c, err := customer.Del(customerId, nil)
	if err != nil {
		return nil, err
	}

	return c, nil
}
