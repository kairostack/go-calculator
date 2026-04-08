// Package floatutil provides utilities for safe floating-point comparisons.
// Float comparisons must use tolerance because direct equality checks fail
// due to IEEE 754 representation limitations (e.g., 0.1 + 0.2 != 0.3).
package floatutil

import "math"

// Epsilon is the default tolerance for float comparisons.
// 1e-9 provides precision suitable for typical calculator operations
// while accounting for floating-point arithmetic errors.
const Epsilon = 1e-9

// Equals compares two float64 values for equality within a tolerance.
// It handles special cases including NaN, infinities, and very large numbers.
//
// The comparison uses both absolute and relative error to ensure
// accuracy across different magnitude ranges:
//   - For small numbers (< 1e9), absolute error is sufficient
//   - For large numbers, relative error prevents false negatives
//
// NaN values are never considered equal, even to other NaN values.
// Positive and negative infinities are equal only to themselves.
func Equals(a, b float64) bool {
	// NaN is never equal to anything, including itself
	if math.IsNaN(a) || math.IsNaN(b) {
		return false
	}

	// Handle infinities: they must be exactly equal
	// (both +Inf or both -Inf)
	if math.IsInf(a, 0) || math.IsInf(b, 0) {
		return a == b
	}

	// Calculate absolute difference
	diff := math.Abs(a - b)

	// Use absolute error check for small differences
	if diff <= Epsilon {
		return true
	}

	// For large numbers, use relative error to avoid precision issues
	// This ensures 1e20 and (1e20 + 1) compare correctly
	return diff <= Epsilon*math.Max(math.Abs(a), math.Abs(b))
}

// EqualsWithTolerance compares two float64 values with a custom epsilon.
// Use this when the default Epsilon is not appropriate for your use case.
func EqualsWithTolerance(a, b, epsilon float64) bool {
	if math.IsNaN(a) || math.IsNaN(b) {
		return false
	}

	diff := math.Abs(a - b)
	if diff <= epsilon {
		return true
	}

	return diff <= epsilon*math.Max(math.Abs(a), math.Abs(b))
}

// IsZero checks if a float64 value is effectively zero within epsilon.
func IsZero(a float64) bool {
	return math.Abs(a) <= Epsilon
}

// GreaterThan checks if a > b within epsilon tolerance.
// Returns false if values are effectively equal.
func GreaterThan(a, b float64) bool {
	return a > b && !Equals(a, b)
}

// LessThan checks if a < b within epsilon tolerance.
// Returns false if values are effectively equal.
func LessThan(a, b float64) bool {
	return a < b && !Equals(a, b)
}

// IsFinite returns true if the value is neither NaN nor +/- Inf.
func IsFinite(a float64) bool {
	return !math.IsNaN(a) && !math.IsInf(a, 0)
}
