package entity

import (
	"time"
)

type StripeSubscription struct {
	Title       string `json:"title_subscription"`
	ID          string `json:"subscription_id"`
	Description string `json:"description_subscription,omitempty"`
	Price       string `json:"price_id"`
	Product     string `json:"product_id"`
}

type ContractParams struct {
	Driver             Driver                 `json:"driver"`
	School             School                 `json:"school"`
	Kid                Kid                    `json:"kid"`
	Responsible        Responsible            `json:"responsible"`
	StripeSubscription StripeSubscription     `json:"stripe"`
	Invoices           map[string]InvoiceInfo `json:"invoices"`
	Amount             float64                `json:"amount" validate:"required"`
	AnualAmount        float64                `json:"anual_amount"`
}

type Contract struct {
	ID                   int     `gorm:"primary_key;auto_increment" json:"id"`
	UUID                 string  `json:"record,omitempty"`
	Status               string  `json:"status" validate:"oneof='currently' 'canceled' 'expired'"`
	StripeSubscriptionID string  `json:"stripe_subscription_id"`
	StripePriceID        string  `json:"stripe_price_id"`
	StripeProductID      string  `json:"stripe_product_id"`
	SigningURL           string  `json:"dropbox_signing_url"`
	DriverCNH            string  `json:"driver_cnh"`
	SchoolCNPJ           string  `json:"school_cnpj"`
	KidRG                string  `json:"kid_rg"`
	ResponsibleCPF       string  `json:"responsible_cpf"`
	CreatedAt            int64   `json:"created_at,omitempty"`
	UpdatedAt            int64   `json:"updated_at,omitempty"`
	ExpireAt             int64   `json:"expire_at"`
	Amount               float64 `json:"amount" validate:"required"`
	AnualAmount          float64 `json:"anual_amount"`

	// Relações
	Driver      Driver      `gorm:"foreignKey:DriverCNH;references:CNH" json:"driver"`
	School      School      `gorm:"foreignKey:SchoolCNPJ;references:CNPJ" json:"school"`
	Kid         Kid         `gorm:"foreignKey:KidRG;references:RG" json:"kid"`
	Responsible Responsible `gorm:"foreignKey:ResponsibleCPF;references:CPF" json:"responsible"`
}

type InvoiceInfo struct {
	ID          string  `json:"id"`
	Status      string  `json:"status"`
	Amount      float64 `json:"amount"`
	AmountCents int64   `json:"amount_cents"`
	Month       string  `json:"month"`
	Date        string  `json:"date"`
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
	URL            string         `json:"url"` // S3 Link to Dropbox Download and use to sign
	UUID           string         `json:"uuid"`
	Time           time.Time      `json:"time"`
	DateTime       string         `json:"date_time"`
	ContractParams ContractParams `json:"contract"`
}

// a way to return currently contracts made little bit of sql queries
