package cli

import (
	"strings"
	"testing"
)

func TestParser_Parse_ValidInput(t *testing.T) {
	parser := NewParser()

	tests := []struct {
		name   string
		args   []string
		wantOp string
		wantA  float64
		wantB  float64
	}{
		{"add integers", []string{"add", "5", "3"}, "add", 5, 3},
		{"subtract integers", []string{"subtract", "10", "4"}, "subtract", 10, 4},
		{"with decimals", []string{"multiply", "2.5", "4"}, "multiply", 2.5, 4},
		{"negative numbers", []string{"add", "-5", "-3"}, "add", -5, -3},
		{"mixed signs", []string{"subtract", "5", "-3"}, "subtract", 5, -3},
		{"uppercase operation", []string{"ADD", "5", "3"}, "add", 5, 3},
		{"with spaces", []string{"  add  ", "  5  ", "  3  "}, "add", 5, 3},
		{"scientific notation", []string{"multiply", "1e3", "2"}, "multiply", 1000, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parser.Parse(tt.args)
			if err != nil {
				t.Fatalf("Parse() unexpected error: %v", err)
			}

			if result.Operation != tt.wantOp {
				t.Errorf("Parse() operation = %q, want %q", result.Operation, tt.wantOp)
			}
			if result.OperandA != tt.wantA {
				t.Errorf("Parse() operandA = %g, want %g", result.OperandA, tt.wantA)
			}
			if result.OperandB != tt.wantB {
				t.Errorf("Parse() operandB = %g, want %g", result.OperandB, tt.wantB)
			}
		})
	}
}

func TestParser_Parse_InvalidInput(t *testing.T) {
	parser := NewParser()

	tests := []struct {
		name string
		args []string
	}{
		{"too few arguments", []string{"add", "5"}},
		{"empty args", []string{}},
		{"single arg", []string{"add"}},
		{"invalid first number", []string{"add", "abc", "3"}},
		{"invalid second number", []string{"add", "5", "xyz"}},
		{"both invalid", []string{"add", "abc", "xyz"}},
		{"empty operation", []string{"", "5", "3"}},
		{"only spaces operation", []string{"   ", "5", "3"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := parser.Parse(tt.args)
			if err == nil {
				t.Error("Parse() expected error, got nil")
			}
		})
	}
}

func TestParseResult_String(t *testing.T) {
	pr := &ParseResult{
		Operation: "add",
		OperandA:  5,
		OperandB:  3,
	}

	result := pr.String()
	if !strings.Contains(result, "add") {
		t.Error("String() should contain operation name")
	}
	if !strings.Contains(result, "5") {
		t.Error("String() should contain first operand")
	}
	if !strings.Contains(result, "3") {
		t.Error("String() should contain second operand")
	}
}
