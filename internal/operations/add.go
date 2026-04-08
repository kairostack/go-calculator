package operations

// AddOperation implements the Operation interface for addition
type AddOperation struct{}

// Execute adds two numbers together
// Returns an error if either input is NaN or infinite
func (a *AddOperation) Execute(x, y float64) (float64, error) {
	if err := validateInputs(x, y); err != nil {
		return 0, err
	}
	return x + y, nil
}

// Name returns the operation identifier
func (a *AddOperation) Name() string {
	return "add"
}

// Description returns a human-readable description
func (a *AddOperation) Description() string {
	return "Adds two numbers together"
}

// Symbol returns the mathematical symbol
func (a *AddOperation) Symbol() string {
	return "+"
}
