package utils

import "testing"

func TestIsCPF(t *testing.T) {
	tests := []struct {
		name     string
		cpf      string
		expected bool
	}{
		{"CPF válido com pontuação", "123.456.789-09", true},
		{"CPF válido sem pontuação", "12345678909", true},
		{"CPF inválido", "123.456.789-00", false},
		{"CPF com letras", "123.456.abc-09", false},
		{"CPF com todos dígitos iguais", "111.111.111-11", false},
		{"CPF com menos dígitos", "123.456.789", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsCPF(tt.cpf)
			if result != tt.expected {
				t.Errorf("IsCPF(%q) = %v, want %v", tt.cpf, result, tt.expected)
			}
		})
	}
}

func TestIsCNPJ(t *testing.T) {
	tests := []struct {
		name     string
		cnpj     string
		expected bool
	}{
		{"CNPJ válido com pontuação", "45.723.174/0001-10", true},
		{"CNPJ válido sem pontuação", "45723174000110", true},
		{"CNPJ inválido", "11.111.111/1111-11", false},
		{"CNPJ com letras", "45.723.abc/0001-10", false},
		{"CNPJ com menos dígitos", "45.723.174/0001", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsCNPJ(tt.cnpj)
			if result != tt.expected {
				t.Errorf("IsCNPJ(%q) = %v, want %v", tt.cnpj, result, tt.expected)
			}
		})
	}
}

func TestIsCNH(t *testing.T) {
	tests := []struct {
		name     string
		cnh      string
		expected bool
	}{
		{"CNH válida", "04302485900", false},
		{"CNH com letras", "04302a85900", false},
		{"CNH inválida", "12345678900", true},
		{"CNH com todos dígitos iguais", "11111111111", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsCNH(tt.cnh)
			if result != tt.expected {
				t.Errorf("IsCNH(%q) = %v, want %v", tt.cnh, result, tt.expected)
			}
		})
	}
}
