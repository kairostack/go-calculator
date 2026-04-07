package operations

import (
	"testing"
)

// MockOperation is a test double for the Operation interface
type MockOperation struct {
	name        string
	symbol      string
	description string
	executeFunc func(a, b float64) (float64, error)
}

func (m *MockOperation) Execute(a, b float64) (float64, error) {
	if m.executeFunc != nil {
		return m.executeFunc(a, b)
	}
	return 0, nil
}

func (m *MockOperation) Name() string        { return m.name }
func (m *MockOperation) Description() string { return m.description }
func (m *MockOperation) Symbol() string      { return m.symbol }

func TestRegistry_RegisterAndGet(t *testing.T) {
	reg := NewRegistry()

	// Register a mock operation
	mockOp := &MockOperation{
		name:        "test",
		symbol:      "?",
		description: "Test operation",
		executeFunc: func(a, b float64) (float64, error) {
			return a + b, nil
		},
	}

	reg.Register(mockOp)

	// Retrieve the operation
	op, err := reg.Get("test")
	if err != nil {
		t.Fatalf("Get() unexpected error: %v", err)
	}

	if op.Name() != "test" {
		t.Errorf("Get() name = %q, want %q", op.Name(), "test")
	}
}

func TestRegistry_Get_NotFound(t *testing.T) {
	reg := NewRegistry()

	_, err := reg.Get("nonexistent")
	if err == nil {
		t.Error("Get() expected error for nonexistent operation, got nil")
	}
}

func TestRegistry_List(t *testing.T) {
	reg := NewRegistry()

	// Register multiple operations
	reg.Register(&MockOperation{name: "op1"})
	reg.Register(&MockOperation{name: "op2"})
	reg.Register(&MockOperation{name: "op3"})

	names := reg.List()
	if len(names) != 3 {
		t.Errorf("List() returned %d names, want 3", len(names))
	}

	// Check all names are present
	nameMap := make(map[string]bool)
	for _, name := range names {
		nameMap[name] = true
	}

	for _, expected := range []string{"op1", "op2", "op3"} {
		if !nameMap[expected] {
			t.Errorf("List() missing expected operation: %s", expected)
		}
	}
}

func TestRegistry_Count(t *testing.T) {
	reg := NewRegistry()

	if reg.Count() != 0 {
		t.Errorf("Count() = %d, want 0", reg.Count())
	}

	reg.Register(&MockOperation{name: "op1"})
	if reg.Count() != 1 {
		t.Errorf("Count() = %d, want 1", reg.Count())
	}
}

func TestRegistry_Unregister(t *testing.T) {
	reg := NewRegistry()

	reg.Register(&MockOperation{name: "op1"})

	// Unregister existing
	if !reg.Unregister("op1") {
		t.Error("Unregister() = false for existing operation")
	}

	// Verify it's gone
	if reg.Count() != 0 {
		t.Errorf("Count() after unregister = %d, want 0", reg.Count())
	}

	// Unregister non-existent
	if reg.Unregister("nonexistent") {
		t.Error("Unregister() = true for nonexistent operation")
	}
}

func TestRegistry_ConcurrentAccess(t *testing.T) {
	reg := NewRegistry()

	// Test concurrent writes
	done := make(chan bool)

	for i := 0; i < 10; i++ {
		go func(n int) {
			reg.Register(&MockOperation{name: string(rune('a' + n))})
			done <- true
		}(i)
	}

	for i := 0; i < 10; i++ {
		<-done
	}

	if reg.Count() != 10 {
		t.Errorf("Count() after concurrent writes = %d, want 10", reg.Count())
	}
}
