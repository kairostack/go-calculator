package operations

import (
	"errors"
	"math"
	"testing"

	calcErrors "github.com/kairostack/go-calculator/internal/errors"
	"github.com/kairostack/go-calculator/pkg/floatutil"
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
			if !floatutil.Equals(result, tt.expected) {
				t.Errorf("Execute(%g, %g) = %g, want %g (diff: %g)",
					tt.a, tt.b, result, tt.expected, result-tt.expected)
			}
		})
	}
}

func TestMultiplyOperation_Execute_InvalidInputs(t *testing.T) {
	mul := &MultiplyOperation{}

	tests := []struct {
		name        string
		a, b        float64
		expectedErr error
	}{
		{"NaN first input", math.NaN(), 5, calcErrors.ErrInputNaN},
		{"NaN second input", 5, math.NaN(), calcErrors.ErrInputNaN},
		{"positive Inf first", math.Inf(1), 5, calcErrors.ErrInputInf},
		{"positive Inf second", 5, math.Inf(1), calcErrors.ErrInputInf},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := mul.Execute(tt.a, tt.b)
			if err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("expected error %v, got %v", tt.expectedErr, err)
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
