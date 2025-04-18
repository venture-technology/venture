package payments

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/customer"
	"github.com/stripe/stripe-go/v79/invoice"
	"github.com/stripe/stripe-go/v79/paymentintent"
	"github.com/stripe/stripe-go/v79/paymentmethod"
	"github.com/stripe/stripe-go/v79/price"
	"github.com/stripe/stripe-go/v79/product"
	"github.com/stripe/stripe-go/v79/subscription"
	"github.com/venture-technology/venture/internal/entity"
)

type StripeAdapter struct {
}

func NewStripeAdapter() *StripeAdapter {
	return &StripeAdapter{}
}

func (su *StripeAdapter) CreateProduct(
	contract *entity.Contract,
) (*stripe.Product, error) {
	stripe.Key = viper.GetString("STRIPE_SECRET_KEY")

	params := &stripe.ProductParams{
		Name: stripe.String(fmt.Sprintf("Assinatura %s - %s - %s - %s",
			contract.DriverCNH,
			contract.ResponsibleCPF,
			contract.KidRG,
			contract.SchoolCNPJ,
		)),
	}

	prodt, err := product.New(params)
	if err != nil {
		return nil, err
	}

	return prodt, nil
}

func (su *StripeAdapter) CreatePrice(
	stripeProductID string,
	amount float64,
) (*stripe.Price, error) {
	stripe.Key = viper.GetString("STRIPE_SECRET_KEY")

	params := &stripe.PriceParams{
		Currency: stripe.String(string("brl")),
		Product:  stripe.String(stripeProductID),
		Recurring: &stripe.PriceRecurringParams{
			Interval: stripe.String("month"),
		},
		UnitAmount: stripe.Int64(int64(amount) * 100),
	}

	pr, err := price.New(params)
	if err != nil {
		return nil, err
	}

	return pr, nil
}

func (su *StripeAdapter) CreateSubscription(
	customerID,
	stripePriceID string,
) (*stripe.Subscription, error) {
	stripe.Key = viper.GetString("STRIPE_SECRET_KEY")

	params := &stripe.SubscriptionParams{
		CancelAt: stripe.Int64(time.Now().AddDate(0, 12, 0).Unix()),
		Customer: stripe.String(customerID),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Price: stripe.String(stripePriceID),
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
	stripe.Key = viper.GetString("STRIPE_SECRET_KEY")
	subscription, err := subscription.Get(subscriptionId, nil)
	if err != nil {
		return nil, err
	}
	return subscription, nil
}

func (su *StripeAdapter) ListSubscriptions(responsible *entity.Responsible) ([]entity.SubscriptionInfo, error) {
	stripe.Key = viper.GetString("STRIPE_SECRET_KEY")

	params := &stripe.SubscriptionListParams{
		Customer: stripe.String(responsible.CustomerId),
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
	stripe.Key = viper.GetString("STRIPE_SECRET_KEY")
	deletedSub, err := subscription.Cancel(contract.StripeSubscriptionID, nil)
	if err != nil {
		return nil, err
	}
	return deletedSub, nil
}

func (su *StripeAdapter) GetInvoice(invoiceId string) (*stripe.Invoice, error) {
	stripe.Key = viper.GetString("STRIPE_SECRET_KEY")
	inv, err := invoice.Get(invoiceId, nil)
	if err != nil {
		return nil, err
	}
	return inv, nil
}

func (su *StripeAdapter) ListInvoices(contractId string) (map[string]entity.InvoiceInfo, error) {
	stripe.Key = viper.GetString("STRIPE_SECRET_KEY")

	params := &stripe.InvoiceListParams{
		Subscription: stripe.String(contractId),
	}

	invoices := make(map[string]entity.InvoiceInfo)
	i := invoice.List(params)
	for i.Next() {
		charge := i.Invoice()
		createdTime := time.Unix(charge.Created, 0)
		month := strings.ToLower(createdTime.Month().String())
		amount := su.GetAmountFromInvoice(charge)
		invoiceInfo := entity.InvoiceInfo{
			ID:          charge.ID,
			Status:      string(charge.Status),
			Amount:      amount,
			AmountCents: int64(amount) * 100,
			Month:       month,
			Date:        createdTime.Format("01/06"),
		}

		invoices[month] = invoiceInfo
	}

	if err := i.Err(); err != nil {
		return nil, err
	}

	return invoices, nil
}

func (su *StripeAdapter) GetAmountFromInvoice(charge *stripe.Invoice) float64 {
	amount := charge.AmountDue
	if charge.AmountDue == 0 {
		amount = charge.AmountRemaining
	}

	return float64(amount) / 100
}

func (su *StripeAdapter) CalculateRemainingValueSubscription(invoices map[string]entity.InvoiceInfo, amount float64) float64 {
	quantity := quantityInvoicesPaid(invoices)
	return (amount * float64(quantity)) * 0.40
}

func (su *StripeAdapter) FineResponsible(
	customerId,
	paymentMethodId string,
	amountFine int64,
) (*stripe.PaymentIntent, error) {
	stripe.Key = viper.GetString("STRIPE_SECRET_KEY")

	params := &stripe.PaymentIntentParams{
		Customer:      stripe.String(customerId),
		Amount:        stripe.Int64(amountFine * 100),
		Currency:      stripe.String("brl"),
		PaymentMethod: stripe.String(paymentMethodId),
		OffSession:    stripe.Bool(true),
		Confirm:       stripe.Bool(true),
	}

	paym, err := paymentintent.New(params)
	if err != nil {
		return nil, err
	}

	return paym, nil
}

func (su *StripeAdapter) CreateCustomer(
	responsible *entity.Responsible,
) (string, error) {
	stripe.Key = viper.GetString("STRIPE_SECRET_KEY")

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

func (su *StripeAdapter) CreatePaymentMethod(
	token string,
) (*stripe.PaymentMethod, error) {
	stripe.Key = viper.GetString("STRIPE_SECRET_KEY")

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

func (su *StripeAdapter) AttachCardToResponsible(
	customerID,
	paymentMethodID string,
) (*stripe.PaymentMethod, error) {
	stripe.Key = viper.GetString("STRIPE_SECRET_KEY")

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

func (su *StripeAdapter) DeleteStripeUser(customerId string) (*stripe.Customer, error) {
	stripe.Key = viper.GetString("STRIPE_SECRET_KEY")

	c, err := customer.Del(customerId, nil)
	if err != nil {
		return nil, err
	}

	return c, nil
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
