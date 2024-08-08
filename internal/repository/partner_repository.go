package repository

import (
	"context"
	"database/sql"

	"github.com/venture-technology/venture/models"
)

type IPartnerRepository interface {
	IsPartner(ctx context.Context, partner *models.Partner) (bool, error)
	CreatePartners(ctx context.Context, partner *models.Partner) error
	GetPartnersByDriver(ctx context.Context) ([]models.Partner, error)
	GetPartnersBySchool(ctx context.Context) ([]models.Partner, error)
	DeletePartner(ctx context.Context) error
}

type PartnerRepository struct {
	db *sql.DB
}

func NewPartnerRepository(db *sql.DB) *PartnerRepository {
	return &PartnerRepository{
		db: db,
	}
}

func (pr *PartnerRepository) IsPartner(ctx context.Context, partner *models.Partner) (bool, error) {
	return true, nil
}

func (pr *PartnerRepository) CreatePartners(ctx context.Context, partner *models.Partner) error {
	return nil
}

func (pr *PartnerRepository) GetPartnersByDriver(ctx context.Context) ([]models.Partner, error) {
	return nil, nil
}

func (pr *PartnerRepository) GetPartnersBySchool(ctx context.Context) ([]models.Partner, error) {
	return nil, nil
}

func (pr *PartnerRepository) DeletePartner(ctx context.Context) error {
	return nil
}
