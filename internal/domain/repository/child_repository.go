package repository

import (
	"github.com/venture-technology/venture/internal/entity"
)

type IChildRepository interface {
	Create(child *entity.Child) error
	Get(rg *string) (*entity.Child, error)
	FindAll(cpf *string) ([]entity.Child, error)
	Update(child *entity.Child, attributes map[string]interface{}) error
	Delete(rg *string) error
}
