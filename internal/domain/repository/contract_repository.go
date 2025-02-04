package repository

import (
	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/entity"
)

type ContractRepository interface {
	Create(contract *entity.Contract) error
	Get(uuid uuid.UUID) (*entity.Contract, error)
	FindAllByCnpj(cnpj *string) ([]entity.Contract, error)
	FindAllByCpf(cpf *string) ([]entity.Contract, error)
	FindAllByCnh(cnh *string) ([]entity.Contract, error)
	Cancel(uuid uuid.UUID) error
	Expired(uuid uuid.UUID) error
	Update(uuid uuid.UUID, attributes map[string]interface{}) error
}
