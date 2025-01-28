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

type StripeAdapter struct {
	config config.Config
}

func NewStripeAdapter(
	config config.Config,
) *StripeAdapter {
	return &StripeAdapter{
		config: config,
	}
}

func (su *StripeAdapter) CreateProduct(contract *entity.Contract) (*stripe.Product, error) {
	stripe.Key = su.config.StripeEnv.SecretKey

	params := &stripe.ProductParams{
		Name:        stripe.String(fmt.Sprintf("Assinatura - Motorista: %s & Responsável: %s por: %s", contract.Driver.CNH, contract.Kid.Responsible.CPF, contract.Kid.RG)),
		Description: stripe.String(fmt.Sprintf("Assinatura - Motorista: %s & Responsável: %s por: %s", contract.Driver.Name, contract.Kid.Responsible.Name, contract.Kid.Name)),
	}

	prodt, err := product.New(params)
	if err != nil {
		return nil, err
	}

	return prodt, nil
}

func (su *StripeAdapter) CreatePrice(contract *entity.Contract) (*stripe.Price, error) {
	stripe.Key = su.config.StripeEnv.SecretKey

	params := &stripe.PriceParams{
		Currency: stripe.String(string("brl")),
		Product:  stripe.String(contract.StripeSubscription.Product),
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

func (su *StripeAdapter) CreateSubscription(contract *entity.Contract) (*stripe.Subscription, error) {
	stripe.Key = su.config.StripeEnv.SecretKey

	params := &stripe.SubscriptionParams{
		CancelAt: stripe.Int64(time.Now().AddDate(0, 12, 0).Unix()),
		Customer: stripe.String(contract.Kid.Responsible.CustomerId),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Price: stripe.String(contract.StripeSubscription.Price),
			},
		},
	}

	subs, err := subscription.New(params)
	if err != nil {
		return nil, err
	}

	return subs, err
}

func (su *StripeAdapter) GetSubscription(subscriptionId string) (*stripe.Subscription, error) {
	stripe.Key = su.config.StripeEnv.SecretKey
	subscription, err := subscription.Get(subscriptionId, nil)
	if err != nil {
		return nil, err
	}
	return subscription, nil
}

func (su *StripeAdapter) ListSubscriptions(contract *entity.Contract) ([]entity.SubscriptionInfo, error) {
	stripe.Key = su.config.StripeEnv.SecretKey

	params := &stripe.SubscriptionListParams{
		Customer: stripe.String(contract.Kid.Responsible.CustomerId),
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

func (su *StripeAdapter) DeleteSubscription(contract *entity.Contract) (*stripe.Subscription, error) {
	stripe.Key = su.config.StripeEnv.SecretKey
	deletedSub, err := subscription.Cancel(contract.StripeSubscription.ID, nil)
	if err != nil {
		return nil, err
	}
	return deletedSub, nil
}

func (su *StripeAdapter) GetInvoice(invoiceId string) (*stripe.Invoice, error) {
	stripe.Key = su.config.StripeEnv.SecretKey
	inv, err := invoice.Get(invoiceId, nil)
	if err != nil {
		return nil, err
	}
	return inv, nil
}

func (su *StripeAdapter) ListInvoices(contractId string) (map[string]entity.InvoiceInfo, error) {
	stripe.Key = su.config.StripeEnv.SecretKey

	params := &stripe.InvoiceListParams{
		Subscription: stripe.String(contractId),
	}

	invoices := make(map[string]entity.InvoiceInfo)
	i := invoice.List(params)
	for i.Next() {
		charge := i.Invoice()
		createdTime := time.Unix(charge.Created, 0)
		month := createdTime.Month().String()
		invoiceInfo := entity.InvoiceInfo{
			ID:              charge.ID,
			Status:          string(charge.Status),
			AmountDue:       charge.AmountDue,
			AmountRemaining: charge.AmountRemaining * 100,
			Month:           month,
			Date:            createdTime.Format("01/06"),
		}
		invoices[month] = invoiceInfo
	}

	if err := i.Err(); err != nil {
		return nil, err
	}

	return invoices, nil
}

func (su *StripeAdapter) CalculateRemainingValueSubscription(invoices map[string]entity.InvoiceInfo, amount float64) float64 {
	quantity := quantityInvoicesPaid(invoices)
	return (amount * float64(quantity)) * 0.40
}

func (su *StripeAdapter) FineResponsible(contract *entity.Contract, amountFine int64) (*stripe.PaymentIntent, error) {
	stripe.Key = su.config.StripeEnv.SecretKey

	params := &stripe.PaymentIntentParams{
		Customer:      stripe.String(contract.Kid.Responsible.CustomerId),
		Amount:        stripe.Int64(amountFine * 100),
		Currency:      stripe.String("brl"),
		PaymentMethod: stripe.String(contract.Kid.Responsible.PaymentMethodId),
		OffSession:    stripe.Bool(true),
		Confirm:       stripe.Bool(true),
	}

	paym, err := paymentintent.New(params)
	if err != nil {
		return nil, err
	}

	return paym, nil
}

func quantityInvoicesPaid(invoices map[string]entity.InvoiceInfo) uint64 {
	qtdMonth := 12
	for _, invoice := range invoices {
		if invoice.Status == "paid" {
			qtdMonth--
		}
	}
	return uint64(qtdMonth)
}
