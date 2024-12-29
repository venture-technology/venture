package application

import (
	"context"

	"github.com/google/uuid"
	"github.com/venture-technology/venture/internal/infra/contracts"
	"github.com/venture-technology/venture/internal/infra/persistence"
)

type GenerateApiKeyAdminService struct {
	repositories *persistence.RedisRepositories
	logger       contracts.Logger
}

func NewGenerateApiKeyAdminService(
	repositories *persistence.RedisRepositories,
	logger contracts.Logger,
) *GenerateApiKeyAdminService {
	return &GenerateApiKeyAdminService{
		repositories: repositories,
		logger:       logger,
	}
}

func (gakas *GenerateApiKeyAdminService) NewApiKey(ctx context.Context, name string) error {
	return gakas.repositories.AdminRepository.NewApiKey(ctx, uuid.New().String(), name)
}
