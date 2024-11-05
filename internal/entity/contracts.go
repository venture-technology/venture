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
	Record             uuid.UUID              `json:"record,omitempty"`
	Status             string                 `json:"status" validate:"oneof='currently' 'canceled' 'expired'"`
	Driver             Driver                 `json:"driver"`
	School             School                 `json:"school"`
	Child              Child                  `json:"child"`
	StripeSubscription StripeSubscription     `json:"stripe"`
	CreatedAt          time.Time              `json:"created_at"`
	ExpireAt           time.Time              `json:"expire_at"`
	Invoices           map[string]InvoiceInfo `json:"invoices"`
	Amount             float64                `json:"amount" validate:"required"`
	Months             int64                  `json:"months,omitempty"`
}

type InvoiceInfo struct {
	ID              string `json:"id"`
	Status          string `json:"status"`
	AmountDue       int64  `json:"amount_due"`
	AmountRemaining int64  `json:"amount_remaining"`
	Month           string `json:"month"`
	Date            string `json:"date"`
}

type InvoiceRemaining struct {
	InvoiceValue float64 `json:"value"`
	Quantity     float64 `json:"quantity"`
	Remaining    float64 `json:"remaining"`
	Fines        float64 `json:"fine"`
}

type SubscriptionInfo struct {
	ID     string `json:"subscription_id"`
	Status string `json:"subscription_status"`
}

func (c *Contract) ValidateAmount() bool {
	return c.Amount != 0
}
