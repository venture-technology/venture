package repository

import (
	"context"
	"database/sql"

	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/pkg/utils"
)

type IDriverRepository interface {
	Create(ctx context.Context, driver *entity.Driver) error
	Get(ctx context.Context, cnh *string) (*entity.Driver, error)
	Update(ctx context.Context, driver *entity.Driver) error
	Delete(ctx context.Context, cnh *string) error

	// podemos ter apenas uma chave pix ou conta de banco registrada, portanto esta ja realiza update
	SavePix(ctx context.Context, driver *entity.Driver) error
	SaveBank(ctx context.Context, driver *entity.Driver) error
}

type DriverRepository struct {
	db *sql.DB
}

func NewDriverRepository(db *sql.DB) *DriverRepository {
	return &DriverRepository{
		db: db,
	}
}

func (dr *DriverRepository) Create(ctx context.Context, driver *entity.Driver) error {
	sqlQuery := `INSERT INTO drivers (amount, qrcode, name, email, password, cpf, cnh, street, number, zip, complement) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := dr.db.Exec(sqlQuery, driver.Amount, driver.QrCode, driver.Name, driver.Email, driver.Password, driver.CPF, driver.CNH, driver.Address.Street, driver.Address.Number, driver.Address.ZIP, driver.Address.Complement)
	return err
}

func (dr *DriverRepository) Get(ctx context.Context, cnh *string) (*entity.Driver, error) {
	sqlQuery := `SELECT id, amount, name, cpf, cnh, qrcode, email, street, number, zip, complement FROM drivers WHERE cnh = $1 LIMIT 1`
	var driver entity.Driver
	err := dr.db.QueryRow(sqlQuery, *cnh).Scan(
		&driver.ID,
		&driver.Amount,
		&driver.Name,
		&driver.CPF,
		&driver.CNH,
		&driver.QrCode,
		&driver.Email,
		&driver.Address.Street,
		&driver.Address.Number,
		&driver.Address.ZIP,
		&driver.Address.Complement,
	)
	if err != nil || err == sql.ErrNoRows {
		return nil, err
	}
	return &driver, nil
}

func (dr *DriverRepository) Update(ctx context.Context, driver *entity.Driver) error {
	sqlQuery := `SELECT name, amount, email, password, street, number, zip, complement FROM drivers WHERE cnh = $1 LIMIT 1`
	var currentDriver entity.Driver
	err := dr.db.QueryRow(sqlQuery, driver.CNH).Scan(
		&currentDriver.Name,
		&currentDriver.Amount,
		&currentDriver.Email,
		&currentDriver.Password,
		&currentDriver.Address.Street,
		&currentDriver.Address.Number,
		&currentDriver.Address.ZIP,
		&currentDriver.Address.Complement,
	)
	if err != nil || err == sql.ErrNoRows {
		return err
	}

	if driver.Name != "" && driver.Name != currentDriver.Name {
		currentDriver.Name = driver.Name
	}

	if driver.Amount != 0 && driver.Amount != currentDriver.Amount {
		currentDriver.Amount = driver.Amount

	}
	if driver.Email != "" && driver.Email != currentDriver.Email {
		currentDriver.Email = driver.Email
	}
	if driver.Password != "" && driver.Password != currentDriver.Password {
		currentDriver.Password = driver.Password
		currentDriver.Password = utils.HashPassword(currentDriver.Password)
	}
	if driver.Address.Street != "" && driver.Address.Street != currentDriver.Address.Street {
		currentDriver.Address.Street = driver.Address.Street
	}
	if driver.Address.Number != "" && driver.Address.Number != currentDriver.Address.Number {
		currentDriver.Address.Number = driver.Address.Number
	}
	if driver.Address.ZIP != "" && driver.Address.ZIP != currentDriver.Address.ZIP {
		currentDriver.Address.ZIP = driver.Address.ZIP
	}
	if driver.Address.Complement != "" && driver.Address.Complement != currentDriver.Address.Complement {
		currentDriver.Address.Complement = driver.Address.Complement
	}

	sqlQueryUpdate := `UPDATE drivers SET name = $1,  amount = $2, email = $3, password = $4, street = $5, number = $6, zip = $7, complement = $8 WHERE cnh = $9`
	_, err = dr.db.ExecContext(ctx, sqlQueryUpdate, currentDriver.Name, currentDriver.Amount, currentDriver.Email, currentDriver.Password, currentDriver.Address.Street, currentDriver.Address.Number, currentDriver.Address.ZIP, currentDriver.Address.Complement, driver.CNH)
	return err
}

func (dr *DriverRepository) Delete(ctx context.Context, cnh *string) error {
	tx, err := dr.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	_, err = tx.Exec("DELETE FROM drivers WHERE cnh = $1", *cnh)
	return err
}

func (dr *DriverRepository) SavePix(ctx context.Context, driver *entity.Driver) error {
	sqlQuery := `UPDATE drivers SET pix_key = $1 WHERE cnh = $2`
	_, err := dr.db.ExecContext(ctx, sqlQuery, driver.Pix.Key, driver.CNH)
	return err
}

func (dr *DriverRepository) SaveBank(ctx context.Context, driver *entity.Driver) error {
	sqlQuery := `UPDATE drivers SET bank_name = $1, agency_number = $2, account_number = $3 WHERE cnh = $4`
	_, err := dr.db.ExecContext(ctx, sqlQuery, driver.Bank.Name, driver.Bank.Agency, driver.Bank.Account, driver.CNH)
	return err
}
