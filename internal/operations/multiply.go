package operations

import (
	"math"

	"github.com/kairostack/go-calculator/internal/errors"
)

// MultiplyOperation implements the Operation interface for multiplication
type MultiplyOperation struct{}

// init registers the MultiplyOperation with the DefaultRegistry
func init() {
	RegisterDefault(&MultiplyOperation{})
}

// Execute multiplies two numbers together
// Returns an error if either input is NaN or infinite, or if overflow occurs
func (m *MultiplyOperation) Execute(x, y float64) (float64, error) {
	if err := validateInputs(x, y); err != nil {
		return 0, err
	}

	// Check for potential overflow
	if math.Abs(x) > 1 && math.Abs(y) > math.MaxFloat64/math.Abs(x) {
		return 0, &errors.CalculatorError{
			Op:      "multiply",
			Err:     "multiplication would overflow",
			Details: "operands too large",
		}
	}

	result := x * y
	if math.IsInf(result, 0) {
		return 0, &errors.CalculatorError{
			Op:      "multiply",
			Err:     "multiplication resulted in infinity",
			Details: "overflow detected",
		}
	}

	return result, nil
}

// Name returns the operation identifier
func (m *MultiplyOperation) Name() string {
	return "multiply"
}

// Description returns a human-readable description
func (m *MultiplyOperation) Description() string {
	return "Multiplies two numbers together"
}

// Symbol returns the mathematical symbol
func (m *MultiplyOperation) Symbol() string {
	return "*"
}
