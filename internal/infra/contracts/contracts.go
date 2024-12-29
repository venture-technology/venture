package contracts

import (
	"github.com/jinzhu/gorm"
	"github.com/redis/go-redis/v9"
	"github.com/venture-technology/venture/internal/entity"
	"go.uber.org/zap"
)

type S3Iface interface {
	Save(path, filename string, file []byte) (string, error)
	List(path string) ([]string, error)
}

type SESIface interface {
	SendEmail(email *entity.Email) error
}

type PostgresIface interface {
	Client() *gorm.DB
	Close() error
}

type RedisIface interface {
	Client() *redis.Client
}

type Logger interface {
	Infof(format string, args ...zap.Field)
	Errorf(format string, args ...zap.Field)
}
