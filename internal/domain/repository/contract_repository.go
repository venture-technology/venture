package repository

import (
	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/entity"
)

type ContractRepository interface {
	Accept(contract *entity.Contract) error
	Cancel(uuid uuid.UUID) error
	Expired(uuid uuid.UUID) error
	Update(uuid uuid.UUID, attributes map[string]interface{}) error
	GetByUUID(id uuid.UUID) (*entity.Contract, error)
	GetBySchool(cnpj string) ([]entity.Contract, error)
	GetByDriver(cnh string) ([]entity.Contract, error)
	GetByResponsible(cpf string) ([]entity.Contract, error)

	// Check if a contract already exists
	ContractAlreadyExist(uuid string) (bool, error)
}
