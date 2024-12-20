package entity

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/venture-technology/venture/pkg/utils"
)

type Driver struct {
	ID           int     `json:"id,omitempty"`
	Name         string  `json:"name,omitempty" validate:"required"`
	Email        string  `json:"email,omitempty" validate:"required"`
	Password     string  `json:"password,omitempty"`
	CPF          string  `json:"cpf,omitempty"`
	CNH          string  `json:"cnh,omitempty" validate:"required"`
	QrCode       string  `json:"qrcode,omitempty"`
	Pix          Pix     `json:"pix,omitempty"`
	Bank         Bank    `json:"bank,omitempty"`
	Address      Address `json:"address,omitempty" validate:"required"`
	Amount       float64 `json:"amount,omitempty" validate:"required"`
	Phone        string  `json:"phone,omitempty" validate:"required" example:"+55 11 123456789"`
	MunicipalId  string  `json:"municipal_id,omitempty" validate:"required"`
	Car          Car     `json:"car,omitempty" validate:"required"`
	ProfileImage string  `json:"profile_image,omitempty"`
}

type Bank struct {
	Agency  string `json:"agency_number,omitempty" validate:"required"`
	Account string `json:"account_number,omitempty" validate:"required"`
	Name    string `json:"bank_name,omitempty" validate:"required"`
}

type Pix struct {
	Key string `json:"pix_key,omitempty" validate:"required"`
}

type Car struct {
	Model string `json:"model,omitempty" validate:"required"`
	Year  string `json:"year,omitempty" validate:"required"`
}

func (d *Driver) ValidateCnh() error {
	status := utils.IsCNH(d.CNH)

	if !status {
		return fmt.Errorf("invalid cnh")
	}

	return nil
}

func (d *Driver) HasPixOrBankAccount() bool {
	return d.Pix != (Pix{}) || d.Bank != (Bank{})
}

func (d *Driver) HasCar() bool {
	return d.Car != (Car{})
}

type ClaimsDriver struct {
	Driver Driver `json:"driver"`
	jwt.StandardClaims
}
