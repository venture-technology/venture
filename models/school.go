package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/venture-technology/venture/pkg/utils"
)

type School struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	CNPJ       string `json:"cnpj"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Street     string `json:"street"`
	Number     string `json:"number"`
	Complement string `json:"complement"`
	ZIP        string `json:"zip"`
}

type ClaimsSchool struct {
	CNPJ string `json:"cnpj"`
	jwt.StandardClaims
}

func (s *School) ValidateCnpj() bool {

	cnpj := utils.IsCNPJ(s.CNPJ)

	return cnpj

}
