# Agent Guide: go-calculator

## Quick Start

```bash
make init     # deps + build + test
make run      # build and run: add 5 3
```

## Architecture

**Strategy Pattern**: Operations implement `Operation` interface:
```go
type Operation interface {
    Execute(a, b float64) (float64, error)
    Name() string
    Symbol() string
    Description() string
}
```

**Registry Pattern**: Thread-safe registration in `internal/operations/registry.go`. Operations self-register via `init()`—no manual registration needed.

## Adding a New Operation

```go
// internal/operations/{name}.go
func init() {
    RegisterDefault(&YourOperation{})  // Critical: auto-registers on import
}
```

1. Create `internal/operations/{name}.go` with `init()` registering the operation
2. Implement the `Operation` interface
3. Add `internal/operations/{name}_test.go` with table-driven tests

**Test pattern** (copy from existing):
```go
func TestOperation(t *testing.T) {
    tests := []struct {
        name    string
        a, b    float64
        want    float64
        wantErr bool
    }{
        {"positive", 2, 3, 8, false},
        {"negative", -2, 3, -8, false},
    }
    // Use pkg/floatutil.Equals() for assertions, never ==
}
```

## Testing

```bash
make test          # Run all tests
make test-race     # With race detector
make coverage      # Generates coverage.html
make bench         # All benchmarks
make bench-operations
make fuzz          # All fuzz tests (30s each)
make fuzz-parser   # Fuzz parser (60s)
make fuzz-operations
```

**Critical**: Use `pkg/floatutil.Equals()` for float assertions, never direct `==`.

## Constraints

- **Go stdlib only** - no external dependencies
- **Exit codes**: `0`=success, `1`=error, `2`=invalid args
- **Output**: Errors → `stderr`, Results → `stdout`
- **Go version**: 1.21

## File Locations

| Purpose | Path |
|---------|------|
| Entry point | `cmd/calculator/main.go` |
| Orchestrator | `internal/calculator/calculator.go` |
| Operations | `internal/operations/*.go` |
| Operation interface | `internal/operations/interface.go` |
| Registry | `internal/operations/registry.go` |
| CLI parser | `pkg/cli/parser.go` |
| Float utilities | `pkg/floatutil/compare.go` |
| Build output | `./build/calculator` |
