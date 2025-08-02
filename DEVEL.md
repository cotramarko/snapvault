# Developer Guide

This document explains how to set up and work on the Snapvault project as a developer or contributor.

## Prerequisites

- Go 1.21.2 or later
- Docker and Docker Compose (for running PostgreSQL locally)
- Git

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/cotramarko/snapvault.git
cd snapvault
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. Set Up Local Database

The project includes a Docker Compose setup for local development:

```bash
docker-compose -f db/docker-compose.yml up -d
```

This will start a PostgreSQL database with:
- Database: `acmedb`
- User: `acmeuser`
- Password: `acmepassword`
- Port: `5432`

The database will be initialized with the schema from `db/schema-and-seed.sql`.

## Development Workflow

### Building the Project

```bash
# Build the main binary
go build -o snapvault main.go

# Or use go run for quick testing
go run main.go [command] [args]

# When running the development db with docker-compose it's easiest to test commands with
go run main.go [command] [args] --url=postgres://acmeuser:acmepassword@localhost:5432/acmedb
```

### Running Tests

Integration tests require the Docker database to be running:

```bash
# Start the database first
docker-compose -f db/docker-compose.yml up -d

# Run integration tests
go test -v ./...

# Shutdown the DB and remove data
docker-compose -f db/docker-compose.yml down -v
```

### Code Organization

- `cmd/` - CLI command implementations (Cobra commands)
- `internal/commands/` - Core business logic for operations
- `internal/engine/` - Database engine and connection management
- `internal/connection/` - Database connection utilities
- `tests/` - Integration tests
- `main.go` - Application entry point

## Contributing

1. Create a feature branch from `main`
2. Make your changes
3. Add or update tests as needed
4. Ensure all tests pass
5. Run `gofmt` on your code
6. Submit a pull request