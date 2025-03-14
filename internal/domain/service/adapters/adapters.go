package adapters

import (
	"github.com/stripe/stripe-go/v79"
	"github.com/venture-technology/venture/internal/domain/service/agreements"
	"github.com/venture-technology/venture/internal/entity"
)

type Adapters struct {
	AddressService   AddressService
	PaymentsService  PaymentsService
	AgreementService AgreementService
}

type AddressService interface {
	GetDistance(origin, destination string) (*float64, error)
}

type PaymentsService interface {
	// this function is used to create price for the driver
	CreatePrice(contract *entity.Contract) (*stripe.Price, error)
	// when driver is created together the product created too for driver work become product
	CreateProduct(contract *entity.Contract) (*stripe.Product, error)
	// when responsible hire driver
	CreateSubscription(contract *entity.Contract) (*stripe.Subscription, error)
	GetSubscription(subscriptionId string) (*stripe.Subscription, error)
	ListSubscriptions(contract *entity.Contract) ([]entity.SubscriptionInfo, error)
	DeleteSubscription(contract *entity.Contract) (*stripe.Subscription, error)
	// get invoice from subcription
	GetInvoice(invoiceId string) (*stripe.Invoice, error)
	ListInvoices(contractId string) (map[string]entity.InvoiceInfo, error)
	// this calc is used to calculate the remaining value of the subscription
	CalculateRemainingValueSubscription(invoices map[string]entity.InvoiceInfo, amount float64) float64
	// fine responsible when cancel subcription
	FineResponsible(contract *entity.Contract, amountFine int64) (*stripe.PaymentIntent, error)
	CreateCustomer(responsible *entity.Responsible) (string, error)
	CreatePaymentMethod(token string) (*stripe.PaymentMethod, error)
	AttachCardToResponsible(customerID, paymentMethodID string) (*stripe.PaymentMethod, error)
	DeleteStripeUser(customerId string) (*stripe.Customer, error)
}

type AgreementService interface {
	SignatureRequest(contract entity.ContractProperty) (agreements.ContractRequest, error)
	// this function is used to get the html file of the agreement
	// the param is variable because can be used to get the html file from different applications
	GetAgreementHtml(path string) ([]byte, error)
	HandleCallbackVerification() (any, error)
}
