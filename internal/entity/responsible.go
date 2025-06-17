package entity

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Responsible struct {
	ID              int       `gorm:"primary_key;auto_increment" json:"id,omitempty"`
	Name            string    `json:"name,omitempty" validate:"required"`
	Email           string    `json:"email,omitempty" validate:"required"`
	Password        string    `json:"password,omitempty" validate:"required"`
	CPF             string    `json:"cpf,omitempty" validate:"required" example:"44000000000"`
	Address         Address   `gorm:"embedded" json:"address,omitempty" validate:"required"`
	CardTokenID     string    `json:"card_token,omitempty"` // CardTokenID is used when the Payment Method is CreditCard.
	CustomerId      string    `json:"customer_id,omitempty"`
	PaymentMethodId string    `json:"payment_method_id,omitempty"`
	Phone           string    `json:"phone,omitempty" validate:"required" example:"+55 11 123456789"` // The Number must be created with this format.
	ProfileImage    string    `json:"profile_image,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
}

type ClaimsResponsible struct {
	Responsible Responsible `json:"responsible"`
	Role        string      `json:"role"`
	jwt.StandardClaims
}

func (r *Responsible) ValidateLogin() error {
	if r.Email == "" {
		return fmt.Errorf("email is required")
	}

	if r.Password == "" {
		return fmt.Errorf("password is required")
	}

	return nil
}
