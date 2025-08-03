# Somana

An OpenAPI-based Go project for Linux host management with a clean, maintainable structure using Gin, GORM, and SQLite.

## Overview

Somana is a simple host management API that allows you to register, monitor, and manage Linux hosts. It provides a RESTful API for host registration, status updates, and health monitoring.

## Project Structure

```
somana/
├── api/                    # OpenAPI specifications
│   └── openapi.yaml       # Main API specification
├── cmd/                   # Application entry points
│   └── server/            # Main server binary
├── internal/              # Private application code
│   ├── database/          # Database connection and configuration
│   ├── generated/         # Generated code from OpenAPI (not in Git)
│   └── services/          # Business logic
├── pkg/                   # Public libraries
├── docs/                  # Generated documentation
├── data/                  # SQLite database files (auto-created)
├── bin/                   # Build artifacts (auto-created)
├── Makefile              # Build automation
├── go.mod                # Go module definition
└── .gitignore           # Git ignore rules
```

## Quick Start

1. **Install dependencies:**
   ```bash
   make deps
   ```

2. **Run the server:**
   ```bash
   make run
   ```

3. **Test the API:**
   ```bash
   # Health check
   curl http://localhost:8080/health
   
   # Register a host
   curl -X POST http://localhost:8080/api/v1/hosts \
     -H "Content-Type: application/json" \
     -d '{
       "hostname": "web-server-01",
       "ip_address": "192.168.1.100",
       "os_name": "Ubuntu",
       "os_version": "22.04.3 LTS",
       "environment": "production"
     }'
   
   # List all hosts
   curl http://localhost:8080/api/v1/hosts
   ```

## Development

### Prerequisites

- Go 1.21+
- `oapi-codegen` for code generation

### Installation

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd somana
   ```

2. **Install oapi-codegen:**
   ```bash
   go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
   ```

3. **Install dependencies:**
   ```bash
   make deps
   ```

### Database Setup

The application uses SQLite with GORM for database operations. The database file will be automatically created in the `data/` directory when the application starts.

Default database location: `data/somana.db`

You can override this using the `DB_PATH` environment variable:
```bash
export DB_PATH=/path/to/your/database.db
```

### Code Generation

This project uses OpenAPI code generation to maintain consistency between the API specification and the Go codebase. The generated code is stored in `internal/generated/` and is automatically cleaned and regenerated during the build process.

To regenerate code after updating the OpenAPI spec:

```bash
make generate
```

The generated files are:
- `internal/generated/types.go` - Generated Go types from OpenAPI schemas
- `internal/generated/server.go` - Generated server interfaces and handlers

## API Endpoints

### Host Management

- `GET /api/v1/hosts` - List all hosts (with optional status/environment filters)
- `POST /api/v1/hosts` - Register a new host
- `GET /api/v1/hosts/:id` - Get a specific host by ID
- `PUT /api/v1/hosts/:id` - Update a host
- `DELETE /api/v1/hosts/:id` - Deregister a host
- `POST /api/v1/hosts/:id/heartbeat` - Update host status/heartbeat

### Health Check

- `GET /health` - Health check endpoint

## Host Model

The Host model includes the following fields:

- `id` - Unique identifier (auto-generated)
- `hostname` - Hostname of the system
- `ip_address` - IP address of the system
- `os_name` - Operating system name (e.g., "Ubuntu", "CentOS")
- `os_version` - Operating system version
- `environment` - Environment (development, staging, production, testing)
- `status` - Current status (online, offline, maintenance)
- `created_at` - Registration timestamp
- `updated_at` - Last update timestamp
- `deleted_at` - Soft delete timestamp (nullable)

## Available Commands

- `make deps` - Install dependencies
- `make generate` - Generate code from OpenAPI spec
- `make build` - Build the application (includes code generation)
- `make run` - Build and run the application
- `make test` - Run tests
- `make clean` - Clean build artifacts and generated files

## Adding New Features

1. **Update the OpenAPI specification** in `api/openapi.yaml`
2. **Implement the business logic** in `internal/services/`
3. **The generated code will automatically update** when you run `make build`

## Architecture

This project follows a clean architecture pattern:

- **OpenAPI Specification** (`api/openapi.yaml`) - Single source of truth for API definition
- **Generated Code** (`internal/generated/`) - Types and interfaces generated from OpenAPI
- **Business Logic** (`internal/services/`) - Implementation of the generated interfaces
- **Database Layer** (`internal/database/`) - Database connection and setup
- **Application Entry** (`cmd/server/`) - Main application entry point

## Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/your-feature`
3. Make your changes
4. Run tests: `make test`
5. Commit your changes: `git commit -am 'Add some feature'`
6. Push to the branch: `git push origin feature/your-feature`
7. Submit a pull request

## License

[Add your license here] 