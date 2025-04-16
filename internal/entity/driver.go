package entity

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
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
	Descriptions    string    `json:"descriptions"`
	Biography       string    `json:"biography"`
}

type Car struct {
	Name     string `gorm:"column:car_name" json:"name,omitempty" validate:"required"`
	Year     string `gorm:"column:car_year" json:"year,omitempty" validate:"required"`
	Capacity uint   `gorm:"column:car_capacity" json:"capacity,omitempty" validate:"required"`
}

type Seats struct {
	Remaining uint  `gorm:"column:seats_remaining" json:"remaining,omitempty"`
	Morning   uint  `gorm:"column:seats_morning" json:"morning,omitempty"`
	Afternoon uint  `gorm:"column:seats_afternoon" json:"afternoon,omitempty"`
	Night     uint  `gorm:"column:seats_night" json:"night,omitempty"`
	Version   int64 `gorm:"column:seats_version" json:"version,omitempty"`
}

type ClaimsDriver struct {
	Driver Driver `json:"driver"`
	Role   string `json:"role"`
	jwt.StandardClaims
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

func (d *Driver) ValidateLogin() error {
	if d.Email == "" {
		return fmt.Errorf("email is required")
	}

	if d.Password == "" {
		return fmt.Errorf("password is required")
	}

	return nil
}

func (d *Driver) ValidateCapacity() error {
	if d.Car.Capacity < 30 {
		return fmt.Errorf("capacity must be greater than 30")
	}

	if d.Car.Capacity > 300 {
		return fmt.Errorf("capacity must be less than 300")
	}

	return nil
}
