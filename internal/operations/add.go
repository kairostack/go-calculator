package operations

import (
	"errors"
	"math"
)

// AddOperation implements the Operation interface for addition
type AddOperation struct{}

// init registers the AddOperation with the DefaultRegistry
func init() {
	RegisterDefault(&AddOperation{})
}

// Execute adds two numbers together
// Returns an error if either input is NaN or infinite, or if overflow occurs
func (a *AddOperation) Execute(x, y float64) (float64, error) {
	if err := validateInputs(x, y); err != nil {
		return 0, err
	}

	result := x + y
	if math.IsInf(result, 0) {
		return 0, errors.New("addition resulted in infinity")
	}

	return result, nil
}

// Name returns the operation identifier
func (a *AddOperation) Name() string {
	return "add"
}

// Description returns a human-readable description
func (a *AddOperation) Description() string {
	return "Adds two numbers together"
}

// Symbol returns the mathematical symbol
func (a *AddOperation) Symbol() string {
	return "+"
}
