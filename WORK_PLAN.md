# Go-Calculator Improvement Initiative - Work Plan

**Version:** 1.0  
**Last Updated:** 2024-04-07  
**Estimated Duration:** 4 weeks  
**Total Effort:** ~80 hours

---

## 1. Executive Summary

### Purpose
This work plan outlines the systematic improvement of the go-calculator codebase to address 28 identified issues across code quality, maintainability, robustness, and testing coverage.

### Goals
- Eliminate floating-point precision bugs and edge cases
- Simplify operation registration via auto-registration pattern
- Add comprehensive overflow/underflow detection
- Improve testability with I/O abstraction
- Achieve 90%+ test coverage with benchmarks and fuzzing

### Scope
| In Scope | Out of Scope |
|----------|--------------|
| Float comparison fixes | New operations |
| Auto-registration refactor | CLI UI redesign |
| Overflow/underflow detection | External API integration |
| Comprehensive testing | Documentation site |
| Error message improvements | Performance optimization beyond benchmarks |

---

## 2. Phase Details

### Phase 1: Critical Fixes (Week 1)
**Theme:** Safety and Correctness

#### Objectives
- Fix floating-point comparison bugs
- Add NaN/Inf validation
- Handle division by zero edge cases
- Implement input validation across all operations

#### Task Breakdown

| ID | Task | File Path | Effort | Dependencies |
|----|------|-----------|--------|--------------|
| 1.1 | Create `pkg/utils/float.go` with tolerance comparison utilities | `pkg/utils/float.go` | 2h | None |
| 1.2 | Add `Equals`, `IsZero`, `IsPositive`, `IsNegative` functions | `pkg/utils/float.go` | 2h | 1.1 |
| 1.3 | Refactor divide operation with epsilon-based zero check | `internal/operations/divide.go` | 2h | 1.2 |
| 1.4 | Add NaN validation to parser | `pkg/cli/parser.go` | 1h | None |
| 1.5 | Add Inf validation to parser | `pkg/cli/parser.go` | 1h | 1.4 |
| 1.6 | Add input validation to Add operation | `internal/operations/add.go` | 1h | None |
| 1.7 | Add input validation to Subtract operation | `internal/operations/subtract.go` | 1h | None |
| 1.8 | Add input validation to Multiply operation | `internal/operations/multiply.go` | 1h | None |
| 1.9 | Write comprehensive unit tests for float utilities | `pkg/utils/float_test.go` | 3h | 1.2 |
| 1.10 | Update existing tests with new validation cases | Various test files | 2h | 1.1-1.8 |

#### Deliverables
- [ ] `pkg/utils/float.go` with floating-point utilities
- [ ] `pkg/utils/float_test.go` with 100% coverage
- [ ] Updated operations with input validation
- [ ] Parser with NaN/Inf detection
- [ ] Updated test suite

#### Acceptance Criteria
```go
// Float comparison works correctly
utils.Equals(0.1+0.2, 0.3) // returns true
utils.IsZero(1e-15)        // returns true with epsilon 1e-9

// Parser rejects invalid inputs
./calculator add NaN 5     // Error: invalid number: NaN
./calculator add Inf 5     // Error: invalid number: Inf

// Division by zero with epsilon
./calculator divide 5 1e-15 // Error: division by zero (within epsilon)
```

---

### Phase 2: Maintainability (Week 2)
**Theme:** Simplicity and Clean Architecture

#### Objectives
- Implement auto-registration via `init()` functions
- Create global DefaultRegistry
- Remove hardcoded operation list from calculator
- Add registry sorting
- Add optional structured logging

#### Task Breakdown

