package operations

import (
	"math"
	"testing"
)

func FuzzAddOperation(f *testing.F) {
	f.Add(123.456, 789.012)
	f.Add(-100.0, 50.0)
	f.Add(0.0, 0.0)
	f.Add(math.MaxFloat64, 1.0)

	f.Fuzz(func(t *testing.T, a, b float64) {
		add := &AddOperation{}
		result, err := add.Execute(a, b)

		// Check for expected errors
		if math.IsNaN(a) || math.IsNaN(b) || math.IsInf(a, 0) || math.IsInf(b, 0) {
			if err == nil {
				t.Error("expected error for NaN/Inf inputs")
			}
			return
		}

		// If no error, verify result is not NaN/Inf unless expected
		if err == nil {
			if math.IsNaN(result) && !math.IsNaN(a+b) {
				t.Error("unexpected NaN result")
			}
		}
	})
}

func FuzzMultiplyOperation(f *testing.F) {
	f.Add(123.456, 789.012)
	f.Add(-100.0, 50.0)
	f.Add(1e100, 1e100) // Potential overflow

	f.Fuzz(func(t *testing.T, a, b float64) {
		mul := &MultiplyOperation{}
		result, err := mul.Execute(a, b)

		// Check for overflow
		if math.IsInf(result, 0) && err == nil {
			t.Error("expected overflow error")
		}
	})
}

func FuzzDivideOperation(f *testing.F) {
	f.Add(100.0, 4.0)
	f.Add(1.0, 3.0)
	f.Add(0.0, 1.0)
	f.Add(1.0, 0.0) // Division by zero

	f.Fuzz(func(t *testing.T, a, b float64) {
		div := &DivideOperation{}
		result, err := div.Execute(a, b)

		// Check for division by zero
		if b == 0 || math.IsNaN(b) || math.IsInf(b, 0) {
			if err == nil {
				t.Error("expected division by zero error")
			}
			return
		}

		// Verify result * b ≈ a (for non-zero b)
		if err == nil && b != 0 {
			reconstructed := result * b
			if math.Abs(reconstructed-a) > 1e-9 {
				t.Errorf("result verification failed: %f * %f = %f, want ≈ %f",
					result, b, reconstructed, a)
			}
		}
	})
}
