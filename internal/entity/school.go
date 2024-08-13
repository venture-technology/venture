package entity

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/venture-technology/venture/pkg/utils"
)

type School struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	CNPJ     string  `json:"cnpj"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Address  Address `json:"address" validate:"required"`
	Phone    string  `json:"phone" validate:"required" example:"+55 11 123456789"`
}

type ClaimsSchool struct {
	CNPJ string `json:"cnpj"`
	jwt.StandardClaims
}

func (s *School) ValidateCnpj() bool {

	cnpj := utils.IsCNPJ(s.CNPJ)

	return cnpj

}
