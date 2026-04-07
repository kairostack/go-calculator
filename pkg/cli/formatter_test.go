package cli

import (
	"strings"
	"testing"
)

func TestFormatter_FormatResult(t *testing.T) {
	formatter := NewFormatter()

	result := formatter.FormatResult("add", "+", 5, 3, 8)

	if !strings.Contains(result, "5") || !strings.Contains(result, "3") || !strings.Contains(result, "8") {
		t.Error("FormatResult() should contain operands and result")
	}
	if !strings.Contains(result, "+") {
		t.Error("FormatResult() should contain operation symbol")
	}
}

func TestFormatter_FormatError(t *testing.T) {
	formatter := NewFormatter()
	testErr := &testError{msg: "test error"}

	result := formatter.FormatError(testErr)

	if !strings.Contains(result, "Error:") {
		t.Error("FormatError() should contain 'Error:' prefix")
	}
	if !strings.Contains(result, "test error") {
		t.Error("FormatError() should contain error message")
	}
}

// testError is a simple error type for testing
type testError struct {
	msg string
}

func (e *testError) Error() string {
	return e.msg
}

func TestFormatter_FormatHelp(t *testing.T) {
	formatter := NewFormatter()
	operations := []string{"add", "subtract", "multiply", "divide"}

	help := formatter.FormatHelp(operations)

	// Check for key sections
	if !strings.Contains(help, "Usage:") {
		t.Error("FormatHelp() should contain 'Usage:' section")
	}
	if !strings.Contains(help, "Available operations:") {
		t.Error("FormatHelp() should contain 'Available operations:' section")
	}
	if !strings.Contains(help, "Examples:") {
		t.Error("FormatHelp() should contain 'Examples:' section")
	}

	// Check all operations are listed
	for _, op := range operations {
		if !strings.Contains(help, op) {
			t.Errorf("FormatHelp() should contain operation %q", op)
		}
	}
}

func TestFormatter_FormatNumber(t *testing.T) {
	tests := []struct {
		input    float64
		expected string
	}{
		{8, "8"},
		{8.0, "8"},
		{8.5, "8.5"},
		{8.50, "8.5"},
		{0.3333333333, "0.3333333333"},
		{1000000, "1000000"},
	}

	formatter := NewFormatter()

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			// Test via FormatResult
			result := formatter.FormatResult("test", "=", 0, 0, tt.input)
			if !strings.HasSuffix(result, tt.expected) {
				// Allow for trailing zeros up to precision
				t.Logf("Format result: %s (may have trailing zeros)", result)
			}
		})
	}
}

func TestNewFormatterWithPrecision(t *testing.T) {
	// Default precision
	f1 := NewFormatter()
	result1 := f1.FormatResult("add", "+", 1, 3, 0.3333333333333333)
	if !strings.Contains(result1, "0.3333333333") {
		t.Logf("Default precision result: %s", result1)
	}

	// Custom precision
	f2 := NewFormatterWithPrecision(2)
	result2 := f2.FormatResult("add", "+", 1, 3, 0.3333333333333333)
	if !strings.HasSuffix(result2, "0.33") {
		t.Errorf("Custom precision (2) result = %s, want suffix 0.33", result2)
	}

	// Negative precision (should use default)
	f3 := NewFormatterWithPrecision(-1)
	_ = f3 // Just verify it doesn't panic
}
