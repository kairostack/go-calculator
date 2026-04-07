package cli

import (
	"fmt"
	"strings"
)

// Formatter handles output formatting for calculator results
type Formatter struct {
	precision int
}

// NewFormatter creates a new result formatter
// Default precision is 10 decimal places
func NewFormatter() *Formatter {
	return &Formatter{precision: 10}
}

// NewFormatterWithPrecision creates a formatter with custom decimal precision
func NewFormatterWithPrecision(precision int) *Formatter {
	if precision < 0 {
		precision = 10
	}
	return &Formatter{precision: precision}
}

// FormatResult formats a calculation result for display
func (f *Formatter) FormatResult(operation, symbol string, a, b, result float64) string {
	return fmt.Sprintf("%g %s %g = %s", a, symbol, b, f.formatNumber(result))
}

// FormatError formats an error message for display
func (f *Formatter) FormatError(err error) string {
	return fmt.Sprintf("Error: %s", err.Error())
}

// FormatHelp generates a help message with available operations
func (f *Formatter) FormatHelp(operations []string) string {
	var sb strings.Builder
	sb.WriteString("Usage: calculator <operation> <number1> <number2>\n\n")
	sb.WriteString("Available operations:\n")

	for _, op := range operations {
		sb.WriteString(fmt.Sprintf("  - %s\n", op))
	}

	sb.WriteString("\nExamples:\n")
	sb.WriteString("  calculator add 5 3       # Result: 8\n")
	sb.WriteString("  calculator subtract 10 4 # Result: 6\n")
	sb.WriteString("  calculator multiply 6 7  # Result: 42\n")
	sb.WriteString("  calculator divide 15 3   # Result: 5\n")

	return sb.String()
}

// formatNumber formats a float64 for display, removing trailing zeros
func (f *Formatter) formatNumber(n float64) string {
	format := fmt.Sprintf("%%.%df", f.precision)
	s := fmt.Sprintf(format, n)

	// Remove trailing zeros
	s = strings.TrimRight(s, "0")
	s = strings.TrimRight(s, ".")

	return s
}
