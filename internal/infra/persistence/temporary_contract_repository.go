package persistence

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/value"
	"github.com/venture-technology/venture/pkg/realtime"
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

func (tcr TempContractRepositoryImpl) Get(uuid uuid.UUID) (*entity.TempContract, error) {
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

func (tcr TempContractRepositoryImpl) Expire(uuids []uuid.UUID) error {
	if len(uuids) == 0 {
		return nil
	}

	err := tcr.Postgres.Client().
		Model(&entity.TempContract{}).
		Where("uuid IN (?)", uuids).
		Update("status", value.TempContractExpired).
		Error

	return err
}

func (tcr TempContractRepositoryImpl) Cancel(uuid uuid.UUID) error {
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
	uuid uuid.UUID,
	attrs map[string]interface{},
) error {
	err := tcr.Postgres.Client().
		Model(&entity.TempContract{}).
		Where("uuid = ?", uuid).
		UpdateColumns(attrs).
		Error

	return err
}

func (tcr TempContractRepositoryImpl) GetExpiredContracts() ([]entity.TempContract, error) {
	var tempContracts []entity.TempContract

	now := realtime.Now().Unix()

	if err := tcr.Postgres.Client().Where("expired_at < ? AND status = ?", now, value.TempContractPending).Find(&tempContracts).Error; err != nil {
		return nil, err
	}
	return tempContracts, nil
}
