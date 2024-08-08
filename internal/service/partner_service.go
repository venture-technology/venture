package service

import (
	"context"

	"github.com/venture-technology/venture/internal/repository"
	"github.com/venture-technology/venture/models"
)

type PartnerService struct {
	partnerrepository repository.IPartnerRepository
}

func NewPartnerService(partnerrepository repository.IPartnerRepository) *PartnerService {
	return &PartnerService{
		partnerrepository: partnerrepository,
	}
}

func (ps *PartnerService) IsPartner(ctx context.Context, partner *models.Partner) (bool, error) {
	return ps.partnerrepository.IsPartner(ctx, partner)
}

func (ps *PartnerService) CreatePartners(ctx context.Context, partner *models.Partner) error {
	return ps.partnerrepository.CreatePartners(ctx, partner)
}

func (ps *PartnerService) GetPartnersByDriver(ctx context.Context) ([]models.Partner, error) {
	return ps.partnerrepository.GetPartnersByDriver(ctx)
}

func (ps *PartnerService) GetPartnersBySchool(ctx context.Context) ([]models.Partner, error) {
	return ps.partnerrepository.GetPartnersBySchool(ctx)
}

func (ps *PartnerService) DeletePartner(ctx context.Context) error {
	return ps.partnerrepository.DeletePartner(ctx)
}
