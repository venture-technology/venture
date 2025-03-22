package persistence

import (
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

func (ir InviteRepositoryImpl) Get(id string) (*entity.Invite, error) {
	var invite entity.Invite
	err := ir.Postgres.Client().Where("id = ?", id).First(&invite).Error
	if err != nil {
		return nil, err
	}
	return &invite, nil
}

func (ir InviteRepositoryImpl) GetByDriver(cnh string) ([]entity.School, error) {
	var schools []entity.School

	err := ir.Postgres.Client().
		Select(`schools.id AS school_id, schools.name, schools.email, schools.phone,
            schools.profile_image, schools.cnpj, schools.created_at, schools.updated_at,
            schools.street, schools.number, schools.zip, schools.complement`).
		Table("invites").
		Joins("JOIN schools ON invites.requester = schools.cnpj").
		Where("invites.guester = ? AND invites.status = 'pending'", cnh).
		Find(&schools).Error

	return schools, err
}

func (ir InviteRepositoryImpl) GetBySchool(cnpj string) ([]entity.Driver, error) {
	var drivers []entity.Driver

	err := ir.Postgres.Client().
		Select(`
			drivers.id, drivers.name, drivers.email, drivers.phone, drivers.profile_image,
			drivers.cpf, drivers.cnh, drivers.qr_code, drivers.municipal_record, drivers.schedule,
			drivers.amount, drivers.created_at, drivers.updated_at,
			drivers.pix_key,
			drivers.street, drivers.number, drivers.zip, drivers.complement,
			drivers.car_name, drivers.car_year, drivers.car_capacity,
			drivers.seats_remaining, drivers.seats_morning, drivers.seats_afternoon, drivers.seats_night
		`).
		Table("invites").
		Joins("JOIN drivers ON invites.guester = drivers.cnh").
		Where("invites.requester = ? AND invites.status = 'pending'", cnpj).
		Find(&drivers).Error

	return drivers, err
}

func (ir InviteRepositoryImpl) Accept(id string) error {
	tx := ir.Postgres.Client().Begin() // Inicia a transação

	if err := tx.Model(&entity.Invite{}).Where("id = ?", id).Updates(map[string]interface{}{
		"status":      "accepted",
		"accepted_at": realtime.Now(),
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	var invite entity.Invite
	if err := tx.Where("id = ?", id).First(&invite).Error; err != nil {
		tx.Rollback()
		return err
	}

	partner := entity.Partner{
		SchoolID: invite.SchoolID,
		DriverID: invite.DriverID,
	}

	if err := tx.Create(&partner).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (ir InviteRepositoryImpl) Decline(id string) error {
	return ir.Postgres.Client().Where("id = ?", id).Delete(&entity.Invite{}).Error
}
