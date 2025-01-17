package entity

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/venture-technology/venture/pkg/utils"
)

type Driver struct {
	ID           int       `gorm:"primary_key;auto_increment" json:"id,omitempty"`
	Name         string    `json:"name,omitempty" validate:"required"`
	Email        string    `json:"email,omitempty" validate:"required"`
	Password     string    `json:"password,omitempty"`
	CPF          string    `json:"cpf,omitempty"`
	CNH          string    `json:"cnh,omitempty" validate:"required"`
	QrCode       string    `json:"qrcode,omitempty"`
	Pix          Pix       `json:"pix,omitempty"`
	Address      Address   `gorm:"embedded" json:"address,omitempty" validate:"required"`
	Amount       float64   `json:"amount,omitempty" validate:"required"`
	Phone        string    `json:"phone,omitempty" validate:"required" example:"+55 11 123456789"`
	MunicipalId  string    `json:"municipal_id,omitempty" validate:"required"`
	Car          Car       `json:"car,omitempty" validate:"required"`
	ProfileImage string    `json:"profile_image,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
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

func (d *Driver) HasCar() bool {
	return d.Car != (Car{})
}

type ClaimsDriver struct {
	Driver Driver `json:"driver"`
	jwt.StandardClaims
}
