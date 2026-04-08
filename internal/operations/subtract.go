package operations

// SubtractOperation implements the Operation interface for subtraction
type SubtractOperation struct{}

// init registers the SubtractOperation with the DefaultRegistry
func init() {
	RegisterDefault(&SubtractOperation{})
}

// Execute subtracts the second number from the first
// Returns an error if either input is NaN or infinite
func (s *SubtractOperation) Execute(x, y float64) (float64, error) {
	if err := validateInputs(x, y); err != nil {
		return 0, err
	}
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
