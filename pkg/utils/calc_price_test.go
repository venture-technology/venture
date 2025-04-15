package utils

import "testing"

func TestCalculateContract(t *testing.T) {
	tests := []struct {
		name     string
		distance float64
		amount   float64
		expected float64
	}{
		{
			name:     "distância menor que 2",
			distance: 1.5,
			amount:   50,
			expected: 200,
		},
		{
			name:     "distância igual a 2",
			distance: 2.0,
			amount:   75,
			expected: 200,
		},
		{
			name:     "distância maior que 2",
			distance: 3.0,
			amount:   100,
			expected: 200 + 100*(3.0-2.0), // 300
		},
		{
			name:     "distância com fração após 2",
			distance: 2.5,
			amount:   80,
			expected: 200 + 80*0.5, // 240
		},
		{
			name:     "distância muito maior que 2",
			distance: 10.0,
			amount:   20,
			expected: 200 + 20*8.0, // 360
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
