package payments

import (
	"fmt"
	"time"

	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/invoice"
	"github.com/stripe/stripe-go/v79/paymentintent"
	"github.com/stripe/stripe-go/v79/price"
	"github.com/stripe/stripe-go/v79/product"
	"github.com/stripe/stripe-go/v79/subscription"
	"github.com/venture-technology/venture/config"
	"github.com/venture-technology/venture/internal/entity"
)

type IStripe interface {
	CreatePrice(contract *entity.Contract) (*stripe.Price, error)
	CreateProduct(contract *entity.Contract) (*stripe.Product, error)
	CreateSubscription(contract *entity.Contract) (*stripe.Subscription, error)
	GetSubscription(subscriptionId string) (*stripe.Subscription, error)
	ListSubscriptions(contract *entity.Contract) ([]entity.SubscriptionInfo, error)
	DeleteSubscription(contract *entity.Contract) (*stripe.Subscription, error)
	GetInvoice(invoiceId string) (*stripe.Invoice, error)
	ListInvoices(contract *entity.Contract) ([]entity.InvoiceInfo, error)
	CalculateRemainingValueSubscription(invoices []entity.InvoiceInfo) *entity.InvoiceRemaining
	FineResponsible(contract *entity.Contract, amountFine int64) (*stripe.PaymentIntent, error)
}

type StripeUseCase struct{}

func NewStripeUseCase() *StripeUseCase {
	return &StripeUseCase{}
}

func (su *StripeUseCase) CreateProduct(contract *entity.Contract) (*stripe.Product, error) {
	conf := config.Get()

	stripe.Key = conf.StripeEnv.SecretKey

	params := &stripe.ProductParams{
		Name:        stripe.String(fmt.Sprintf("Assinatura - Motorista: %s & Responsável: %s por: %s", contract.Driver.CNH, contract.Child.Responsible.CPF, contract.Child.RG)),
		Description: stripe.String(fmt.Sprintf("Assinatura - Motorista: %s & Responsável: %s por: %s", contract.Driver.Name, contract.Child.Responsible.Name, contract.Child.Name)),
	}

	prodt, err := product.New(params)
	if err != nil {
		return nil, err
	}

	return prodt, nil
}

func (su *StripeUseCase) CreatePrice(contract *entity.Contract) (*stripe.Price, error) {
	conf := config.Get()

	stripe.Key = conf.StripeEnv.SecretKey

	params := &stripe.PriceParams{
		Currency: stripe.String(string("brl")),
		Product:  stripe.String(contract.StripeSubscription.ProductSubscriptionId),
		Recurring: &stripe.PriceRecurringParams{
			Interval: stripe.String("month"),
		},
		UnitAmount: stripe.Int64(int64(contract.Amount) * 100),
	}

	pr, err := price.New(params)
	if err != nil {
		return nil, err
	}

	return pr, nil
}

func (su *StripeUseCase) CreateSubscription(contract *entity.Contract) (*stripe.Subscription, error) {
	conf := config.Get()

	stripe.Key = conf.StripeEnv.SecretKey

	params := &stripe.SubscriptionParams{
		CancelAt: stripe.Int64(time.Now().AddDate(0, 12, 0).Unix()),
		Customer: stripe.String(contract.Child.Responsible.CustomerId),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Price: stripe.String(contract.StripeSubscription.PriceSubscriptionId),
			},
		},
	}

	subs, err := subscription.New(params)
	if err != nil {
		return nil, err
	}

	return subs, err
}

func (su *StripeUseCase) GetSubscription(subscriptionId string) (*stripe.Subscription, error) {
	conf := config.Get()

	stripe.Key = conf.StripeEnv.SecretKey

	subscription, err := subscription.Get(subscriptionId, nil)
	if err != nil {
		return nil, err
	}
	return subscription, nil
}

func (su *StripeUseCase) ListSubscriptions(contract *entity.Contract) ([]entity.SubscriptionInfo, error) {
	conf := config.Get()

	stripe.Key = conf.StripeEnv.SecretKey

	params := &stripe.SubscriptionListParams{
		Customer: stripe.String(contract.Child.Responsible.CustomerId),
	}
	params.Filters.AddFilter("limit", "", "10")

	var subscriptions []entity.SubscriptionInfo

	i := subscription.List(params)

	for i.Next() {
		s := i.Subscription()
		subscriptions = append(subscriptions, entity.SubscriptionInfo{
			ID:     s.ID,
			Status: string(s.Status),
		})
	}

	if err := i.Err(); err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func (su *StripeUseCase) DeleteSubscription(contract *entity.Contract) (*stripe.Subscription, error) {
	conf := config.Get()

	stripe.Key = conf.StripeEnv.SecretKey

	deletedSub, err := subscription.Cancel(contract.StripeSubscription.SubscriptionId, nil)
	if err != nil {
		return nil, err
	}
	return deletedSub, nil
}

func (su *StripeUseCase) GetInvoice(invoiceId string) (*stripe.Invoice, error) {
	conf := config.Get()

	stripe.Key = conf.StripeEnv.SecretKey

	inv, err := invoice.Get(invoiceId, nil)
	if err != nil {
		return nil, err
	}
	return inv, nil
}

func (su *StripeUseCase) ListInvoices(contract *entity.Contract) ([]entity.InvoiceInfo, error) {
	conf := config.Get()

	stripe.Key = conf.StripeEnv.SecretKey

	params := &stripe.InvoiceListParams{
		Subscription: stripe.String(contract.StripeSubscription.SubscriptionId),
	}

	var invoices []entity.InvoiceInfo

	i := invoice.List(params)

	for i.Next() {

		charge := i.Invoice()
		invoices = append(invoices, entity.InvoiceInfo{
			ID:              charge.ID,
			Status:          string(charge.Status),
			AmountDue:       charge.AmountDue,
			AmountRemaining: charge.AmountRemaining * 100,
		})

	}

	if err := i.Err(); err != nil {
		return nil, err
	}

	return invoices, nil
}

func (su *StripeUseCase) CalculateRemainingValueSubscription(invoices []entity.InvoiceInfo) *entity.InvoiceRemaining {
	invoice := entity.InvoiceRemaining{
		InvoiceValue: float64(invoices[0].AmountDue / 100),
		Quantity:     float64(len(invoices)),
	}

	invoice.Remaining = invoice.InvoiceValue * (12 - invoice.Quantity)

	invoice.Fines = invoice.Remaining * 0.40

	return &invoice
}

func (su *StripeUseCase) FineResponsible(contract *entity.Contract, amountFine int64) (*stripe.PaymentIntent, error) {
	conf := config.Get()

	stripe.Key = conf.StripeEnv.SecretKey

	params := &stripe.PaymentIntentParams{
		Customer:      stripe.String(contract.Child.Responsible.CustomerId),
		Amount:        stripe.Int64(amountFine * 100),
		Currency:      stripe.String("brl"),
		PaymentMethod: stripe.String(contract.Child.Responsible.PaymentMethodId),
		OffSession:    stripe.Bool(true),
		Confirm:       stripe.Bool(true),
	}

	paym, err := paymentintent.New(params)
	if err != nil {
		return nil, err
	}

	return paym, nil
}