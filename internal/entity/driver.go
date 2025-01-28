package entity

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/venture-technology/venture/pkg/utils"
)

type Driver struct {
	ID           int       `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	Name         string    `json:"name,omitempty" validate:"required"`
	Email        string    `json:"email,omitempty" validate:"required"`
	Password     string    `json:"password,omitempty"`
	CPF          string    `json:"cpf,omitempty"`
	CNH          string    `json:"cnh,omitempty" validate:"required"`
	QrCode       string    `json:"qrcode,omitempty"`
	Pix          Pix       `gorm:"embedded" json:"pix,omitempty"`
	Address      Address   `gorm:"embedded" json:"address,omitempty" validate:"required"`
	Amount       float64   `gorm:"type:numeric(10,2)" json:"amount,omitempty" validate:"required"`
	Phone        string    `json:"phone,omitempty" validate:"required" example:"+55 11 123456789"`
	MunicipalID  string    `json:"municipal_id,omitempty" validate:"required"`
	Car          Car       `gorm:"embedded" json:"car,omitempty" validate:"required"`
	ProfileImage string    `json:"profile_image,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	Schedule     string    `json:"schedule,omitempty"`
	Seats        Seats     `gorm:"embedded" json:"seats,omitempty"`
}

type Pix struct {
	Key string `json:"pix_key,omitempty" validate:"required"`
}

type Car struct {
	Name     string `json:"name,omitempty" validate:"required"`
	Year     string `json:"year,omitempty" validate:"required"`
	Capacity int64  `json:"capacity,omitempty" validate:"required"`
}

type Seats struct {
	Remaining int64 `json:"remaining,omitempty"`
	Morning   int64 `json:"morning,omitempty"`
	Afternoon int64 `json:"afternoon,omitempty"`
	Night     int64 `json:"night,omitempty"`
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
