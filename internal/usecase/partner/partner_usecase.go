package partner

import (
	"context"

	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/repository"
)

type PartnerUseCase struct {
	partnerRepository repository.IPartnerRepository
}

func NewPartnerUseCase(pr repository.IPartnerRepository) *PartnerUseCase {
	return &PartnerUseCase{
		partnerRepository: pr,
	}
}

func (pu *PartnerUseCase) IsPartner(ctx context.Context, cnh, cnpj *string) (bool, error) {
	return pu.partnerRepository.IsPartner(ctx, cnh, cnpj)
}

func (pu *PartnerUseCase) Get(ctx context.Context, id *string) (*entity.Partner, error) {
	return pu.partnerRepository.Get(ctx, id)
}

func (pu *PartnerUseCase) FindAllByCnh(ctx context.Context, cnh *string) ([]entity.Partner, error) {
	return pu.partnerRepository.FindAllByCnh(ctx, cnh)
}

func (pu *PartnerUseCase) FindAllByCnpj(ctx context.Context, cnpj *string) ([]entity.Partner, error) {
	return pu.partnerRepository.FindAllByCnpj(ctx, cnpj)
}

func (pu *PartnerUseCase) Delete(ctx context.Context, id *string) error {
	return pu.partnerRepository.Delete(ctx, id)
}