| ID | Task | File Path | Effort | Dependencies |
|----|------|-----------|--------|--------------|
| 2.1 | Add `init()` function to Add operation for auto-registration | `internal/operations/add.go` | 1h | None |
| 2.2 | Add `init()` function to Subtract operation | `internal/operations/subtract.go` | 1h | None |
| 2.3 | Add `init()` function to Multiply operation | `internal/operations/multiply.go` | 1h | None |
| 2.4 | Add `init()` function to Divide operation | `internal/operations/divide.go` | 1h | None |
| 2.5 | Create global DefaultRegistry variable | `internal/calculator/registry.go` | 2h | None |
| 2.6 | Add `init()` safe initialization with sync.Once | `internal/calculator/registry.go` | 2h | 2.5 |
| 2.7 | Simplify Calculator to use DefaultRegistry | `internal/calculator/calculator.go` | 3h | 2.5-2.6 |
| 2.8 | Implement registry sorting by symbol | `internal/calculator/registry.go` | 2h | None |
| 2.9 | Add ListOperationsSorted() method | `internal/calculator/registry.go` | 1h | 2.8 |
| 2.10 | Create logger interface and implementation | `pkg/logger/logger.go` | 3h | None |
| 2.11 | Integrate logging into calculator | `internal/calculator/calculator.go` | 2h | 2.10 |
| 2.12 | Add tests for auto-registration behavior | `internal/calculator/registry_test.go` | 3h | 2.1-2.6 |
| 2.13 | Add tests for logging integration | `pkg/logger/logger_test.go` | 2h | 2.10 |

#### Deliverables
- [ ] Auto-registration in all operation files
- [ ] Global DefaultRegistry with thread-safe init
- [ ] Simplified calculator constructor
- [ ] Sorted registry operations list
- [ ] Structured logging package
- [ ] Updated test suite

#### Acceptance Criteria
```go
// Auto-registration works
import _ "github.com/kairostack/go-calculator/internal/operations"
// Operations are automatically registered

// DefaultRegistry is globally accessible
calc := calculator.New() // Uses DefaultRegistry automatically

// Operations are sorted
ops := registry.ListOperationsSorted()
// Returns: [add, divide, multiply, subtract]

// Logging integration
logger := logger.New(logger.Config{Level: logger.Debug})
calc := calculator.New(calculator.WithLogger(logger))
```

---

### Phase 3: Robustness (Week 3)
**Theme:** Error Handling and Testability

#### Objectives
- Add overflow/underflow detection
- Improve error messages with context
- Abstract CLI I/O for testability
- Create comprehensive integration tests

#### Task Breakdown

| ID | Task | File Path | Effort | Dependencies |
|----|------|-----------|--------|--------------|
| 3.1 | Create `pkg/utils/overflow.go` with detection logic | `pkg/utils/overflow.go` | 3h | None |
| 3.2 | Implement overflow detection for Add | `internal/operations/add.go` | 2h | 3.1 |
| 3.3 | Implement overflow detection for Subtract | `internal/operations/subtract.go` | 2h | 3.1 |
| 3.4 | Implement overflow detection for Multiply | `internal/operations/multiply.go` | 2h | 3.1 |
| 3.5 | Implement underflow detection | `pkg/utils/overflow.go` | 2h | 3.1 |
| 3.6 | Create custom error types with context | `internal/calculator/errors.go` | 3h | None |
| 3.7 | Update error messages across all operations | Various operation files | 3h | 3.6 |
| 3.8 | Create I/O abstraction interfaces | `pkg/cli/io.go` | 2h | None |
| 3.9 | Implement stdin/stdout I/O provider | `pkg/cli/io.go` | 2h | 3.8 |
| 3.10 | Implement test I/O provider (buffers) | `pkg/cli/io_test.go` | 2h | 3.8 |
| 3.11 | Refactor main.go to use I/O abstraction | `cmd/calculator/main.go` | 3h | 3.8-3.10 |
| 3.12 | Create testable CLI runner | `pkg/cli/runner.go` | 4h | 3.8-3.11 |
| 3.13 | Write integration tests | `tests/integration/calculator_test.go` | 4h | 3.12 |
| 3.14 | Add overflow/underflow unit tests | `pkg/utils/overflow_test.go` | 3h | 3.1-3.5 |

#### Deliverables
- [ ] Overflow/underflow detection utilities
- [ ] Custom error types with context
- [ ] I/O abstraction layer
- [ ] Testable CLI runner
- [ ] Integration test suite
- [ ] Updated unit tests

