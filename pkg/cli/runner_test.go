package cli

import (
	"bytes"
	"strings"
	"testing"

	"github.com/kairostack/go-calculator/internal/calculator"
	"github.com/kairostack/go-calculator/internal/operations"
)

func TestRunner_Run(t *testing.T) {
	// Reset the default registry for consistent test state
	operations.DefaultRegistry = operations.NewRegistry()
	operations.DefaultRegistry.Register(&operations.AddOperation{})
	operations.DefaultRegistry.Register(&operations.SubtractOperation{})
	operations.DefaultRegistry.Register(&operations.MultiplyOperation{})
	operations.DefaultRegistry.Register(&operations.DivideOperation{})

	calc := calculator.New()

	tests := []struct {
		name           string
		args           []string
		wantExitCode   int
		wantOutput     string
		wantErrContain string
	}{
		{
			name:         "successful addition",
			args:         []string{"calc", "add", "5", "3"},
			wantExitCode: 0,
			wantOutput:   "5 + 3 = 8",
		},
		{
			name:         "successful subtraction",
			args:         []string{"calc", "subtract", "10", "4"},
			wantExitCode: 0,
			wantOutput:   "10 - 4 = 6",
		},
		{
			name:         "successful multiplication",
			args:         []string{"calc", "multiply", "6", "7"},
			wantExitCode: 0,
			wantOutput:   "6 * 7 = 42",
		},
		{
			name:         "successful division",
			args:         []string{"calc", "divide", "15", "3"},
			wantExitCode: 0,
			wantOutput:   "15 / 3 = 5",
		},
		{
			name:         "help flag",
			args:         []string{"calc", "--help"},
			wantExitCode: 0,
			wantOutput:   "Available operations:",
		},
		{
			name:         "help flag short",
			args:         []string{"calc", "-h"},
			wantExitCode: 0,
			wantOutput:   "Available operations:",
		},
		{
			name:         "version flag",
			args:         []string{"calc", "--version"},
			wantExitCode: 0,
			wantOutput:   "calculator version",
		},
		{
			name:         "version flag short",
			args:         []string{"calc", "-v"},
			wantExitCode: 0,
			wantOutput:   "calculator version",
		},
		{
			name:         "no arguments shows help",
			args:         []string{"calc"},
			wantExitCode: 0,
			wantOutput:   "Available operations:",
		},
		{
			name:           "insufficient arguments",
			args:           []string{"calc", "add", "5"},
			wantExitCode:   1,
			wantErrContain: "insufficient arguments",
		},
		{
			name:           "invalid operation",
			args:           []string{"calc", "invalid", "5", "3"},
			wantExitCode:   1,
			wantErrContain: "operation not found",
		},
		{
			name:           "invalid first operand",
			args:           []string{"calc", "add", "abc", "3"},
			wantExitCode:   1,
			wantErrContain: "invalid",
		},
		{
			name:           "invalid second operand",
			args:           []string{"calc", "add", "5", "xyz"},
			wantExitCode:   1,
			wantErrContain: "invalid",
		},
		{
			name:           "division by zero",
			args:           []string{"calc", "divide", "5", "0"},
			wantExitCode:   1,
			wantErrContain: "division by zero",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			runner := NewRunner(calc, nil, &stdout, &stderr)

			got := runner.Run(tt.args)

			if got != tt.wantExitCode {
				t.Errorf("Run() exit code = %d, want %d", got, tt.wantExitCode)
			}

			if tt.wantOutput != "" && !strings.Contains(stdout.String(), tt.wantOutput) {
				t.Errorf("stdout = %q, want to contain %q", stdout.String(), tt.wantOutput)
			}

			if tt.wantErrContain != "" && !strings.Contains(stderr.String(), tt.wantErrContain) {
				t.Errorf("stderr = %q, want to contain %q", stderr.String(), tt.wantErrContain)
			}
		})
	}
}

func TestRunner_Run_DecimalResults(t *testing.T) {
	operations.DefaultRegistry = operations.NewRegistry()
	operations.DefaultRegistry.Register(&operations.AddOperation{})
	operations.DefaultRegistry.Register(&operations.DivideOperation{})

	calc := calculator.New()

	tests := []struct {
		name       string
		args       []string
		wantOutput string
	}{
		{
			name:       "decimal addition",
			args:       []string{"calc", "add", "5.5", "3.2"},
			wantOutput: "5.5 + 3.2 = 8.7",
		},
		{
			name:       "decimal division",
			args:       []string{"calc", "divide", "10", "4"},
			wantOutput: "10 / 4 = 2.5",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout bytes.Buffer
			runner := NewRunner(calc, nil, &stdout, nil)

			exitCode := runner.Run(tt.args)

			if exitCode != 0 {
				t.Errorf("Run() exit code = %d, want 0", exitCode)
			}

			if !strings.Contains(stdout.String(), tt.wantOutput) {
				t.Errorf("stdout = %q, want to contain %q", stdout.String(), tt.wantOutput)
			}
		})
	}
}

func TestRunner_NewRunner(t *testing.T) {
	calc := calculator.New()
	var stdout, stderr bytes.Buffer

	runner := NewRunner(calc, nil, &stdout, &stderr)

	if runner.calculator != calc {
		t.Error("NewRunner() did not set calculator correctly")
	}
	if runner.formatter == nil {
		t.Error("NewRunner() did not initialize formatter")
	}
	if runner.parser == nil {
		t.Error("NewRunner() did not initialize parser")
	}
	if runner.stdout != &stdout {
		t.Error("NewRunner() did not set stdout correctly")
	}
	if runner.stderr != &stderr {
		t.Error("NewRunner() did not set stderr correctly")
	}
}
