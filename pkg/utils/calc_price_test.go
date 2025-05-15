package utils

import "testing"

func TestCalculateContract(t *testing.T) {
	tests := []struct {
		name     string
		distance float64
		amount   int64 // em centavos
		expected int64 // em centavos
	}{
		{
			name:     "distância menor que 2",
			distance: 1.5,
			amount:   5000, // R$50,00
			expected: 20000,
		},
		{
			name:     "distância igual a 2",
			distance: 2.0,
			amount:   7500, // R$75,00
			expected: 20000,
		},
		{
			name:     "distância maior que 2",
			distance: 3.0,
			amount:   10000,           // R$100,00
			expected: 20000 + 10000*1, // 30000
		},
		{
			name:     "distância com fração após 2",
			distance: 2.5,
			amount:   8000,                    // R$80,00
			expected: 20000 + int64(8000*0.5), // 24000
		},
		{
			name:     "distância muito maior que 2",
			distance: 10.0,
			amount:   2000,           // R$20,00
			expected: 20000 + 2000*8, // 36000
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CalculateContract(tt.distance, tt.amount)
			if result != tt.expected {
				t.Errorf("CalculateContract() = %v, want %v", result, tt.expected)
			}
		})
	}
}
