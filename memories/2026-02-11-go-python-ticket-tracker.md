# Topic: Go Ticket Tracker Server and Python CLI Client

## Metadata
- **Date**: 2026-02-11
- **Tags**: go, python, sqlite, openapi, uv, testing, integration
- **Status**: Completed
- **Related Memories**: 2026-02-11-gemini-cli-mcporter-integration.md

## Context/Problem
The goal was to create a robust, tested, and documented ticket tracking system consisting of a Go backend and a Python CLI frontend. The Go server needed to use SQLite and expose an API based on an OpenAPI 3.0 specification. The Python client needed to be built using `uv` for dependency management, with auto-generated Pydantic models from the server's OpenAPI spec.

## Solution/Findings
### Go Server Implementation
- **Architecture**: Follows standard Go layout (`cmd/`, `internal/`).
- **Persistence**: Switched from `github.com/mattn/go-sqlite3` (CGO) to `modernc.org/sqlite` (pure Go) to avoid build hangs in limited environments.
- **Features**: 
  - CRUD operations for tickets.
  - Advanced filtering in `GET /tickets` (`exclude_area`, `exclude_status`).
  - Handling of long `resolution_message` fields (tested up to 6KB).
- **Testing**: Integration tests using `httptest` and in-memory SQLite.

### Python Client Implementation
- **Tools**: Managed by `uv`.
- **Codegen**: Used `datamodel-code-generator` to create Pydantic v2 models from `openapi.yaml`.
- **Client**: Built with `httpx` for asynchronous-capable synchronous requests.
- **CLI**: Implemented with `typer` and `rich` for a polished terminal experience (tables, panels).
- **Testing**: 
  - `pytest` with `respx` for mocking.
  - A unique README-driven test suite that extracts and validates Python code examples directly from `README.md`.
  - Edge case testing for `null` vs `""` and empty filter strings.

## Actionable Insights/Next Steps
- **Go Server**: To allow "never initialized" (NULL) states for the `resolution_message`, the field type should be changed to `*string` or `sql.NullString`.
- **Python Client**: The `uv run ticket-client` command is registered in `pyproject.toml` for easy access.
- **Testing**: Run tests with `go test ./...` in the Go directory and `uv run pytest` in the Python directory.

## References
- Go Source: `workspace/go/go-ticket-tracker/`
- Python Source: `workspace/python/ticket-client/`
- OpenAPI Spec: `workspace/go/go-ticket-tracker/openapi.yaml`
