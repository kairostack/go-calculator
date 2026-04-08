package main

import (
	"os"

	"github.com/kairostack/go-calculator/internal/calculator"
	"github.com/kairostack/go-calculator/pkg/cli"
)

func main() {
	calc := calculator.New()
	runner := cli.NewRunner(calc, os.Stdin, os.Stdout, os.Stderr)
	os.Exit(runner.Run(os.Args))
}
