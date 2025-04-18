package persistence

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/value"
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

	// valida seats > 0
	if driver.Seats.Remaining <= 0 {
		return fmt.Errorf("driver %s has no remaining seats", driver.CNH)
	}

	// update de acordo com o turno
	createSeats := map[string]func(driver entity.Driver) error{
		value.MorningShift: func(driver entity.Driver) error {
			if driver.Seats.Morning <= 0 {
				return fmt.Errorf("no morning seats available for driver %s", driver.CNH)
			}
			attributes := map[string]interface{}{
				"updated_at":      realtime.Now().UTC(),
				"seats_remaining": driver.Seats.Remaining - 1,
				"seats_morning":   driver.Seats.Morning - 1,
				"seats_version":   driver.Seats.Version + 1,
			}
			result := tx.Model(&entity.Driver{}).
				Where("cnh = ? AND seats_version = ?", contract.DriverCNH, driver.Seats.Version).
				UpdateColumns(attributes)
			if result.Error != nil {
				return result.Error
			}
			if result.RowsAffected == 0 {
				return fmt.Errorf("seats_version mismatch: possible concurrent update")
			}
			return nil
		},
		value.AfternoonShift: func(driver entity.Driver) error {
			if driver.Seats.Afternoon <= 0 {
				return fmt.Errorf("no afternoon seats available for driver %s", driver.CNH)
			}
			attributes := map[string]interface{}{
				"updated_at":      realtime.Now().UTC(),
				"seats_remaining": driver.Seats.Remaining - 1,
				"seats_afternoon": driver.Seats.Afternoon - 1,
				"seats_version":   driver.Seats.Version + 1,
			}
			result := tx.Model(&entity.Driver{}).
				Where("cnh = ? AND seats_version = ?", contract.DriverCNH, driver.Seats.Version).
				UpdateColumns(attributes)
			if result.Error != nil {
				return result.Error
			}
			if result.RowsAffected == 0 {
				return fmt.Errorf("seats_version mismatch: possible concurrent update")
			}
			return nil
		},
		value.NightShift: func(driver entity.Driver) error {
			if driver.Seats.Night <= 0 {
				return fmt.Errorf("no night seats available for driver %s", driver.CNH)
			}
			attributes := map[string]interface{}{
				"updated_at":      realtime.Now().UTC(),
				"seats_remaining": driver.Seats.Remaining - 1,
				"seats_night":     driver.Seats.Night - 1,
				"seats_version":   driver.Seats.Version + 1,
			}
			result := tx.Model(&entity.Driver{}).
				Where("cnh = ? AND seats_version = ?", contract.DriverCNH, driver.Seats.Version).
				UpdateColumns(attributes)
			if result.Error != nil {
				return result.Error
			}
			if result.RowsAffected == 0 {
				return fmt.Errorf("seats_version mismatch: possible concurrent update")
			}
			return nil
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
		Update("status", value.TempContractAccepted).Error; err != nil {
		return err
	}

	return tx.Commit().Error
}

func (cr ContractRepositoryImpl) Cancel(id uuid.UUID) error {
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

	err := tx.Model(&entity.Contract{}).
		Where("uuid = ?", id).
		UpdateColumns(map[string]interface{}{
			"status":     value.ContractCanceled,
			"updated_at": realtime.Now().UTC().Unix(),
		}).Debug().Error

	if err != nil {
		return err
	}

	var contract entity.Contract
	if err := tx.Where("uuid = ?", id).First(&contract).Error; err != nil {
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

	addSeats := map[string]func(driver entity.Driver) error{
		"morning": func(driver entity.Driver) error {
			attributes := make(map[string]interface{})
			attributes["updated_at"] = realtime.Now().UTC()
			attributes["seats_remaining"] = driver.Seats.Remaining + 1
			attributes["seats_morning"] = driver.Seats.Morning + 1
			return tx.Model(&entity.Driver{}).
				Where("cnh = ?", contract.DriverCNH).
				UpdateColumns(attributes).Error
		},
		"afternoon": func(driver entity.Driver) error {
			attributes := make(map[string]interface{})
			attributes["updated_at"] = realtime.Now().UTC()
			attributes["seats_remaining"] = driver.Seats.Remaining + 1
			attributes["seats_afternoon"] = driver.Seats.Afternoon + 1
			return tx.Model(&entity.Driver{}).
				Where("cnh = ?", contract.DriverCNH).
				UpdateColumns(attributes).Error
		},
		"night": func(driver entity.Driver) error {
			attributes := make(map[string]interface{})
			attributes["updated_at"] = realtime.Now().UTC()
			attributes["seats_remaining"] = driver.Seats.Remaining + 1
			attributes["seats_night"] = driver.Seats.Night + 1
			return tx.Model(&entity.Driver{}).
				Where("cnh = ?", contract.DriverCNH).
				UpdateColumns(attributes).Error
		},
	}

	updateFunc, exists := addSeats[kid.Shift]
	if !exists {
		return fmt.Errorf("invalid shift: %s", kid.Shift)
	}

	if err := updateFunc(driver); err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (cr ContractRepositoryImpl) Expired(id uuid.UUID) error {
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

	err := tx.Model(&entity.Contract{}).Where("uuid = ?", id).Update("status", value.TempContractExpired).Error
	if err != nil {
		return err
	}

	var contract entity.Contract
	if err := tx.Where("uuid = ?", id).First(&contract).Error; err != nil {
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

	addSeats := map[string]func(driver entity.Driver) error{
		"morning": func(driver entity.Driver) error {
			attributes := make(map[string]interface{})
			attributes["updated_at"] = realtime.Now().UTC()
			attributes["seats_remaining"] = driver.Seats.Remaining + 1
			attributes["seats_morning"] = driver.Seats.Morning + 1
			return tx.Model(&entity.Driver{}).
				Where("cnh = ?", contract.DriverCNH).
				UpdateColumns(attributes).Error
		},
		"afternoon": func(driver entity.Driver) error {
			attributes := make(map[string]interface{})
			attributes["updated_at"] = realtime.Now().UTC()
			attributes["seats_remaining"] = driver.Seats.Remaining + 1
			attributes["seats_afternoon"] = driver.Seats.Afternoon + 1
			return tx.Model(&entity.Driver{}).
				Where("cnh = ?", contract.DriverCNH).
				UpdateColumns(attributes).Error
		},
		"night": func(driver entity.Driver) error {
			attributes := make(map[string]interface{})
			attributes["updated_at"] = realtime.Now().UTC()
			attributes["seats_remaining"] = driver.Seats.Remaining + 1
			attributes["seats_night"] = driver.Seats.Night + 1
			return tx.Model(&entity.Driver{}).
				Where("cnh = ?", contract.DriverCNH).
				UpdateColumns(attributes).Error
		},
	}

	updateFunc, exists := addSeats[kid.Shift]
	if !exists {
		return fmt.Errorf("invalid shift: %s", kid.Shift)
	}

	if err := updateFunc(driver); err != nil {
		return err
	}

	return nil
}

func (cr ContractRepositoryImpl) Update(id uuid.UUID, attributes map[string]interface{}) error {
	err := cr.Postgres.Client().Model(&entity.Contract{}).Where("uuid = ?", id).Updates(attributes).Error
	if err != nil {
		return err
	}
	return nil
}

func (cr ContractRepositoryImpl) GetByUUID(id uuid.UUID) (*entity.Contract, error) {
	var contract entity.Contract

	err := cr.Postgres.Client().
		Preload("Responsible").
		Preload("Driver").
		Preload("Kid").
		Preload("School").
		Where("uuid = ? AND status = ?", id, value.ContractCurrently).
		First(&contract).Error

	if err != nil {
		return nil, err
	}

	return &contract, nil
}

func (cr ContractRepositoryImpl) GetBySchool(cnpj string) ([]entity.Contract, error) {
	var contracts []entity.Contract

	err := cr.Postgres.Client().
		Preload("Responsible").
		Preload("Driver").
		Preload("Kid").
		Preload("School").
		Where("school_cnpj = ? AND status = ?", cnpj, value.ContractCurrently).
		Find(&contracts).Error

	if err != nil {
		return nil, err
	}

	return contracts, nil
}

func (cr ContractRepositoryImpl) GetByDriver(cnh string) ([]entity.Contract, error) {
	var contracts []entity.Contract

	err := cr.Postgres.Client().
		Preload("Responsible").
		Preload("Driver").
		Preload("Kid").
		Preload("School").
		Where("driver_cnh = ? AND status = ?", cnh, value.ContractCurrently).
		Find(&contracts).Error

	if err != nil {
		return nil, err
	}

	return contracts, nil
}

func (cr ContractRepositoryImpl) GetByResponsible(cpf string) ([]entity.Contract, error) {
	var contracts []entity.Contract

	err := cr.Postgres.Client().
		Preload("Responsible").
		Preload("Driver").
		Preload("Kid").
		Preload("School").
		Where("responsible_cpf = ? AND status = ?", cpf, value.ContractCurrently).
		Find(&contracts).Error

	if err != nil {
		return nil, err
	}

	return contracts, nil
}

func (cr ContractRepositoryImpl) ContractAlreadyExist(uuid string) (bool, error) {
	var count int64

	err := cr.Postgres.Client().Model(&entity.Contract{}).Where("uuid = ?", uuid).Count(&count).Error
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func (cr ContractRepositoryImpl) GetByKid(rg string) (*entity.Contract, error) {
	var contracts entity.Contract

	err := cr.Postgres.Client().
		Preload("Responsible").
		Preload("Driver").
		Preload("Kid").
		Preload("School").
		Where("kid_rg = ? AND status = ?", rg, value.ContractCurrently).
		Find(&contracts).Error

	if err != nil {
		return nil, err
	}

	return &contracts, nil
}

func (cr ContractRepositoryImpl) KidHasEnableContract(rg string) (bool, error) {
	var count int64

	err := cr.Postgres.Client().
		Model(&entity.Contract{}).
		Where("kid_rg = ? AND status = ?", rg, value.ContractCurrently).
		Count(&count).Error
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func (cr ContractRepositoryImpl) ResponsibleHasEnableContract(cpf string) (bool, error) {
	var count int64

	err := cr.Postgres.Client().
		Model(&entity.Contract{}).
		Where("responsible_cpf = ? AND status = ?", cpf, value.ContractCurrently).
		Count(&count).Error
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func (cr ContractRepositoryImpl) DriverHasEnableContract(cnh string) (bool, error) {
	var count int64

	err := cr.Postgres.Client().
		Model(&entity.Contract{}).
		Where("driver_cnh = ? AND status = ?", cnh, value.ContractCurrently).
		Count(&count).Error
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func (cr ContractRepositoryImpl) SchoolHasEnableContract(cnpj string) (bool, error) {
	var count int64

	err := cr.Postgres.Client().
		Model(&entity.Contract{}).
		Where("school_cnpj = ? AND status = ?", cnpj, value.ContractCurrently).
		Count(&count).Error
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

func (cr ContractRepositoryImpl) GetNumberOfEnableContractsByDriver(cnh string) (int64, error) {
	var count int64

	err := cr.Postgres.Client().
		Model(&entity.Contract{}).
		Where("driver_cnh = ? AND status = ?", cnh, value.ContractCurrently).
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (cr ContractRepositoryImpl) PartnerHasEnableContract(id string) ([]entity.Contract, error) {
	var contracts []entity.Contract

	err := cr.Postgres.Client().
		Joins("JOIN partners ON partners.driver_cnh = contracts.driver_cnh AND partners.school_cnpj = contracts.school_cnpj").
		Where("partners.id = ? AND status = ?", id, value.ContractCurrently).
		Find(&contracts).Error

	if err != nil {
		return []entity.Contract{}, err
	}

	return contracts, nil
}
