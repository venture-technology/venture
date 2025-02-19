package entity

import (
	"time"

	"github.com/venture-technology/venture/pkg/utils"
)

type School struct {
	ID           int       `gorm:"primary_key;auto_increment" json:"id,omitempty"`
	Name         string    `json:"name,omitempty" validate:"required"`
	CNPJ         string    `json:"cnpj,omitempty" validate:"required"`
	Email        string    `json:"email,omitempty" validate:"required"`
	Password     string    `json:"password,omitempty" validate:"required"`
	Address      Address   `gorm:"embedded" json:"address,omitempty" validate:"required"`
	Phone        string    `json:"phone,omitempty" validate:"required" example:"+55 11 123456789"`
	ProfileImage string    `json:"profile_image,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
}

func (s *School) ValidateCnpj() bool {
	cnpj := utils.IsCNPJ(s.CNPJ)

	return cnpj
}
