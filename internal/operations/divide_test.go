package operations

import (
	"errors"
	"math"
	"testing"

	calcErrors "github.com/kairostack/go-calculator/internal/errors"
	"github.com/kairostack/go-calculator/pkg/floatutil"
)

func TestDivideOperation_Execute(t *testing.T) {
	div := &DivideOperation{}

	tests := []struct {
		name        string
		a, b        float64
		expected    float64
		expectError bool
	}{
		{"positive numbers", 15, 3, 5, false},
		{"negative numbers", -15, -3, 5, false},
		{"mixed signs", 15, -3, -5, false},
		{"decimals", 7.5, 2.5, 3, false},
		{"result fraction", 10, 4, 2.5, false},
		{"divide by zero", 5, 0, 0, true},
		{"zero dividend", 0, 5, 0, false},
		{"one as divisor", 5, 1, 5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := div.Execute(tt.a, tt.b)

			if tt.expectError {
				if err == nil {
					t.Errorf("Execute(%g, %g) expected error, got nil", tt.a, tt.b)
					return
				}
				if !errors.Is(err, calcErrors.ErrDivisionByZero) {
					t.Errorf("Execute(%g, %g) error = %v, want ErrDivisionByZero", tt.a, tt.b, err)
				}
				return
			}

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

func TestDivideOperation_Execute_EdgeCases(t *testing.T) {
	div := &DivideOperation{}

	tests := []struct {
		name        string
		a, b        float64
		expectedErr error
	}{
		{"NaN first input", math.NaN(), 5, calcErrors.ErrInputNaN},
		{"NaN second input", 5, math.NaN(), calcErrors.ErrInputNaN},
		{"NaN divisor (zero check)", 5, math.NaN(), calcErrors.ErrInputNaN},
		{"positive Inf first", math.Inf(1), 5, calcErrors.ErrInputInf},
		{"positive Inf second", 5, math.Inf(1), calcErrors.ErrInputInf},
		{"negative Inf first", math.Inf(-1), 5, calcErrors.ErrInputInf},
		{"negative Inf second", 5, math.Inf(-1), calcErrors.ErrInputInf},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := div.Execute(tt.a, tt.b)
			if err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("expected error %v, got %v", tt.expectedErr, err)
			}
		})
	}
}

func TestDivideOperation_Execute_OverflowUnderflow(t *testing.T) {
	div := &DivideOperation{}

	tests := []struct {
		name    string
		x, y    float64
		wantErr bool
	}{
		{"overflow: MaxFloat64 / small", math.MaxFloat64, math.SmallestNonzeroFloat64, true},
		{"underflow: small / MaxFloat64", math.SmallestNonzeroFloat64, math.MaxFloat64, true},
		{"valid division", 10, 2, false},
		{"zero dividend", 0, 5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := div.Execute(tt.x, tt.y)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error, got result %g", result)
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestDivideOperation_Name(t *testing.T) {
	div := &DivideOperation{}
	if got := div.Name(); got != "divide" {
		t.Errorf("Name() = %q, want %q", got, "divide")
	}
}

func TestDivideOperation_Description(t *testing.T) {
	div := &DivideOperation{}
	expected := "Divides the first number by the second"
	if got := div.Description(); got != expected {
		t.Errorf("Description() = %q, want %q", got, expected)
	}
}

func TestDivideOperation_Symbol(t *testing.T) {
	div := &DivideOperation{}
	if got := div.Symbol(); got != "/" {
		t.Errorf("Symbol() = %q, want %q", got, "/")
	}
}
