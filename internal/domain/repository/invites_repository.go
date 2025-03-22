package repository

import (
	"github.com/venture-technology/venture/internal/entity"
)

type InviteRepository interface {
	Create(invite *entity.Invite) error
	Get(id string) (*entity.Invite, error)
	GetByDriver(cnh string) ([]entity.School, error)
	GetBySchool(cnpj string) ([]entity.Driver, error)
	Accept(id string) error
	Decline(id string) error
}
