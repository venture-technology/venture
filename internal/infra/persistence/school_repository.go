package persistence

import (
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/pkg/realtime"
)

type SchoolRepositoryImpl struct {
	Postgres contracts.PostgresIface
}

func (sr SchoolRepositoryImpl) Create(school *entity.School) error {
	return sr.Postgres.Client().Create(school).Error
}

func (sr SchoolRepositoryImpl) Get(cnpj string) (*entity.School, error) {
	var school entity.School
	err := sr.Postgres.Client().Where("cnpj = ?", cnpj).First(&school).Error
	if err != nil {
		return nil, err
	}
	return &school, nil
}

func (sr SchoolRepositoryImpl) FindAll() ([]entity.School, error) {
	var schools []entity.School
	err := sr.Postgres.Client().Find(&schools).Error
	return schools, err
}

func (sr SchoolRepositoryImpl) Update(cnpj string, attributes map[string]interface{}) error {
	attributes["updated_at"] = realtime.Now().UTC()

	err := sr.Postgres.Client().
		Model(&entity.School{}).
		Where("cnpj = ?", cnpj).
		UpdateColumns(attributes).
		Error

	return err
}

func (sr SchoolRepositoryImpl) Delete(cnpj string) error {
	return sr.Postgres.Client().Delete(&entity.School{}, "cnpj = ?", cnpj).Error
}

func (sr SchoolRepositoryImpl) GetByEmail(email string) (*entity.School, error) {
	var school entity.School
	err := sr.Postgres.Client().Where("email = ?", email).First(&school).Error
	if err != nil {
		return nil, err
	}
	return &school, nil
}
