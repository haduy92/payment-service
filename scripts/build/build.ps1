# Payment Service - PowerShell Build Script
# Windows equivalent of Makefile commands

param(
    [Parameter(Position=0)]
    [string]$Command = "help"
)

function Show-Help {
    Write-Host "Payment Service - Available Commands:" -ForegroundColor Green
    Write-Host ""
    Write-Host "Convenience scripts (from root directory):" -ForegroundColor Cyan
    Write-Host "  .\build.ps1 run                          - Same as: .\scripts\build\build.ps1 run"
    Write-Host "  .\build.bat run                          - Same as: scripts\build\build.bat run"
    Write-Host ""
    Write-Host "Basic commands:" -ForegroundColor Blue
    Write-Host "  .\scripts\build\build.ps1 help          - Show this help message"
    Write-Host "  .\scripts\build\build.ps1 build         - Build the application"
    Write-Host "  .\scripts\build\build.ps1 run           - Run the application"
    Write-Host "  .\scripts\build\build.ps1 test          - Run tests"
    Write-Host "  .\scripts\build\build.ps1 test-coverage - Run tests with coverage report"
    Write-Host ""
    Write-Host "Documentation:" -ForegroundColor Blue
    Write-Host "  .\scripts\build\build.ps1 docs          - Generate Swagger documentation"
    Write-Host "  .\scripts\build\build.ps1 docs-serve    - Generate docs and start server"
    Write-Host "  .\scripts\build\build.ps1 install-swag  - Install swag CLI tool"
    Write-Host ""
    Write-Host "Docker commands:" -ForegroundColor Blue
    Write-Host "  .\scripts\build\build.ps1 docker-build  - Build Docker image"
    Write-Host "  .\scripts\build\build.ps1 docker-run    - Run Docker container"
    Write-Host "  .\scripts\build\build.ps1 docker-dev    - Run development environment"
    Write-Host ""
    Write-Host "Code quality:" -ForegroundColor Blue
    Write-Host "  .\scripts\build\build.ps1 fmt           - Format Go code"
    Write-Host "  .\scripts\build\build.ps1 vet           - Run go vet"
    Write-Host "  .\scripts\build\build.ps1 lint          - Run all linting tools"
    Write-Host "  .\scripts\build\build.ps1 ci-test       - Run tests as in CI"
    Write-Host ""
    Write-Host "Development tools:" -ForegroundColor Blue
    Write-Host "  .\scripts\build\build.ps1 install-air   - Install air for hot reloading"
    Write-Host "  .\scripts\build\build.ps1 clean         - Clean build artifacts"
    Write-Host ""
}

function Build-App {
    Write-Host "🔨 Building the application..." -ForegroundColor Blue
    New-Item -ItemType Directory -Force -Path "bin" | Out-Null
    go build -o bin/payment-service.exe cmd/server/main.go
    if ($LASTEXITCODE -eq 0) {
        Write-Host "✅ Build completed successfully!" -ForegroundColor Green
    } else {
        Write-Host "❌ Build failed!" -ForegroundColor Red
    }
}

function Run-App {
    Write-Host "🚀 Running the application..." -ForegroundColor Blue
    go run cmd/server/main.go
}

function Test-App {
    Write-Host "🧪 Running tests..." -ForegroundColor Blue
    go test -v ./...
}

function Test-Coverage {
    Write-Host "🧪 Running tests with coverage..." -ForegroundColor Blue
    go test -v -coverprofile=coverage.out ./...
    if ($LASTEXITCODE -eq 0) {
        go tool cover -html=coverage.out -o coverage.html
        Write-Host "✅ Coverage report generated: coverage.html" -ForegroundColor Green
    }
}

