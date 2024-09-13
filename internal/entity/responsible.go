package entity

import "github.com/dgrijalva/jwt-go"

type Responsible struct {
	ID              int        `json:"id,omitempty"`
	Name            string     `json:"name,omitempty" validate:"required"`
	Email           string     `json:"email,omitempty" validate:"required"`
	Password        string     `json:"password,omitempty" validate:"required"`
	CPF             string     `json:"cpf,omitempty" validate:"required" example:"44000000000"` // sem pontuação
	Status          string     `json:"status,omitempty" validate:"oneof='ok' 'active' 'blocked' 'banned'"`
	Address         Address    `json:"address,omitempty" validate:"required"`
	CreditCard      CreditCard `json:"card,omitempty"`
	CustomerId      string     `json:"customer_id,omitempty"`
	PaymentMethodId string     `json:"payment_method_id,omitempty"`
	Phone           string     `json:"phone,omitempty" validate:"required" example:"+55 11 123456789"` // o telefone deve seguir este mesmo formato
}

type CreditCard struct {
	CardToken string `json:"card_token,omitempty"`
	Default   bool   `json:"default,omitempty"`
}

func (r *Responsible) IsCreditCardEmpty() bool {
	return r.CreditCard == (CreditCard{})
}

func (r *Responsible) HasPaymentMethod() bool {
	return r.PaymentMethodId != ""
}

type ClaimsResponsible struct {
	Responsible Responsible `json:"responsible"`
	jwt.StandardClaims
}
