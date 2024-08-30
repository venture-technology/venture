package partner

import "github.com/venture-technology/venture/internal/repository"

type PartnerUseCase struct {
	partnerRepository repository.IPartnerRepository
}

func NewPartnerUseCase(pr repository.IPartnerRepository) *PartnerUseCase {
	return &PartnerUseCase{
		partnerRepository: pr,
	}
}
