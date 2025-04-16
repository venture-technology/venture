package repository

import (
	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/entity"
)

type TempContractRepository interface {
	Create(contract *entity.TempContract) error
	GetByResponsible(cpf *string) ([]entity.TempContract, error)
	//para verificar se todos os objetos do contrato existem
	//
	//{responsible, driver e kid}
	GetByEveryone(contract *entity.TempContract) (bool, error)
	GetByDriver(cnh *string) ([]entity.TempContract, error)
	Cancel(uuid uuid.UUID) error
	Update(uuid uuid.UUID, attrs map[string]interface{}) error
	Expire(uuids []uuid.UUID) error
	GetExpiredContracts() ([]entity.TempContract, error)
}
