package entity

type Address struct {
	Street     string `json:"street" validate:"required"`
	Number     string `json:"number" validate:"required"`
	ZIP        string `json:"zip" validate:"required"`
	Complement string `json:"complement,omitempty"`
}
