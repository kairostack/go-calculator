package operations

import (
	"math"

	"github.com/kairostack/go-calculator/internal/errors"
)

// SubtractOperation implements the Operation interface for subtraction
type SubtractOperation struct{}

// init registers the SubtractOperation with the DefaultRegistry
func init() {
	RegisterDefault(&SubtractOperation{})
}

// Execute subtracts the second number from the first
// Returns an error if either input is NaN or infinite, or if overflow occurs
func (s *SubtractOperation) Execute(x, y float64) (float64, error) {
	if err := validateInputs(x, y); err != nil {
		return 0, err
	}

	result := x - y
	if math.IsInf(result, 0) {
		return 0, &errors.CalculatorError{
			Op:      "subtract",
			Err:     "subtraction resulted in infinity",
			Details: "overflow detected",
		}
	}

	return result, nil
}

// Name returns the operation identifier
func (s *SubtractOperation) Name() string {
	return "subtract"
}

// Description returns a human-readable description
func (s *SubtractOperation) Description() string {
	return "Subtracts the second number from the first"
}

// Symbol returns the mathematical symbol
func (s *SubtractOperation) Symbol() string {
	return "-"
}
