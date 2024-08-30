package repository

import (
	"context"
	"database/sql"

	"github.com/venture-technology/venture/internal/entity"
)

type IPartnerRepository interface {
	// Get()
	// FindAllByCnpj()
	// FindAllByCnh()
	IsPartner(ctx context.Context, cnh, cnpj *string) (bool, error)
	// Delete()
}

type PartnerRepository struct {
	db *sql.DB
}

func NewPartnerRepository(db *sql.DB) *PartnerRepository {
	return &PartnerRepository{
		db: db,
	}
}

func (pr *PartnerRepository) IsPartner(ctx context.Context, cnh, cnpj *string) (bool, error) {
	sqlQuery := `SELECT record, driver_id, school_id, created_at FROM partners WHERE driver_id = $1 AND school_id = $2 LIMIT 1`
	var partner entity.Partner
	err := pr.db.QueryRow(sqlQuery, *cnh, *cnpj).Scan(
		&partner.Record,
		&partner.Driver.CNH,
		&partner.School.CNPJ,
		&partner.CreatedAt,
	)
	if err != nil && err != sql.ErrNoRows {
		return true, err
	}

	if err == sql.ErrNoRows {
		return false, nil
	}

	return false, nil
}
