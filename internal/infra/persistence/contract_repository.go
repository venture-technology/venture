package persistence

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/pkg/realtime"
)

type ContractRepositoryImpl struct {
	Postgres contracts.PostgresIface
}

func (cr ContractRepositoryImpl) Accept(contract *entity.Contract) error {
	tx := cr.Postgres.Client().Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		} else if tx.Error != nil {
			tx.Rollback()
		}
	}()

	// creation of contract
	if err := tx.Create(&contract).Error; err != nil {
		return err
	}

	var kid entity.Kid
	if err := tx.Where("rg = ?", contract.KidRG).First(&kid).Error; err != nil {
		return err
	}

	var driver entity.Driver
	if err := tx.Where("cnh = ?", contract.DriverCNH).First(&driver).Error; err != nil {
		return err
	}

	// update of driver's seats
	createSeats := map[string]func(driver entity.Driver) error{
		"morning": func(driver entity.Driver) error {
			attributes := make(map[string]interface{})
			attributes["updated_at"] = realtime.Now().UTC()
			attributes["seats_remaining"] = driver.Seats.Remaining - 1
			attributes["seats_morning"] = driver.Seats.Morning - 1
			return tx.Model(&entity.Driver{}).
				Where("cnh = ?", contract.DriverCNH).
				UpdateColumns(attributes).Error
		},
		"afternoon": func(driver entity.Driver) error {
			attributes := make(map[string]interface{})
			attributes["updated_at"] = realtime.Now().UTC()
			attributes["seats_remaining"] = driver.Seats.Remaining - 1
			attributes["seats_afternoon"] = driver.Seats.Afternoon - 1
			return tx.Model(&entity.Driver{}).
				Where("cnh = ?", contract.DriverCNH).
				UpdateColumns(attributes).Error
		},
		"night": func(driver entity.Driver) error {
			attributes := make(map[string]interface{})
			attributes["updated_at"] = realtime.Now().UTC()
			attributes["seats_remaining"] = driver.Seats.Remaining - 1
			attributes["seats_night"] = driver.Seats.Night - 1
			return tx.Model(&entity.Driver{}).
				Where("cnh = ?", contract.DriverCNH).
				UpdateColumns(attributes).Error
		},
	}

	updateFunc, exists := createSeats[kid.Shift]
	if !exists {
		return fmt.Errorf("invalid shift: %s", kid.Shift)
	}

	if err := updateFunc(driver); err != nil {
		return err
	}

	// change status of temp_contract
	if err := tx.Model(&entity.TempContract{}).
		Where("uuid = ?", contract.UUID).
		Update("status = accepted").Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

func (cr ContractRepositoryImpl) Get(id uuid.UUID) (*entity.Contract, error) {
	var contract entity.Contract
	err := cr.Postgres.Client().
		Where("record = ?", id).
		Preload("Driver", "cnh = ?", "driver_id").
		Preload("Kid", "rg = ?", "kid_id").
		Preload("Responsible", "cpf = ?", "responsible_id").
		Preload("School", "cnpj = ?", "school_id").
		First(&contract).Error

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

func (cr ContractRepositoryImpl) Update(id uuid.UUID, attributes map[string]interface{}) error {
	err := cr.Postgres.Client().Model(&entity.Contract{}).Where("record = ?", id).Updates(attributes).Error
	if err != nil {
		return err
	}
	return nil
}
