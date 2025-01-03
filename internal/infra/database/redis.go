package database

import (
	"github.com/redis/go-redis/v9"
	"github.com/venture-technology/venture/config"
)

type RedisImpl struct {
	c *redis.Client
}

func NewRedisImpl(config config.Config) RedisImpl {
	redisImpl := RedisImpl{}

	opts := &redis.Options{
		Addr:     config.Cache.Address,
		Password: config.Cache.Password,
	}

	redisImpl.c = redis.NewClient(opts)

	return redisImpl
}

func (r RedisImpl) Client() *redis.Client {
	return r.c
}
