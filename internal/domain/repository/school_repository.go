package repository

import (
	"github.com/venture-technology/venture/internal/entity"
)

type ISchoolRepository interface {
	Create(school *entity.School) error
	Get(cnpj string) (*entity.School, error)
	FindAll() ([]entity.School, error)
	Update(cnpj string, attributes map[string]interface{}) error
	Delete(cnpj string) error
	FindByEmail(email string) (*entity.School, error)
}
