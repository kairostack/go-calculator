# Agent Guide: go-calculator

## Quick Start

```bash
make init     # deps + build + test
make run      # build and run example: add 5 3
./build/calculator add 5 3
```

## Architecture Patterns

**Strategy Pattern**: Each operation implements `Operation` interface:
```go
type Operation interface {
    Execute(a, b float64) (float64, error)
    Symbol() string
    Description() string
}
```

**Registry Pattern**: Thread-safe operation registration in `internal/calculator/registry.go`.

**Clean Architecture**: `cmd/` → `internal/` → `pkg/`

## Adding a New Operation

1. Create `internal/operations/{name}.go`
2. Implement the `Operation` interface
3. Register in `internal/calculator/calculator.go`:
   ```go
   calc.RegisterOperation(operations.NewPower())
   ```
4. Add `internal/operations/{name}_test.go` with table-driven tests

**Example test pattern** (copy from existing):
```go
func TestOperation(t *testing.T) {
    tests := []struct {
        name     string
        a, b     float64
        want     float64
        wantErr  bool
    }{
        {"positive", 2, 3, 8, false},
        {"negative", -2, 3, -8, false},
    }
    // ... table-driven test execution
}
```

## Testing

```bash
# All tests
make test

# Single package
go test ./internal/operations/...

# Single test
go test -run TestAdd ./internal/operations/...

# With race detector
make test-race

# Coverage report
make coverage   # generates coverage.html
```

## Makefile Reference

| Command | Purpose |
|---------|---------|
| `make build` | Build to `./build/calculator` |
| `make test` | Run all tests |
| `make test-race` | Run with race detector |
| `make coverage` | Generate HTML coverage report |
| `make run` | Build and run example (`add 5 3`) |
| `make lint` | Run golangci-lint |
| `make fmt` | Run `go fmt` |
| `make vet` | Run `go vet` |
| `make init` | Full setup: deps + build + test |

## Constraints

- **Go stdlib only** - no external dependencies
- **Exit codes**: `0`=success, `1`=error, `2`=invalid args
- **Output**: Errors → `stderr`, Results → `stdout`
- **Module**: `github.com/kairostack/go-calculator`
- **Go version**: 1.21

## File Locations

| Purpose | Path |
|---------|------|
| Entry point | `cmd/calculator/main.go` |
| Orchestrator | `internal/calculator/calculator.go` |
| Operations | `internal/operations/*.go` |
| Operation interface | `internal/operations/interface.go` |
| Registry | `internal/calculator/registry.go` |
| CLI parser | `pkg/cli/parser.go` |
| CLI formatter | `pkg/cli/formatter.go` |
| Build output | `./build/calculator` |
