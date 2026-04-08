package operations

import (
	"math"
	"testing"
)

func TestEdgeCases_MaxFloat64_Multiplication(t *testing.T) {
	mul := &MultiplyOperation{}

	// Test overflow scenarios
	_, err := mul.Execute(math.MaxFloat64, 2)
	if err == nil {
		t.Error("expected overflow error for MaxFloat64 * 2")
	}

	_, err = mul.Execute(math.MaxFloat64, math.MaxFloat64)
	if err == nil {
		t.Error("expected overflow error for MaxFloat64 * MaxFloat64")
	}

	// Test large but valid multiplication
	result, err := mul.Execute(math.MaxFloat64/2, 1.5)
	if err != nil {
		t.Errorf("unexpected error for valid large multiplication: %v", err)
	}
	if math.IsInf(result, 0) {
		t.Error("result should not be infinity for valid large multiplication")
	}
}

func TestEdgeCases_MinFloat64_Division(t *testing.T) {
	div := &DivideOperation{}

	// Test underflow scenario (smallest number divided by a large number should underflow)
	_, err := div.Execute(math.SmallestNonzeroFloat64, math.MaxFloat64)
	if err == nil {
		t.Error("expected underflow error for smallest / max")
	}

	// Test valid small division (dividing by small number produces larger result)
	result, err := div.Execute(math.SmallestNonzeroFloat64, 0.5)
	if err != nil {
		t.Errorf("unexpected error for valid small division: %v", err)
	}
	if result == 0 {
		t.Error("result should not be zero for valid small division")
	}
}

func TestEdgeCases_AdditionOverflow(t *testing.T) {
	add := &AddOperation{}

	// Test overflow in addition
	_, err := add.Execute(math.MaxFloat64, math.MaxFloat64)
	if err == nil {
		t.Error("expected overflow error for MaxFloat64 + MaxFloat64")
	}

	// Test near overflow
	result, err := add.Execute(math.MaxFloat64/2, math.MaxFloat64/2)
	if err != nil {
		t.Errorf("unexpected error for valid near-overflow addition: %v", err)
	}
	if math.IsInf(result, 0) {
		t.Error("result should not be infinity for valid addition")
	}
}

func TestEdgeCases_DivisionOverflow(t *testing.T) {
	div := &DivideOperation{}

	// Test overflow in division
	_, err := div.Execute(math.MaxFloat64, math.SmallestNonzeroFloat64)
	if err == nil {
		t.Error("expected overflow error for MaxFloat64 / smallest")
	}
}

func TestEdgeCases_InfinityResults(t *testing.T) {
	mul := &MultiplyOperation{}
	add := &AddOperation{}
	div := &DivideOperation{}

	// Multiplication overflow - MaxFloat64/2 * 3 should overflow
	_, err := mul.Execute(math.MaxFloat64/2, 3)
	if err == nil {
		t.Error("expected overflow error for large multiplication")
	}

	// Addition overflow - MaxFloat64 + large number should overflow
	_, err = add.Execute(math.MaxFloat64, math.MaxFloat64)
	if err == nil {
		t.Error("expected overflow error for large addition")
	}

	// Division overflow - MaxFloat64 / 0.5 should overflow
	_, err = div.Execute(math.MaxFloat64, 0.5)
	if err == nil {
		t.Error("expected overflow error for large division result")
	}
}

func TestEdgeCases_ZeroOperations(t *testing.T) {
	mul := &MultiplyOperation{}
	div := &DivideOperation{}

	// Zero multiplication
	result, err := mul.Execute(0, math.MaxFloat64)
	if err != nil {
		t.Errorf("unexpected error for 0 * MaxFloat64: %v", err)
	}
	if result != 0 {
		t.Errorf("expected 0, got %g", result)
	}

	// Zero division (dividend is zero)
	result, err = div.Execute(0, 5)
	if err != nil {
		t.Errorf("unexpected error for 0 / 5: %v", err)
	}
	if result != 0 {
		t.Errorf("expected 0, got %g", result)
	}
}
