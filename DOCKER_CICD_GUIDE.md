# Payment## 🚀 Quick Demo Steps

### 1. Initial Setup
```bash
# Clone and run locally
git clone <your-repo>
cd payment-service

# Using Make (Linux/macOS - Recommended)
make run

# Using organized scripts
# Windows PowerShell
.\scripts\build\build.ps1 run

# Linux/macOS
./scripts/build/build.sh run

# Using convenience scripts (calls organized scripts)
# Windows
.\build.ps1 run

# Linux/macOS
./build.sh run

# Or direct Go command (all platforms)
go run cmd/server/main.go
```- CI/CD Demo Guide

This guide shows how to demonstrate the CI/CD pipeline with this Payment Service.

## 🎯 Demo Scenario

This project demonstrates a simple CI/CD workflow:
**Code Change → Pull Request → Automated Testing → Merge → Deployment**

## � Quick Demo Steps

### 1. Initial Setup
```bash
# Clone and run locally
git clone <your-repo>
cd payment-service
go run cmd/server/main.go
```

### 2. Make a Change
```bash
# Create a feature branch
git checkout -b add-new-feature

# Make a small change (e.g., update a comment in main.go)
# Edit cmd/server/main.go and change a comment

# Commit the change
git add .
git commit -m "feat: update API description"
git push origin add-new-feature
```

### 3. Create Pull Request
1. Go to GitHub and create a PR from `add-new-feature` to `main`
2. Watch the CI pipeline run:
   - ✅ Tests
   - ✅ Build
   - ✅ Code quality checks

### 4. Merge to Main
1. Merge the PR
2. Watch the deployment pipeline:
   - ✅ Docker build
   - ✅ "Production" deployment simulation

## 🐳 Docker Demo

### Build and Run Container
```bash
# Build the Docker image
docker build -t payment-service .

# Run the container
docker run -p 8080:8080 payment-service

# Test the API
curl http://localhost:8080/health
curl http://localhost:8080/swagger/
```

### Development with Hot Reload
```bash
# Start development environment
make docker-dev

# Make changes to code - they'll automatically reload!
```

## 🔄 CI/CD Pipeline Overview

### On Pull Request:
- Runs tests
- Checks code quality
- Builds application
- Reports status back to PR

### On Merge to Main:
- Runs full test suite
- Builds Docker image
- Pushes to container registry
- Simulates deployment

### Pipeline Jobs:
1. **Test** - Unit tests, linting, format checks
2. **Build** - Compile application and create artifacts  
3. **Docker** - Build and push container image
4. **Deploy** - Simulate production deployment

## � Demo Checklist

- [ ] Repository set up on GitHub
- [ ] GitHub Actions enabled
- [ ] Local development environment working
- [ ] Can create and merge PRs
- [ ] Pipeline runs on PR creation
- [ ] Pipeline runs on merge to main
- [ ] Docker images build successfully
- [ ] API documentation accessible

## �️ Customization for Your Demo

### Change API Response
Edit `cmd/server/main.go` to modify the health check response:
```go
fmt.Fprint(w, `{"status":"ok","service":"payment-service","version":"v2.0"}`)
```

### Add New Endpoint
Add a simple endpoint in `internal/handler/payment.go`:
```go
r.Get("/demo", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprint(w, `{"message":"Demo endpoint working!"}`)
})
```

### Trigger Pipeline
```bash
git add .
git commit -m "feat: add demo endpoint"
git push origin feature-branch
# Create PR and watch pipeline run!
```

This creates a complete, working demonstration of modern CI/CD practices with Go and Docker!
