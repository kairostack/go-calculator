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
	"fmt"
)

// CalculatorError represents a custom error for calculator operations
type CalculatorError struct {
	Op      string
	Err     string
	Details string
}

// Error implements the error interface
func (e *CalculatorError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("calculator error [%s]: %s (%s)", e.Op, e.Err, e.Details)
	}
	return fmt.Sprintf("calculator error [%s]: %s", e.Op, e.Err)
}

// Common calculator errors
var (
	// ErrDivisionByZero is returned when attempting to divide by zero
	ErrDivisionByZero = &CalculatorError{
		Op:  "divide",
		Err: "division by zero",
	}

	// ErrInvalidOperation is returned when an operation name is not recognized
	ErrInvalidOperation = &CalculatorError{
		Op:  "unknown",
		Err: "invalid operation",
	}

	// ErrInsufficientArguments is returned when not enough arguments are provided
	ErrInsufficientArguments = &CalculatorError{
		Op:  "input",
		Err: "insufficient arguments",
	}

	// ErrInvalidNumber is returned when a number cannot be parsed
	ErrInvalidNumber = &CalculatorError{
		Op:  "input",
		Err: "invalid number format",
	}

	// ErrInvalidInput is returned when input validation fails
	ErrInvalidInput = &CalculatorError{
		Op:  "input",
		Err: "invalid input",
	}

	// ErrInputNaN is returned when an input is NaN
	ErrInputNaN = &CalculatorError{
		Op:  "input",
		Err: "input is NaN",
	}

	// ErrInputInf is returned when an input is infinite
	ErrInputInf = &CalculatorError{
		Op:  "input",
		Err: "input is infinite",
	}
)

// NewCalculatorError creates a new CalculatorError with operation context
func NewCalculatorError(op, err string) *CalculatorError {
	return &CalculatorError{
		Op:  op,
		Err: err,
	}
}

// Wrap wraps an existing error with calculator context
func Wrap(op string, err error) *CalculatorError {
	return &CalculatorError{
		Op:      op,
		Err:     err.Error(),
		Details: "wrapped error",
	}
}

// OperationError provides detailed context for operation failures
type OperationError struct {
	Op   string
	A, B float64
	Err  error
}

// Error implements the error interface with context
func (e *OperationError) Error() string {
	return fmt.Sprintf("operation %s(%g, %g) failed: %v", e.Op, e.A, e.B, e.Err)
}

// Unwrap allows error inspection with errors.Is and errors.As
func (e *OperationError) Unwrap() error {
	return e.Err
}
