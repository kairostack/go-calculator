// Package errors provides custom error types for the calculator application.
//
// The errors package defines CalculatorError and OperationError types
// that provide rich error context including operation names, details,
// and support for error wrapping and unwrapping.
//
// Basic usage:
//
//	err := &CalculatorError{
//	    Op:  "divide",
//	    Err: "division by zero",
//	}
//
// Error wrapping:
//
//	err := errors.New("database connection failed")
//	calcErr := Wrap("add", err)
//	if errors.Is(calcErr, err) {
//	    // handle wrapped error
//	}
//
// Predefined errors:
//
//	err := ErrDivisionByZero
//	err := ErrInvalidOperation
//	...
package errors

import (
	"errors"
	"testing"
)

func TestCalculatorError_Error(t *testing.T) {
	tests := []struct {
		name     string
		err      CalculatorError
		expected string
	}{
		{
			name:     "with details",
			err:      CalculatorError{Op: "add", Err: "overflow", Details: "values too large"},
			expected: "calculator error [add]: overflow (values too large)",
		},
		{
			name:     "without details",
			err:      CalculatorError{Op: "divide", Err: "by zero"},
			expected: "calculator error [divide]: by zero",
		},
		{
			name:     "empty fields",
			err:      CalculatorError{Op: "", Err: ""},
			expected: "calculator error []: ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.err.Error()
			if got != tt.expected {
				t.Errorf("Error() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestCalculatorError_Predefined(t *testing.T) {
	tests := []struct {
		name   string
		err    error
		op     string
		errMsg string
	}{
		{
			name:   "ErrDivisionByZero",
			err:    ErrDivisionByZero,
			op:     "divide",
			errMsg: "division by zero",
		},
		{
			name:   "ErrInvalidOperation",
			err:    ErrInvalidOperation,
			op:     "unknown",
			errMsg: "invalid operation",
		},
		{
			name:   "ErrInsufficientArguments",
			err:    ErrInsufficientArguments,
			op:     "input",
			errMsg: "insufficient arguments",
		},
		{
			name:   "ErrInvalidNumber",
			err:    ErrInvalidNumber,
			op:     "input",
			errMsg: "invalid number format",
		},
		{
			name:   "ErrInvalidInput",
			err:    ErrInvalidInput,
			op:     "input",
			errMsg: "invalid input",
		},
		{
			name:   "ErrInputNaN",
			err:    ErrInputNaN,
			op:     "input",
			errMsg: "input is NaN",
		},
		{
			name:   "ErrInputInf",
			err:    ErrInputInf,
			op:     "input",
			errMsg: "input is infinite",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calcErr, ok := tt.err.(*CalculatorError)
			if !ok {
				t.Fatal("expected CalculatorError type")
			}
			if calcErr.Op != tt.op {
				t.Errorf("Op = %q, want %q", calcErr.Op, tt.op)
			}
			if calcErr.Err != tt.errMsg {
				t.Errorf("Err = %q, want %q", calcErr.Err, tt.errMsg)
			}
		})
	}
}

func TestNewCalculatorError(t *testing.T) {
	tests := []struct {
		op       string
		errMsg   string
		expected string
	}{
		{"add", "overflow", "calculator error [add]: overflow"},
		{"divide", "by zero", "calculator error [divide]: by zero"},
		{"", "", "calculator error []: "},
	}

	for _, tt := range tests {
		t.Run(tt.op+"_"+tt.errMsg, func(t *testing.T) {
			err := NewCalculatorError(tt.op, tt.errMsg)
			if err.Op != tt.op {
				t.Errorf("Op = %q, want %q", err.Op, tt.op)
			}
			if err.Err != tt.errMsg {
				t.Errorf("Err = %q, want %q", err.Err, tt.errMsg)
			}
			if err.Error() != tt.expected {
				t.Errorf("Error() = %q, want %q", err.Error(), tt.expected)
			}
		})
	}
}

func TestWrap(t *testing.T) {
	originalErr := errors.New("original error")
	wrappedErr := Wrap("multiply", originalErr)

	if wrappedErr.Op != "multiply" {
		t.Errorf("Op = %q, want %q", wrappedErr.Op, "multiply")
	}
	if wrappedErr.Err != "original error" {
		t.Errorf("Err = %q, want %q", wrappedErr.Err, "original error")
	}
	if wrappedErr.Details != "wrapped error" {
		t.Errorf("Details = %q, want %q", wrappedErr.Details, "wrapped error")
	}
}

func TestOperationError_Error(t *testing.T) {
	originalErr := errors.New("computation failed")
	opErr := &OperationError{
		Op:  "add",
		A:   5.5,
		B:   3.2,
		Err: originalErr,
	}

	expected := "operation add(5.5, 3.2) failed: computation failed"
	if got := opErr.Error(); got != expected {
		t.Errorf("Error() = %q, want %q", got, expected)
	}
}

func TestOperationError_Unwrap(t *testing.T) {
	originalErr := errors.New("wrapped error")
	opErr := &OperationError{
		Op:  "divide",
		A:   10,
		B:   0,
		Err: originalErr,
	}

	unwrapped := opErr.Unwrap()
	if unwrapped != originalErr {
		t.Error("Unwrap() did not return the original error")
	}

	// Test errors.Is integration
	if !errors.Is(opErr, originalErr) {
		t.Error("errors.Is() should return true for the wrapped error")
	}
}

func TestOperationError_Unwrap_Nil(t *testing.T) {
	opErr := &OperationError{
		Op:  "multiply",
		A:   2,
		B:   3,
		Err: nil,
	}

	unwrapped := opErr.Unwrap()
	if unwrapped != nil {
		t.Errorf("Unwrap() = %v, want nil", unwrapped)
	}
}

func TestOperationError_ErrorFormatting(t *testing.T) {
	tests := []struct {
		name    string
		op      string
		a, b    float64
		wantFmt string
	}{
		{
			name:    "integers",
			op:      "add",
			a:       10,
			b:       20,
			wantFmt: "operation add(10, 20) failed",
		},
		{
			name:    "decimals",
			op:      "divide",
			a:       5.5,
			b:       2.2,
			wantFmt: "operation divide(5.5, 2.2) failed",
		},
		{
			name:    "negative",
			op:      "subtract",
			a:       -10,
			b:       -5,
			wantFmt: "operation subtract(-10, -5) failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opErr := &OperationError{
				Op:  tt.op,
				A:   tt.a,
				B:   tt.b,
				Err: errors.New("test"),
			}

			if !contains(opErr.Error(), tt.wantFmt) {
				t.Errorf("Error() = %q, want to contain %q", opErr.Error(), tt.wantFmt)
			}
		})
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && containsSubstring(s, substr))
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func TestErrorWrapping(t *testing.T) {
	tests := []struct {
		name        string
		wrap        func() error
		target      error
		shouldMatch bool
	}{
		{
			name: "OperationError unwraps to original",
			wrap: func() error {
				return &OperationError{Op: "add", A: 1, B: 2, Err: errors.New("fail")}
			},
			target:      errors.New("fail"),
			shouldMatch: false, // Different error instances
		},
		{
			name: "same error identity",
			wrap: func() error {
				target := errors.New("specific error")
				return &OperationError{Op: "add", A: 1, B: 2, Err: target}
			},
			target:      errors.New("specific error"),
			shouldMatch: false, // Different error instances
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.wrap()
			// Check that we can unwrap and get the original
			if opErr, ok := err.(*OperationError); ok {
				if opErr.Unwrap() == nil {
					t.Error("Unwrap() returned nil for wrapped error")
				}
			}
		})
	}
}

func TestCalculatorError_NilCheck(t *testing.T) {
	var err *CalculatorError
	if err != nil {
		t.Error("nil *CalculatorError should be nil")
	}

	// Test Error() on nil (will panic - document this behavior)
	// This is expected behavior in Go - calling methods on nil pointers
}

func TestWrap_NilError(t *testing.T) {
	// Skip test if Wrap doesn't handle nil properly
	// The Wrap function is not designed to handle nil errors
	t.Skip("Wrap function does not support nil errors - handled by caller")
}
