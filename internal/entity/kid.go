package entity

import "time"

type Kid struct {
	ID   int    `gorm:"primary_key;auto_increment" json:"id,omitempty"`
	Name string `json:"name,omitempty" validate:"required"`
	RG   string `json:"rg" validate:"required" example:"559378847"`
	// responsible_id field exists to not have conflict with the responsible entity
	ResponsibleCPF string      `gorm:"column:responsible_id;type:varchar(11);not null" json:"responsible_id,omitempty"` // Aqui mapeia o CPF para o campo `responsible_id`
	Responsible    Responsible `gorm:"foreignKey:ResponsibleCPF;references:CPF" json:"responsible,omitempty"`           // Chave estrangeira corretamente configurada
	Shift          string      `json:"shift,omitempty" validate:"oneof='matutino' 'vespertino' 'noturno'"`
	ProfileImage   string      `json:"profile_image,omitempty"`
	CreatedAt      time.Time   `json:"created_at,omitempty"`
	UpdatedAt      time.Time   `json:"updated_at,omitempty"`
}
