package repository

import (
	"github.com/venture-technology/venture/internal/entity"
)

type DriverRepository interface {
	Create(driver *entity.Driver) error
	Get(cnh string) (*entity.Driver, error)
	Update(cnh string, attributes map[string]interface{}) error
	Delete(cnh string) error
	GetByEmail(email string) (*entity.Driver, error)
}
