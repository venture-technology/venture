package usecase

import (
	"testing"
)

func TestCalculateContract(t *testing.T) {
	t.Run("testing value increase 200", func(t *testing.T) {
		distance := 4.0
		amount := 200.0
		result := CalculateContract(distance, amount)
		if result != 600 {
			t.Errorf("Expected 600, got %f", result)
		}
	})

	t.Run("testing value below 200", func(t *testing.T) {
		distance := 1.5
		amount := 100.0
		result := CalculateContract(distance, amount)
		if result != 200 {
			t.Errorf("Expected 200, got %f", result)
		}
	})
}
