package entity

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/venture-technology/venture/pkg/utils"
)

type School struct {
	ID       int     `json:"id,omitempty"`
	Name     string  `json:"name,omitempty" validate:"required"`
	CNPJ     string  `json:"cnpj" validate:"required"`
	Email    string  `json:"email,omitempty" validate:"required"`
	Password string  `json:"password,omitempty" validate:"required"`
	Address  Address `json:"address,omitempty" validate:"required"`
	Phone    string  `json:"phone,omitempty" validate:"required" example:"+55 11 123456789"`
}

type ClaimsSchool struct {
	CNPJ string `json:"cnpj"`
	jwt.StandardClaims
}

func (s *School) ValidateCnpj() bool {

	cnpj := utils.IsCNPJ(s.CNPJ)

	return cnpj

}
