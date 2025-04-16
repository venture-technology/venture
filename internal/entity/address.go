package entity

import "fmt"

type Address struct {
	Street       string `json:"street,omitempty" validate:"required"`
	Number       string `json:"number,omitempty" validate:"required"`
	Zip          string `json:"zip,omitempty" validate:"required"`
	Complement   string `json:"complement,omitempty"`
	City         string `json:"city,omitempty" validate:"required"`
	State        string `json:"provincy,omitempty" validate:"required"`
	Neighborhood string `json:"neighborhood,omitempty" validate:"required"`
}

func (a Address) GetFullAddress() string {
	if a.Complement == "" {
		return fmt.Sprintf(
			"%s, %s - %s, %s - %s, %s",
			a.Street,
			a.Number,
			a.Neighborhood,
			a.City,
			a.State,
			a.Zip,
		)
	}

	return fmt.Sprintf(
		"%s, %s - %s, %s (%s) - %s, %s",
		a.Street,
		a.Number,
		a.Complement,
		a.Neighborhood,
		a.City,
		a.State,
		a.Zip,
	)
}
