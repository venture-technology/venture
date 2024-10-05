package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/entity"
)

type IContractRepository interface {
	Create(ctx context.Context, contract *entity.Contract) error
	Get(ctx context.Context, id uuid.UUID) (*entity.Contract, error)
	FindAllByCnpj(ctx context.Context, cnpj *string) ([]entity.Contract, error)
	FindAllByCpf(ctx context.Context, cpf *string) ([]entity.Contract, error)
	FindAllByCnh(ctx context.Context, cnh *string) ([]entity.Contract, error)
	Cancel(ctx context.Context, id uuid.UUID) error
	Expired(ctx context.Context, id uuid.UUID) error
	GetSimpleContractByTitle(ctx context.Context, title *string) (*entity.Contract, error)
}

type ContractRepository struct {
	db *sql.DB
}

func NewContractRepository(db *sql.DB) *ContractRepository {
	return &ContractRepository{
		db,
	}
}

func (cr *ContractRepository) Create(ctx context.Context, contract *entity.Contract) error {
	sqlQuery := `INSERT INTO contracts (
    record,
    title_stripe_subscription,
    description_stripe_subscription,
    id_stripe_subscription,
    id_price_subscription,
    id_product_subscription,
    school_id,
    driver_id,
    responsible_id,
    child_id,
    status
  ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	_, err := cr.db.Exec(sqlQuery, contract.Record, contract.StripeSubscription.Title, contract.StripeSubscription.Description, contract.StripeSubscription.ID, contract.StripeSubscription.Price, contract.StripeSubscription.Product, contract.School.CNPJ, contract.Driver.CNH, contract.Child.Responsible.CPF, contract.Child.RG, contract.Status)
	return err
}

func (cr *ContractRepository) Get(ctx context.Context, id uuid.UUID) (*entity.Contract, error) {
	sqlQuery := `
		SELECT
			c.record, c.title_stripe_subscription, c.description_stripe_subscription, c.id_stripe_subscription, c.id_price_subscription, c.id_product_subscription, c.created_at, c.expire_at, c.status, c.amount,
			d.name AS driver_name, d.email AS driver_email, d.qrcode AS driver_qrcode, d.phone AS driver_phone,
			s.name AS school_name, s.email AS school_email, s.phone AS school_phone,
			ch.name AS child_name, ch.rg AS child_rg, ch.responsible_id AS child_responsible_id, ch.shift AS child_shift,
			r.name AS responsible_name, r.email AS responsible_email, r.phone AS responsible_phone
		FROM
			contracts c
		JOIN
			drivers d ON c.driver_id = d.cnh
		JOIN
			schools s ON c.school_id = s.cnpj
		JOIN
			children ch ON c.child_id = ch.rg
		JOIN
			responsible r ON ch.responsible_id = r.cpf
		WHERE
			c.record = $1
		LIMIT 1;
	`
	var contract entity.Contract
	err := cr.db.QueryRowContext(ctx, sqlQuery, id).Scan(
		&contract.Record,
		&contract.StripeSubscription.Title,
		&contract.StripeSubscription.Description,
		&contract.StripeSubscription.ID,
		&contract.StripeSubscription.Price,
		&contract.StripeSubscription.Product,
		&contract.CreatedAt,
		&contract.ExpireAt,
		&contract.Status,
		&contract.Amount,
		&contract.Driver.Name,
		&contract.Driver.Email,
		&contract.Driver.QrCode,
		&contract.Driver.Phone,
		&contract.School.Name,
		&contract.School.Email,
		&contract.School.Phone,
		&contract.Child.Name,
		&contract.Child.RG,
		&contract.Child.Responsible.CPF,
		&contract.Child.Shift,
		&contract.Child.Responsible.Name,
		&contract.Child.Responsible.Email,
		&contract.Child.Responsible.Phone,
	)
	if err != nil {
		return nil, err
	}
	return &contract, nil
}

func (cr *ContractRepository) FindAllByCnpj(ctx context.Context, cnpj *string) ([]entity.Contract, error) {
	sqlQuery := `
		SELECT
			c.record, c.title_stripe_subscription, c.description_stripe_subscription, c.id_stripe_subscription, c.id_price_subscription, c.id_product_subscription, c.created_at, c.expire_at, c.status,
			d.name AS driver_name, d.email AS driver_email, d.qrcode AS driver_qrcode, d.phone AS driver_phone,
			s.name AS school_name, s.email AS school_email, s.phone AS school_phone,
			ch.name AS child_name, ch.rg AS child_rg, ch.responsible_id AS child_responsible_id, ch.shift AS child_shift,
			r.name AS responsible_name, r.email AS responsible_email, r.phone AS responsible_phone
		FROM
			contracts c
		JOIN
			drivers d ON c.driver_id = d.cnh
		JOIN
			schools s ON c.school_id = s.cnpj
		JOIN
			children ch ON c.child_id = ch.rg
		JOIN
			responsibles r ON ch.responsible_id = r.cpf
		WHERE
			s.cnpj = $1;
	`

	rows, err := cr.db.QueryContext(ctx, sqlQuery, cnpj)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contracts []entity.Contract
	for rows.Next() {
		var contract entity.Contract
		err := rows.Scan(
			&contract.Record,
			&contract.StripeSubscription.Title,
			&contract.StripeSubscription.Description,
			&contract.StripeSubscription.ID,
			&contract.StripeSubscription.Price,
			&contract.StripeSubscription.Product,
			&contract.CreatedAt,
			&contract.ExpireAt,
			&contract.Status,
			&contract.Driver.Name,
			&contract.Driver.Email,
			&contract.Driver.QrCode,
			&contract.Driver.Phone,
			&contract.School.Name,
			&contract.School.Email,
			&contract.School.Phone,
			&contract.Child.Name,
			&contract.Child.RG,
			&contract.Child.Responsible.CPF,
			&contract.Child.Shift,
			&contract.Child.Responsible.Name,
			&contract.Child.Responsible.Email,
			&contract.Child.Responsible.Phone,
		)
		if err != nil {
			return nil, err
		}
		contracts = append(contracts, contract)
	}

	return contracts, nil
}

func (cr *ContractRepository) FindAllByCpf(ctx context.Context, cpf *string) ([]entity.Contract, error) {
	sqlQuery := `
		SELECT
			c.record, c.title_stripe_subscription, c.description_stripe_subscription, c.id_stripe_subscription, c.id_price_subscription, c.id_product_subscription, c.created_at, c.expire_at, c.status,
			d.name AS driver_name, d.email AS driver_email, d.qrcode AS driver_qrcode, d.phone AS driver_phone,
			s.name AS school_name, s.email AS school_email, s.phone AS school_phone,
			ch.name AS child_name, ch.rg AS child_rg, ch.responsible_id AS child_responsible_id, ch.shift AS child_shift,
			r.name AS responsible_name, r.email AS responsible_email, r.phone AS responsible_phone
		FROM
			contracts c
		JOIN
			drivers d ON c.driver_id = d.cnh
		JOIN
			schools s ON c.school_id = s.cnpj
		JOIN
			children ch ON c.child_id = ch.rg
		JOIN
			responsibles r ON ch.responsible_id = r.cpf
		WHERE
			r.cpf = $1;
	`

	rows, err := cr.db.QueryContext(ctx, sqlQuery, cpf)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contracts []entity.Contract
	for rows.Next() {
		var contract entity.Contract
		err := rows.Scan(
			&contract.Record,
			&contract.StripeSubscription.Title,
			&contract.StripeSubscription.Description,
			&contract.StripeSubscription.ID,
			&contract.StripeSubscription.Price,
			&contract.StripeSubscription.Product,
			&contract.CreatedAt,
			&contract.ExpireAt,
			&contract.Status,
			&contract.Driver.Name,
			&contract.Driver.Email,
			&contract.Driver.QrCode,
			&contract.Driver.Phone,
			&contract.School.Name,
			&contract.School.Email,
			&contract.School.Phone,
			&contract.Child.Name,
			&contract.Child.RG,
			&contract.Child.Responsible.CPF,
			&contract.Child.Shift,
			&contract.Child.Responsible.Name,
			&contract.Child.Responsible.Email,
			&contract.Child.Responsible.Phone,
		)
		if err != nil {
			return nil, err
		}
		contracts = append(contracts, contract)
	}

	return contracts, nil
}

func (cr *ContractRepository) FindAllByCnh(ctx context.Context, cnh *string) ([]entity.Contract, error) {
	sqlQuery := `
		SELECT
			c.record, c.title_stripe_subscription, c.description_stripe_subscription, c.id_stripe_subscription, c.id_price_subscription, c.id_product_subscription, c.created_at, c.expire_at, c.status,
			d.name AS driver_name, d.email AS driver_email, d.qrcode AS driver_qrcode, d.phone AS driver_phone,
			s.name AS school_name, s.email AS school_email, s.phone AS school_phone,
			ch.name AS child_name, ch.rg AS child_rg, ch.responsible_id AS child_responsible_id, ch.shift AS child_shift,
			r.name AS responsible_name, r.email AS responsible_email, r.phone AS responsible_phone
		FROM
			contracts c
		JOIN
			drivers d ON c.driver_id = d.cnh
		JOIN
			schools s ON c.school_id = s.cnpj
		JOIN
			children ch ON c.child_id = ch.rg
		JOIN
			responsibles r ON ch.responsible_id = r.cpf
		WHERE
			d.cnh = $1;
	`

	rows, err := cr.db.QueryContext(ctx, sqlQuery, cnh)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contracts []entity.Contract
	for rows.Next() {
		var contract entity.Contract
		err := rows.Scan(
			&contract.Record,
			&contract.StripeSubscription.Title,
			&contract.StripeSubscription.Description,
			&contract.StripeSubscription.ID,
			&contract.StripeSubscription.Price,
			&contract.StripeSubscription.Product,
			&contract.CreatedAt,
			&contract.ExpireAt,
			&contract.Status,
			&contract.Driver.Name,
			&contract.Driver.Email,
			&contract.Driver.QrCode,
			&contract.Driver.Phone,
			&contract.School.Name,
			&contract.School.Email,
			&contract.School.Phone,
			&contract.Child.Name,
			&contract.Child.RG,
			&contract.Child.Responsible.CPF,
			&contract.Child.Shift,
			&contract.Child.Responsible.Name,
			&contract.Child.Responsible.Email,
			&contract.Child.Responsible.Phone,
		)
		if err != nil {
			return nil, err
		}
		contracts = append(contracts, contract)
	}

	return contracts, nil
}

func (cr *ContractRepository) Cancel(ctx context.Context, id uuid.UUID) error {
	sqlQueryUpdate := `UPDATE contract SET status = 'canceled' WHERE id = $1`
	_, err := cr.db.ExecContext(ctx, sqlQueryUpdate, id)
	return err
}

func (cr *ContractRepository) Expired(ctx context.Context, id uuid.UUID) error {
	sqlQueryUpdate := `UPDATE contract SET status = 'expired' WHERE id = $1`
	_, err := cr.db.ExecContext(ctx, sqlQueryUpdate, id)
	return err
}

func (cr *ContractRepository) GetSimpleContractByTitle(ctx context.Context, title *string) (*entity.Contract, error) {
	sqlQuery := `SELECT title_stripe_subscription FROM contracts WHERE title_stripe_subscription = $1`

	var contract entity.Contract
	err := cr.db.QueryRow(sqlQuery, *title).Scan(
		&contract.StripeSubscription.Title,
	)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &contract, nil
}
