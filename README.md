# Payment Service

A Go-based payment service implementing clean architecture with idempotent payment processing, plus a worker pool demonstration.

## Features

- RESTful API using go-chi router
- Idempotent payment processing (prevents duplicate charges)
- In-memory transaction storage
- Clean architecture pattern
- Comprehensive unit tests
- **Worker Pool Demo**: Demonstrates concurrent task processing with limited workers

## Project Structure

This project follows clean architecture principles with the following organized structure:

### Core Application
- `cmd/server/` - Payment service application entry point - Part 1
- `cmd/worker/` - Worker pool demonstration program - Part 2
- `internal/` - Internal application code (entities, use cases, repositories, handlers)
- `docs/` - Generated API documentation
- `api/` - OpenAPI specifications

### Scripts & Tools
- `scripts/build/` - Build scripts for different platforms
- `scripts/setup/` - Environment setup scripts  
- `scripts/docker/` - Docker configuration files
- `Makefile` - Make commands (Linux/macOS)
- `build.*` - Convenience scripts that call organized scripts

### Configuration
- `Dockerfile` - Production container image
- `.air.toml` - Hot reload configuration
- `.github/workflows/` - CI/CD pipeline configuration

### Convenience Scripts
For ease of use, convenience scripts are provided in the root directory:
- `build.ps1` → calls `scripts/build/build.ps1`
- `build.bat` → calls `scripts/build/build.bat`  
- `build.sh` → calls `scripts/build/build.sh`

```
payment-service/
├── cmd/
│   ├── server/
│   │   └── main.go                 # Payment service entry point
│   └── worker/
│       └── main.go                 # Worker pool demo entry point
├── internal/
│   ├── entity/
│   │   └── payment.go              # Business entities
│   ├── usecase/
│   │   ├── interfaces.go           # Use case interfaces
│   │   ├── payment.go              # Payment business logic
│   │   └── payment_test.go         # Unit tests
│   ├── repository/
│   │   └── payment.go              # Data storage layer
│   └── handler/
│       ├── payment.go              # HTTP handlers
│       └── payment_test.go         # Handler tests
├── scripts/
│   ├── build/
│   │   ├── build.ps1               # PowerShell build script
│   │   ├── build.bat               # Windows batch script
│   │   └── build.sh                # Linux/macOS build script
│   ├── setup/
│   │   ├── setup.ps1               # Windows setup script
│   │   └── setup.sh                # Linux/macOS setup script
│   └── docker/
│       ├── docker-compose.yml      # Production Docker setup
│       ├── docker-compose.dev.yml  # Development Docker setup
│       └── Dockerfile.dev          # Development Dockerfile
├── .github/
│   └── workflows/
│       └── ci-cd.yml               # GitHub Actions pipeline
├── docs/                           # Generated API documentation (auto-generated)
├── Dockerfile                      # Production Dockerfile
├── Makefile                        # Make commands (Linux/macOS)
├── go.mod                          # Go module definition
└── README.md                       # This file
```

## API Endpoints

### POST /pay
Processes a payment request with idempotency support.

**Request Body:**
```json
{
  "user_id": "user123",
  "amount": 100.50,
  "transaction_id": "txn_unique_id"
}
```

**Response:**
```json
{
  "transaction_id": "txn_unique_id",
  "user_id": "user123",
  "amount": 100.50,
  "status": "completed",
  "message": "Payment processed successfully"
}
```

### GET /health
Health check endpoint.

**Response:**
```json
{
  "status": "ok",
  "service": "payment-service"
}
```

## Worker Pool Demo

This project includes a worker pool demonstration program that showcases concurrent task processing in Go.

### Features
- **Concurrent Processing**: 100 tasks processed concurrently
- **Limited Workers**: Only 5 workers run simultaneously 
- **Ordered Results**: All results are collected and displayed in original order
- **Task Simulation**: Each task squares a number with simulated processing time

### Running the Worker Pool Demo

```bash
# Using Make (Linux/macOS)
make run-worker

# Using PowerShell script (Windows)
.\scripts\build\build.ps1 run-worker

# Using batch script (Windows)
scripts\build\build.bat run-worker

# Using shell script (Linux/macOS)
./scripts/build/build.sh run-worker

# Using direct Go command (all platforms)
go run cmd/worker/main.go
```

### Building the Worker Pool Demo

```bash
# Using Make (Linux/macOS)
make build-worker

# Using PowerShell script (Windows)
.\scripts\build\build.ps1 build-worker

# Using batch script (Windows)
scripts\build\build.bat build-worker

# Using shell script (Linux/macOS)
./scripts/build/build.sh build-worker

# Using direct Go command (all platforms)
go build -o bin/worker-pool cmd/worker/main.go
```

