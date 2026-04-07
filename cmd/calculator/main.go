package main

import (
	"fmt"
	"os"

	"github.com/kairostack/go-calculator/internal/calculator"
	"github.com/kairostack/go-calculator/pkg/cli"
)

const (
	exitSuccess = 0
	exitError   = 1
)

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	// Create calculator and formatter
	calc := calculator.New()
	formatter := cli.NewFormatter()

	// Show help if no arguments
	if len(args) == 0 {
		fmt.Print(formatter.FormatHelp(calc.ListOperations()))
		return exitSuccess
	}

	// Handle help flag
	if args[0] == "--help" || args[0] == "-h" || args[0] == "help" {
		fmt.Print(formatter.FormatHelp(calc.ListOperations()))
		return exitSuccess
	}

	// Parse arguments
	parser := cli.NewParser()
	parseResult, err := parser.Parse(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, formatter.FormatError(err))
		return exitError
	}

	// Execute calculation
	result, symbol, err := calc.ExecuteOperation(
		parseResult.Operation,
		parseResult.OperandA,
		parseResult.OperandB,
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, formatter.FormatError(err))
		return exitError
	}

	// Format and print result
	output := formatter.FormatResult(
		parseResult.Operation,
		symbol,
		parseResult.OperandA,
		parseResult.OperandB,
		result,
	)
	fmt.Println(output)

	return exitSuccess
}
