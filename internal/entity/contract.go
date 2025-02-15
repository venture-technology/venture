package entity

import (
	"time"

	"github.com/google/uuid"
)

type StripeSubscription struct {
	Title       string `json:"title_subscription"`
	ID          string `json:"subscription_id"`
	Description string `json:"description_subscription,omitempty"`
	Price       string `json:"price_id"`
	Product     string `json:"product_id"`
}

type Contract struct {
	ID                 int                    `gorm:"primary_key;auto_increment" json:"id"`
	Record             uuid.UUID              `json:"record,omitempty"`
	Status             string                 `json:"status" validate:"oneof='currently' 'canceled' 'expired'"`
	Driver             Driver                 `json:"driver"`
	School             School                 `json:"school"`
	Kid                Kid                    `json:"kid"`
	Responsible        Responsible            `json:"responsible"`
	StripeSubscription StripeSubscription     `json:"stripe"`
	Invoices           map[string]InvoiceInfo `json:"invoices"`
	Amount             float64                `json:"amount" validate:"required"`
	AnualAmount        float64                `json:"anual_amount"`
	Months             int64                  `json:"months,omitempty"`
	CreatedAt          time.Time              `json:"created_at,omitempty"`
	UpdatedAt          time.Time              `json:"updated_at,omitempty"`
	ExpireAt           time.Time              `json:"expire_at"`
	ContractUrl        string                 `json:"contract_url,omitempty"`
}

type InvoiceInfo struct {
	ID              string `json:"id"`
	Status          string `json:"status"`
	AmountDue       int64  `json:"amount_due"`
	AmountRemaining int64  `json:"amount_remaining"`
	Month           string `json:"month"`
	Date            string `json:"date"`
}

type SubscriptionInfo struct {
	ID     string `json:"subscription_id"`
	Status string `json:"subscription_status"`
}

func (c *Contract) ValidateAmount() bool {
	return c.Amount != 0
}

// contract property is a contract struct with dropbox data, like signed url, pdf contract and etc.
type ContractProperty struct {
	URL      string    `json:"url"` // S3 Link to Dropbox Download and use to sign
	UUID     string    `json:"uuid"`
	Time     time.Time `json:"time"`
	DateTime string    `json:"date_time"`
	Contract Contract  `json:"contract"`
}
