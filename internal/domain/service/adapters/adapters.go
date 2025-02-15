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
	CreatePrice(contract *entity.Contract) (*stripe.Price, error)
	CreateProduct(contract *entity.Contract) (*stripe.Product, error)
	CreateSubscription(contract *entity.Contract) (*stripe.Subscription, error)
	GetSubscription(subscriptionId string) (*stripe.Subscription, error)
	ListSubscriptions(contract *entity.Contract) ([]entity.SubscriptionInfo, error)
	DeleteSubscription(contract *entity.Contract) (*stripe.Subscription, error)
	GetInvoice(invoiceId string) (*stripe.Invoice, error)
	ListInvoices(contractId string) (map[string]entity.InvoiceInfo, error)

	// this calc is used to calculate the remaining value of the subscription
	CalculateRemainingValueSubscription(invoices map[string]entity.InvoiceInfo, amount float64) float64

	FineResponsible(contract *entity.Contract, amountFine int64) (*stripe.PaymentIntent, error)
}

type AgreementService interface {
	SignatureRequest(contract entity.ContractProperty) (agreements.ContractRequest, error)
	// this function is used to get the html file of the agreement
	// the param is variable because can be used to get the html file from different applications
	GetAgreementHtml(path string) ([]byte, error)
}