#### Acceptance Criteria
```go
// Overflow detection
./calculator multiply 1e308 10
// Error: overflow: result exceeds maximum float64 value

// Underflow detection
./calculator multiply 1e-308 1e-308
// Error: underflow: result is too small to represent

// Improved error messages
./calculator divide 5 0
// Error: division by zero: cannot divide 5 by 0

// Testable runner
runner := cli.NewRunner(testIO)
runner.Run([]string{"add", "5", "3"})
// Output captured in testIO.Out buffer
```

---

### Phase 4: Quality Expansion (Week 4)
**Theme:** Performance and Coverage

#### Objectives
- Add benchmark tests for all operations
- Implement fuzzing tests
- Create example functions for documentation
- Add edge case tests

#### Task Breakdown

| ID | Task | File Path | Effort | Dependencies |
|----|------|-----------|--------|--------------|
| 4.1 | Create benchmark tests for Add | `internal/operations/add_test.go` | 2h | None |
| 4.2 | Create benchmark tests for Subtract | `internal/operations/subtract_test.go` | 2h | None |
| 4.3 | Create benchmark tests for Multiply | `internal/operations/multiply_test.go` | 2h | None |
| 4.4 | Create benchmark tests for Divide | `internal/operations/divide_test.go` | 2h | None |
| 4.5 | Create benchmark for registry operations | `internal/calculator/registry_test.go` | 2h | None |
| 4.6 | Create benchmark for parser | `pkg/cli/parser_test.go` | 2h | None |
| 4.7 | Implement fuzz tests for Add | `internal/operations/add_test.go` | 3h | None |
| 4.8 | Implement fuzz tests for Subtract | `internal/operations/subtract_test.go` | 3h | None |
| 4.9 | Implement fuzz tests for Multiply | `internal/operations/multiply_test.go` | 3h | None |
| 4.10 | Implement fuzz tests for Divide | `internal/operations/divide_test.go` | 3h | None |
| 4.11 | Add example functions for all operations | `internal/operations/example_test.go` | 3h | None |
| 4.12 | Add edge case test suite (boundary values) | `tests/edge/edge_cases_test.go` | 4h | None |
| 4.13 | Add stress tests for concurrent registry access | `internal/calculator/registry_test.go` | 3h | None |
| 4.14 | Generate coverage report and documentation | `coverage.html` | 2h | All above |

#### Deliverables
- [ ] Benchmark tests for all operations
- [ ] Fuzzing tests for all operations
- [ ] Example functions for documentation
- [ ] Edge case test suite
- [ ] Stress tests for concurrency
- [ ] Coverage report (target: 90%+)

#### Acceptance Criteria
```bash
# Benchmarks run successfully
$ go test -bench=. ./internal/operations/...
BenchmarkAdd-8           1000000000    0.312 ns/op
BenchmarkSubtract-8      1000000000    0.315 ns/op
BenchmarkMultiply-8      1000000000    0.318 ns/op
BenchmarkDivide-8        1000000000    0.320 ns/op

# Fuzzing runs for 60s without issues
$ go test -fuzz=FuzzAdd -fuzztime=60s ./internal/operations/...

# Examples compile and run
$ go test -run Example ./internal/operations/...

# Coverage meets target
$ make coverage
ok      github.com/kairostack/go-calculator/internal/operations    0.456s  coverage: 94.2% of statements
```

---

## 3. Risk Assessment

### High-Risk Changes

#### Float Comparison Changes (Phase 1)
| Aspect | Assessment |
|--------|------------|
| **Risk Level** | Medium-High |
| **Impact** | Could change behavior of existing calculations |
| **Mitigation** | <ul><li>Define epsilon as configurable constant</li><li>Document tolerance behavior</li><li>Add comprehensive regression tests</li><li>Test with existing calculation examples</li></ul> |
| **Rollback Plan** | Revert to exact comparison if tolerance causes issues |

