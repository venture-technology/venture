package persistence

import (
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/pkg/realtime"
)

type ChildRepositoryImpl struct {
	Postgres contracts.PostgresIface
}

func (cr ChildRepositoryImpl) Create(child *entity.Child) error {
	return cr.Postgres.Client().Create(child).Error
}

func (cr ChildRepositoryImpl) Get(rg *string) (*entity.Child, error) {
	var child entity.Child
	err := cr.Postgres.Client().Where("rg = ?", *rg).First(&child).Error
	if err != nil {
		return nil, err
	}
	return &child, nil
}

func (cr ChildRepositoryImpl) FindAll(cpf *string) ([]entity.Child, error) {
	var children []entity.Child
	err := cr.Postgres.Client().Where("responsible_id = ?", *cpf).Find(&children).Error
	if err != nil {
		return nil, err
	}
	return children, nil
}

func (cr ChildRepositoryImpl) Update(rg string, attributes map[string]interface{}) error {
	attributes["updated_at"] = realtime.Now().UTC()

	err := cr.Postgres.Client().
		Model(&entity.Child{}).
		Where("rg = ?", rg).
		UpdateColumns(attributes).
		Error

	return err
}

func (cr ChildRepositoryImpl) Delete(rg *string) error {
	return cr.Postgres.Client().Where("rg = ?", *rg).Delete(&entity.Child{}).Error
}
