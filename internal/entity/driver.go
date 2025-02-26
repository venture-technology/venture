package entity

import (
	"fmt"
	"time"

	"github.com/venture-technology/venture/pkg/utils"
)

type Driver struct {
	ID              int       `gorm:"primaryKey;autoIncrement" json:"id,omitempty"`
	Name            string    `json:"name,omitempty" validate:"required"`
	Email           string    `json:"email,omitempty" validate:"required"`
	Password        string    `json:"password,omitempty"`
	CPF             string    `json:"cpf,omitempty"`
	CNH             string    `json:"cnh,omitempty" validate:"required"`
	QrCode          string    `json:"qrcode,omitempty"`
	PixKey          string    `json:"pix_key,omitempty"`
	Address         Address   `gorm:"embedded" json:"address,omitempty" validate:"required"`
	Amount          float64   `gorm:"type:numeric(10,2)" json:"amount,omitempty" validate:"required"`
	Phone           string    `json:"phone,omitempty" validate:"required" example:"+55 11 123456789"`
	MunicipalRecord string    `json:"municipal_record,omitempty" validate:"required"`
	Car             Car       `gorm:"embedded" json:"car,omitempty" validate:"required"`
	ProfileImage    string    `json:"profile_image,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	Schedule        string    `json:"schedule,omitempty"`
	Seats           Seats     `gorm:"embedded" json:"seats,omitempty"`
	Accessibility   bool      `json:"accessibility"`
	City            string    `json:"city"`
	States          string    `json:"states"`
}

type Car struct {
	Name     string `gorm:"column:car_name" json:"name,omitempty" validate:"required"`
	Year     string `gorm:"column:car_year" json:"year,omitempty" validate:"required"`
	Capacity int64  `gorm:"column:car_capacity" json:"capacity,omitempty" validate:"required"`
}

type Seats struct {
	Remaining int64 `gorm:"column:seats_remaining" json:"remaining,omitempty"`
	Morning   int64 `gorm:"column:seats_morning" json:"morning,omitempty"`
	Afternoon int64 `gorm:"column:seats_afternoon" json:"afternoon,omitempty"`
	Night     int64 `gorm:"column:seats_night" json:"night,omitempty"`
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