#### Auto-Registration with init() (Phase 2)
| Aspect | Assessment |
|--------|------------|
| **Risk Level** | Medium |
| **Impact** | init() order can be unpredictable; testing requires side-effect management |
| **Mitigation** | <ul><li>Use sync.Once for thread-safe initialization</li><li>Provide manual registration option as fallback</li><li>Document import requirements</li><li>Add integration tests covering import scenarios</li></ul> |
| **Rollback Plan** | Keep explicit registration as alternative in calculator constructor |

#### Overflow Detection (Phase 3)
| Aspect | Assessment |
|--------|------------|
| **Risk Level** | Medium |
| **Impact** | May reject valid calculations that previously worked (edge cases) |
| **Mitigation** | <ul><li>Test against IEEE 754 standard cases</li><li>Make overflow check configurable</li><li>Test with extreme but valid values</li><li>Document overflow/underflow thresholds</li></ul> |
| **Rollback Plan** | Disable overflow checks via configuration flag |

#### NaN/Inf Validation (Phase 1)
| Aspect | Assessment |
|--------|------------|
| **Risk Level** | Low-Medium |
| **Impact** | Rejects inputs that might be used intentionally in edge cases |
| **Mitigation** | <ul><li>Ensure error messages are clear</li><li>Document that NaN/Inf are not supported</li><li>Add validation at parser level (early fail)</li></ul> |
| **Rollback Plan** | Remove validation if business requirement changes |

---

## 4. Testing Strategy

### Test Pyramid

```
       /\
      /  \     Fuzzing (Phase 4)
     /    \    ~5% of tests
    /------\
   /        \   Integration (Phase 3)
  /          \  ~15% of tests
 /------------\
/              \ Unit Tests (All Phases)
/                \ ~80% of tests
------------------
```

### Phase-Specific Testing

#### Phase 1 Testing
- [ ] Unit tests for all float utility functions
- [ ] Table-driven tests for boundary values
- [ ] Parser validation tests (valid/invalid inputs)
- [ ] Regression tests for existing calculations

#### Phase 2 Testing
- [ ] Test auto-registration with blank imports
- [ ] Test manual registration (fallback)
- [ ] Concurrent access tests for registry
- [ ] Test operation sorting
- [ ] Logging output verification

#### Phase 3 Testing
- [ ] Overflow boundary tests (near MaxFloat64)
- [ ] Underflow boundary tests (near SmallestNonzeroFloat64)
- [ ] I/O abstraction tests with mock implementations
- [ ] Integration tests with real file descriptors
- [ ] Error message format verification

#### Phase 4 Testing
- [ ] Benchmark all operations (measure, not verify)
- [ ] Fuzz all operations for 60+ seconds
- [ ] Edge case tests for boundary values
- [ ] Stress tests for concurrent registry access
- [ ] Example tests for documentation

### Test Execution Commands

```bash
# Phase 1: Float utilities and validation
go test ./pkg/utils/... -v
go test ./pkg/cli/... -v

# Phase 2: Auto-registration and registry
go test ./internal/calculator/... -v -race
go test ./pkg/logger/... -v

# Phase 3: Overflow and integration
go test ./pkg/utils/... -v
go test ./tests/integration/... -v

# Phase 4: Performance and coverage
go test -bench=. ./...
go test -fuzz=FuzzAdd -fuzztime=60s ./internal/operations/...
make coverage
```

---

## 5. Timeline

### Gantt Chart Representation

```
Week 1: [Phase 1: Critical Fixes]
Mon     Tue     Wed     Thu     Fri
[1.1-1.3][1.4-1.8][1.9   ][1.10  ][Buffer ]
Float   Parser  Tests   Tests   Review

Week 2: [Phase 2: Maintainability]
Mon     Tue     Wed     Thu     Fri
[2.1-2.4][2.5-2.7][2.8-2.11][2.12-2.13][Buffer ]
Auto-reg Registry  Logging  Tests   Review

Week 3: [Phase 3: Robustness]
Mon     Tue     Wed     Thu     Fri
[3.1-3.5][3.6-3.9][3.10-3.12][3.13-3.14][Buffer ]
Overflow Errors   I/O      Integ   Review

Week 4: [Phase 4: Quality Expansion]
Mon     Tue     Wed     Thu     Fri
[4.1-4.6][4.7-4.10][4.11-4.13][4.14   ][Buffer ]
Bench   Fuzz     Edge    Cover   Final
```

