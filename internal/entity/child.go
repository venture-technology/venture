package entity

import "time"

type Child struct {
	ID           int         `gorm:"primary_key;auto_increment" json:"id,omitempty"`
	Name         string      `json:"name,omitempty" validate:"required"`
	RG           string      `json:"rg" validate:"required" example:"559378847"`
	Responsible  Responsible `json:"responsible,omitempty" validate:"required"`
	Shift        string      `json:"shift,omitempty" validate:"oneof='matutino' 'vespertino' 'noturno'"`
	ProfileImage string      `json:"profile_image,omitempty"`
	CreatedAt    time.Time   `json:"created_at,omitempty"`
	UpdatedAt    time.Time   `json:"updated_at,omitempty"`
}
