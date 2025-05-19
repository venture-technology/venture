package entity

import (
	"fmt"
	"strings"
)

type Address struct {
	Street       string `json:"street,omitempty" validate:"required"`
	Number       string `json:"number,omitempty" validate:"required"`
	Zip          string `json:"zip,omitempty" validate:"required"`
	Complement   string `json:"complement,omitempty"`
	City         string `json:"city,omitempty" validate:"required"`
	State        string `json:"state,omitempty" validate:"required"`
	Neighborhood string `json:"neighborhood,omitempty" validate:"required"`
}

func (a Address) GetFullAddress() string {
	parts := []string{}

	if a.Street != "" {
		parts = append(parts, a.Street)
	}

	if a.Number != "" {
		parts = append(parts, a.Number)
	}

	if a.Complement != "" {
		parts = append(parts, fmt.Sprintf("(%s)", a.Complement))
	}

	if a.Neighborhood != "" {
		parts = append(parts, a.Neighborhood)
	}

	if a.City != "" {
		parts = append(parts, a.City)
	}

	if a.State != "" {
		parts = append(parts, a.State)
	}

	if a.Zip != "" {
		parts = append(parts, a.Zip)
	}

	return strings.Join(parts, ", ")
}
