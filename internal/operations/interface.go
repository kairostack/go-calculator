package operations

// Operation defines the interface for all calculator operations
// This is the Strategy pattern - each operation implements this interface
type Operation interface {
	// Execute performs the calculation with the given operands
	// Returns the result or an error if the operation fails
	Execute(a, b float64) (float64, error)

	// Name returns the operation identifier (e.g., "add", "subtract")
	Name() string

	// Description returns a human-readable description of the operation
	Description() string

	// Symbol returns the mathematical symbol for the operation
	Symbol() string
}
