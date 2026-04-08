package floatutil

import "testing"

func BenchmarkEqualsDifferentValues(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Equals(123.456789, 123.456788)
	}
}

func BenchmarkEqualsSameValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Equals(123.456, 123.456)
	}
}

func BenchmarkEqualsLargeNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Equals(1e150, 1e150+1e135)
	}
}

func BenchmarkEqualsWithTolerance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EqualsWithTolerance(123.456789, 123.456788, 1e-6)
	}
}
