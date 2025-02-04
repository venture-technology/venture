package entity

import "time"

type Partner struct {
	Record    int       `gorm:"primary_key;auto_increment" json:"record"`
	SchoolID  string    `json:"school_id"`
	DriverID  string    `json:"driver_id"`
	Driver    Driver    `json:"driver"`
	School    School    `json:"school"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
