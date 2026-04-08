package floatutil

import (
	"math"
	"testing"
)

func TestEquals(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected bool
	}{
		// Basic equality cases
		{"identical positive", 5.0, 5.0, true},
		{"identical negative", -3.5, -3.5, true},
		{"identical zero", 0.0, 0.0, true},
		{"zero and negative zero", 0.0, -0.0, true},

		// Within epsilon tolerance
		{"within epsilon", 1.0, 1.0 + Epsilon/2, true},
		// Note: At exact epsilon boundary, the relative error check may fail
		// due to floating-point representation, so we test slightly within epsilon
		{"slightly within epsilon", 1.0, 1.0 + Epsilon*0.9, true},

		// Outside epsilon tolerance
		{"outside epsilon", 1.0, 1.0 + Epsilon*10, false},
		{"clearly different", 5.0, 3.0, false},

		// Floating-point precision issues (the classic problem)
		{"float precision issue", 0.1 + 0.2, 0.3, true},
		{"decimal addition", 0.3 + 0.6, 0.9, true},

		// Large numbers - test relative error handling
		{"large numbers equal", 1e15, 1e15, true},
		{"large numbers with small diff", 1e15, 1e15 + 1, true},

		// Small numbers
		{"very small numbers", 1e-12, 1e-12, true},
		{"small with epsilon diff", 1e-12, 1e-12 + Epsilon/10, true},

		// NaN handling
		{"NaN with number", math.NaN(), 5.0, false},
		{"number with NaN", 5.0, math.NaN(), false},
		{"NaN with NaN", math.NaN(), math.NaN(), false},

		// Infinity handling
		{"positive Inf equal", math.Inf(1), math.Inf(1), true},
		{"negative Inf equal", math.Inf(-1), math.Inf(-1), true},
		{"opposite Infinities", math.Inf(1), math.Inf(-1), false},
		{"Inf with number", math.Inf(1), 5.0, false},
		{"number with Inf", 5.0, math.Inf(1), false},

		// Mixed signs
		{"positive and negative", 5.0, -5.0, false},
		{"near zero opposite signs", Epsilon / 2, -Epsilon / 2, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Equals(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Equals(%g, %g) = %v, want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestEqualsWithTolerance(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		epsilon  float64
		expected bool
	}{
		{"custom tolerance success", 10.0, 10.5, 1.0, true},
		{"custom tolerance fail", 10.0, 100.0, 0.5, false},
		{"very strict tolerance", 1.0, 1.000001, 1e-10, false},
		{"very loose tolerance", 1.0, 100.0, 1000.0, true},
		{"NaN with custom tolerance", math.NaN(), 5.0, 1.0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EqualsWithTolerance(tt.a, tt.b, tt.epsilon)
			if result != tt.expected {
				t.Errorf("EqualsWithTolerance(%g, %g, %g) = %v, want %v",
					tt.a, tt.b, tt.epsilon, result, tt.expected)
			}
		})
	}
}

func TestIsZero(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		expected bool
	}{
		{"exactly zero", 0.0, true},
		{"negative zero", -0.0, true},
		{"within epsilon", Epsilon / 2, true},
		{"at epsilon boundary", Epsilon, true},
		{"outside epsilon", Epsilon * 10, false},
		{"positive number", 1.0, false},
		{"negative number", -1.0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsZero(tt.value)
			if result != tt.expected {
				t.Errorf("IsZero(%g) = %v, want %v", tt.value, result, tt.expected)
			}
		})
	}
}

func TestGreaterThan(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected bool
	}{
		{"clearly greater", 5.0, 3.0, true},
		{"clearly less", 3.0, 5.0, false},
		{"equal values", 5.0, 5.0, false},
		{"within epsilon", 5.0, 5.0 + Epsilon/2, false},
		{"just outside epsilon", 5.0, 5.0 + Epsilon*10, false},
		{"negative numbers", -3.0, -5.0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GreaterThan(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("GreaterThan(%g, %g) = %v, want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestLessThan(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected bool
	}{
		{"clearly less", 3.0, 5.0, true},
		{"clearly greater", 5.0, 3.0, false},
		{"equal values", 5.0, 5.0, false},
		{"within epsilon", 5.0 + Epsilon/2, 5.0, false},
		{"negative numbers", -5.0, -3.0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LessThan(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("LessThan(%g, %g) = %v, want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestIsFinite(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		expected bool
	}{
		{"zero", 0.0, true},
		{"positive number", 5.0, true},
		{"negative number", -5.0, true},
		{"large number", 1e308, true},
		{"small number", 1e-308, true},
		{"positive infinity", math.Inf(1), false},
		{"negative infinity", math.Inf(-1), false},
		{"NaN", math.NaN(), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsFinite(tt.value)
			if result != tt.expected {
				t.Errorf("IsFinite(%g) = %v, want %v", tt.value, result, tt.expected)
			}
		})
	}
}

// BenchmarkEquals benchmarks the Equals function
func BenchmarkEquals(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Equals(0.1+0.2, 0.3)
	}
}

// BenchmarkDirectComparison compares with direct == operator
func BenchmarkDirectComparison(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = 0.1+0.2 == 0.3
	}
}
