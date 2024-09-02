package repository

import (
	"context"
	"database/sql"

	"github.com/venture-technology/venture/internal/entity"
)

type IPartnerRepository interface {
	Get(ctx context.Context, id *string) (*entity.Partner, error)
	FindAllByCnpj(ctx context.Context, cnpj *string) ([]entity.Partner, error)
	FindAllByCnh(ctx context.Context, cnh *string) ([]entity.Partner, error)
	IsPartner(ctx context.Context, cnh, cnpj *string) (bool, error)
	Delete(ctx context.Context, id *string) error
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

func (pr *PartnerRepository) Get(ctx context.Context, id *string) (*entity.Partner, error) {
	sqlQuery := `
        SELECT 
            p.record, p.created_at,
            d.name AS driver_name, d.email AS driver_email, d.qrcode AS driver_qrcode, d.phone AS driver_phone,
            s.name AS school_name, s.email AS school_email
        FROM 
            partners p
        JOIN 
            drivers d ON p.driver_id = d.cnh
        JOIN 
            schools s ON p.school_id = s.cnpj
        WHERE 
            p.record = $1
        LIMIT 1;
    `
	var partner entity.Partner
	err := pr.db.QueryRowContext(ctx, sqlQuery, *id).Scan(
		&partner.Record,
		&partner.CreatedAt,
		&partner.Driver.Name,
		&partner.Driver.Email,
		&partner.Driver.QrCode,
		&partner.Driver.Phone,
		&partner.School.Name,
		&partner.School.Email,
	)
	if err != nil {
		return nil, err
	}
	return &partner, nil
}

func (pr *PartnerRepository) FindAllByCnpj(ctx context.Context, cnpj *string) ([]entity.Partner, error) {
	sqlQuery := `
        SELECT 
            p.record, p.created_at,
            d.name AS driver_name, d.email AS driver_email, d.qrcode AS driver_qrcode, d.phone AS driver_phone,
            s.name AS school_name, s.email AS school_email
        FROM 
            partners p
        JOIN 
            drivers d ON p.driver_id = d.cnh
        JOIN 
            schools s ON p.school_id = s.cnpj
        WHERE 
            p.school_id = $1;
    `
	rows, err := pr.db.QueryContext(ctx, sqlQuery, *cnpj)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var partners []entity.Partner
	for rows.Next() {
		var partner entity.Partner
		if err := rows.Scan(
			&partner.Record,
			&partner.CreatedAt,
			&partner.Driver.Name,
			&partner.Driver.Email,
			&partner.Driver.QrCode,
			&partner.Driver.Phone,
			&partner.School.Name,
			&partner.School.Email,
		); err != nil {
			return nil, err
		}
		partners = append(partners, partner)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return partners, nil
}

func (pr *PartnerRepository) FindAllByCnh(ctx context.Context, cnh *string) ([]entity.Partner, error) {
	sqlQuery := `
        SELECT 
            p.record, p.created_at,
            d.name AS driver_name, d.email AS driver_email, d.qrcode AS driver_qrcode, d.phone AS driver_phone,
            s.name AS school_name, s.email AS school_email
        FROM 
            partners p
        JOIN 
            drivers d ON p.driver_id = d.cnh
        JOIN 
            schools s ON p.school_id = s.cnpj
        WHERE 
            p.driver_id = $1;
    `
	rows, err := pr.db.QueryContext(ctx, sqlQuery, *cnh)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var partners []entity.Partner
	for rows.Next() {
		var partner entity.Partner
		if err := rows.Scan(
			&partner.Record,
			&partner.CreatedAt,
			&partner.Driver.Name,
			&partner.Driver.Email,
			&partner.Driver.QrCode,
			&partner.Driver.Phone,
			&partner.School.Name,
			&partner.School.Email,
		); err != nil {
			return nil, err
		}
		partners = append(partners, partner)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return partners, nil
}

func (pr *PartnerRepository) Delete(ctx context.Context, id *string) error {
	sqlQuery := `DELETE FROM partners WHERE record = $1;`
	_, err := pr.db.ExecContext(ctx, sqlQuery, *id)
	if err != nil {
		return err
	}
	return nil
}
