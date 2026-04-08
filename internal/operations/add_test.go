package operations

import (
	"errors"
	"math"
	"testing"

	calcErrors "github.com/kairostack/go-calculator/internal/errors"
	"github.com/kairostack/go-calculator/pkg/floatutil"
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
			if !floatutil.Equals(result, tt.expected) {
				t.Errorf("Execute(%g, %g) = %g, want %g (diff: %g)",
					tt.a, tt.b, result, tt.expected, result-tt.expected)
			}
		})
	}
}

func TestAddOperation_Execute_InvalidInputs(t *testing.T) {
	add := &AddOperation{}

	tests := []struct {
		name        string
		a, b        float64
		expectedErr error
	}{
		{"NaN first input", math.NaN(), 5, calcErrors.ErrInputNaN},
		{"NaN second input", 5, math.NaN(), calcErrors.ErrInputNaN},
		{"NaN both inputs", math.NaN(), math.NaN(), calcErrors.ErrInputNaN},
		{"positive Inf first", math.Inf(1), 5, calcErrors.ErrInputInf},
		{"positive Inf second", 5, math.Inf(1), calcErrors.ErrInputInf},
		{"negative Inf first", math.Inf(-1), 5, calcErrors.ErrInputInf},
		{"negative Inf second", 5, math.Inf(-1), calcErrors.ErrInputInf},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := add.Execute(tt.a, tt.b)
			if err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("expected error %v, got %v", tt.expectedErr, err)
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
