package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type IAdminRepository interface {
	NewApiKey(ctx context.Context, id *uuid.UUID) error
}

type AdminRepository struct {
	db *sql.DB
}

func NewAdminRepository(db *sql.DB) *AdminRepository {
	return &AdminRepository{
		db: db,
	}
}

func (ar *AdminRepository) NewApiKey(ctx context.Context, id *uuid.UUID) error {
	sqlQuery := `INSERT INTO api_keys (id) VALUES ($1)`
	_, err := ar.db.Exec(sqlQuery, id)
	return err

}