function Generate-Docs {
    Write-Host "📚 Generating Swagger documentation..." -ForegroundColor Blue
    & "$env:USERPROFILE\go\bin\swag.exe" init -g cmd/server/main.go -o docs
    if ($LASTEXITCODE -eq 0) {
        Write-Host "✅ Documentation generated successfully!" -ForegroundColor Green
    } else {
        Write-Host "❌ Documentation generation failed. Make sure swag is installed." -ForegroundColor Red
        Write-Host "Run: .\build.ps1 install-swag" -ForegroundColor Yellow
    }
}

function Docs-Serve {
    Generate-Docs
    if ($LASTEXITCODE -eq 0) {
        Write-Host "📚 Starting server with documentation..." -ForegroundColor Blue
        Write-Host "API documentation will be available at: http://localhost:8080/swagger/" -ForegroundColor Green
        Run-App
    }
}

function Install-Swag {
    Write-Host "🔧 Installing swag CLI tool..." -ForegroundColor Blue
    go install github.com/swaggo/swag/cmd/swag@latest
    if ($LASTEXITCODE -eq 0) {
        Write-Host "✅ Swag installed successfully!" -ForegroundColor Green
    }
}

function Install-Air {
    Write-Host "🔧 Installing air for hot reloading..." -ForegroundColor Blue
    go install github.com/cosmtrek/air@latest
    if ($LASTEXITCODE -eq 0) {
        Write-Host "✅ Air installed successfully!" -ForegroundColor Green
    }
}

function Docker-Build {
    Write-Host "🐳 Building Docker image..." -ForegroundColor Blue
    docker build -t payment-service:latest .
}

function Docker-Run {
    Write-Host "🐳 Running Docker container..." -ForegroundColor Blue
    docker run -p 8080:8080 payment-service:latest
}

function Docker-Dev {
    Write-Host "🐳 Starting development environment..." -ForegroundColor Blue
    docker-compose -f scripts/docker/docker-compose.dev.yml up --build
}

function Format-Code {
    Write-Host "🎨 Formatting Go code..." -ForegroundColor Blue
    go fmt ./...
}

function Vet-Code {
    Write-Host "🔍 Running go vet..." -ForegroundColor Blue
    go vet ./...
}

function Lint-Code {
    Write-Host "🔍 Running linting tools..." -ForegroundColor Blue
    Format-Code
    Vet-Code
}

function CI-Test {
    Write-Host "🧪 Running CI-style tests..." -ForegroundColor Blue
    go mod verify
    go vet ./...
    $formatted = go fmt ./...
    if ($formatted) {
        Write-Host "❌ Code is not formatted. Run: .\build.ps1 fmt" -ForegroundColor Red
        return
    }
    go test -v -race -coverprofile=coverage.out ./...
}

function Clean-Artifacts {
    Write-Host "🧹 Cleaning build artifacts..." -ForegroundColor Blue
    Remove-Item -Path "bin" -Recurse -Force -ErrorAction SilentlyContinue
    Remove-Item -Path "docs" -Recurse -Force -ErrorAction SilentlyContinue
    Remove-Item -Path "coverage.out" -Force -ErrorAction SilentlyContinue
    Remove-Item -Path "coverage.html" -Force -ErrorAction SilentlyContinue
    Write-Host "✅ Cleanup completed!" -ForegroundColor Green
}

# Main command dispatcher
switch ($Command.ToLower()) {
    "help" { Show-Help }
    "build" { Build-App }
    "run" { Run-App }
    "test" { Test-App }
    "test-coverage" { Test-Coverage }
    "docs" { Generate-Docs }
    "docs-serve" { Docs-Serve }
    "install-swag" { Install-Swag }
    "install-air" { Install-Air }
    "docker-build" { Docker-Build }
    "docker-run" { Docker-Run }
    "docker-dev" { Docker-Dev }
    "fmt" { Format-Code }
    "vet" { Vet-Code }
    "lint" { Lint-Code }
    "ci-test" { CI-Test }
    "clean" { Clean-Artifacts }
    default { 
        Write-Host "❌ Unknown command: $Command" -ForegroundColor Red
        Write-Host ""
        Show-Help 
    }
}
