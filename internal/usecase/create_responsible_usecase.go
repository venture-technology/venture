package usecase

import (
	"fmt"

	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/customer"
	"github.com/stripe/stripe-go/v79/paymentmethod"
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

	paymentMethod, err := cruc.CreatePaymentMethod(responsible.CardToken)
	if err != nil {
		return err
	}

	responsible.PaymentMethodId = paymentMethod.ID

	_, err = cruc.AttachCardToResponsible(responsible.CustomerId, responsible.PaymentMethodId)

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

func (cruc *CreateResponsibleUseCase) CreatePaymentMethod(
	token string,
) (*stripe.PaymentMethod, error) {
	stripe.Key = cruc.config.StripeEnv.SecretKey

	params := &stripe.PaymentMethodParams{
		Type: stripe.String(string(stripe.PaymentMethodTypeCard)),
		Card: &stripe.PaymentMethodCardParams{
			Token: stripe.String(token),
		},
	}

	pm, err := paymentmethod.New(params)
	if err != nil {
		fmt.Println("Erro ao criar método de pagamento:", err)
		return nil, err
	}

	return pm, nil
}

func (cruc *CreateResponsibleUseCase) AttachCardToResponsible(
	customerID,
	paymentMethodID string,
) (*stripe.PaymentMethod, error) {
	stripe.Key = cruc.config.StripeEnv.SecretKey

	params := &stripe.PaymentMethodAttachParams{
		Customer: &customerID,
	}
	pm, err := paymentmethod.Attach(paymentMethodID, params)
	if err != nil {
		return nil, err
	}

	updateParams := &stripe.CustomerParams{
		InvoiceSettings: &stripe.CustomerInvoiceSettingsParams{
			DefaultPaymentMethod: stripe.String(pm.ID),
		},
	}

	_, err = customer.Update(customerID, updateParams)
	if err != nil {
		return nil, err
	}

	return pm, nil
}
