package calculator

import "testing"

func BenchmarkCalculateAdd(b *testing.B) {
	calc := New()
	for i := 0; i < b.N; i++ {
		calc.Calculate("add", 123.456, 789.012)
	}
}

func BenchmarkCalculateMultiply(b *testing.B) {
	calc := New()
	for i := 0; i < b.N; i++ {
		calc.Calculate("multiply", 123.456, 789.012)
	}
}

// Benchmark registry lookup overhead
func BenchmarkRegistryLookup(b *testing.B) {
	calc := New()
	for i := 0; i < b.N; i++ {
		calc.Calculate("add", 1.0, 1.0)
	}
}
