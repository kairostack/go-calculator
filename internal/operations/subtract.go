package operations

// SubtractOperation implements the Operation interface for subtraction
type SubtractOperation struct{}

// Execute subtracts the second number from the first
func (s *SubtractOperation) Execute(x, y float64) (float64, error) {
	return x - y, nil
}

// Name returns the operation identifier
func (s *SubtractOperation) Name() string {
	return "subtract"
}

// Description returns a human-readable description
func (s *SubtractOperation) Description() string {
	return "Subtracts the second number from the first"
}

// Symbol returns the mathematical symbol
func (s *SubtractOperation) Symbol() string {
	return "-"
}
