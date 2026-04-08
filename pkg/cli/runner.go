package cli

import (
	"fmt"
	"io"

	"github.com/kairostack/go-calculator/internal/calculator"
)

// Runner orchestrates CLI execution with I/O abstraction for testability
type Runner struct {
	calculator *calculator.Calculator
	formatter  *Formatter
	parser     *Parser
	stdin      io.Reader
	stdout     io.Writer
	stderr     io.Writer
}

// NewRunner creates a new CLI runner with the given dependencies
func NewRunner(calc *calculator.Calculator, stdin io.Reader, stdout, stderr io.Writer) *Runner {
	return &Runner{
		calculator: calc,
		formatter:  NewFormatter(),
		parser:     NewParser(),
		stdin:      stdin,
		stdout:     stdout,
		stderr:     stderr,
	}
}

// Run executes the calculator with the provided arguments
// Returns exit code: 0 for success, 1 for error
func (r *Runner) Run(args []string) int {
	// Handle help flag
	if len(args) > 1 && (args[1] == "--help" || args[1] == "-h" || args[1] == "help") {
		fmt.Fprint(r.stdout, r.formatter.FormatHelp(r.calculator.ListOperations()))
		return 0
	}

	// Handle version flag
	if len(args) > 1 && (args[1] == "--version" || args[1] == "-v") {
		fmt.Fprintf(r.stdout, "calculator version %s\n", Version)
		return 0
	}

	// Show help if no arguments (just program name)
	if len(args) <= 1 {
		fmt.Fprint(r.stdout, r.formatter.FormatHelp(r.calculator.ListOperations()))
		return 0
	}

	// Parse arguments (skip program name at args[0])
	parseResult, err := r.parser.Parse(args[1:])
	if err != nil {
		fmt.Fprintln(r.stderr, r.formatter.FormatError(err))
		return 1
	}

	// Execute calculation
	result, symbol, err := r.calculator.ExecuteOperation(
		parseResult.Operation,
		parseResult.OperandA,
		parseResult.OperandB,
	)
	if err != nil {
		fmt.Fprintln(r.stderr, r.formatter.FormatError(err))
		return 1
	}

	// Format and print result
	output := r.formatter.FormatResult(
		parseResult.Operation,
		symbol,
		parseResult.OperandA,
		parseResult.OperandB,
		result,
	)
	fmt.Fprintln(r.stdout, output)

	return 0
}
