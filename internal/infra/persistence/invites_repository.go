package persistence

import (
	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/pkg/realtime"
)

type InviteRepositoryImpl struct {
	Postgres contracts.PostgresIface
}

func (ir InviteRepositoryImpl) Create(invite *entity.Invite) error {
	return ir.Postgres.Client().Create(invite).Error
}

func (ir InviteRepositoryImpl) Get(uuid uuid.UUID) (*entity.Invite, error) {
	var invite entity.Invite
	err := ir.Postgres.Client().Where("uuid = ?", uuid).First(&invite).Error
	if err != nil {
		return nil, err
	}
	return &invite, nil
}

func (ir InviteRepositoryImpl) FindAllByCnh(cnh string) ([]entity.Invite, error) {
	var invites []entity.Invite
	err := ir.Postgres.Client().Where("guester = ? AND status = 'pending'", cnh).Find(&invites).Error
	return invites, err
}

func (ir InviteRepositoryImpl) FindAllByCnpj(cnpj string) ([]entity.Invite, error) {
	var invites []entity.Invite
	err := ir.Postgres.Client().Where("requester = ? AND status = 'pending'", cnpj).Find(&invites).Error
	return invites, err
}

func (ir InviteRepositoryImpl) Accept(uuid uuid.UUID) error {
	return ir.Postgres.Client().Model(&entity.Invite{}).Where("uuid = ?", uuid).Updates(map[string]interface{}{
		"status":      "accepted",
		"accepted_at": realtime.Now(),
	}).Error
}

func (ir InviteRepositoryImpl) Decline(uuid uuid.UUID) error {
	return ir.Postgres.Client().Model(&entity.Invite{}).Where("uuid = ?", uuid).Update("status", "declined").Error
}
