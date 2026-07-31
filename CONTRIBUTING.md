# Contributing to Pranor

Welcome! We are excited that you want to contribute to the Pranor ecosystem. This document outlines the guidelines, code style standards, and processes for contributing to all repositories within the Pranor ecosystem.

---

## Code Style & Standards

### Go Development
- **Formatting**: Always format your Go code using `gofmt` or `goimports` before committing.
- **Linting**: Ensure your code compiles cleanly and passes all standard Go linters (`golangci-lint`).
- **Comment Integrity**: Preserve all existing comments and documentation that are unrelated to your edits. Document public functions, types, and packages following idiomatic Go standards.

### API Consistency Guidelines
All REST endpoints inside Pranor services must adhere to the following API rules:
- **Prefix**: Every endpoint must be versioned and start with the `/api/v1/` prefix.
- **Standardized Error Format**: JSON error responses must include both `"error"` (internal code) and `"message"` (human-readable explanation) fields:
  ```json
  {
    "error": "resource_not_found",
    "message": "The requested service record was not found."
  }
  ```
- **Deprecation**: If an API route is deprecated, annotate the handler with a `@deprecated` tag and ensure it sets a deprecation header (`X-Deprecated: true` or `Deprecated: true`).

### Quality Gates & Test Coverage
- **Coverage**: Every repository enforces a test coverage threshold (baseline-ratcheted or 60.0% default) verified by the coverage checker on pull requests. Make sure to write unit and integration tests for new features.
- **Backward Compatibility**: Pre-merge checks will analyze specs for breaking API changes. Avoid removing fields or breaking JSON structure compatibility.
- **Performance SLA**: Merges are blocked if latency runs exceed SLA p99 benchmarks.

---

## Pull Request Process

1. **Branch Naming**:
   - Use descriptive branch prefixes: `feature/` for new functionality, `bugfix/` for fixes, `docs/` for documentation updates, and `test/` for test suite additions.
2. **Local Validation**:
   - Before opening a PR, run tests locally for your component: `go test -v ./...`
   - Run the local CI checkers inside the `pranor-repo` directory if applicable:
     - Coverage: `go run scripts/check_coverage.go`
     - API Consistency: `go run scripts/check_api_consistency.go`
3. **Submitting**:
   - Target the `main` branch.
   - Provide a clear PR summary detailing what was changed, why, and how you verified it.

---

## How to Add a Compiler Stdlib Module

The `Pranor` standard library provides built-in functions accessible across services.
To add a new built-in function or package:
1. **Define the signature**: Open the semantic analyzer in the compiler and register the new module/function name, return types, and parameter bounds.
2. **Implement the runtime**: Add the underlying Go logic execution code in the interpreter/runtime engine.
3. **Verify Domain Boundaries**: Make sure helper methods or internal namespaces are properly restricted and cannot be invoked from outside their domains.
4. **Write Tests**: Add table-driven semantic analysis and compiler execution tests verifying parameter limits.

---

## How to Write a WASM Plugin

Pranor supports running serverless functions compiled to WebAssembly (WASM).
To write a custom WASM plugin:
1. **Compilation**: Compile your Go, Rust, or C code into a WASM module targeting the WebAssembly System Interface (WASI).
2. **Deployment**: Annotate your service definition header comments with runtime details:
   ```go
   // runtime: wasm
   ```
3. **API Contracts**: WASM modules communicate via stdin/stdout stream pipelines or standard POSIX file descriptors managed by the `Pranor Deploy` container isolation layer. Ensure your entrypoints handle input arguments and output results correctly.
