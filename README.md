# Somana

An OpenAPI-based Go project with a clean, maintainable structure for building CRUD APIs using Gin and GORM with SQLite.

## Project Structure

```
somana/
├── api/                    # OpenAPI specifications
│   └── openapi.yaml       # Main API specification
├── cmd/                   # Application entry points
│   └── server/            # Main server binary
├── internal/              # Private application code
│   ├── database/          # Database connection and configuration
│   ├── handlers/          # HTTP request handlers
│   ├── models/            # GORM data models
│   └── services/          # Business logic
├── pkg/                   # Public libraries
├── docs/                  # Generated documentation
├── data/                  # SQLite database files (auto-created)
├── Makefile              # Build automation
├── go.mod                # Go module definition
└── .gitignore           # Git ignore rules
```

## Quick Start

1. **Install dependencies:**
   ```bash
   make deps
   ```

2. **Generate documentation:**
   ```bash
   make generate
   ```

3. **Run the server:**
   ```bash
   make run
   ```

4. **Run tests:**
   ```bash
   make test
   ```

## Development

### Prerequisites

- Go 1.21+
- `go-swagger` for code generation (optional)

### Database Setup

The application uses SQLite with GORM for database operations. The database file will be automatically created in the `data/` directory when the application starts.

Default database location: `data/somana.db`

You can override this using the `DB_PATH` environment variable:
```bash
export DB_PATH=/path/to/your/database.db
```

### Code Generation

This project uses OpenAPI code generation to maintain consistency between the API specification and the Go codebase.

To regenerate code after updating the OpenAPI spec:

```bash
make generate
```

### API Documentation

API documentation is automatically generated from the OpenAPI specification and available at:

- Swagger UI: `http://localhost:8080/swagger/`
- OpenAPI JSON: `http://localhost:8080/openapi.json`

## API Endpoints

### Resources

- `GET /api/v1/resources` - Get all resources
- `POST /api/v1/resources` - Create a new resource
- `GET /api/v1/resources/:id` - Get a resource by ID
- `PUT /api/v1/resources/:id` - Update a resource
- `DELETE /api/v1/resources/:id` - Delete a resource
- `GET /api/v1/resources/active` - Get active resources only

### Health Check

- `GET /health` - Health check endpoint

## Adding New Resources

1. **Define your resource in `api/openapi.yaml`**
2. **Create corresponding GORM model in `internal/models/`**
3. **Create service logic in `internal/services/`**
4. **Create handlers in `internal/handlers/`**
5. **Register routes in `cmd/server/main.go`**
6. **Add model to database migration in `internal/database/database.go`**

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