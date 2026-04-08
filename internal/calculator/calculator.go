package calculator

import (
	"io"
	"log"

	"github.com/kairostack/go-calculator/internal/operations"
)

// Calculator orchestrates calculation operations
// Uses the Strategy pattern via the operations registry
type Calculator struct {
	registry *operations.Registry
	logger   *log.Logger
}

// New creates a new Calculator with all operations pre-registered
// Uses the DefaultRegistry which is populated via init() functions
func New() *Calculator {
	return &Calculator{
		registry: operations.DefaultRegistry,
		logger:   log.New(io.Discard, "[calculator] ", log.LstdFlags), // No-op by default
	}
}

// NewWithRegistry creates a Calculator with a custom registry
// Useful for testing with mock operations
func NewWithRegistry(registry *operations.Registry) *Calculator {
	return &Calculator{
		registry: registry,
		logger:   log.New(io.Discard, "[calculator] ", log.LstdFlags), // No-op by default
	}
}

// NewWithLogger creates a Calculator with a custom logger
func NewWithLogger(logger *log.Logger) *Calculator {
	calc := New()
	calc.logger = logger
	return calc
}

// Calculate performs the specified operation on two operands
// Returns the result or an error if the operation fails
func (c *Calculator) Calculate(operationName string, a, b float64) (float64, error) {
	c.logger.Printf("Executing operation: %s(%f, %f)", operationName, a, b)

	op, err := c.registry.Get(operationName)
	if err != nil {
		c.logger.Printf("Operation not found: %s", operationName)
		return 0, err
	}

	result, err := op.Execute(a, b)
	if err != nil {
		c.logger.Printf("Operation %s failed: %v", operationName, err)
		return 0, err
	}

	c.logger.Printf("Operation %s result: %f", operationName, result)
	return result, nil
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
	c.logger.Printf("Executing operation: %s(%f, %f)", operationName, a, b)

	op, err := c.registry.Get(operationName)
	if err != nil {
		c.logger.Printf("Operation not found: %s", operationName)
		return 0, "", err
	}

	result, err := op.Execute(a, b)
	if err != nil {
		c.logger.Printf("Operation %s failed: %v", operationName, err)
		return 0, op.Symbol(), err
	}

	c.logger.Printf("Operation %s result: %f", operationName, result)
	return result, op.Symbol(), nil
}

// ValidateOperation checks if an operation exists without executing it
func (c *Calculator) ValidateOperation(name string) bool {
	_, err := c.registry.Get(name)
	return err == nil
}
