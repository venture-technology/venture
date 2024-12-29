package entity

import "time"

type Partner struct {
	ID        int       `gorm:"primary_key;auto_increment" json:"record"`
	Driver    Driver    `json:"driver"`
	School    School    `json:"school"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
