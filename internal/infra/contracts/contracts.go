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
	// List files in a bucket
	List(bucket, path string) ([]string, error)

	// copy a file from a bucket
	Copy(bucket, objectKey string) ([]byte, error)

	// Save a file like anytype file
	Save(bucket, path, filename string, file []byte, contentType string) (string, error)
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
	ConvertHTMLtoPDF(htmlFile []byte, contractProperty value.CreateContractParams) ([]byte, error)
}

// Interface to use Simple Queue Service
type Queue interface {
	// Send message to queue
	SendMessage(queue, message string) error

	// Send message to fifo queue using message group ID
	SendFifoMessage(queue, message, group string) error

	// Gets the messages from the queue, need the CreateMessage format.
	PullMessages(queue string) ([]*value.CreateMessage, error)

	// Delete a specific message of queue by identifier
	DeleteMessage(queue, identifier string) error
}

type Workers interface {
	Enqueue(requestParams *value.CreateContractParams) error
}
