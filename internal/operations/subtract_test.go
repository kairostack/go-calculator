package operations

import (
	"errors"
	"math"
	"testing"

	calcErrors "github.com/kairostack/go-calculator/internal/errors"
	"github.com/kairostack/go-calculator/pkg/floatutil"
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
			if !floatutil.Equals(result, tt.expected) {
				t.Errorf("Execute(%g, %g) = %g, want %g (diff: %g)",
					tt.a, tt.b, result, tt.expected, result-tt.expected)
			}
		})
	}
}

func TestSubtractOperation_Execute_InvalidInputs(t *testing.T) {
	sub := &SubtractOperation{}

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
			_, err := sub.Execute(tt.a, tt.b)
			if err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("expected error %v, got %v", tt.expectedErr, err)
			}
		})
	}
}

func TestSubtractOperation_Execute_Overflow(t *testing.T) {
	sub := &SubtractOperation{}

	tests := []struct {
		name    string
		a, b    float64
		wantErr bool
	}{
		{"MaxFloat64 - (-MaxFloat64)", math.MaxFloat64, -math.MaxFloat64, true},
		{"(-MaxFloat64) - MaxFloat64", -math.MaxFloat64, math.MaxFloat64, true},
		{"valid large", math.MaxFloat64 / 2, -math.MaxFloat64 / 4, false},
		{"valid subtraction", 1e200, -1e200, false}, // 2e200 is still finite
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := sub.Execute(tt.a, tt.b)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected overflow error, got result %g", result)
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
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
