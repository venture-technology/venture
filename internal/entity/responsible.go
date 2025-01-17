package entity

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Responsible struct {
	ID              int        `gorm:"primary_key;auto_increment" json:"id,omitempty"`
	Name            string     `json:"name,omitempty" validate:"required"`
	Email           string     `json:"email,omitempty" validate:"required"`
	Password        string     `json:"password,omitempty" validate:"required"`
	CPF             string     `json:"cpf,omitempty" validate:"required" example:"44000000000"` // sem pontuação
	Address         Address    `gorm:"embedded" json:"address,omitempty" validate:"required"`
	CreditCard      CreditCard `json:"card,omitempty"`
	CustomerId      string     `json:"customer_id,omitempty"`
	PaymentMethodId string     `json:"payment_method_id,omitempty"`
	Phone           string     `json:"phone,omitempty" validate:"required" example:"+55 11 123456789"` // o telefone deve seguir este mesmo formato
	ProfileImage    string     `json:"profile_image,omitempty"`
	CreatedAt       time.Time  `json:"created_at,omitempty"`
	UpdatedAt       time.Time  `json:"updated_at,omitempty"`
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
