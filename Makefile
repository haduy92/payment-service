# Payment Service Makefile

.PHONY: help build build-worker run run-worker test clean docs docker docker-dev docker-clean

# Default target
help: ## Show this help message
	@echo "Available commands:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-20s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the payment service (default)
	@mkdir -p bin
	go build -o bin/payment-service cmd/server/main.go

build-worker: ## Build the worker pool demo
	@mkdir -p bin
	go build -o bin/worker-pool cmd/worker/main.go

run: ## Run the payment service (default)
	@echo "API will be available at: http://localhost:8080"
	@echo "API documentation: http://localhost:8080/swagger/"
	go run cmd/server/main.go

run-worker: ## Run the worker pool demo
	go run cmd/worker/main.go

test: ## Run tests
	go test -v ./...

test-coverage: ## Run tests with coverage
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

clean: ## Clean build artifacts
	rm -rf bin/
	rm -rf docs/
	rm -f coverage.out coverage.html

docs: ## Generate Swagger documentation
	swag init -g cmd/server/main.go -o docs

docs-serve: docs run ## Generate docs and start the server
	@echo "API documentation available at: http://localhost:8080/swagger/"

install-swag: ## Install swag CLI tool
	go install github.com/swaggo/swag/cmd/swag@latest

install-air: ## Install air for hot reloading
	go install github.com/cosmtrek/air@latest

# Docker commands
docker-build: ## Build Docker image
	docker build -t payment-service:latest .

docker-run: ## Run Docker container
	docker run -p 8080:8080 payment-service:latest

docker-dev: ## Run development environment with hot reloading
	docker-compose -f scripts/docker/docker-compose.dev.yml up --build

docker-prod: ## Run production environment
	docker-compose -f scripts/docker/docker-compose.yml up --build -d

docker-down: ## Stop Docker containers
	docker-compose -f scripts/docker/docker-compose.yml down
	docker-compose -f scripts/docker/docker-compose.dev.yml down

docker-clean: ## Clean Docker images and containers
	docker-compose -f scripts/docker/docker-compose.yml down --rmi all --volumes --remove-orphans
	docker-compose -f scripts/docker/docker-compose.dev.yml down --rmi all --volumes --remove-orphans

# Linting and formatting
fmt: ## Format Go code
	go fmt ./...

vet: ## Run go vet
	go vet ./...

lint: fmt vet ## Run all linting tools

# CI/CD helpers
ci-test: ## Run tests as in CI
	go mod verify
	go vet ./...
	gofmt -s -l .
	go test -v -race -coverprofile=coverage.out ./...

.DEFAULT_GOAL := help
