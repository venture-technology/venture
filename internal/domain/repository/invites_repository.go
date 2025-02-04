package repository

import (
	"github.com/venture-technology/venture/internal/entity"
)

type InviteRepository interface {
	Create(invite *entity.Invite) error
	Get(id string) (*entity.Invite, error)
	FindAllByCnh(cnh string) ([]entity.School, error)
	FindAllByCnpj(cnpj string) ([]entity.Driver, error)
	Accept(id string) error
	Decline(id string) error
}
