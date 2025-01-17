package repository

import (
	"github.com/venture-technology/venture/internal/entity"
)

type PartnerRepository interface {
	Get(id string) (*entity.Partner, error)
	FindAllByCnpj(cnpj string) ([]entity.Partner, error)
	FindAllByCnh(cnh string) ([]entity.Partner, error)
	ArePartner(cnh, cnpj string) (bool, error)
	Delete(id string) error
}
