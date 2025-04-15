package utils

import "testing"

func TestBuildAddress(t *testing.T) {
	tests := []struct {
		name       string
		street     string
		number     string
		complement string
		zip        string
		expected   string
	}{
		{
			name:       "complemento vazio",
			street:     "Rua das Flores",
			number:     "123",
			complement: "",
			zip:        "01234567",
			expected:   "Rua das Flores, 123, 01234567",
		},
		{
			name:       "complemento preenchido",
			street:     "Av. Brasil",
			number:     "456",
			complement: "Apto 101",
			zip:        "98765432",
			expected:   "Av. Brasil, 456, Apto 101, 98765432",
		},
		{
			name:       "todos os campos vazios exceto zip",
			street:     "",
			number:     "",
			complement: "",
			zip:        "00000000",
			expected:   ", , 00000000",
		},
		{
			name:       "complemento com espaço",
			street:     "Rua Teste",
			number:     "1",
			complement: " ",
			zip:        "11111111",
			expected:   "Rua Teste, 1,  , 11111111",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BuildAddress(tt.street, tt.number, tt.complement, tt.zip)
			if result != tt.expected {
				t.Errorf("BuildAddress() = %q, want %q", result, tt.expected)
			}
		})
	}
}
