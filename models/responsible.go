package models

import "github.com/dgrijalva/jwt-go"

type Responsible struct {
	ID              int        `json:"id"`
	Name            string     `json:"name" validate:"required"`
	Email           string     `json:"email" validate:"required"`
	Password        string     `json:"password,omitempty" validate:"required"`
	CPF             string     `json:"cpf" validate:"required" example:"44000000000"` // sem pontuação
	Street          string     `json:"street" validate:"required"`
	Number          string     `json:"number" validate:"required"`
	ZIP             string     `json:"zip" validate:"required"`
	Complement      string     `json:"complement,omitempty"`
	Status          string     `json:"status" validate:"oneof='ok' 'active' 'blocked' 'banned'"`
	CreditCard      CreditCard `json:"card,omitempty"`
	CustomerId      string     `json:"customer_id,omitempty"`
	PaymentMethodId string     `json:"payment_method_id,omitempty"`
	Phone           string     `json:"phone" validate:"required" example:"+55 11 123456789"` // o telefone deve seguir este mesmo formato
}

type CreditCard struct {
	CardToken string `json:"card_token,omitempty"`
	Default   bool   `json:"default,omitempty"`
}

func (r *Responsible) IsCreditCardEmpty() bool {
	return r.CreditCard == (CreditCard{})
}

type ClaimsResponsible struct {
	CPF string `json:"cpf"`
	jwt.StandardClaims
}
