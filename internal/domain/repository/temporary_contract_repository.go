package repository

import (
	"github.com/venture-technology/venture/internal/entity"
)

type TempContractRepository interface {
	Create(contract *entity.TempContract) error
	FindAllByResponsible(cpf *string) ([]entity.TempContract, error)
	//para verificar se todos os objetos do contrato existem
	//
	//{responsible, driver e kid}
	GetByEveryone(contract *entity.TempContract) (bool, error)
	FindAllByDriver(cnh *string) ([]entity.TempContract, error)
	Expire(uuid string) error
	Cancel(uuid string) error
	Update(uuid string, attrs map[string]interface{}) error
}

