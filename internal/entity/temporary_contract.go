package entity

import (
	"database/sql"
	"time"
)

type TempContract struct {
	ID                  int           `gorm:"primary_key;auto_increment" json:"id"`
	SigningURL          string        `json:"signing_url"`
	Status              string        `json:"status"`
	DriverCNH           string        `json:"driver_cnh"`
	SchoolCNPJ          string        `json:"school_cnpj"`
	KidRG               string        `json:"kid_rg"`
	ResponsibleCPF      string        `json:"responsible_cpf"`
	SignatureRequestID  string        `json:"signature_request_id"`
	UUID                string        `json:"uuid"`
	CreatedAt           int64         `json:"created_at"`               // epoch time
	ExpiredAt           int64         `json:"expires_at"`               // epoch time
	DriverSignedAt      sql.NullInt64 `json:"driver_assigned_at"`       // epoch time
	ResponsibleSignedAt sql.NullInt64 `json:"responsibles_assigned_at"` // epoch time
}

// struct to must be used to create a pre contract at create contract route to checks.
type TemporaryContract struct {
	AmountCents         int64         `json:"amount_cents"`
	AmountAnualCents    int64         `json:"amount_anual_cents"`
	DriverAmount        int64         `json:"driver_amount"`
	CreatedAt           int64         `json:"created_at"`
	ExpireAt            int64         `json:"expire_at"`
	UUID                string        `json:"uuid"`
	ResponsibleCPF      string        `json:"responsible_cpf`
	ResponsibleName     string        `json:"responsible_name`
	ResponsibleAddr     string        `json:"responsible_addr`
	ResponsibleEmail    string        `json:"responsible_email`
	ResponsiblePhone    string        `json:"responsible_phone`
	KidRG               string        `json:"kid_rg"`
	KidName             string        `json:"kid_name`
	KidShift            string        `json:"kid_shift`
	DriverName          string        `json:"driver_name`
	DriverEmail         string        `json:"driver_email"`
	DriverCNH           string        `json:"driver_cnh`
	SchoolCNPJ          string        `json:"school_cnpj"`
	SchoolName          string        `json:"school_name`
	SchoolAddr          string        `json:"school_addr`
	FileURL             string        `json:"file_url"`
	DateTime            string        `json:"date_time"`
	Time                time.Time     `json:"time"`
	DriverSignedAt      sql.NullInt64 `json:"driver_signed_at"`
	ResponsibleSignedAt sql.NullInt64 `json:"responsible_signed_at"`
}
