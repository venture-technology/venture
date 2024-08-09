package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/venture-technology/venture/pkg/utils"
)

type Driver struct {
	ID         int     `json:"id"`
	Amount     float64 `json:"amount"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Password   string  `json:"password"`
	CPF        string  `json:"cpf"`
	CNH        string  `json:"cnh"`
	QrCode     string  `json:"qrcode"`
	Street     string  `json:"street"`
	Number     string  `json:"number"`
	ZIP        string  `json:"zip"`
	Complement string  `json:"complement"`
	Price      float64 `json:"price"`
	Phone      string  `json:"phone" validate:"required" example:"+55 11 123456789"`
}

func (d *Driver) ValidateCnh() bool {
	return utils.IsCNH(d.CNH)
}

type ClaimsDriver struct {
	CNH string `json:"cnh"`
	jwt.StandardClaims
}
