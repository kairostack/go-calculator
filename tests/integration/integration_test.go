package integration

import (
	"bytes"
	"strings"
	"testing"

	"github.com/kairostack/go-calculator/internal/calculator"
	"github.com/kairostack/go-calculator/pkg/cli"
)

func TestIntegration_Add(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		wantOut  string
		wantErr  string
		wantCode int
	}{
		{"add 2 3", []string{"calc", "add", "2", "3"}, "5", "", 0},
		{"invalid op", []string{"calc", "invalid", "1", "2"}, "", "operation", 1},
		{"missing args", []string{"calc", "add"}, "", "arguments", 1},
		{"help flag", []string{"calc", "--help"}, "Usage", "", 0},
		{"no args", []string{"calc"}, "Usage", "", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calc := calculator.New()
			var stdout, stderr bytes.Buffer
			runner := cli.NewRunner(calc, nil, &stdout, &stderr)

			code := runner.Run(tt.args)

			if code != tt.wantCode {
				t.Errorf("exit code = %d, want %d", code, tt.wantCode)
			}
			if tt.wantOut != "" && !strings.Contains(stdout.String(), tt.wantOut) {
				t.Errorf("stdout = %q, want to contain %q", stdout.String(), tt.wantOut)
			}
			if tt.wantErr != "" && !strings.Contains(stderr.String(), tt.wantErr) {
				t.Errorf("stderr = %q, want to contain %q", stderr.String(), tt.wantErr)
			}
		})
	}
}

func TestIntegration_Subtract(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		wantOut  string
		wantCode int
	}{
		{"subtract 5 3", []string{"calc", "subtract", "5", "3"}, "2", 0},
		{"subtract negative", []string{"calc", "subtract", "5", "-3"}, "8", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calc := calculator.New()
			var stdout, stderr bytes.Buffer
			runner := cli.NewRunner(calc, nil, &stdout, &stderr)

			code := runner.Run(tt.args)

			if code != tt.wantCode {
				t.Errorf("exit code = %d, want %d", code, tt.wantCode)
			}
			if tt.wantOut != "" && !strings.Contains(stdout.String(), tt.wantOut) {
				t.Errorf("stdout = %q, want to contain %q", stdout.String(), tt.wantOut)
			}
		})
	}
}

func TestIntegration_Multiply(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		wantOut  string
		wantCode int
	}{
		{"multiply 6 7", []string{"calc", "multiply", "6", "7"}, "42", 0},
		{"multiply by zero", []string{"calc", "multiply", "5", "0"}, "0", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calc := calculator.New()
			var stdout, stderr bytes.Buffer
			runner := cli.NewRunner(calc, nil, &stdout, &stderr)

			code := runner.Run(tt.args)

			if code != tt.wantCode {
				t.Errorf("exit code = %d, want %d", code, tt.wantCode)
			}
			if tt.wantOut != "" && !strings.Contains(stdout.String(), tt.wantOut) {
				t.Errorf("stdout = %q, want to contain %q", stdout.String(), tt.wantOut)
			}
		})
	}
}

func TestIntegration_Divide(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		wantOut  string
		wantErr  string
		wantCode int
	}{
		{"divide 15 3", []string{"calc", "divide", "15", "3"}, "5", "", 0},
		{"divide by zero", []string{"calc", "divide", "5", "0"}, "", "division by zero", 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calc := calculator.New()
			var stdout, stderr bytes.Buffer
			runner := cli.NewRunner(calc, nil, &stdout, &stderr)

			code := runner.Run(tt.args)

			if code != tt.wantCode {
				t.Errorf("exit code = %d, want %d", code, tt.wantCode)
			}
			if tt.wantOut != "" && !strings.Contains(stdout.String(), tt.wantOut) {
				t.Errorf("stdout = %q, want to contain %q", stdout.String(), tt.wantOut)
			}
			if tt.wantErr != "" && !strings.Contains(stderr.String(), tt.wantErr) {
				t.Errorf("stderr = %q, want to contain %q", stderr.String(), tt.wantErr)
			}
		})
	}
}

func TestIntegration_InvalidNumber(t *testing.T) {
	calc := calculator.New()
	var stdout, stderr bytes.Buffer
	runner := cli.NewRunner(calc, nil, &stdout, &stderr)

	args := []string{"calc", "add", "abc", "3"}
	code := runner.Run(args)

	if code != 1 {
		t.Errorf("exit code = %d, want 1", code)
	}
	if !strings.Contains(stderr.String(), "invalid") {
		t.Errorf("stderr = %q, want to contain 'invalid'", stderr.String())
	}
}

func TestIntegration_Version(t *testing.T) {
	calc := calculator.New()
	var stdout, stderr bytes.Buffer
	runner := cli.NewRunner(calc, nil, &stdout, &stderr)

	// Test --version flag
	args := []string{"calc", "--version"}
	code := runner.Run(args)

	if code != 0 {
		t.Errorf("exit code = %d, want 0", code)
	}
	if !strings.Contains(stdout.String(), "version") {
		t.Errorf("stdout = %q, want to contain 'version'", stdout.String())
	}

	// Test -v flag
	stdout.Reset()
	args = []string{"calc", "-v"}
	code = runner.Run(args)

	if code != 0 {
		t.Errorf("exit code = %d, want 0", code)
	}
	if !strings.Contains(stdout.String(), "version") {
		t.Errorf("stdout = %q, want to contain 'version'", stdout.String())
	}
}
