@echo off
REM Payment Service - Windows Batch Script
REM Simple commands for Windows Command Prompt users

if "%1"=="" goto help
if "%1"=="help" goto help
if "%1"=="run" goto run
if "%1"=="test" goto test
if "%1"=="build" goto build
if "%1"=="docs" goto docs
goto unknown

:help
echo Payment Service - Available Commands:
echo.
echo Convenience scripts (from root directory):
echo   build.bat run    - Same as: scripts\build\build.bat run
echo   build.ps1 run    - Same as: scripts\build\build.ps1 run
echo.
echo Basic commands:
echo   scripts\build\build.bat help   - Show this help
echo   scripts\build\build.bat run    - Run the application
echo   scripts\build\build.bat test   - Run tests
echo   scripts\build\build.bat build  - Build the application
echo   scripts\build\build.bat docs   - Generate documentation
echo.
echo For more commands, use scripts\build\build.ps1 (PowerShell script)
goto end

:run
echo Running the application...
go run cmd/server/main.go
goto end

:test
echo Running tests...
go test ./...
goto end

:build
echo Building the application...
if not exist bin mkdir bin
go build -o bin/payment-service.exe cmd/server/main.go
goto end

:docs
echo Generating documentation...
%USERPROFILE%\go\bin\swag.exe init -g cmd/server/main.go -o docs
goto end

:unknown
echo Unknown command: %1
echo Run 'build.bat help' for available commands
goto end

:end
