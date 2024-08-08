package models

import "github.com/dgrijalva/jwt-go"

type Responsible struct {
	ID              int        `json:"id"`
	Name            string     `json:"name"`
	Email           string     `json:"email"`
	Password        string     `json:"password"`
	CPF             string     `json:"cpf"`
	Street          string     `json:"street"`
	Number          string     `json:"number"`
	ZIP             string     `json:"zip"`
	Complement      string     `json:"complement"`
	Status          string     `json:"status" validate:"oneof=OK ACTIVE BLOCKED BANNED"`
	CreditCard      CreditCard `json:"card"`
	CustomerId      string     `json:"customer_id,omitempty"`
	PaymentMethodId string     `json:"payment_method_id,omitempty"`
	Phone           string     `json:"phone"`
}

type CreditCard struct {
	CardToken string `json:"card_token"`
	CPF       string `json:"cpf"`
	Default   bool   `json:"default"`
}

type ClaimsResponsible struct {
	CPF string `json:"cpf"`
	jwt.StandardClaims
}