### Detailed Schedule

| Week | Day | Phase | Task(s) | Owner | Hours |
|------|-----|-------|---------|-------|-------|
| 1 | Mon | 1 | Float utilities creation | TBD | 4 |
| 1 | Tue | 1 | Parser NaN/Inf validation | TBD | 3 |
| 1 | Wed | 1 | Operation input validation | TBD | 3 |
| 1 | Thu | 1 | Float utility tests | TBD | 3 |
| 1 | Fri | 1 | Test updates & buffer | TBD | 3 |
| 2 | Mon | 2 | Auto-registration init() | TBD | 4 |
| 2 | Tue | 2 | DefaultRegistry & sync.Once | TBD | 4 |
| 2 | Wed | 2 | Calculator simplification | TBD | 4 |
| 2 | Thu | 2 | Logging & registry sorting | TBD | 4 |
| 2 | Fri | 2 | Auto-registration tests | TBD | 4 |
| 3 | Mon | 3 | Overflow/underflow detection | TBD | 4 |
| 3 | Tue | 3 | Custom error types | TBD | 4 |
| 3 | Wed | 3 | I/O abstraction layer | TBD | 4 |
| 3 | Thu | 3 | Integration tests | TBD | 4 |
| 3 | Fri | 3 | Buffer & overflow tests | TBD | 4 |
| 4 | Mon | 4 | Benchmark tests | TBD | 4 |
| 4 | Tue | 4 | Fuzzing tests | TBD | 4 |
| 4 | Wed | 4 | Example functions & edge cases | TBD | 4 |
| 4 | Thu | 4 | Stress tests & coverage | TBD | 4 |
| 4 | Fri | 4 | Final review & documentation | TBD | 4 |

---

## 6. Success Metrics

### Completion Criteria

| Phase | Metric | Target | Measurement |
|-------|--------|--------|-------------|
| 1 | Float utility coverage | 100% | `go test -cover` |
| 1 | Parser validation coverage | 100% | `go test -cover` |
| 1 | NaN/Inf rejection rate | 100% | Test cases |
| 2 | Auto-registration tests | Pass | `go test -race` |
| 2 | Registry sort order | Verified | Unit test output |
| 2 | Logging integration | Verified | Log output inspection |
| 3 | Overflow detection accuracy | 100% | Boundary tests |
| 3 | Integration test coverage | >80% | `go test -cover` |
| 3 | Error message clarity | Verified | Manual inspection |
| 4 | Benchmark baseline | Established | `go test -bench` |
| 4 | Fuzzing run time | 60s/issue-free | `go test -fuzz` |
| 4 | Overall coverage | >90% | `make coverage` |
| 4 | Example tests | All pass | `go test -run Example` |

### Quality Gates

Each phase must pass these gates before proceeding:

1. **Phase 1 Gate**
   - [ ] All float utilities tested
   - [ ] NaN/Inf correctly rejected
   - [ ] Division by epsilon handled
   - [ ] Code review completed

2. **Phase 2 Gate**
   - [ ] Auto-registration works with blank import
   - [ ] Manual registration still available
   - [ ] No race conditions detected
   - [ ] Operations sorted alphabetically

3. **Phase 3 Gate**
   - [ ] Overflow/underflow detected correctly
   - [ ] Error messages include context
   - [ ] I/O abstraction tested
   - [ ] Integration tests pass

4. **Phase 4 Gate**
   - [ ] Benchmarks establish baseline
   - [ ] Fuzzing runs 60s without crash
   - [ ] Edge cases documented and tested
   - [ ] Coverage report shows >90%

### Final Deliverables Checklist

