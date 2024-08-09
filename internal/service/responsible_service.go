package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/customer"
	"github.com/stripe/stripe-go/v79/paymentmethod"
	"github.com/venture-technology/venture/config"
	"github.com/venture-technology/venture/internal/repository"
	"github.com/venture-technology/venture/models"
	"github.com/venture-technology/venture/pkg/utils"
)

type ResponsibleService struct {
	responsiblerepository repository.IResponsibleRepository
}

func NewResponsibleService(responsiblerepository repository.IResponsibleRepository) *ResponsibleService {
	return &ResponsibleService{
		responsiblerepository: responsiblerepository,
	}
}

func (rs *ResponsibleService) CreateResponsible(ctx context.Context, responsible *models.Responsible) error {

	responsible.Password = utils.HashPassword(responsible.Password)

	return rs.responsiblerepository.CreateResponsible(ctx, responsible)
}

func (rs *ResponsibleService) GetResponsible(ctx context.Context, cpf *string) (*models.Responsible, error) {
	log.Printf("param read school -> cpf: %s", *cpf)
	return rs.responsiblerepository.GetResponsible(ctx, cpf)
}

func (rs *ResponsibleService) UpdateResponsible(ctx context.Context, currentResponsible, responsible *models.Responsible) error {
	log.Printf("input received to update school -> name: %s, cpf: %s, email: %s", responsible.Name, responsible.CPF, responsible.Email)
	return rs.responsiblerepository.UpdateResponsible(ctx, currentResponsible, responsible)
}

func (rs *ResponsibleService) DeleteResponsible(ctx context.Context, cpf *string) error {
	log.Printf("trying delete your infos --> %v", *cpf)
	return rs.responsiblerepository.DeleteResponsible(ctx, cpf)
}

func (rs *ResponsibleService) AuthResponsible(ctx context.Context, responsible *models.Responsible) (*models.Responsible, error) {
	responsible.Password = utils.HashPassword((responsible.Password))
	return rs.responsiblerepository.AuthResponsible(ctx, responsible)
}

func (rs *ResponsibleService) ParserJwtResponsible(ctx *gin.Context) (interface{}, error) {

	cpf, found := ctx.Get("cpf")

	if !found {
		return nil, fmt.Errorf("error while veryfing token")
	}

	return cpf, nil

}

func (rs *ResponsibleService) CreateTokenJWTResponsible(ctx context.Context, responsible *models.Responsible) (string, error) {

	conf := config.Get()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"cpf": responsible.CPF,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	jwt, err := token.SignedString([]byte(conf.Server.Secret))

	if err != nil {
		return "", err
	}

	return jwt, nil

}

func (rs *ResponsibleService) CreateCustomer(ctx context.Context, responsible *models.Responsible) (*stripe.Customer, error) {

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

func (rs *ResponsibleService) UpdateCustomer(ctx context.Context, responsible *models.Responsible) (*stripe.Customer, error) {

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

func (rs *ResponsibleService) DeleteCustomer(ctx context.Context, customerId string) (*stripe.Customer, error) {

	conf := config.Get()

	stripe.Key = conf.StripeEnv.SecretKey

	c, err := customer.Del(customerId, nil)
	if err != nil {
		return nil, err
	}

	return c, nil

}

func (rs *ResponsibleService) CreatePaymentMethod(ctx context.Context, cardToken *string) (*stripe.PaymentMethod, error) {

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

func (rs *ResponsibleService) AttachPaymentMethod(ctx context.Context, customerId, paymentMethodId *string, isDefault bool) (*stripe.PaymentMethod, error) {

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

func (rs *ResponsibleService) UpdatePaymentMethodDefault(ctx context.Context, customerId, paymentMethodId *string) (*stripe.Customer, error) {

	conf := config.Get()

	stripe.Key = conf.StripeEnv.SecretKey

	params := &stripe.CustomerParams{
		InvoiceSettings: &stripe.CustomerInvoiceSettingsParams{
			DefaultPaymentMethod: stripe.String(*paymentMethodId),
		},
	}

	updatedCustomer, err := customer.Update(*customerId, params)

	if err != nil {
		return nil, err
	}

	return updatedCustomer, nil

}

func (rs *ResponsibleService) SaveCreditCard(ctx context.Context, cpf, cardToken, paymentMethodId *string) error {
	return rs.responsiblerepository.SaveCreditCard(ctx, cpf, cardToken, paymentMethodId)
}
