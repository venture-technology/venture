package repository

import (
	"context"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type IAdminRepository interface {
	NewApiKey(ctx context.Context, id, name string) error
	GetApiKey(ctx context.Context, id string) (string, error)
}

type AdminRepository struct {
	rdb    *redis.Client
	logger *zap.Logger
}

func NewAdminRepository(rdb *redis.Client, logger *zap.Logger) *AdminRepository {
	return &AdminRepository{
		rdb:    rdb,
		logger: logger,
	}
}

func (ar *AdminRepository) NewApiKey(ctx context.Context, id, name string) error {
	return ar.rdb.Set(ctx, id, name, 0).Err()
}

func (ar *AdminRepository) GetApiKey(ctx context.Context, id string) (string, error) {
	return ar.rdb.Get(ctx, id).Result()
}
