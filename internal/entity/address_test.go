package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/venture-technology/venture/internal/entity"
)

func TestGetFullAddress_WithComplement(t *testing.T) {
	address := entity.Address{
		Street:       "Rua das Flores",
		Number:       "123",
		Complement:   "Apto 45",
		City:         "São Paulo",
		State:        "SP",
		Neighborhood: "Jardins",
		Zip:          "01423-001",
	}

	expected := "Rua das Flores, 123 - Apto 45, Jardins (São Paulo) - SP, 01423-001"
	assert.Equal(t, expected, address.GetFullAddress())
}

func TestGetFullAddress_WithoutComplement(t *testing.T) {
	address := entity.Address{
		Street:       "Avenida Paulista",
		Number:       "1000",
		Complement:   "",
		City:         "São Paulo",
		State:        "SP",
		Neighborhood: "Bela Vista",
		Zip:          "01310-000",
	}

	expected := "Avenida Paulista, 1000 - Bela Vista, São Paulo - SP, 01310-000"
	assert.Equal(t, expected, address.GetFullAddress())
}
