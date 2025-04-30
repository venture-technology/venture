package contracts

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/value"
	"go.uber.org/zap"
)

// Interace to communicate with S3 Bucket Services.
type S3Iface interface {
	// Save a file like png, to more performance
	Save(bucket, path, filename string, file []byte) (string, error)

	// List files in a bucket
	List(bucket, path string) ([]string, error)

	// Save a file like anytype file
	SaveWithType(bucket, path, filaneme string, file []byte, contentType string) (string, error)

	// Save file like html
	HTML() string

	// Save file like pdf
	PDF() string

	// Save file like png
	PNG() string
}

// Interface to use Simple Email Service
type SESIface interface {
	// SendEmail send email to user
	SendEmail(email *entity.Email) error
}

type PostgresIface interface {
	Client() *gorm.DB
	Close() error
}

// Interface to use Redis
type Cacher interface {
	Set(key string, value any, expiration time.Duration) error
	Get(key string) (string, error)
	Expire(key string, expiration time.Duration) (bool, error)
}

type Logger interface {
	Infof(format string, args ...zap.Field)
	Errorf(format string, args ...zap.Field)
}

// Interface to convert files
type Converters interface {
	ConvertPDFtoHTML(htmlFile []byte, contractProperty entity.ContractProperty) ([]byte, error)
}

// Interface to use Simple Queue Service
type Queue interface {
	// Send message to queue
	SendMessage(queue, message string) error

	// Gets the messages from the queue, need the CreateMessage format.
	PullMessages(queue string) ([]value.CreateMessage, error)

	// Delete a specific message of queue by identifier
	DeleteMessage(queue, identifier string) error
}
