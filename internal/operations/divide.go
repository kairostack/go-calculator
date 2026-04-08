package operations

import (
	"errors"
	"math"

	calcErrors "github.com/kairostack/go-calculator/internal/errors"
)

// DivideOperation implements the Operation interface for division
type DivideOperation struct{}

// init registers the DivideOperation with the DefaultRegistry
func init() {
	RegisterDefault(&DivideOperation{})
}

// Execute divides the first number by the second
// Returns ErrDivisionByZero if the divisor is zero, NaN, or infinite
// Also validates that inputs are finite numbers and checks for underflow/overflow
func (d *DivideOperation) Execute(x, y float64) (float64, error) {
	// Validate inputs first
	if err := validateInputs(x, y); err != nil {
		return 0, err
	}

	// Check for division by zero or invalid divisor values
	if y == 0 || math.IsNaN(y) || math.IsInf(y, 0) {
		return 0, calcErrors.ErrDivisionByZero
	}

	result := x / y
	// Check for underflow: when a non-zero number divided by a large number produces zero
	if x != 0 && result == 0 && math.Abs(y) > 1 {
		return 0, errors.New("division underflow - result too small")
	}
	if math.IsInf(result, 0) {
		return 0, errors.New("division overflow - result too large")
	}

	return result, nil
}

// Name returns the operation identifier
func (d *DivideOperation) Name() string {
	return "divide"
}

// Description returns a human-readable description
func (d *DivideOperation) Description() string {
	return "Divides the first number by the second"
}

// Symbol returns the mathematical symbol
func (d *DivideOperation) Symbol() string {
	return "/"
}
