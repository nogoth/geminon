# Ticket Tracker Client

Python CLI client for the Go Ticket Tracker Server.

## Prerequisites

- [uv](https://github.com/astral-sh/uv)
- Running Go Ticket Tracker Server

## Usage

### 1. Start the Go Server
```bash
cd ../../go/go-ticket-tracker
go run cmd/server/main.go
```

### 2. Run the CLI Client
```bash
cd workspace/python/ticket-client

# Create a ticket
uv run ticket-client create "Issue Title" --description "Details" --area "IT"

# List tickets (with filters)
uv run ticket-client list --exclude-status "Closed"

# Get ticket details
uv run ticket-client get 1

# Resolve a ticket
uv run ticket-client resolve 1 "Fixed the issue"
```

## Development

Models are generated from the Go server's `openapi.yaml`:
```bash
uv run datamodel-codegen --input openapi.yaml --output src/ticket_client/models.py
```
