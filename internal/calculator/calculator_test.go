package calculator

import (
	"testing"

	"github.com/kairostack/go-calculator/internal/operations"
)

func TestNew(t *testing.T) {
	calc := New()

	// Verify all operations are registered
	expectedOps := []string{"add", "subtract", "multiply", "divide"}
	registeredOps := calc.ListOperations()

	if len(registeredOps) != len(expectedOps) {
		t.Errorf("New() registered %d operations, want %d", len(registeredOps), len(expectedOps))
	}

	// Verify each expected operation exists
	for _, opName := range expectedOps {
		if !calc.ValidateOperation(opName) {
			t.Errorf("New() missing operation: %s", opName)
		}
	}
}

func TestCalculator_Calculate(t *testing.T) {
	calc := New()

	tests := []struct {
		name        string
		operation   string
		a, b        float64
		expected    float64
		expectError bool
	}{
		{"add", "add", 5, 3, 8, false},
		{"subtract", "subtract", 10, 4, 6, false},
		{"multiply", "multiply", 6, 7, 42, false},
		{"divide", "divide", 15, 3, 5, false},
		{"divide by zero", "divide", 5, 0, 0, true},
		{"invalid operation", "power", 5, 3, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.Calculate(tt.operation, tt.a, tt.b)

			if tt.expectError {
				if err == nil {
					t.Errorf("Calculate() expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("Calculate() unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("Calculate() = %g, want %g", result, tt.expected)
			}
		})
	}
}

func TestCalculator_GetOperation(t *testing.T) {
	calc := New()

	tests := []struct {
		name        string
		operation   string
		expectError bool
	}{
		{"existing add", "add", false},
		{"existing subtract", "subtract", false},
		{"existing multiply", "multiply", false},
		{"existing divide", "divide", false},
		{"nonexistent", "power", true},
		{"empty", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op, err := calc.GetOperation(tt.operation)

			if tt.expectError {
				if err == nil {
					t.Error("GetOperation() expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("GetOperation() unexpected error: %v", err)
			}

			if op.Name() != tt.operation {
				t.Errorf("GetOperation() name = %q, want %q", op.Name(), tt.operation)
			}
		})
	}
}

func TestCalculator_ListOperations(t *testing.T) {
	calc := New()

	ops := calc.ListOperations()

	if len(ops) == 0 {
		t.Error("ListOperations() returned empty list")
	}

	// Verify all expected operations are present
	expected := map[string]bool{"add": false, "subtract": false, "multiply": false, "divide": false}
	for _, op := range ops {
		if _, exists := expected[op]; exists {
			expected[op] = true
		}
	}

	for op, found := range expected {
		if !found {
			t.Errorf("ListOperations() missing operation: %s", op)
		}
	}
}

func TestCalculator_GetOperationDetails(t *testing.T) {
	calc := New()

	details, err := calc.GetOperationDetails("add")
	if err != nil {
		t.Fatalf("GetOperationDetails() unexpected error: %v", err)
	}

	requiredFields := []string{"name", "symbol", "description"}
	for _, field := range requiredFields {
		if _, ok := details[field]; !ok {
			t.Errorf("GetOperationDetails() missing field: %s", field)
		}
	}

	if details["name"] != "add" {
		t.Errorf("GetOperationDetails() name = %q, want %q", details["name"], "add")
	}

	// Test nonexistent operation
	_, err = calc.GetOperationDetails("nonexistent")
	if err == nil {
		t.Error("GetOperationDetails() expected error for nonexistent operation")
	}
}

func TestCalculator_ExecuteOperation(t *testing.T) {
	calc := New()

	result, symbol, err := calc.ExecuteOperation("add", 5, 3)
	if err != nil {
		t.Fatalf("ExecuteOperation() unexpected error: %v", err)
	}

	if result != 8 {
		t.Errorf("ExecuteOperation() result = %g, want 8", result)
	}

	if symbol != "+" {
		t.Errorf("ExecuteOperation() symbol = %q, want +", symbol)
	}

	// Test with error
	_, symbol, err = calc.ExecuteOperation("divide", 5, 0)
	if err == nil {
		t.Error("ExecuteOperation() expected error for division by zero")
	}
	// Symbol should still be returned even on error
	if symbol != "/" {
		t.Errorf("ExecuteOperation() symbol on error = %q, want /", symbol)
	}

	// Test nonexistent operation
	_, _, err = calc.ExecuteOperation("nonexistent", 5, 3)
	if err == nil {
		t.Error("ExecuteOperation() expected error for nonexistent operation")
	}
}

func TestCalculator_ValidateOperation(t *testing.T) {
	calc := New()

	tests := []struct {
		operation string
		expected  bool
	}{
		{"add", true},
		{"subtract", true},
		{"multiply", true},
		{"divide", true},
		{"power", false},
		{"", false},
		{"ADD", false}, // case-sensitive validation
	}

	for _, tt := range tests {
		t.Run(tt.operation, func(t *testing.T) {
			result := calc.ValidateOperation(tt.operation)
			if result != tt.expected {
				t.Errorf("ValidateOperation(%q) = %v, want %v", tt.operation, result, tt.expected)
			}
		})
	}
}

func TestNewWithRegistry(t *testing.T) {
	// Create a custom registry with only add operation
	registry := operations.NewRegistry()
	registry.Register(&operations.AddOperation{})

	calc := NewWithRegistry(registry)

	// Should only have add operation
	if !calc.ValidateOperation("add") {
		t.Error("NewWithRegistry() should have add operation")
	}

	if calc.ValidateOperation("subtract") {
		t.Error("NewWithRegistry() should not have subtract operation")
	}
}
