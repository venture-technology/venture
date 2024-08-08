package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/venture-technology/venture/models"
)

type IContractRepository interface {
	CreateContract(ctx context.Context, contract *models.Contract) error
	FindContractsByCpf(ctx context.Context, cpf, status *string) ([]models.Contract, error)
	GetContract(ctx context.Context, record uuid.UUID) (*models.Contract, error)
	UpdateStatusContract(ctx context.Context, record uuid.UUID, status string) error
}

type ContractRepository struct {
	db *sql.DB
}

func NewContractRepository(db *sql.DB) *ContractRepository {
	return &ContractRepository{
		db: db,
	}
}

func (cr *ContractRepository) CreateContract(ctx context.Context, contract *models.Contract) error {

	sqlQuery := `INSERT INTO contracts (
		record,
		title_stripe_subscription,
		description_stripe_subscription,
		id_stripe_subsciption,
		id_price_subscription,
		id_product_subscription,
		school_id,
		driver_id
		responsible_id,
		child_id,
		created_at
		expire_at,
		status
	)
	VALUES ($1, $2. $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, "currently")`

	_, err := cr.db.Exec(sqlQuery,
		contract.Record,
		contract.StripeSubscription.Title,
		contract.Description,
		contract.StripeSubscription.SubscriptionId,
		contract.StripeSubscription.PriceSubscriptionId,
		contract.StripeSubscription.ProductSubscriptionId,
		contract.School.ID,
		contract.Driver.ID,
		contract.Child.Responsible.ID,
		contract.Child.ID,
		contract.ExpireAt,
	)

	return err

}

func (cr *ContractRepository) FindContractsByCpf(ctx context.Context, cpf, status *string) ([]models.Contract, error) {

	sqlQuery := `SELECT 
		record,
		title_stripe_subscription,
		description_stripe_subscription,
		id_stripe_subsciption,
		id_price_subscription,
		id_product_subscription,
		school_id,
		driver_id
		responsible_id,
		child_id,
		created_at,
		expire_at, 
		status FROM contracts WHERE cpf_responsible = $1 AND status = $2`

	rows, err := cr.db.Query(sqlQuery, *cpf, *status)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contracts []models.Contract

	for rows.Next() {
		var contract models.Contract

		err := rows.Scan(
			&contract.Record,
			&contract.StripeSubscription.Title,
			&contract.Description,
			&contract.StripeSubscription.SubscriptionId,
			&contract.StripeSubscription.PriceSubscriptionId,
			&contract.StripeSubscription.ProductSubscriptionId,
			&contract.School.ID,
			&contract.Driver.ID,
			&contract.Child.Responsible.ID,
			&contract.Child.ID,
			&contract.CreatedAt,
			&contract.ExpireAt,
		)

		if err != nil {
			return nil, err
		}

		contracts = append(contracts, contract)
	}

	if err := rows.Scan(); err != nil {
		return nil, err
	}

	return contracts, nil

}

func (cr *ContractRepository) UpdateStatusContract(ctx context.Context, record uuid.UUID, status string) error {

	sqlQuery := `UPDATE contracts SET status = $1 WHERE record = $2`
	_, err := cr.db.ExecContext(ctx, sqlQuery, status)
	return err

}
