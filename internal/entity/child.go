package entity

type Child struct {
	ID          int         `json:"id,omitempty"`
	Name        string      `json:"name" validate:"required"`
	RG          string      `json:"rg" validate:"required" example:"559378847"`
	Responsible Responsible `json:"responsible,omitempty" validate:"required"`
	Shift       string      `json:"shift" validate:"oneof='matutino' 'vespertino' 'noturno'"`
}
