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
	Record             uuid.UUID          `json:"record,omitempty"`
	Status             string             `json:"status" validate:"oneof='currently' 'canceled' 'expired'"`
	Driver             Driver             `json:"driver"`
	School             School             `json:"school"`
	Child              Child              `json:"child"`
	StripeSubscription StripeSubscription `json:"stripe"`
	CreatedAt          time.Time          `json:"created_at"`
	ExpireAt           time.Time          `json:"expire_at"`
	Invoices           []InvoiceInfo      `json:"invoices"`
	Amount             float64            `json:"amount"`
	Months             int64              `json:"months,omitempty"`
}

type InvoiceInfo struct {
	ID              string `json:"invoice_info_id"`
	Status          string `json:"invoice_info_status"`
	AmountDue       int64  `json:"invoice_info_amount_due"`
	AmountRemaining int64  `json:"invoice_info_amount_remaining"`
}

type InvoiceRemaining struct {
	InvoiceValue float64 `json:"invoice_value"`
	Quantity     float64 `json:"invoice_quantity"`
	Remaining    float64 `json:"invoice_remaining"`
	Fines        float64 `json:"invoice_fine"`
}

type SubscriptionInfo struct {
	ID     string `json:"subscription_id"`
	Status string `json:"subscription_status"`
}
