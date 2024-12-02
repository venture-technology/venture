package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/pkg/utils"
	"go.uber.org/zap"
)

type IResponsibleRepository interface {
	Create(ctx context.Context, responsible *entity.Responsible) error
	Get(ctx context.Context, cpf *string) (*entity.Responsible, error)
	Update(ctx context.Context, currentResponsible, responsible *entity.Responsible) error
	Delete(ctx context.Context, cpf *string) error
	SaveCard(ctx context.Context, cpf, cardToken, paymentMethodId *string) error
	Auth(ctx context.Context, responsible *entity.Responsible) (*entity.Responsible, error)
	FindByEmail(ctx context.Context, email *string) (*entity.Responsible, error)
}

type ResponsibleRepository struct {
	db     *sql.DB
	logger *zap.Logger
}

func NewResponsibleRepository(db *sql.DB, logger *zap.Logger) *ResponsibleRepository {
	return &ResponsibleRepository{
		db:     db,
		logger: logger,
	}
}

func (rr *ResponsibleRepository) Create(ctx context.Context, responsible *entity.Responsible) error {
	sqlQuery := `INSERT INTO responsible (
		name, 
		email, 
		password, 
		cpf, 
		street, 
		number, 
		zip, 
		complement, 
		status, 
		customer_id, 
		phone, 
		profile_image
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
	_, err := rr.db.Exec(sqlQuery, responsible.Name, responsible.Email, responsible.Password, responsible.CPF, responsible.Address.Street, responsible.Address.Number, responsible.Address.ZIP, responsible.Address.Complement, responsible.Status, responsible.CustomerId, responsible.Phone, responsible.ProfileImage)
	return err
}

func (rr *ResponsibleRepository) Get(ctx context.Context, cpf *string) (*entity.Responsible, error) {
	sqlQuery := `SELECT 
		id, 
		name, 
		cpf, 
		email, 
		street, 
		number, 
		zip, 
		status, 
		complement, 
		COALESCE(card_token, '') AS card_token, 
		COALESCE(payment_method_id, '') AS payment_method_id, 
		customer_id, 
		phone, 
		profile_image FROM responsible WHERE cpf = $1 LIMIT 1`
	var responsible entity.Responsible
	err := rr.db.QueryRow(sqlQuery, *cpf).Scan(
		&responsible.ID,
		&responsible.Name,
		&responsible.CPF,
		&responsible.Email,
		&responsible.Address.Street,
		&responsible.Address.Number,
		&responsible.Address.ZIP,
		&responsible.Status,
		&responsible.Address.Complement,
		&responsible.CreditCard.CardToken,
		&responsible.PaymentMethodId,
		&responsible.CustomerId,
		&responsible.Phone,
		&responsible.ProfileImage,
	)
	if err != nil || err == sql.ErrNoRows {
		return nil, err
	}
	return &responsible, nil
}

func (rr *ResponsibleRepository) Update(ctx context.Context, currentResponsible, responsible *entity.Responsible) error {
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

	if responsible.Address.Street != "" && responsible.Address.Street != currentResponsible.Address.Street {
		currentResponsible.Address.Street = responsible.Address.Street
	}

	if responsible.Address.Number != "" && responsible.Address.Number != currentResponsible.Address.Number {
		currentResponsible.Address.Number = responsible.Address.Number
	}

	if responsible.Address.ZIP != "" && responsible.Address.ZIP != currentResponsible.Address.ZIP {
		currentResponsible.Address.ZIP = responsible.Address.ZIP
	}

	if responsible.Address.Complement != "" && responsible.Address.Complement != currentResponsible.Address.Complement {
		currentResponsible.Address.Complement = responsible.Address.Complement
	}

	if responsible.ProfileImage != "" && responsible.ProfileImage != currentResponsible.ProfileImage {
		currentResponsible.ProfileImage = responsible.ProfileImage
	}

	sqlQueryUpdate := `UPDATE responsible SET 
		name = $1, 
		email = $2, 
		password = $3, 
		street = $4, 
		number = $5, 
		zip = $6, 
		complement = $7, 
		profile_image = $8 WHERE cpf = $9`
	_, err := rr.db.ExecContext(ctx, sqlQueryUpdate, currentResponsible.Name, currentResponsible.Email, currentResponsible.Password, currentResponsible.Address.Street, currentResponsible.Address.Number, currentResponsible.Address.ZIP, currentResponsible.Address.Complement, currentResponsible.ProfileImage, responsible.CPF)
	return err
}

func (rr *ResponsibleRepository) Delete(ctx context.Context, cpf *string) error {
	tx, err := rr.db.Begin()
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

func (rr *ResponsibleRepository) SaveCard(ctx context.Context, cpf, cardToken, paymentMethodId *string) error {
	sqlQueryUpdate := `UPDATE responsible SET card_token = $1, payment_method_id = $2 WHERE cpf = $3`
	_, err := rr.db.ExecContext(ctx, sqlQueryUpdate, cardToken, paymentMethodId, cpf)
	return err
}

func (rr *ResponsibleRepository) Auth(ctx context.Context, responsible *entity.Responsible) (*entity.Responsible, error) {
	sqlQuery := `SELECT 
		id, 
		name, 
		cpf, 
		email, 
		street, 
		number, 
		zip, 
		status, 
		complement, 
		card_token, 
		payment_method_id, 
		customer_id, 
		phone, 
		password, 
		profile_image FROM responsible WHERE email = $1 LIMIT 1`
	var responsibleData entity.Responsible
	err := rr.db.QueryRow(sqlQuery, responsible.Email).Scan(
		&responsibleData.ID,
		&responsibleData.Name,
		&responsibleData.CPF,
		&responsibleData.Email,
		&responsibleData.Address.Street,
		&responsibleData.Address.Number,
		&responsibleData.Address.ZIP,
		&responsibleData.Status,
		&responsibleData.Address.Complement,
		&responsibleData.CreditCard.CardToken,
		&responsibleData.PaymentMethodId,
		&responsibleData.CustomerId,
		&responsibleData.Phone,
		&responsibleData.Password,
		&responsibleData.ProfileImage,
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

func (rr *ResponsibleRepository) FindByEmail(ctx context.Context, email *string) (*entity.Responsible, error) {
	sqlQuery := `SELECT 
		id, 
		name, 
		cpf, 
		email, 
		street, 
		number, 
		zip, 
		status, 
		complement, 
		COALESCE(card_token, '') AS card_token, 
		COALESCE(payment_method_id, '') AS payment_method_id, 
		customer_id, 
		phone, 
		profile_image FROM responsible WHERE email = $1 LIMIT 1`
	var responsible entity.Responsible
	err := rr.db.QueryRow(sqlQuery, *email).Scan(
		&responsible.ID,
		&responsible.Name,
		&responsible.CPF,
		&responsible.Email,
		&responsible.Address.Street,
		&responsible.Address.Number,
		&responsible.Address.ZIP,
		&responsible.Status,
		&responsible.Address.Complement,
		&responsible.CreditCard.CardToken,
		&responsible.PaymentMethodId,
		&responsible.CustomerId,
		&responsible.Phone,
		&responsible.ProfileImage,
	)
	if err != nil || err == sql.ErrNoRows {
		return nil, err
	}
	return &responsible, nil
}
