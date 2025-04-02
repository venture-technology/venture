package persistence

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/value"
)

type TempContractRepositoryImpl struct {
	Postgres contracts.PostgresIface
}

func (tcr TempContractRepositoryImpl) Create(tempContract *entity.TempContract) error {
	err := tcr.Postgres.Client().Create(&tempContract).Error
	if err != nil {
		return err
	}
	return nil
}

func (tcr TempContractRepositoryImpl) Get(uuid string) (*entity.TempContract, error) {
	var tempContract entity.TempContract
	err := tcr.Postgres.Client().
		Where("uuid = ?", uuid).
		First(&tempContract).Error

	if err != nil {
		return nil, err
	}
	return &tempContract, nil
}

func (tcr TempContractRepositoryImpl) GetByEveryone(tempContract *entity.TempContract) (bool, error) {
	var count int64
	now := time.Now().Unix()
	oneYearAgo := now - (365 * 24 * 60 * 60)

	query := tcr.Postgres.Client().
		Model(&entity.TempContract{}).
		Where("kid_rg = ?", tempContract.KidRG).
		Where("responsible_cpf = ?", tempContract.ResponsibleCPF).
		Where("(status = 'pending' AND expired_at > ?) OR (status = 'accepted' AND created_at >= ?)",
			now,        // Para "pending", só listar se ainda não expirou
			oneYearAgo, // Para "accepted", só listar se foi criado nos últimos 365 dias
		).
		Count(&count).
		Error

	if query != nil {
		return false, query
	}

	return count > 0, nil
}

func (tcr TempContractRepositoryImpl) Expire(uuid string) error {
	err := tcr.Postgres.Client().
		Model(&entity.TempContract{}).
		Where("uuid = ?", uuid).
		Update("status", value.TempContractExpired).
		Error

	if err != nil {
		return err
	}
	return nil
}

func (tcr TempContractRepositoryImpl) Cancel(uuid string) error {
	err := tcr.Postgres.Client().
		Model(&entity.TempContract{}).
		Where("uuid = ?", uuid).
		Update("status", value.TempContractCanceled).
		Error

	if err != nil {
		return err
	}
	return nil
}

func (tcr TempContractRepositoryImpl) Accept(uuid string) error {
	err := tcr.Postgres.Client().
		Model(&entity.TempContract{}).
		Where("uuid = ?", uuid).
		Update("status", value.TempContractAccepted).
		Error

	if err != nil {
		return err
	}
	return nil
}

func (tcr TempContractRepositoryImpl) GetByResponsible(cpf *string) ([]entity.TempContract, error) {
	var contracts []entity.TempContract

	if cpf == nil {
		return nil, gorm.ErrRecordNotFound
	}

	err := tcr.Postgres.Client().Where("responsible_cpf = ?", cpf).Find(&contracts).Error
	if err != nil {
		return nil, err
	}

	return contracts, nil
}

func (tcr TempContractRepositoryImpl) GetByDriver(cnh *string) ([]entity.TempContract, error) {
	var contracts []entity.TempContract

	if cnh == nil {
		return nil, gorm.ErrRecordNotFound
	}

	err := tcr.Postgres.Client().Where("driver_cnh = ?", cnh).Find(&contracts).Error
	if err != nil {
		return nil, err
	}

	return contracts, nil
}

func (tcr TempContractRepositoryImpl) Update(
	uuid string,
	attrs map[string]interface{},
) error {
	err := tcr.Postgres.Client().
		Model(&entity.TempContract{}).
		Where("uuid = ?", uuid).
		UpdateColumns(attrs).
		Error

	return err
}