### What the Demo Demonstrates
1. **Worker Pool Pattern**: Creates a fixed number of worker goroutines
2. **Channel Communication**: Uses channels to distribute tasks and collect results
3. **Synchronization**: Uses sync.WaitGroup to coordinate worker completion
4. **Ordered Output**: Maintains task order despite concurrent processing
5. **Resource Management**: Limits concurrent operations to prevent resource exhaustion

## Getting Started

### Prerequisites
- Go 1.21 or higher

### Installation

1. Install dependencies:
```bash
go mod tidy
```

2. Generate API documentation:
```bash
# Install swag CLI tool (first time only)
go install github.com/swaggo/swag/cmd/swag@latest

# Generate swagger docs
swag init -g cmd/server/main.go -o docs
```

3. Run the service:
```bash
go run cmd/server/main.go
```

The service will start on `http://localhost:8080`.

### Quick Setup

**Automated Setup Scripts:**

*Linux/macOS:*
```bash
chmod +x scripts/setup/setup.sh
./scripts/setup/setup.sh
```

*Windows PowerShell:*
```powershell
.\scripts\setup\setup.ps1
```

**Manual Setup:**

1. Install dependencies:
```bash
go mod tidy
```

2. Generate API documentation:
```bash
# Install swag CLI tool (first time only)
go install github.com/swaggo/swag/cmd/swag@latest

# Generate swagger docs
swag init -g cmd/server/main.go -o docs
```

3. Run the service:
```bash
# Using Make (Linux/macOS)
make run

# Using PowerShell script (Windows)
.\scripts\build\build.ps1 run

# Using direct Go command (all platforms)
go run cmd/server/main.go
```

The service will start on `http://localhost:8080`.

## Docker

### Quick Start with Docker

1. **Build and run with Docker:**
```bash
# Build the Docker image
docker build -t payment-service .

# Run the container
docker run -p 8080:8080 payment-service
```

2. **Using Docker Compose (recommended):**
```bash
# Production environment
docker-compose -f scripts/docker/docker-compose.yml up -d

# Development environment with hot reloading
docker-compose -f scripts/docker/docker-compose.dev.yml up
```

### Docker Commands

```bash
# Build image
make docker-build

# Run container
make docker-run

# Development with hot reloading
make docker-dev

# Stop containers
make docker-down

# Clean up
make docker-clean
```

**Or using scripts:**
```bash
# Windows PowerShell
.\scripts\build\build.ps1 docker-dev

# Linux/macOS
./scripts/build/build.sh docker-dev
```

### Multi-platform Docker Images

The CI/CD pipeline automatically builds Docker images for:
- `linux/amd64`
- `linux/arm64`

Images are published to GitHub Container Registry at `ghcr.io/your-username/payment-service`.

## CI/CD Pipeline

This project includes a simple GitHub Actions CI/CD pipeline for demonstration:

### Pipeline Flow
1. **Create PR** → Pipeline runs tests and builds
2. **Merge to main** → Pipeline deploys to "production"

### Automated Testing
- **Code Quality**: `go vet`, `go fmt` checks
- **Unit Tests**: Full test suite with race condition detection
- **Build Verification**: Ensures code compiles successfully

### Build & Deploy
- **Docker Images**: Automatic container builds
- **Deployment Simulation**: Demonstrates production deployment flow

### Pipeline Triggers
- **Push**: `main` and `develop` branches
- **Pull Requests**: All PRs to `main`

### Demo Workflow
1. Make changes to the code
2. Create a pull request
3. Watch the CI pipeline run tests and build
4. Merge to main branch
5. Watch the deployment pipeline trigger

### Workflow File
- `.github/workflows/ci-cd.yml` - Main CI/CD pipeline

## API Documentation

