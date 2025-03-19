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
	GetByUUID(id uuid.UUID) (*entity.EnableContract, error)
	GetBySchool(cnpj string) ([]entity.EnableContract, error)
	GetByDriver(cnh string) ([]entity.EnableContract, error)
	GetByResponsible(cpf string) ([]entity.EnableContract, error)
}
