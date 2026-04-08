package operations

import (
	"math"

	"github.com/kairostack/go-calculator/internal/errors"
)

// validateInputs checks that both inputs are valid finite numbers.
// Returns specific error types for NaN and infinite values.
func validateInputs(a, b float64) error {
	if math.IsNaN(a) || math.IsNaN(b) {
		return errors.ErrInputNaN
	}
	if math.IsInf(a, 0) || math.IsInf(b, 0) {
		return errors.ErrInputInf
	}
	return nil
}
