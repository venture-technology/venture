package admin

import (
	"context"

	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/repository"
)

type AdminUseCase struct {
	adminRepository repository.IAdminRepository
}

func NewAdminUseCase(adminRepository repository.IAdminRepository) *AdminUseCase {
	return &AdminUseCase{
		adminRepository: adminRepository,
	}
}

func (au *AdminUseCase) NewApiKey(ctx context.Context, name string) error {
	return au.adminRepository.NewApiKey(ctx, uuid.New().String(), name)
}
