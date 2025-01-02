package repository

import (
	"github.com/venture-technology/venture/internal/entity"
)

type IDriverRepository interface {
	Create(driver *entity.Driver) error
	Get(cnh string) (*entity.Driver, error)
	Update(driver *entity.Driver, attributes map[string]interface{}) error
	Delete(cnh string) error
	FindByEmail(email string) (*entity.Driver, error)
}
