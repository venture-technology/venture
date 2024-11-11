package responsible

import (
	"context"
	"fmt"

	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/customer"
	"github.com/stripe/stripe-go/v79/paymentmethod"
	"github.com/venture-technology/venture/config"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/repository"
	"go.uber.org/zap"
)

type ResponsibleUseCase struct {
	responsibleRepository repository.IResponsibleRepository
	logger                *zap.Logger
}

func NewResponsibleUseCase(responsibleRepository repository.IResponsibleRepository, logger *zap.Logger) *ResponsibleUseCase {
	return &ResponsibleUseCase{
		responsibleRepository: responsibleRepository,
		logger:                logger,
	}
}

func (ru *ResponsibleUseCase) Create(ctx context.Context, responsible *entity.Responsible) error {
	return ru.responsibleRepository.Create(ctx, responsible)
}

func (ru *ResponsibleUseCase) Get(ctx context.Context, cpf *string) (*entity.Responsible, error) {
	return ru.responsibleRepository.Get(ctx, cpf)
}

func (ru *ResponsibleUseCase) Update(ctx context.Context, currentResponsible, responsible *entity.Responsible) error {
	return ru.responsibleRepository.Update(ctx, currentResponsible, responsible)
}

func (ru *ResponsibleUseCase) Delete(ctx context.Context, cpf *string) error {
	return ru.responsibleRepository.Delete(ctx, cpf)
}

func (ru *ResponsibleUseCase) SaveCard(ctx context.Context, cpf, cardToken, paymentMethodId *string) error {
	return ru.responsibleRepository.SaveCard(ctx, cpf, cardToken, paymentMethodId)
}

func (ru *ResponsibleUseCase) CreateCustomer(ctx context.Context, responsible *entity.Responsible) (*stripe.Customer, error) {
	conf := config.Get()

	stripe.Key = conf.StripeEnv.SecretKey

	params := &stripe.CustomerParams{
		Name:  stripe.String(responsible.Name),
		Email: stripe.String(responsible.Email),
		Phone: stripe.String(responsible.Phone),
	}

	resp, err := customer.New(params)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (ru *ResponsibleUseCase) UpdateCustomer(ctx context.Context, responsible *entity.Responsible) (*stripe.Customer, error) {
	conf := config.Get()

	stripe.Key = conf.StripeEnv.SecretKey

	params := &stripe.CustomerParams{
		Email: &responsible.Email,
		Phone: &responsible.Phone,
	}

	updatedCustomer, err := customer.Update(responsible.CustomerId, params)

	if err != nil {
		return nil, err
	}

	return updatedCustomer, nil
}

func (ru *ResponsibleUseCase) DeleteCustomer(ctx context.Context, customerId string) (*stripe.Customer, error) {
	conf := config.Get()

	stripe.Key = conf.StripeEnv.SecretKey

	c, err := customer.Del(customerId, nil)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (ru *ResponsibleUseCase) CreatePaymentMethod(ctx context.Context, cardToken *string) (*stripe.PaymentMethod, error) {
	conf := config.Get()

	stripe.Key = conf.StripeEnv.SecretKey

	params := &stripe.PaymentMethodParams{
		Type: stripe.String(string(stripe.PaymentMethodTypeCard)),
		Card: &stripe.PaymentMethodCardParams{
			Token: stripe.String(*cardToken),
		},
	}

	pm, err := paymentmethod.New(params)
	if err != nil {
		fmt.Println("Erro ao criar método de pagamento:", err)
		return nil, err
	}

	return pm, nil
}

func (ru *ResponsibleUseCase) AttachPaymentMethod(ctx context.Context, customerId, paymentMethodId *string, isDefault bool) (*stripe.PaymentMethod, error) {
	conf := config.Get()

	stripe.Key = conf.StripeEnv.SecretKey

	params := &stripe.PaymentMethodAttachParams{
		Customer: customerId,
	}
	pm, err := paymentmethod.Attach(*paymentMethodId, params)
	if err != nil {
		return nil, err
	}

	if isDefault {

		updateParams := &stripe.CustomerParams{
			InvoiceSettings: &stripe.CustomerInvoiceSettingsParams{
				DefaultPaymentMethod: stripe.String(pm.ID),
			},
		}

		_, err := customer.Update(*customerId, updateParams)
		if err != nil {
			return nil, err
		}

	}

	return pm, nil
}