This project uses [swaggo/swag](https://github.com/swaggo/swag) to automatically generate OpenAPI/Swagger documentation from Go annotations.

### Accessing the Documentation

Once the service is running, you can access the interactive API documentation at:

- **Swagger UI**: [http://localhost:8080/swagger/](http://localhost:8080/swagger/)
- **JSON Specification**: [http://localhost:8080/swagger/doc.json](http://localhost:8080/swagger/doc.json)
- **YAML Specification**: Available in `docs/swagger.yaml`

### Generating Documentation

To regenerate the API documentation after making changes to the API annotations:

```bash
# Using swag directly
swag init -g cmd/server/main.go -o docs

# Or using make (if available)
make docs
```

### Available API Endpoints

1. **POST /pay** - Process a payment
   - Supports idempotent payment processing
   - Requires: `user_id`, `amount`, `transaction_id`

2. **GET /health** - Health check endpoint
   - Returns service status

3. **GET /swagger/** - Interactive API documentation

### API Documentation Features

- **Interactive Testing**: Test endpoints directly from the Swagger UI
- **Request/Response Examples**: See example payloads and responses
- **Schema Validation**: View required fields and data types
- **Error Codes**: Understand possible error responses

The service will start on port 8080.

### Running Tests

Run all tests:
```bash
go test ./...
```

Run tests with coverage:
```bash
go test -cover ./...
```

### Example Usage

Process a payment:
```bash
curl -X POST http://localhost:8080/pay \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user123",
    "amount": 100.50,
    "transaction_id": "txn_001"
  }'
```

Retry the same payment (idempotent):
```bash
curl -X POST http://localhost:8080/pay \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user123",
    "amount": 100.50,
    "transaction_id": "txn_001"
  }'
```

Check service health:
```bash
curl http://localhost:8080/health
```

## Idempotency

The service ensures idempotency by checking the `transaction_id`. If a payment with the same `transaction_id` is submitted multiple times, the service will:

1. Return the existing payment result
2. Not charge the user multiple times
3. Maintain the same response format

## Error Handling

The service validates incoming requests and returns appropriate HTTP status codes:

- `400 Bad Request`: Invalid request data (empty user_id, invalid amount, etc.)
- `500 Internal Server Error`: Server-side errors
- `200 OK`: Successful payment processing

### Development Workflow

When working on this project:

1. **Make API Changes**: Update handlers, use cases, or entities as needed
2. **Update Swagger Annotations**: Add or modify swagger comments in your code
3. **Regenerate Documentation**: Run `swag init -g cmd/server/main.go -o docs`
4. **Test Changes**: Use the Swagger UI to test your endpoints interactively
5. **Run Tests**: Execute `go test ./...` to ensure everything works

### Available Commands

**Using Make (Linux/macOS - Recommended):**
```bash
make help           # Show available commands
make build          # Build the payment service (default)
make build-worker   # Build the worker pool demo
make run            # Run the payment service (default)
make run-worker     # Run the worker pool demo
make test           # Run tests
make test-coverage  # Run tests with coverage report
make docs           # Generate Swagger documentation
make docs-serve     # Generate docs and start server
make docker-build   # Build Docker image
make docker-dev     # Run development environment
```

**Using Scripts (Cross-platform):**

*Windows PowerShell:*
```powershell
.\scripts\build\build.ps1 help          # Show available commands
.\scripts\build\build.ps1 run           # Run the payment service (default)
.\scripts\build\build.ps1 run-worker    # Run the worker pool demo
.\scripts\build\build.ps1 build         # Build the payment service (default)
.\scripts\build\build.ps1 build-worker  # Build the worker pool demo
.\scripts\build\build.ps1 test          # Run tests
.\scripts\build\build.ps1 docs          # Generate documentation

# Or use convenience script:
.\build.ps1 run                          # Calls scripts\build\build.ps1
```

*Windows Command Prompt:*
```cmd
scripts\build\build.bat help             # Show available commands
scripts\build\build.bat run              # Run the payment service (default)
scripts\build\build.bat run-worker       # Run the worker pool demo
scripts\build\build.bat build            # Build the payment service (default)
scripts\build\build.bat build-worker     # Build the worker pool demo

# Or use convenience script:
build.bat run                            # Calls scripts\build\build.bat
```

*Linux/macOS:*
```bash
./scripts/build/build.sh help            # Show available commands
./scripts/build/build.sh run             # Run the payment service (default)
./scripts/build/build.sh run-worker      # Run the worker pool demo
./scripts/build/build.sh build           # Build the payment service (default)
./scripts/build/build.sh build-worker    # Build the worker pool demo
./scripts/build/build.sh test            # Run tests

# Or use convenience script:
./build.sh run                           # Calls scripts/build/build.sh
```

**Direct Go Commands (Universal):**
```bash
# Run the payment service (default)
go run cmd/server/main.go

# Run the worker pool demo
go run cmd/worker/main.go

# Run tests
go test ./...

# Build payment service binary
go build -o bin/payment-service cmd/server/main.go

# Build worker pool demo binary
go build -o bin/worker-pool cmd/worker/main.go

# Generate documentation (requires swag)
swag init -g cmd/server/main.go -o docs
```

## Development

### Project Structure
- **Entity Layer**: Defines core business objects
- **Use Case Layer**: Contains business logic and rules
- **Repository Layer**: Handles data persistence
- **Handler Layer**: Manages HTTP requests and responses

### Adding Features
1. Add new entities in `internal/entity/`
2. Define interfaces in `internal/usecase/interfaces.go`
3. Implement business logic in `internal/usecase/`
4. Add data storage logic in `internal/repository/`
5. Create HTTP handlers in `internal/handler/`
6. Add swagger annotations for new endpoints
7. Regenerate documentation with `make docs`
8. Write comprehensive tests

## License

This project is licensed under the MIT License.
