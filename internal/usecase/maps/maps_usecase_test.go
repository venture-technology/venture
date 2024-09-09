package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/venture-technology/venture/internal/usecase"
)

func TestMapsUseCase_CalculatePrice(t *testing.T) {

	origin := "Rua Masato Sakai, 180, 008538300"
	destination := "Avenida Barão de Alagoas, 223, 08120000"

	km, err := usecase.GetDistance(origin, destination)

	if err != nil {
		t.Error(err)
	}

	value := usecase.CalculateContract(*km, 200)

	assert.Equal(t, 740.0, value)
}
