package operations

import (
	"fmt"
	"sort"
	"sync"

	"github.com/kairostack/go-calculator/internal/errors"
)

var (
	// DefaultRegistry is the global default registry that is pre-populated
	// via init() functions in each operation file
	DefaultRegistry = NewRegistry()
	registerOnce    sync.Once
)

// RegisterDefault registers an operation to the DefaultRegistry.
// This is called from init() functions in each operation file.
// Thread-safe and idempotent.
func RegisterDefault(op Operation) {
	registerOnce.Do(func() {
		DefaultRegistry = NewRegistry()
	})
	DefaultRegistry.Register(op)
}

// Registry maintains a thread-safe collection of operations
// Uses the Registry pattern for dynamic operation lookup
type Registry struct {
	mu         sync.RWMutex
	operations map[string]Operation
}

// NewRegistry creates a new operation registry
func NewRegistry() *Registry {
	return &Registry{
		operations: make(map[string]Operation),
	}
}

// Register adds an operation to the registry
// Thread-safe for concurrent access
func (r *Registry) Register(op Operation) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.operations[op.Name()] = op
}

// Get retrieves an operation by name
// Returns ErrInvalidOperation if the operation is not found
func (r *Registry) Get(name string) (Operation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if op, ok := r.operations[name]; ok {
		return op, nil
	}

	return nil, &errors.CalculatorError{
		Op:      name,
		Err:     "operation not found",
		Details: fmt.Sprintf("available operations: %v", r.List()),
	}
}

// List returns all registered operation names, sorted alphabetically
func (r *Registry) List() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()

	names := make([]string, 0, len(r.operations))
	for name := range r.operations {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

// Count returns the number of registered operations
func (r *Registry) Count() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.operations)
}

// Unregister removes an operation from the registry
func (r *Registry) Unregister(name string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.operations[name]; exists {
		delete(r.operations, name)
		return true
	}
	return false
}
