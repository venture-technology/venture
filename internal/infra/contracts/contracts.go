package contracts

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/venture-technology/venture/internal/entity"
	"go.uber.org/zap"
)

type S3Iface interface {
	// Save a file like png, to more performance
	Save(path, filename string, file []byte) (string, error)
	List(path string) ([]string, error)
	// Save a file like anytype file
	SaveWithType(path, filaneme string, file []byte, contentType string) (string, error)

	// types to return

	HTML() string
	PDF() string
	PNG() string
}

type SESIface interface {
	SendEmail(email *entity.Email) error
}

type PostgresIface interface {
	Client() *gorm.DB
	Close() error
}

type Cacher interface {
	Set(key string, value any, expiration time.Duration) error
	Get(key string) (string, error)
	Expire(key string, expiration time.Duration) (bool, error)
}

type Logger interface {
	Infof(format string, args ...zap.Field)
	Errorf(format string, args ...zap.Field)
}

type Converters interface {
	ConvertPDFtoHTML(htmlFile []byte, contractProperty entity.ContractProperty) ([]byte, error)
}
