package contract

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/domain/payments"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/repository"
	"github.com/venture-technology/venture/internal/usecase"
)

type ContractUseCase struct {
	contractRepository repository.IContractRepository
	childRepository    repository.IChildRepository
	driverRepository   repository.IDriverRepository
	schoolRepository   repository.ISchoolRepository
	stripe             payments.IStripe
}

func NewContractUseCase(cou repository.IContractRepository, cr repository.IChildRepository, dr repository.IDriverRepository, sr repository.ISchoolRepository, st payments.IStripe) *ContractUseCase {
	return &ContractUseCase{
		contractRepository: cou,
		childRepository:    cr,
		driverRepository:   dr,
		schoolRepository:   sr,
		stripe:             st,
	}
}

// we create the contract, checking whether the person responsible has a payment method, calculating the distance between the school and the person responsible's residence, creating the product, the price and the signature on the stripe, and finally, creating the contract in the database
func (cou *ContractUseCase) Create(ctx context.Context, contract *entity.Contract) error {

	contract.StripeSubscription.Title = fmt.Sprintf("%s - %s - %s - %s", contract.Driver.Name, contract.School.Name, contract.Child.Responsible.Name, contract.Child.Name)

	simpleContract, err := cou.contractRepository.GetSimpleContractByTitle(ctx, &contract.StripeSubscription.Title)
	if err != nil {
		return err
	}

	if simpleContract.StripeSubscription.Title == contract.StripeSubscription.Title {
		return fmt.Errorf("contract already exists")
	}

	hasPaymentMethod := contract.Child.Responsible.HasPaymentMethod()

	log.Print(hasPaymentMethod, contract.Child.Responsible.PaymentMethodId)

	if !hasPaymentMethod {
		return fmt.Errorf("responsible %s doesnt have a payment method", contract.Child.Responsible.CPF)
	}

	hasPixOrBankAccount := contract.Driver.HasPixOrBankAccount()

	if !hasPixOrBankAccount {
		return fmt.Errorf("driver %s need pix or bank account register", contract.Driver.CNH)
	}

	hasCar := contract.Driver.HasCar()

	if !hasCar {
		return fmt.Errorf("driver %s need car register", contract.Driver.CNH)
	}

	distance, err := usecase.GetDistance(fmt.Sprintf("%s,%s,%s", contract.Child.Responsible.Address.Street, contract.Child.Responsible.Address.Number, contract.Child.Responsible.Address.ZIP), fmt.Sprintf("%s,%s,%s", contract.School.Address.Street, contract.School.Address.Number, contract.School.Address.ZIP))
	if err != nil {
		return err
	}

	contract.Amount = usecase.CalculateContract(*distance, float64(contract.Driver.Amount))

	prodt, err := cou.stripe.CreateProduct(contract)
	if err != nil {
		return err
	}

	contract.StripeSubscription.ProductSubscriptionId = prodt.ID

	pr, err := cou.stripe.CreatePrice(contract)
	if err != nil {
		return err
	}

	contract.StripeSubscription.PriceSubscriptionId = pr.ID

	subs, err := cou.stripe.CreateSubscription(contract)
	if err != nil {
		return err
	}

	contract.StripeSubscription.SubscriptionId = subs.ID

	id, err := uuid.NewV7()
	if err != nil {
		return err
	}

	contract.Record = id

	log.Print(contract)

	err = cou.contractRepository.Create(ctx, contract)
	log.Print(err)
	if err != nil {
		return err
	}

	return nil
}

func (cou *ContractUseCase) Get(ctx context.Context, id uuid.UUID) (*entity.Contract, error) {
	contract, err := cou.contractRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	invoices, err := cou.stripe.ListInvoices(contract)
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

		invoices, err := cou.stripe.ListInvoices(&contract)
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

	values := cou.stripe.CalculateRemainingValueSubscription(contract.Invoices)

	_, err = cou.stripe.FineResponsible(contract, int64(values.Fines))
	if err != nil {
		return err
	}

	return nil
}

func (cou *ContractUseCase) GetInvoice(ctx context.Context, invoice *string) (*entity.InvoiceInfo, error) {
	inv, err := cou.stripe.GetInvoice(*invoice)
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
