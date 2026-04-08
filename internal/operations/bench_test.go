package operations

import "testing"

func BenchmarkAdd(b *testing.B) {
	add := &AddOperation{}
	for i := 0; i < b.N; i++ {
		add.Execute(123.456, 789.012)
	}
}

func BenchmarkSubtract(b *testing.B) {
	sub := &SubtractOperation{}
	for i := 0; i < b.N; i++ {
		sub.Execute(1000.0, 500.0)
	}
}

func BenchmarkMultiply(b *testing.B) {
	mul := &MultiplyOperation{}
	for i := 0; i < b.N; i++ {
		mul.Execute(123.456, 789.012)
	}
}

func BenchmarkDivide(b *testing.B) {
	div := &DivideOperation{}
	for i := 0; i < b.N; i++ {
		div.Execute(1000.0, 4.0)
	}
}

// Benchmark with different input sizes
func BenchmarkMultiplyLargeNumbers(b *testing.B) {
	mul := &MultiplyOperation{}
	for i := 0; i < b.N; i++ {
		mul.Execute(1e150, 1e150)
	}
}

func BenchmarkDivideSmallNumbers(b *testing.B) {
	div := &DivideOperation{}
	for i := 0; i < b.N; i++ {
		div.Execute(1e-150, 1e150)
	}
}
