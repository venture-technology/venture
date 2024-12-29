package persistence

import (
	"fmt"

	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
)

type ResponsibleRepositoryImpl struct {
	Postgres contracts.PostgresIface
}

func (rr ResponsibleRepositoryImpl) Create(responsible *entity.Responsible) error {
	return rr.Postgres.Client().Create(responsible).Error
}

func (rr ResponsibleRepositoryImpl) Get(cpf string) (*entity.Responsible, error) {
	var responsible entity.Responsible
	err := rr.Postgres.Client().Where("cpf = ?", cpf).First(&responsible).Error
	if err != nil {
		return nil, err
	}
	return &responsible, nil
}

func (rr ResponsibleRepositoryImpl) Update(responsible *entity.Responsible, attributes map[string]interface{}) error {
	return rr.Postgres.Client().Model(responsible).Updates(attributes).Error
}

func (rr ResponsibleRepositoryImpl) Delete(cpf string) error {
	return rr.Postgres.Client().Delete(&entity.Responsible{}, "cpf = ?", cpf).Error
}

func (rr ResponsibleRepositoryImpl) SaveCard(cpf, cardToken, paymentMethodId string) error {
	return rr.Postgres.Client().Model(&entity.Responsible{}).Where("cpf = ?", cpf).Updates(map[string]interface{}{
		"card_token":        cardToken,
		"payment_method_id": paymentMethodId,
	}).Error
}

func (rr ResponsibleRepositoryImpl) Auth(responsible *entity.Responsible) (*entity.Responsible, error) {
	var responsibleData entity.Responsible
	err := rr.Postgres.Client().Where("email = ?", responsible.Email).First(&responsibleData).Error
	if err != nil {
		return nil, err
	}
	match := responsibleData.Password == responsible.Password
	if !match {
		return nil, fmt.Errorf("email or password wrong")
	}
	responsibleData.Password = ""
	return &responsibleData, nil
}

func (rr ResponsibleRepositoryImpl) FindByEmail(email string) (*entity.Responsible, error) {
	var responsible entity.Responsible
	err := rr.Postgres.Client().Where("email = ?", email).First(&responsible).Error
	if err != nil {
		return nil, err
	}
	return &responsible, nil
}
