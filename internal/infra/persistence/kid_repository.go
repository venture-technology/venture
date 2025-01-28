package persistence

import (
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/pkg/realtime"
)

type KidRepositoryImpl struct {
	Postgres contracts.PostgresIface
}

func (cr KidRepositoryImpl) Create(kid *entity.Kid) error {
	return cr.Postgres.Client().Create(kid).Error
}

func (cr KidRepositoryImpl) Get(rg *string) (*entity.Kid, error) {
	var kid entity.Kid
	err := cr.Postgres.Client().
		Preload("Responsible").
		Where("rg = ?", *rg).
		First(&kid).Error
	if err != nil {
		return nil, err
	}
	return &kid, nil
}
func (cr KidRepositoryImpl) FindAll(cpf *string) ([]entity.Kid, error) {
	var kids []entity.Kid
	err := cr.Postgres.Client().Where("responsible_id = ?", *cpf).Find(&kids).Error
	if err != nil {
		return nil, err
	}
	return kids, nil
}

func (cr KidRepositoryImpl) Update(rg string, attributes map[string]interface{}) error {
	attributes["updated_at"] = realtime.Now().UTC()

	err := cr.Postgres.Client().
		Model(&entity.Kid{}).
		Where("rg = ?", rg).
		UpdateColumns(attributes).
		Error

	return err
}

func (cr KidRepositoryImpl) Delete(rg *string) error {
	return cr.Postgres.Client().Where("rg = ?", *rg).Delete(&entity.Kid{}).Error
}
