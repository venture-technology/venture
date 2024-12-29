package repository

import (
	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/entity"
)

type IInviteRepository interface {
	Create(invite *entity.Invite) error
	Get(uuid uuid.UUID) (*entity.Invite, error)
	FindAllByCnh(cnh string) ([]entity.Invite, error)
	FindAllByCnpj(cnpj string) ([]entity.Invite, error)
	Accept(uuid uuid.UUID) error
	Decline(uuid uuid.UUID) error
}
