#!/bin/bash

# Payment Service - Build Script
# Linux/macOS equivalent of Makefile commands

show_help() {
    echo "Payment Service - Available Commands:"
    echo ""
    echo "Basic commands:"
    echo "  ./scripts/build/build.sh help          - Show this help message"
    echo "  ./scripts/build/build.sh build         - Build the application"
    echo "  ./scripts/build/build.sh run           - Run the application"
    echo "  ./scripts/build/build.sh test          - Run tests"
    echo "  ./scripts/build/build.sh test-coverage - Run tests with coverage report"
    echo ""
    echo "Documentation:"
    echo "  ./scripts/build/build.sh docs          - Generate Swagger documentation"
    echo "  ./scripts/build/build.sh docs-serve    - Generate docs and start server"
    echo "  ./scripts/build/build.sh install-swag  - Install swag CLI tool"
    echo ""
    echo "Docker commands:"
    echo "  ./scripts/build/build.sh docker-build  - Build Docker image"
    echo "  ./scripts/build/build.sh docker-run    - Run Docker container"
    echo "  ./scripts/build/build.sh docker-dev    - Run development environment"
    echo ""
    echo "Code quality:"
    echo "  ./scripts/build/build.sh fmt           - Format Go code"
    echo "  ./scripts/build/build.sh vet           - Run go vet"
    echo "  ./scripts/build/build.sh lint          - Run all linting tools"
    echo "  ./scripts/build/build.sh ci-test       - Run tests as in CI"
    echo ""
    echo "Development tools:"
    echo "  ./scripts/build/build.sh install-air   - Install air for hot reloading"
    echo "  ./scripts/build/build.sh clean         - Clean build artifacts"
    echo ""
}

build_app() {
    echo "🔨 Building the application..."
    mkdir -p bin
    go build -o bin/payment-service cmd/server/main.go
    if [ $? -eq 0 ]; then
        echo "✅ Build completed successfully!"
    else
        echo "❌ Build failed!"
    fi
}

run_app() {
    echo "🚀 Running the application..."
    go run cmd/server/main.go
}

test_app() {
    echo "🧪 Running tests..."
    go test -v ./...
}

test_coverage() {
    echo "🧪 Running tests with coverage..."
    go test -v -coverprofile=coverage.out ./...
    if [ $? -eq 0 ]; then
        go tool cover -html=coverage.out -o coverage.html
        echo "✅ Coverage report generated: coverage.html"
    fi
}

generate_docs() {
    echo "📚 Generating Swagger documentation..."
    swag init -g cmd/server/main.go -o docs
    if [ $? -eq 0 ]; then
        echo "✅ Documentation generated successfully!"
    else
        echo "❌ Documentation generation failed. Make sure swag is installed."
        echo "Run: ./scripts/build/build.sh install-swag"
    fi
}

docs_serve() {
    generate_docs
    if [ $? -eq 0 ]; then
        echo "📚 Starting server with documentation..."
        echo "API documentation will be available at: http://localhost:8080/swagger/"
        run_app
    fi
}

install_swag() {
    echo "🔧 Installing swag CLI tool..."
    go install github.com/swaggo/swag/cmd/swag@latest
    if [ $? -eq 0 ]; then
        echo "✅ Swag installed successfully!"
    fi
}

install_air() {
    echo "🔧 Installing air for hot reloading..."
    go install github.com/cosmtrek/air@latest
    if [ $? -eq 0 ]; then
        echo "✅ Air installed successfully!"
    fi
}

docker_build() {
    echo "🐳 Building Docker image..."
    docker build -t payment-service:latest .
}

docker_run() {
    echo "🐳 Running Docker container..."
    docker run -p 8080:8080 payment-service:latest
}

docker_dev() {
    echo "🐳 Starting development environment..."
    docker-compose -f scripts/docker/docker-compose.dev.yml up --build
}

format_code() {
    echo "🎨 Formatting Go code..."
    go fmt ./...
}

vet_code() {
    echo "🔍 Running go vet..."
    go vet ./...
}

lint_code() {
    echo "🔍 Running linting tools..."
    format_code
    vet_code
}

ci_test() {
    echo "🧪 Running CI-style tests..."
    go mod verify
    go vet ./...
    if [ "$(gofmt -s -l . | wc -l)" -gt 0 ]; then
        echo "❌ Code is not formatted. Run: ./scripts/build/build.sh fmt"
        return 1
    fi
    go test -v -race -coverprofile=coverage.out ./...
}

clean_artifacts() {
    echo "🧹 Cleaning build artifacts..."
    rm -rf bin/
    rm -rf docs/
    rm -f coverage.out coverage.html
    echo "✅ Cleanup completed!"
}

# Main command dispatcher
case "${1:-help}" in
    "help") show_help ;;
    "build") build_app ;;
    "run") run_app ;;
    "test") test_app ;;
    "test-coverage") test_coverage ;;
    "docs") generate_docs ;;
    "docs-serve") docs_serve ;;
    "install-swag") install_swag ;;
    "install-air") install_air ;;
    "docker-build") docker_build ;;
    "docker-run") docker_run ;;
    "docker-dev") docker_dev ;;
    "fmt") format_code ;;
    "vet") vet_code ;;
    "lint") lint_code ;;
    "ci-test") ci_test ;;
    "clean") clean_artifacts ;;
    *) 
        echo "❌ Unknown command: $1"
        echo ""
        show_help 
        ;;
esac
