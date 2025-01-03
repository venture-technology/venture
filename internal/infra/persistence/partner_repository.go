package persistence

import (
	"github.com/jinzhu/gorm"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
)

type PartnerRepositoryImpl struct {
	Postgres contracts.PostgresIface
}

func (pr PartnerRepositoryImpl) ArePartner(cnh, cnpj string) (bool, error) {
	var partner entity.Partner
	err := pr.Postgres.Client().Where("driver_id = ? AND school_id = ?", cnh, cnpj).First(&partner).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	return err == gorm.ErrRecordNotFound, nil
}

func (pr PartnerRepositoryImpl) Get(id string) (*entity.Partner, error) {
	var partner entity.Partner
	err := pr.Postgres.Client().Where("record = ?", id).First(&partner).Error
	if err != nil {
		return nil, err
	}
	return &partner, nil
}

func (pr PartnerRepositoryImpl) FindAllByCnpj(cnpj string) ([]entity.Partner, error) {
	var partners []entity.Partner
	err := pr.Postgres.Client().Where("school_id = ?", cnpj).Find(&partners).Error
	return partners, err
}

func (pr PartnerRepositoryImpl) FindAllByCnh(cnh string) ([]entity.Partner, error) {
	var partners []entity.Partner
	err := pr.Postgres.Client().Where("driver_id = ?", cnh).Find(&partners).Error
	return partners, err
}

func (pr PartnerRepositoryImpl) Delete(id string) error {
	return pr.Postgres.Client().Delete(&entity.Partner{}, "record = ?", id).Error
}
