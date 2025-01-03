package persistence

import (
	"context"

	"github.com/venture-technology/venture/internal/infra/contracts"
)

type AdminRepositoryImpl struct {
	Redis contracts.RedisIface
}

func (ar AdminRepositoryImpl) NewApiKey(ctx context.Context, id, name string) error {
	return ar.Redis.Client().Set(ctx, id, name, 0).Err()
}

func (ar AdminRepositoryImpl) GetApiKey(ctx context.Context, id string) (string, error) {
	return ar.Redis.Client().Get(ctx, id).Result()
}
