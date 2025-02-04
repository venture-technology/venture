package repository

import (
	"github.com/venture-technology/venture/internal/entity"
)

type KidRepository interface {
	Create(kid *entity.Kid) error
	Get(rg *string) (*entity.Kid, error)
	FindAll(cpf *string) ([]entity.Kid, error)
	Update(rg string, attributes map[string]interface{}) error
	Delete(rg *string) error
}
