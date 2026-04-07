# Go Calculator

A robust command-line calculator written in Go, implementing clean architecture principles with the Strategy pattern for operations.

## Features

- **Strategy Pattern**: Each operation (add, subtract, multiply, divide) implements a common interface
- **Registry Pattern**: Dynamic operation registration and lookup
- **Idiomatic Error Handling**: Proper error types for division by zero, invalid operations, etc.
- **Thread-Safe**: Registry uses mutex for concurrent access
- **Well-Tested**: Comprehensive unit tests with >80% coverage
- **Clean Architecture**: Separation of concerns across packages

## Project Structure

```
calculator/
├── cmd/calculator/         # Application entry point
│   └── main.go
├── internal/
│   ├── calculator/         # Core calculator logic
│   │   ├── calculator.go
│   │   └── calculator_test.go
│   ├── operations/         # Operation implementations (Strategy pattern)
│   │   ├── interface.go    # Operation interface
│   │   ├── registry.go     # Operation registry
│   │   ├── add.go
│   │   ├── subtract.go
│   │   ├── multiply.go
│   │   ├── divide.go
│   │   └── *_test.go       # Unit tests
│   └── errors/             # Custom error types
│       └── errors.go
├── pkg/cli/                # CLI utilities
│   ├── parser.go           # Argument parsing
│   ├── parser_test.go
│   ├── formatter.go        # Output formatting
│   └── formatter_test.go
├── go.mod
├── Makefile
└── README.md
```

## Installation

```bash
# Clone the repository
git clone https://github.com/kairostack/go-calculator.git
cd go-calculator

# Build the binary
make build

# Or install to GOPATH/bin
make install
```

## Usage

### Basic Usage

```bash
# Addition
./build/calculator add 5 3
# Output: 5 + 3 = 8

# Subtraction
./build/calculator subtract 10 4
# Output: 10 - 4 = 6

# Multiplication
./build/calculator multiply 6 7
# Output: 6 * 7 = 42

# Division
./build/calculator divide 15 3
# Output: 15 / 3 = 5
```

### Help

```bash
./build/calculator
# Shows available operations and examples
```

### Decimal Numbers

```bash
./build/calculator add 5.5 3.2
# Output: 5.5 + 3.2 = 8.7

./build/calculator divide 7.5 2.5
# Output: 7.5 / 2.5 = 3
```

### Negative Numbers

```bash
./build/calculator add -5 -3
# Output: -5 + -3 = -8

./build/calculator multiply -4 5
# Output: -4 * 5 = -20
```

## Error Handling

The calculator provides clear error messages:

```bash
# Division by zero
./build/calculator divide 5 0
# Error: calculator error [divide]: division by zero

# Invalid operation
./build/calculator power 2 3
# Error: calculator error [power]: operation not found

# Invalid arguments
./build/calculator add abc 3
# Error: calculator error [parse]: invalid first operand ('abc' is not a valid number)
```

## Development

### Prerequisites

- Go 1.21 or higher
- Make (optional, for using Makefile)

### Running Tests

```bash
# Run all tests
make test

# Run with race detection
make test-race

# Generate coverage report
make coverage

# View coverage summary
make coverage-summary
```

### Project Initialization

```bash
# Run the init command - sets up everything
make init
```

This will:
1. Download dependencies
2. Build the binary
3. Run all tests

### Code Organization

- **`cmd/`** - Application entry points
- **`internal/`** - Private application code
  - **`operations/`** - Strategy pattern implementation for operations
  - **`calculator/`** - Core calculator orchestration
  - **`errors/`** - Custom error types
- **`pkg/`** - Public libraries that can be imported
  - **`cli/`** - CLI parsing and formatting utilities

## Architecture

### Strategy Pattern

Each operation implements the `Operation` interface:

```go
type Operation interface {
    Execute(a, b float64) (float64, error)
    Name() string
    Description() string
    Symbol() string
}
```

This allows easy addition of new operations without modifying existing code (Open/Closed Principle).

### Registry Pattern

Operations are registered in a thread-safe registry:

```go
registry := operations.NewRegistry()
registry.Register(&operations.AddOperation{})

op, err := registry.Get("add")
```

### Error Handling

Custom error types provide context:

```go
type CalculatorError struct {
    Op      string
    Err     string
    Details string
}
```

## Testing

The project includes comprehensive tests:

- **Unit tests** for each operation
- **Integration tests** for the calculator
- **Table-driven tests** for multiple scenarios
- **Concurrent access tests** for thread safety

Run tests with:

```bash
make test
```

Generate coverage report:

```bash
make coverage
# Opens coverage.html in browser
```

## Makefile Targets

| Target | Description |
|--------|-------------|
| `make init` | Initialize project (deps, build, test) |
| `make build` | Build the calculator binary |
| `make test` | Run all tests |
| `make test-race` | Run tests with race detector |
| `make coverage` | Generate coverage report |
| `make run` | Run example calculation |
| `make install` | Install to GOPATH/bin |
| `make clean` | Remove build artifacts |
| `make fmt` | Format Go code |
| `make vet` | Run go vet |
| `make lint` | Run golangci-lint |
| `make help` | Show help |

## License

MIT License - see LICENSE file for details.

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Examples

### Adding a New Operation

To add a new operation (e.g., `power`):

1. Create `internal/operations/power.go`:

```go
package operations

type PowerOperation struct{}

func (p *PowerOperation) Execute(a, b float64) (float64, error) {
    result := 1.0
    for i := 0; i < int(b); i++ {
        result *= a
    }
    return result, nil
}

func (p *PowerOperation) Name() string        { return "power" }
func (p *PowerOperation) Description() string { return "Raises first number to power of second" }
func (p *PowerOperation) Symbol() string      { return "^" }
```

2. Register in `internal/calculator/calculator.go`:

```go
calc.registry.Register(&operations.PowerOperation{})
```

3. Add tests in `internal/operations/power_test.go`

No other code changes needed!