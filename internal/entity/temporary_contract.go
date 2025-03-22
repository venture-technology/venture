package entity

import (
	"database/sql"
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
