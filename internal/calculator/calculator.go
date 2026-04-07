package calculator

import (
	"github.com/kairostack/go-calculator/internal/operations"
)

// Calculator orchestrates calculation operations
// Uses the Strategy pattern via the operations registry
type Calculator struct {
	registry *operations.Registry
}

// New creates a new Calculator with all operations pre-registered
func New() *Calculator {
	calc := &Calculator{
		registry: operations.NewRegistry(),
	}

	// Register all available operations
	calc.registry.Register(&operations.AddOperation{})
	calc.registry.Register(&operations.SubtractOperation{})
	calc.registry.Register(&operations.MultiplyOperation{})
	calc.registry.Register(&operations.DivideOperation{})

	return calc
}

// NewWithRegistry creates a Calculator with a custom registry
// Useful for testing with mock operations
func NewWithRegistry(registry *operations.Registry) *Calculator {
	return &Calculator{
		registry: registry,
	}
}

// Calculate performs the specified operation on two operands
// Returns the result or an error if the operation fails
func (c *Calculator) Calculate(operationName string, a, b float64) (float64, error) {
	op, err := c.registry.Get(operationName)
	if err != nil {
		return 0, err
	}

	return op.Execute(a, b)
}

// GetOperation retrieves an operation by name
func (c *Calculator) GetOperation(name string) (operations.Operation, error) {
	return c.registry.Get(name)
}

// ListOperations returns all available operation names
func (c *Calculator) ListOperations() []string {
	return c.registry.List()
}

// GetOperationDetails returns details about a specific operation
func (c *Calculator) GetOperationDetails(name string) (map[string]string, error) {
	op, err := c.registry.Get(name)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"name":        op.Name(),
		"symbol":      op.Symbol(),
		"description": op.Description(),
	}, nil
}

// ExecuteOperation is a convenience method that parses and executes in one call
// Returns both the result and the operation symbol for formatting
func (c *Calculator) ExecuteOperation(operationName string, a, b float64) (float64, string, error) {
	op, err := c.registry.Get(operationName)
	if err != nil {
		return 0, "", err
	}

	result, err := op.Execute(a, b)
	if err != nil {
		return 0, op.Symbol(), err
	}

	return result, op.Symbol(), nil
}

// ValidateOperation checks if an operation exists without executing it
func (c *Calculator) ValidateOperation(name string) bool {
	_, err := c.registry.Get(name)
	return err == nil
}
