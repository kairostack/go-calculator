package cli

import (
	"strconv"
	"testing"
	"unicode"
)

func FuzzParseArgs(f *testing.F) {
	// Seed corpus with valid inputs
	f.Add("add", "123.456", "789.012")
	f.Add("multiply", "-100", "50.5")
	f.Add("divide", "0", "1")

	f.Fuzz(func(t *testing.T, op, aStr, bStr string) {
		// Skip invalid operation names (should contain only letters)
		for _, r := range op {
			if !unicode.IsLetter(r) {
				t.Skip()
			}
		}

		args := []string{"calc", op, aStr, bStr}
		parser := NewParser()

		result, err := parser.Parse(args)

		// If parsing succeeds, verify the values
		if err == nil {
			// Verify operation name matches
			if result.Operation != op {
				t.Errorf("operation name mismatch: got %s, want %s", result.Operation, op)
			}

			// Verify numbers can be parsed
			expectedA, _ := strconv.ParseFloat(aStr, 64)
			expectedB, _ := strconv.ParseFloat(bStr, 64)

			if result.OperandA != expectedA {
				t.Errorf("first number mismatch: got %f, want %f", result.OperandA, expectedA)
			}
			if result.OperandB != expectedB {
				t.Errorf("second number mismatch: got %f, want %f", result.OperandB, expectedB)
			}
		}
	})
}
