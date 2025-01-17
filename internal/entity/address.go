package entity

type Address struct {
	Street     string `json:"street,omitempty" validate:"required"`
	Number     string `json:"number,omitempty" validate:"required"`
	Zip        string `json:"zip,omitempty" validate:"required"`
	Complement string `json:"complement,omitempty"`
}
