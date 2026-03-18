package usecase

import (
	"go-app/internal/domain"
	"testing"
)

func TestGreet(t *testing.T) {
	uc := NewGreeterUseCase()

	tests := []struct {
		name     string
		input    domain.GreetRequest
		expected string
	}{
		{
			name:     "saluda con nombre simple",
			input:    domain.GreetRequest{Name: "Juan"},
			expected: "Hola, Juan",
		},
		{
			name:     "saluda con nombre compuesto",
			input:    domain.GreetRequest{Name: "Maria Jose"},
			expected: "Hola, Maria Jose",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := uc.Greet(tt.input)
			if result.Message != tt.expected {
				t.Errorf("esperado %q, obtenido %q", tt.expected, result.Message)
			}
		})
	}
}
