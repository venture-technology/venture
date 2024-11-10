package maps

import (
	"context"
	"testing"

	"github.com/venture-technology/venture/internal/domain/adapter"
)

func TestMapsUseCase_CalculatePrice(t *testing.T) {

	t.Run("when calculate price is sucess", func(t *testing.T) {
		googleAdapter := adapter.NewGoogleAdapter()

		origin := "Rua Masato Sakai, 180, 008538300"
		destination := "Avenida Barão de Alagoas, 223, 08120000"

		useCase := NewMapsUseCase(googleAdapter)

		_, err := useCase.CalculatePrice(context.Background(), origin, destination, 185.0)
		if err != nil {
			t.Errorf("Error: %s", err)
		}
	})

	t.Run("when calculate price is fail", func(t *testing.T) {
		googleAdapter := adapter.NewGoogleAdapter()

		origin := "Rua Vander Marcelo Freitas Juvenesso, 892374983, 472389472"
		destination := "DAHSUIDAJHIe Xuyrhauds, 783425893, 842763842"

		useCase := NewMapsUseCase(googleAdapter)

		_, err := useCase.CalculatePrice(context.Background(), origin, destination, 185.0)
		if err == nil {
			t.Errorf("Error: %s", err)
		}
	})
}
