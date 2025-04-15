package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

const CacheTTL time.Duration = 0

type CacheImpl struct {
	client *redis.Client
}

func NewCacheImpl() *CacheImpl {
	return &CacheImpl{
		client: redis.NewClient(&redis.Options{
			Addr:     viper.GetString("REDIS_ADDRESS"),
			Password: viper.GetString("REDIS_PASSWORD"),
		}),
	}
}

func (c CacheImpl) Set(key string, value any, expiration time.Duration) error {
	serializedValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	statusCmd := c.client.Set(context.Background(), key, serializedValue, expiration)
	_, err = statusCmd.Result()
	if err != nil {
		return err
	}

	return nil
}

func (c CacheImpl) Get(key string) (string, error) {
	statusCmd := c.client.Get(context.Background(), key)

	result, err := statusCmd.Result()
	if err != nil {
		return "", err
	}

	return result, nil
}

func (c CacheImpl) Expire(key string, expiration time.Duration) (bool, error) {
	statusCmd := c.client.Expire(context.Background(), key, expiration)
	err := statusCmd.Err()
	if err != nil {
		return false, err
	}

	result, err := statusCmd.Result()
	if err != nil {
		return false, err
	}

	return result, nil
}
