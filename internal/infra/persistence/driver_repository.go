package persistence

import (
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/pkg/realtime"
)

type DriverRepositoryImpl struct {
	Postgres contracts.PostgresIface
}

func (dr DriverRepositoryImpl) Create(driver *entity.Driver) error {
	return dr.Postgres.Client().Create(driver).Error
}

func (dr DriverRepositoryImpl) Get(cnh string) (*entity.Driver, error) {
	var driver entity.Driver
	result := dr.Postgres.Client().Where("cnh = ?", cnh).First(&driver)
	if result.Error != nil {
		return nil, result.Error
	}
	return &driver, nil
}

func (dr DriverRepositoryImpl) Update(cnh string, attributes map[string]interface{}) error {
	attributes["updated_at"] = realtime.Now().UTC()

	err := dr.Postgres.Client().
		Model(&entity.Driver{}).
		Where("cnh = ?", cnh).
		UpdateColumns(attributes).
		Error

	return err
}

func (dr DriverRepositoryImpl) Delete(cnh string) error {
	return dr.Postgres.Client().Where("cnh = ?", cnh).Delete(&entity.Driver{}).Error
}

func (dr DriverRepositoryImpl) GetByEmail(email string) (*entity.Driver, error) {
	var driver entity.Driver
	result := dr.Postgres.Client().Where("email = ?", email).First(&driver)
	return &driver, result.Error
}