- [ ] All 28 tasks completed
- [ ] All tests passing (unit, integration, fuzz)
- [ ] Coverage report >90%
- [ ] Benchmark results documented
- [ ] No race conditions (verified with `-race`)
- [ ] Code review completed
- [ ] Documentation updated
- [ ] CHANGELOG.md updated

---

## 7. Appendix

### A. Reference Commands

```bash
# Development workflow
make init          # Setup dependencies
make test          # Run all tests
make test-race     # Run with race detector
make coverage      # Generate coverage report
make lint          # Run linter
make build         # Build calculator

# Specific test commands
go test -v ./pkg/utils/...                    # Float utilities
go test -v ./internal/calculator/...          # Calculator & registry
go test -v ./internal/operations/...          # Operations
go test -v ./pkg/cli/...                      # CLI & parser
go test -v ./tests/integration/...            # Integration tests
go test -v ./tests/edge/...                   # Edge cases

# Benchmarking
go test -bench=. ./internal/operations/...
go test -bench=BenchmarkAdd -benchmem ./internal/operations/...

# Fuzzing
go test -fuzz=FuzzAdd -fuzztime=60s ./internal/operations/...
go test -fuzz=FuzzMultiply -parallel=4 ./internal/operations/...
```

### B. File Structure (Expected After Completion)

```
cmd/calculator/
  main.go                       # Updated with I/O abstraction

internal/
  calculator/
    calculator.go               # Simplified with DefaultRegistry
    registry.go                 # Added sorting, global instance
    registry_test.go            # Added concurrency tests
    errors.go                   # NEW: Custom error types
  operations/
    interface.go                # Unchanged
    add.go                      # Updated: init(), validation, overflow
    add_test.go                 # Updated: benchmarks, fuzz, examples
    subtract.go                 # Updated: init(), validation, overflow
    subtract_test.go            # Updated: benchmarks, fuzz, examples
    multiply.go                 # Updated: init(), validation, overflow
    multiply_test.go            # Updated: benchmarks, fuzz, examples
    divide.go                   # Updated: init(), validation, epsilon check
    divide_test.go              # Updated: benchmarks, fuzz, examples
    example_test.go             # NEW: Example functions

pkg/
  cli/
    parser.go                   # Updated: NaN/Inf validation
    parser_test.go              # Updated: validation tests
    formatter.go                # Unchanged
    io.go                       # NEW: I/O abstraction
    io_test.go                  # NEW: I/O tests
    runner.go                   # NEW: Testable runner
  logger/
    logger.go                   # NEW: Structured logging
    logger_test.go              # NEW: Logger tests
  utils/
    float.go                    # NEW: Float utilities
    float_test.go               # NEW: Float tests
    overflow.go                 # NEW: Overflow detection
    overflow_test.go            # NEW: Overflow tests

tests/
  integration/
    calculator_test.go          # NEW: Integration tests
  edge/
    edge_cases_test.go          # NEW: Edge case tests

WORK_PLAN.md                    # This document
```

### C. Change Log Template

```markdown
## [Unreleased] - Improvement Initiative

### Added
- Float comparison utilities with configurable epsilon
- NaN and Inf validation in parser
- Overflow and underflow detection
- Auto-registration via init() functions
- Global DefaultRegistry with thread-safe initialization
- Registry sorting by operation symbol
- Structured logging support
- I/O abstraction for testability
- Comprehensive benchmark tests
- Fuzzing tests for all operations
- Example functions for documentation
- Edge case and stress tests

### Changed
- Division by zero check now uses epsilon tolerance
- Calculator constructor simplified to use DefaultRegistry
- Error messages improved with context
- All operations now validate inputs

### Fixed
- Floating-point precision issues in comparisons
- Race conditions in registry access
- Inconsistent error message formatting
```

---

## Sign-off

| Role | Name | Date | Signature |
|------|------|------|-----------|
| Technical Lead | | | |
| Engineering Manager | | | |
| QA Lead | | | |
| Product Owner | | | |

---

*This work plan is a living document. Updates should be tracked with version numbers and change descriptions.*
