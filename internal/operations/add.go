package operations

// AddOperation implements the Operation interface for addition
type AddOperation struct{}

// Execute adds two numbers together
func (a *AddOperation) Execute(x, y float64) (float64, error) {
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
