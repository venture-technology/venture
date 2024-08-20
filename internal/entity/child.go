package entity

type Child struct {
	ID          int         `json:"id"`
	Name        string      `json:"name" validate:"required"`
	RG          string      `json:"rg" validate:"required" example:"559378847"`
	Responsible Responsible `json:"responsible" validate:"required"`
	Shift       string      `json:"shift" validate:"oneof='matutino' 'vespertino' 'noturno'"`
}
