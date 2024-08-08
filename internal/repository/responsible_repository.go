package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/venture-technology/venture/models"
	"github.com/venture-technology/venture/pkg/utils"
)

type IResponsibleRepository interface {
	CreateResponsible(ctx context.Context, responsible *models.Responsible) error
	GetResponsible(ctx context.Context, cpf *string) (*models.Responsible, error)
	UpdateResponsible(ctx context.Context, currentResponsible, responsible *models.Responsible) error
	DeleteResponsible(ctx context.Context, cpf *string) error
	AuthResponsible(ctx context.Context, responsible *models.Responsible) (*models.Responsible, error)
	SaveCreditCard(ctx context.Context, cpf, cardToken, paymentMethodId *string) error
}

type ResponsibleRepository struct {
	db *sql.DB
}

func NewResponsibleRepository(conn *sql.DB) *ResponsibleRepository {
	return &ResponsibleRepository{
		db: conn,
	}
}

func (rer *ResponsibleRepository) CreateResponsible(ctx context.Context, responsible *models.Responsible) error {
	sqlQuery := `INSERT INTO responsible (name, email, password, cpf, street, number, zip, complement, status, customer_id, phone) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := rer.db.Exec(sqlQuery, responsible.Name, responsible.Email, responsible.Password, responsible.CPF, responsible.Street, responsible.Number, responsible.ZIP, responsible.Complement, responsible.Status, responsible.CustomerId, responsible.Phone)
	return err
}

func (rer *ResponsibleRepository) GetResponsible(ctx context.Context, cpf *string) (*models.Responsible, error) {
	sqlQuery := `SELECT id, name, cpf, email, street, number, zip, status, complement, COALESCE(card_token, '') AS card_token, COALESCE(payment_method_id, '') AS payment_method_id, customer_id, phone FROM responsible WHERE cpf = $1 LIMIT 1`
	var responsible models.Responsible
	err := rer.db.QueryRow(sqlQuery, *cpf).Scan(
		&responsible.ID,
		&responsible.Name,
		&responsible.CPF,
		&responsible.Email,
		&responsible.Street,
		&responsible.Number,
		&responsible.ZIP,
		&responsible.Status,
		&responsible.Complement,
		&responsible.CreditCard.CardToken,
		&responsible.PaymentMethodId,
		&responsible.CustomerId,
		&responsible.Phone,
	)
	if err != nil || err == sql.ErrNoRows {
		return nil, err
	}
	return &responsible, nil
}

func (rer *ResponsibleRepository) UpdateResponsible(ctx context.Context, currentResponsible, responsible *models.Responsible) error {

	if responsible.Name != "" && responsible.Name != currentResponsible.Name {
		currentResponsible.Name = responsible.Name
	}

	if responsible.Email != "" && responsible.Email != currentResponsible.Email {
		currentResponsible.Email = responsible.Email
	}

	if responsible.Password != "" && responsible.Password != currentResponsible.Password {
		currentResponsible.Password = responsible.Password
		currentResponsible.Password = utils.HashPassword(currentResponsible.Password)
	}

	if responsible.Street != "" && responsible.Street != currentResponsible.Street {
		currentResponsible.Street = responsible.Street
	}

	if responsible.Number != "" && responsible.Number != currentResponsible.Number {
		currentResponsible.Number = responsible.Number
	}

	if responsible.ZIP != "" && responsible.ZIP != currentResponsible.ZIP {
		currentResponsible.ZIP = responsible.ZIP
	}

	if responsible.Complement != "" && responsible.Complement != currentResponsible.Complement {
		currentResponsible.Complement = responsible.Complement
	}

	sqlQueryUpdate := `UPDATE responsible SET name = $1, email = $2, password = $3, street = $4, number = $5, zip = $6, complement = $7 WHERE cpf = $8`
	_, err := rer.db.ExecContext(ctx, sqlQueryUpdate, currentResponsible.Name, currentResponsible.Email, currentResponsible.Password, currentResponsible.Street, currentResponsible.Number, currentResponsible.ZIP, currentResponsible.Complement, responsible.CPF)
	return err
}

func (rer *ResponsibleRepository) DeleteResponsible(ctx context.Context, cpf *string) error {
	tx, err := rer.db.Begin()
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
	_, err = tx.Exec("DELETE FROM responsible WHERE cpf = $1", *cpf)
	return err
}

func (rer *ResponsibleRepository) AuthResponsible(ctx context.Context, responsible *models.Responsible) (*models.Responsible, error) {
	sqlQuery := `SELECT id, name, cpf, email, street, number, zip, status, complement, card_token, payment_method_id, customer_id, phone, password FROM responsible WHERE email = $1 LIMIT 1`
	var responsibleData models.Responsible
	err := rer.db.QueryRow(sqlQuery, responsible.Email).Scan(
		&responsibleData.ID,
		&responsibleData.Name,
		&responsibleData.CPF,
		&responsibleData.Email,
		&responsibleData.Street,
		&responsibleData.Number,
		&responsibleData.ZIP,
		&responsibleData.Status,
		&responsibleData.Complement,
		&responsibleData.CreditCard.CardToken,
		&responsibleData.PaymentMethodId,
		&responsibleData.CustomerId,
		&responsibleData.Phone,
		&responsibleData.Password,
	)
	if err != nil || err == sql.ErrNoRows {
		return nil, err
	}
	match := responsibleData.Password == responsible.Password
	if !match {
		return nil, fmt.Errorf("email or password wrong")
	}
	responsibleData.Password = ""
	return &responsibleData, nil
}

func (rer *ResponsibleRepository) SaveCreditCard(ctx context.Context, cpf, cardToken, paymentMethodId *string) error {

	sqlQueryUpdate := `UPDATE responsible SET card_token = $1, payment_method_id = $2 WHERE cpf = $3`
	_, err := rer.db.ExecContext(ctx, sqlQueryUpdate, cardToken, paymentMethodId, cpf)
	return err

}
