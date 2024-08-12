package repository

import (
	"context"
	"database/sql"

	"github.com/venture-technology/venture/internal/entity"
)

type IEmailRepository interface {
	CreateRecord(ctx context.Context, email *entity.Email) error
}

type EmailRepository struct {
	db *sql.DB
}

func NewEmailRepository(db *sql.DB) *EmailRepository {
	return &EmailRepository{
		db: db,
	}
}

func (er *EmailRepository) CreateRecord(ctx context.Context, email *entity.Email) error {
	sqlQuery := `INSERT INTO email_records (recipient, subject, body) VALUES ($1, $2, $3)`
	_, err := er.db.Exec(sqlQuery, email.Recipient, email.Subject, email.Body)
	return err
}
