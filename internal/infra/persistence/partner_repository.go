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
	if err != nil {
		return nil, err
	}

	for i := range partners {
		var driver entity.Driver
		err := pr.Postgres.Client().Where("cnh = ?", partners[i].DriverID).First(&driver).Error
		if err != nil {
			continue
		}
		partners[i].Driver = driver
	}

	return partners, nil
}

func (pr PartnerRepositoryImpl) FindAllByCnh(cnh string) ([]entity.Partner, error) {
	var partners []entity.Partner

	err := pr.Postgres.Client().Where("driver_id = ?", cnh).Find(&partners).Error
	if err != nil {
		return nil, err
	}

	for i := range partners {
		var school entity.School
		err := pr.Postgres.Client().Where("cnpj = ?", partners[i].SchoolID).First(&school).Error
		if err != nil {
			continue
		}
		partners[i].School = school
	}

	return partners, nil
}

func (pr PartnerRepositoryImpl) Delete(id string) error {
	tx := pr.Postgres.Client().Begin()
	if tx.Error != nil {
		return tx.Error
	}

	var partner entity.Partner
	if err := tx.Where("record = ?", id).First(&partner).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("requester = ? AND guester = ?", partner.SchoolID, partner.DriverID).
		Delete(&entity.Invite{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&entity.Partner{}, "record = ?", id).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
