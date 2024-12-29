package persistence

import (
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
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

func (sr SchoolRepositoryImpl) Update(school *entity.School, attributes map[string]interface{}) error {
	return sr.Postgres.Client().Model(school).Updates(attributes).Error
}

func (sr SchoolRepositoryImpl) Delete(cnpj string) error {
	return sr.Postgres.Client().Delete(&entity.School{}, "cnpj = ?", cnpj).Error
}

func (sr SchoolRepositoryImpl) FindByEmail(email string) (*entity.School, error) {
	var school entity.School
	err := sr.Postgres.Client().Where("email = ?", email).First(&school).Error
	if err != nil {
		return nil, err
	}
	return &school, nil
}
