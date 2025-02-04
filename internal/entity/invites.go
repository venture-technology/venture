package entity

import "time"

type Invite struct {
	ID         int       `gorm:"primary_key;auto_increment" json:"id,omitempty"`
	School     School    `json:"school,omitempty"`
	Driver     Driver    `json:"driver,omitempty"`
	SchoolID   string    `gorm:"column:requester" json:"school_id,omitempty"`
	DriverID   string    `gorm:"column:guester" json:"driver_id,omitempty"`
	Status     string    `json:"status,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
	AcceptedAt time.Time `json:"accepted_at,omitempty"`
}
