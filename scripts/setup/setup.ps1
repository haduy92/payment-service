# Payment Service Setup Script (PowerShell)
# This script sets up the development environment for the Payment Service project

Write-Host "🚀 Setting up Payment Service development environment..." -ForegroundColor Green

# Check if Go is installed
try {
    $goVersion = go version
    Write-Host "✅ Go is installed: $goVersion" -ForegroundColor Green
} catch {
    Write-Host "❌ Go is not installed. Please install Go 1.21 or higher." -ForegroundColor Red
    exit 1
}

# Install dependencies
Write-Host "📦 Installing Go dependencies..." -ForegroundColor Blue
go mod download
go mod tidy

# Install swag CLI tool
Write-Host "🔧 Installing swag CLI tool..." -ForegroundColor Blue
go install github.com/swaggo/swag/cmd/swag@latest

# Generate Swagger documentation
Write-Host "📚 Generating Swagger documentation..." -ForegroundColor Blue
& "$env:USERPROFILE\go\bin\swag.exe" init -g cmd/server/main.go -o docs

# Build the application
Write-Host "🔨 Building the application..." -ForegroundColor Blue
go build -o bin/payment-service.exe cmd/server/main.go

# Run tests
Write-Host "🧪 Running tests..." -ForegroundColor Blue
go test -v ./...

Write-Host ""
Write-Host "✅ Setup completed successfully!" -ForegroundColor Green
Write-Host ""
Write-Host "To start the service:"
Write-Host "  go run cmd/server/main.go"
Write-Host ""
Write-Host "To view API documentation:"
Write-Host "  1. Start the service"
Write-Host "  2. Open http://localhost:8080/swagger/ in your browser"
Write-Host ""
Write-Host "Available make commands (if make is installed):"
Write-Host "  make help          # Show available commands"
Write-Host "  make run           # Run the application"
Write-Host "  make docs          # Generate documentation"
Write-Host "  make test          # Run tests"
Write-Host ""
