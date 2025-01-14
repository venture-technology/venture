package usecase

import (
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/customer"
	"github.com/venture-technology/venture/config"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type CreateResponsibleUseCase struct {
	repositories *persistence.PostgresRepositories
	logger       contracts.Logger
	config       config.Config
}

func NewCreateResponsibleUseCase(
	repositories *persistence.PostgresRepositories,
	logger contracts.Logger,
	config config.Config,
) *CreateResponsibleUseCase {
	return &CreateResponsibleUseCase{
		repositories: repositories,
		logger:       logger,
		config:       config,
	}
}

func (cruc *CreateResponsibleUseCase) CreateResponsible(responsible *entity.Responsible) error {
	var err error
	responsible.CustomerId, err = cruc.CreateCustomer(responsible)
	if err != nil {
		return err
	}
	return cruc.repositories.ResponsibleRepository.Create(responsible)
}

func (cruc *CreateResponsibleUseCase) CreateCustomer(
	responsible *entity.Responsible,
) (string, error) {
	stripe.Key = cruc.config.StripeEnv.SecretKey
	params := &stripe.CustomerParams{
		Name:  stripe.String(responsible.Name),
		Email: stripe.String(responsible.Email),
		Phone: stripe.String(responsible.Phone),
	}
	resp, err := customer.New(params)
	if err != nil {
		return "", err
	}
	return resp.ID, nil
}
