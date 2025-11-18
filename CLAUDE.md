# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Go SDK for the Devin API v1. The SDK provides strongly typed access to sessions, secrets, knowledge, and playbooks domains through a unified client interface with automatic retries, logging, and bearer authentication.

## Development Commands

### Basic Operations
```bash
# Run all tests
make test

# Run specific test
go test -run TestName ./path/to/package -v

# Lint code (golangci-lint with govet, revive, gosec)
make lint

# Build all packages
make build
```

### API Integration Tests (requires DEVIN_API_KEY in .env)
```bash
make api-session    # Test sessions endpoints
make api-secret     # Test secrets endpoints
make api-playbook   # Test playbooks endpoints
make api-knowledge  # Test knowledge endpoints
```

### Running Examples
```bash
# Set API key and run example
DEVIN_API_KEY=xxx go run ./examples/sessions/list
DEVIN_API_KEY=xxx DEVIN_SESSION_ID=devin-123 go run ./examples/sessions/send_message
```

## Architecture

### Three-Layer Design

1. **Root Package** (`client.go`): `devin.NewClient(apiKey, ...opts)` returns a `*Client` with four service fields: `Sessions`, `Secrets`, `Knowledge`, `Playbooks`. All configuration (base URL, HTTP client, retry, logger) flows through functional options.

2. **Domain Services** (`pkg/<domain>/service.go`): Each domain (sessions, secrets, knowledge, playbooks) has a `Service` struct that wraps `httpclient.Transport` and a validator function. Domain-specific operations are in separate files (e.g., `list_sessions.go`, `create_session.go`).

3. **HTTP Foundation** (`internal/httpclient/httpclient.go`): `Executor` implements retry logic (408/425/429/5xx with exponential backoff), bearer auth, and request/response handling. Exposes `DoJSON` and `DoBytes` methods consumed by all services.

### Key Components

- **Validation**: Uses Gin's `binding.Validator` with struct tags like `binding:"required"`
- **Query Encoding**: `go-querystring` for URL query parameters
- **Testing**: `internal/testtransport.RoundTripFunc` and `internal/testutil.NewExecutor` for mocking HTTP without real servers
- **Examples**: Each endpoint has a runnable `package main` in `examples/<domain>/<endpoint>`

## Coding Conventions

### File Organization
- Domain services: `pkg/<domain>/<operation>.go` (e.g., `list_sessions.go`, `create_secret.go`)
- Tests: `<operation>_test.go` next to implementation (table-driven with `go-cmp` for diffs)
- Examples: `examples/<domain>/<operation>/main.go`

### Testing Pattern
```go
// Table-driven tests using testutil.NewExecutor
func TestServiceOperation(t *testing.T) {
    tests := []struct {
        name    string
        // test cases...
    }{
        // ...
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Use testutil.NewExecutor with RoundTripFunc
            // Compare with go-cmp
        })
    }
}
```

### Adding New Endpoints
1. Define request/response structs in `pkg/<domain>/<operation>.go`
2. Implement service method on `*Service`
3. Add table-driven test in `<operation>_test.go`
4. Create runnable example in `examples/<domain>/<operation>/main.go`
5. Update README.md with usage example

## Import Path
All imports use `github.com/gassara-kys/go-devin`. Never use relative imports or alternative paths.

## Quality Gates
- Run `gofmt -w .` before committing
- Ensure `make lint` passes (golangci-lint config in `.golangci.yml`)
- Ensure `make test` passes all tests
- Examples must read `DEVIN_API_KEY` from environment, never hardcode

## Notes
- No direct API integration tests in `make test` - all tests use mocked transport
- Module uses Go 1.24 (see `.golangci.yml`)
- Retry configuration: 3 attempts, 500ms initial backoff, 2s max backoff (configurable via `WithRetry`)
- Default timeout: 30s (configurable via `WithTimeout`)
