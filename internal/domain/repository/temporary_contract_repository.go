package repository

import (
	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/entity"
)

type TempContractRepository interface {
	Create(contract *entity.TempContract) error
	Get(uuid uuid.UUID) (*entity.TempContract, error)
	FindAllByResponsible(cpf *string) ([]entity.TempContract, error)
	//para verificar se todos os objetos do contrato existem
	//
	//{responsible, driver e kid}
	GetByEveryone(contract *entity.TempContract) (bool, error)
	Expire(uuid uuid.UUID) error
	Cancel(uuid uuid.UUID) error
	Accept(uuid uuid.UUID) error
}
