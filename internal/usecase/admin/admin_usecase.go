package admin

import (
	"context"

	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/repository"
	"go.uber.org/zap"
)

type AdminUseCase struct {
	adminRepository repository.IAdminRepository
	logger          *zap.Logger
}

func NewAdminUseCase(adminRepository repository.IAdminRepository, logger *zap.Logger) *AdminUseCase {
	return &AdminUseCase{
		adminRepository: adminRepository,
		logger:          logger,
	}
}

func (au *AdminUseCase) NewApiKey(ctx context.Context, name string) error {
	return au.adminRepository.NewApiKey(ctx, uuid.New().String(), name)
}
