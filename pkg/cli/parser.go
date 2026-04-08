package cli

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/kairostack/go-calculator/internal/errors"
)

// Version is the current version of the calculator CLI
const Version = "1.0.0"

// ParseResult holds the parsed command-line arguments
type ParseResult struct {
	Operation string
	OperandA  float64
	OperandB  float64
}

// Parser handles command-line argument parsing
type Parser struct{}

// NewParser creates a new argument parser
func NewParser() *Parser {
	return &Parser{}
}

// Parse converts command-line arguments into a structured result
// Expected format: [operation] [number1] [number2]
// Example: add 5 3
func (p *Parser) Parse(args []string) (*ParseResult, error) {
	if len(args) < 3 {
		return nil, &errors.CalculatorError{
			Op:      "parse",
			Err:     "insufficient arguments",
			Details: fmt.Sprintf("expected: <operation> <number1> <number2>, got %d arguments", len(args)),
		}
	}

	operation := strings.ToLower(strings.TrimSpace(args[0]))
	if operation == "" {
		return nil, &errors.CalculatorError{
			Op:  "parse",
			Err: "operation cannot be empty",
		}
	}

	a, err := strconv.ParseFloat(strings.TrimSpace(args[1]), 64)
	if err != nil {
		return nil, &errors.CalculatorError{
			Op:      "parse",
			Err:     "invalid first operand",
			Details: fmt.Sprintf("'%s' is not a valid number", args[1]),
		}
	}

	if !isValidNumber(a) {
		return nil, &errors.CalculatorError{
			Op:      "parse",
			Err:     "invalid first operand",
			Details: fmt.Sprintf("'%s' is not a valid finite number", args[1]),
		}
	}

	b, err := strconv.ParseFloat(strings.TrimSpace(args[2]), 64)
	if err != nil {
		return nil, &errors.CalculatorError{
			Op:      "parse",
			Err:     "invalid second operand",
			Details: fmt.Sprintf("'%s' is not a valid number", args[2]),
		}
	}

	if !isValidNumber(b) {
		return nil, &errors.CalculatorError{
			Op:      "parse",
			Err:     "invalid second operand",
			Details: fmt.Sprintf("'%s' is not a valid finite number", args[2]),
		}
	}

	return &ParseResult{
		Operation: operation,
		OperandA:  a,
		OperandB:  b,
	}, nil
}

// isValidNumber checks if the parsed float is a finite number (not NaN or +/- Inf).
// This prevents silent failures from special floating-point values.
func isValidNumber(n float64) bool {
	return !math.IsNaN(n) && !math.IsInf(n, 0)
}

// ParseResultString returns a formatted string representation of the parse result
func (pr *ParseResult) String() string {
	return fmt.Sprintf("Operation: %s, A: %g, B: %g", pr.Operation, pr.OperandA, pr.OperandB)
}
