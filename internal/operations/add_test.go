package operations

import (
	"testing"
)

func TestAddOperation_Execute(t *testing.T) {
	add := &AddOperation{}

	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 5, 3, 8},
		{"negative numbers", -5, -3, -8},
		{"mixed signs", 5, -3, 2},
		{"with zero", 5, 0, 5},
		{"decimals", 5.5, 3.2, 8.7},
		{"large numbers", 1000000, 2000000, 3000000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := add.Execute(tt.a, tt.b)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("Execute(%g, %g) = %g, want %g", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestAddOperation_Name(t *testing.T) {
	add := &AddOperation{}
	if got := add.Name(); got != "add" {
		t.Errorf("Name() = %q, want %q", got, "add")
	}
}

func TestAddOperation_Description(t *testing.T) {
	add := &AddOperation{}
	expected := "Adds two numbers together"
	if got := add.Description(); got != expected {
		t.Errorf("Description() = %q, want %q", got, expected)
	}
}

func TestAddOperation_Symbol(t *testing.T) {
	add := &AddOperation{}
	if got := add.Symbol(); got != "+" {
		t.Errorf("Symbol() = %q, want %q", got, "+")
	}
}
