#!/bin/bash

# Payment Service Setup Script
# This script sets up the development environment for the Payment Service project

echo "🚀 Setting up Payment Service development environment..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.21 or higher."
    exit 1
fi

echo "✅ Go is installed: $(go version)"

# Install dependencies
echo "📦 Installing Go dependencies..."
go mod download
go mod tidy

# Install swag CLI tool
echo "🔧 Installing swag CLI tool..."
go install github.com/swaggo/swag/cmd/swag@latest

# Generate Swagger documentation
echo "📚 Generating Swagger documentation..."
swag init -g cmd/server/main.go -o docs

# Build the application
echo "🔨 Building the application..."
go build -o bin/payment-service cmd/server/main.go

# Run tests
echo "🧪 Running tests..."
go test -v ./...

echo ""
echo "✅ Setup completed successfully!"
echo ""
echo "To start the service:"
echo "  go run cmd/server/main.go"
echo ""
echo "To view API documentation:"
echo "  1. Start the service"
echo "  2. Open http://localhost:8080/swagger/ in your browser"
echo ""
echo "Available make commands:"
echo "  make help          # Show available commands"
echo "  make run           # Run the application"
echo "  make docs          # Generate documentation"
echo "  make test          # Run tests"
echo ""
