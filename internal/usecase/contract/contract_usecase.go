package contract

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/invoice"
	"github.com/stripe/stripe-go/v79/paymentintent"
	"github.com/stripe/stripe-go/v79/price"
	"github.com/stripe/stripe-go/v79/product"
	"github.com/stripe/stripe-go/v79/subscription"
	"github.com/venture-technology/venture/config"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/repository"
)

type ContractUseCase struct {
	contractRepository repository.IContractRepository
}

func NewContractUseCase(cr repository.IContractRepository) *ContractUseCase {
	return &ContractUseCase{
		contractRepository: cr,
	}
}

func (cu *ContractUseCase) Create() {

}

func (cu *ContractUseCase) Get() {

}

func (cu *ContractUseCase) FindAllByCnh() {

}

func (cu *ContractUseCase) FindAllByCnpj() {

}

func (cu *ContractUseCase) FindAllByCpf() {

}

func (cu *ContractUseCase) Cancel() {

}

func CalculateContract(distance, amount float64) float64 {

	if distance < 2 {
		return 200
	}

	diff := distance - 2

	return 200 + (amount * diff)

}

func CreateProduct(contract *entity.Contract) (*stripe.Product, error) {

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

func CreatePrice(contract *entity.Contract) (*stripe.Price, error) {

	conf := config.Get()

	stripe.Key = conf.StripeEnv.SecretKey

	params := &stripe.PriceParams{
		Currency: stripe.String(string("brl")),
		Product:  stripe.String(contract.StripeSubscription.ProductSubscriptionId),
		Recurring: &stripe.PriceRecurringParams{
			Interval: stripe.String("month"),
		},
		UnitAmount: stripe.Int64(contract.Amount * 100),
	}

	pr, err := price.New(params)

	if err != nil {
		return nil, err
	}

	return pr, nil

}

func CreateSubscription(contract *entity.Contract) (*stripe.Subscription, error) {

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

func GetSubscription(subscriptionId string) (*stripe.Subscription, error) {

	conf := config.Get()

	stripe.Key = conf.StripeEnv.SecretKey

	subscription, err := subscription.Get(subscriptionId, nil)
	if err != nil {
		return nil, err
	}
	return subscription, nil

}

func ListSubscriptions(contract *entity.Contract) ([]entity.SubscriptionInfo, error) {

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

func DeleteSubscription(contract *entity.Contract) (*stripe.Subscription, error) {

	conf := config.Get()

	stripe.Key = conf.StripeEnv.SecretKey

	deletedSub, err := subscription.Cancel(contract.StripeSubscription.SubscriptionId, nil)
	if err != nil {
		return nil, err
	}
	return deletedSub, nil

}

func GetInvoice(invoiceId string) (*stripe.Invoice, error) {

	conf := config.Get()

	stripe.Key = conf.StripeEnv.SecretKey

	inv, err := invoice.Get(invoiceId, nil)
	if err != nil {
		return nil, err
	}
	return inv, nil

}

func ListInvoices(contract *entity.Contract) ([]entity.InvoiceInfo, error) {

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

func CalculateRemainingValueSubscription(invoices []entity.InvoiceInfo) *entity.InvoiceRemaining {

	invoice := entity.InvoiceRemaining{
		InvoiceValue: float64(invoices[0].AmountDue / 100),
		Quantity:     float64(len(invoices)),
	}

	invoice.Remaining = invoice.InvoiceValue * (12 - invoice.Quantity)

	invoice.Fines = invoice.Remaining * 0.40

	return &invoice

}

func FineResponsible(contract *entity.Contract, amountFine int64) (*stripe.PaymentIntent, error) {

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

func GetDistance(origin, destination string) (*float64, error) {

	conf := config.Get()

	endpoint := conf.GoogleCloudSecret.EndpointMatrixDistance

	params := url.Values{
		"units":        {"metric"},
		"origins":      {origin},
		"destinations": {destination},
		"key":          {conf.GoogleCloudSecret.ApiKey},
	}

	url := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	log.Print(url)

	resp, err := http.Get(url)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	var data entity.DistanceMatrixResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Print(err.Error())
		return nil, err
	}

	if data.Status != "OK" {
		log.Print("Erro na API:", data.Status)
	}

	distance := data.Rows[0].Elements[0].Distance.Text

	distance = strings.TrimSpace(strings.Replace(distance, "km", "", 1))

	kmFloat, err := strconv.ParseFloat(distance, 64)
	if err != nil {
		return nil, err
	}

	return &kmFloat, err

}
