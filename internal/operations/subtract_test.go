package operations

import (
	"testing"
)

func TestSubtractOperation_Execute(t *testing.T) {
	sub := &SubtractOperation{}

	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"positive numbers", 10, 3, 7},
		{"negative numbers", -5, -3, -2},
		{"mixed signs", 5, -3, 8},
		{"with zero", 5, 0, 5},
		{"decimals", 5.5, 3.2, 2.3},
		{"result zero", 5, 5, 0},
		{"subtract larger", 3, 5, -2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := sub.Execute(tt.a, tt.b)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result != tt.expected {
				t.Errorf("Execute(%g, %g) = %g, want %g", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestSubtractOperation_Name(t *testing.T) {
	sub := &SubtractOperation{}
	if got := sub.Name(); got != "subtract" {
		t.Errorf("Name() = %q, want %q", got, "subtract")
	}
}

func TestSubtractOperation_Description(t *testing.T) {
	sub := &SubtractOperation{}
	expected := "Subtracts the second number from the first"
	if got := sub.Description(); got != expected {
		t.Errorf("Description() = %q, want %q", got, expected)
	}
}

func TestSubtractOperation_Symbol(t *testing.T) {
	sub := &SubtractOperation{}
	if got := sub.Symbol(); got != "-" {
		t.Errorf("Symbol() = %q, want %q", got, "-")
	}
}
