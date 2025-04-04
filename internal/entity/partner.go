package entity

import "time"

type Partner struct {
	ID         int       `gorm:"primary_key;auto_increment" json:"id"`
	SchoolCNPJ string    `json:"school_cnpj"`
	DriverCNH  string    `json:"driver_cnh"`
	Driver     Driver    `json:"driver"`
	School     School    `json:"school"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}
