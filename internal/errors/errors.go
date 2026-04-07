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
