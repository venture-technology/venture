package contract

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/invoice"
	"github.com/stripe/stripe-go/v79/paymentintent"
	"github.com/stripe/stripe-go/v79/price"
	"github.com/stripe/stripe-go/v79/product"
	"github.com/stripe/stripe-go/v79/subscription"
	"github.com/venture-technology/venture/config"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/repository"
	"github.com/venture-technology/venture/internal/usecase"
)

type ContractUseCase struct {
	contractRepository repository.IContractRepository
	childRepository    repository.IChildRepository
	driverRepository   repository.IDriverRepository
	schoolRepository   repository.ISchoolRepository
}

func NewContractUseCase(cou repository.IContractRepository, cr repository.IChildRepository, dr repository.IDriverRepository, sr repository.ISchoolRepository) *ContractUseCase {
	return &ContractUseCase{
		contractRepository: cou,
		childRepository:    cr,
		driverRepository:   dr,
		schoolRepository:   sr,
	}
}

// we create the contract, checking whether the person responsible has a payment method, calculating the distance between the school and the person responsible's residence, creating the product, the price and the signature on the stripe, and finally, creating the contract in the database
func (cou *ContractUseCase) Create(ctx context.Context, contract *entity.Contract) error {

	// get responsible data
	responsible, err := cou.childRepository.FindResponsibleByChild(ctx, &contract.Child.RG)
	if err != nil {
		return err
	}

	if responsible.PaymentMethodId == "" {
		return fmt.Errorf("responsible %s doesnt have a payment method", responsible.CPF)
	}

	contract.Child.Responsible = *responsible

	// get driver data
	driver, err := cou.driverRepository.Get(ctx, &contract.Driver.CNH)
	if err != nil {
		return err
	}

	contract.Driver = *driver

	// get school data
	school, err := cou.schoolRepository.Get(ctx, &contract.School.CNPJ)
	if err != nil {
		return err
	}

	contract.School = *school

	distance, err := usecase.GetDistance(fmt.Sprintf("%s,%s,%s", contract.Child.Responsible.Address.Street, contract.Child.Responsible.Address.Number, contract.Child.Responsible.Address.ZIP), fmt.Sprintf("%s,%s,%s", contract.School.Address.Street, contract.School.Address.Number, contract.School.Address.ZIP))
	if err != nil {
		return err
	}

	contract.Amount = usecase.CalculateContract(*distance, float64(contract.Driver.Amount))

	prodt, err := CreateProduct(contract)
	if err != nil {
		return err
	}

	contract.StripeSubscription.ProductSubscriptionId = prodt.ID

	pr, err := CreatePrice(contract)
	if err != nil {
		return err
	}

	contract.StripeSubscription.PriceSubscriptionId = pr.ID

	subs, err := CreateSubscription(contract)
	if err != nil {
		return err
	}

	contract.StripeSubscription.SubscriptionId = subs.ID

	id, err := uuid.NewV7()
	if err != nil {
		return err
	}

	contract.Record = id

	err = cou.contractRepository.Create(ctx, contract)
	if err != nil {
		// cancelo assinatura se registro não funcionar no banco
		_, err := DeleteSubscription(contract)

		if err != nil {
			return fmt.Errorf("erro ao criar contrato e ao deletar assinatura: %v", err)
		}

		return err
	}

	return nil

}

func (cou *ContractUseCase) Get(ctx context.Context, id uuid.UUID) (*entity.Contract, error) {

	contract, err := cou.contractRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	invoices, err := ListInvoices(contract)
	if err != nil {
		return nil, err
	}

	contract.Invoices = invoices

	return contract, nil

}

func (cou *ContractUseCase) FindAllByCnpj(ctx context.Context, cnpj *string) ([]entity.Contract, error) {
	return cou.contractRepository.FindAllByCnpj(ctx, cnpj)
}

func (cou *ContractUseCase) FindAllByCpf(ctx context.Context, cpf *string) ([]entity.Contract, error) {
	return cou.contractRepository.FindAllByCpf(ctx, cpf)
}

func (cou *ContractUseCase) FindAllByCnh(ctx context.Context, cnh *string) ([]entity.Contract, error) {

	contracts, err := cou.contractRepository.FindAllByCnh(ctx, cnh)

	if err != nil {
		return nil, err
	}

	for _, contract := range contracts {

		invoices, err := ListInvoices(&contract)
		if err != nil {
			return nil, err
		}

		contract.Invoices = invoices

	}

	return contracts, nil

}

func (cou *ContractUseCase) Cancel(ctx context.Context, id uuid.UUID) error {

	contract, err := cou.contractRepository.Get(ctx, id)
	if err != nil {
		return err
	}

	err = cou.contractRepository.Cancel(ctx, id)
	if err != nil {
		return err
	}

	values := CalculateRemainingValueSubscription(contract.Invoices)

	_, err = FineResponsible(contract, int64(values.Fines))

	if err != nil {
		return err
	}

	return nil

}

func (cou *ContractUseCase) GetInvoice(ctx context.Context, invoice *string) (*entity.InvoiceInfo, error) {

	inv, err := GetInvoice(*invoice)
	if err != nil {
		return nil, err
	}

	return &entity.InvoiceInfo{
		ID:              inv.ID,
		Status:          string(inv.Status),
		AmountDue:       inv.AmountDue,
		AmountRemaining: inv.AmountRemaining * 100,
	}, nil
}

func (cou *ContractUseCase) Expired(ctx context.Context, id uuid.UUID) error {
	return cou.contractRepository.Expired(ctx, id)
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
		UnitAmount: stripe.Int64(int64(contract.Amount) * 100),
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
