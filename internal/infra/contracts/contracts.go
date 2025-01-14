package contracts

import (
	"time"

	"github.com/jinzhu/gorm"
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

type CacheIface interface {
	Set(key string, value any, expiration time.Duration) error
	Get(key string) (string, error)
	Expire(key string, expiration time.Duration) (bool, error)
}

type Logger interface {
	Infof(format string, args ...zap.Field)
	Errorf(format string, args ...zap.Field)
}
