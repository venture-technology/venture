package persistence

import (
	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
)

type ContractRepositoryImpl struct {
	Postgres contracts.PostgresIface
}

func (cr ContractRepositoryImpl) Create(contract *entity.Contract) error {
	err := cr.Postgres.Client().Create(&contract).Error
	if err != nil {
		return err
	}
	return nil
}

func (cr ContractRepositoryImpl) Get(id uuid.UUID) (*entity.Contract, error) {
	var contract entity.Contract
	err := cr.Postgres.Client().Where("record = ?", id).First(&contract).Error
	if err != nil {
		return nil, err
	}
	return &contract, nil
}

func (cr ContractRepositoryImpl) FindAllByCnpj(cnpj *string) ([]entity.Contract, error) {
	var contracts []entity.Contract
	err := cr.Postgres.Client().Where("school_id = ?", cnpj).Find(&contracts).Error
	if err != nil {
		return nil, err
	}
	return contracts, nil
}

func (cr ContractRepositoryImpl) FindAllByCpf(cpf *string) ([]entity.Contract, error) {
	var contracts []entity.Contract
	err := cr.Postgres.Client().Where("responsible_id = ?", cpf).Find(&contracts).Error
	if err != nil {
		return nil, err
	}
	return contracts, nil
}

func (cr ContractRepositoryImpl) FindAllByCnh(cnh *string) ([]entity.Contract, error) {
	var contracts []entity.Contract
	err := cr.Postgres.Client().Where("driver_id = ?", cnh).Find(&contracts).Error
	if err != nil {
		return nil, err
	}
	return contracts, nil
}

func (cr ContractRepositoryImpl) Cancel(id uuid.UUID) error {
	err := cr.Postgres.Client().Model(&entity.Contract{}).Where("record = ?", id).Update("status", "canceled").Error
	if err != nil {
		return err
	}
	return nil
}

func (cr ContractRepositoryImpl) Expired(id uuid.UUID) error {
	err := cr.Postgres.Client().Model(&entity.Contract{}).Where("record = ?", id).Update("status", "expired").Error
	if err != nil {
		return err
	}
	return nil
}

func (cr ContractRepositoryImpl) GetSimpleContractByTitle(title *string) (*entity.Contract, error) {
	var contract entity.Contract
	err := cr.Postgres.Client().Where("title_stripe_subscription = ?", *title).First(&contract).Error
	if err != nil {
		return nil, err
	}
	return &contract, nil
}

func (cr ContractRepositoryImpl) Update(id uuid.UUID, attributes map[string]interface{}) error {
	err := cr.Postgres.Client().Model(&entity.Contract{}).Where("record = ?", id).Updates(attributes).Error
	if err != nil {
		return err
	}
	return nil
}
