package operations

import (
	"testing"
)

func TestMultiplyOperation_Execute(t *testing.T) {
	mul := &MultiplyOperation{}

	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 6, 7, 42},
		{"negative numbers", -5, -3, 15},
		{"mixed signs", 5, -3, -15},
		{"with zero", 5, 0, 0},
		{"with one", 5, 1, 5},
		{"decimals", 2.5, 4, 10},
		{"large numbers", 1000, 1000, 1000000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := mul.Execute(tt.a, tt.b)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("Execute(%g, %g) = %g, want %g", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestMultiplyOperation_Name(t *testing.T) {
	mul := &MultiplyOperation{}
	if got := mul.Name(); got != "multiply" {
		t.Errorf("Name() = %q, want %q", got, "multiply")
	}
}

func TestMultiplyOperation_Description(t *testing.T) {
	mul := &MultiplyOperation{}
	expected := "Multiplies two numbers together"
	if got := mul.Description(); got != expected {
		t.Errorf("Description() = %q, want %q", got, expected)
	}
}

func TestMultiplyOperation_Symbol(t *testing.T) {
	mul := &MultiplyOperation{}
	if got := mul.Symbol(); got != "*" {
		t.Errorf("Symbol() = %q, want %q", got, "*")
	}
}
