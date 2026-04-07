package operations

// MultiplyOperation implements the Operation interface for multiplication
type MultiplyOperation struct{}

// Execute multiplies two numbers together
func (m *MultiplyOperation) Execute(x, y float64) (float64, error) {
	return x * y, nil
}

// Name returns the operation identifier
func (m *MultiplyOperation) Name() string {
	return "multiply"
}

// Description returns a human-readable description
func (m *MultiplyOperation) Description() string {
	return "Multiplies two numbers together"
}

// Symbol returns the mathematical symbol
func (m *MultiplyOperation) Symbol() string {
	return "*"
}
