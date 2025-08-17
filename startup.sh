#!/bin/bash

# Startup script for DesktopBuilder application
# This script checks and starts PostgreSQL and Redis if needed, then runs the Go backend

set -e  # Exit on any error

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Function to check if a service is running
is_service_running() {
    local service_name="$1"
    systemctl is-active --quiet "$service_name" 2>/dev/null
}

# Function to check if PostgreSQL is running
check_postgres() {
    print_status "Checking PostgreSQL status..."

    # Check if PostgreSQL service is running
    if is_service_running "postgresql"; then
        print_success "PostgreSQL is already running"
        return 0
    fi

    # Try alternative service names
    if is_service_running "postgres"; then
        print_success "PostgreSQL is already running"
        return 0
    fi

    # Check if PostgreSQL is running on the expected port
    if netstat -tuln 2>/dev/null | grep -q ":5432 "; then
        print_success "PostgreSQL is running on port 5432"
        return 0
    fi

    return 1
}

# Function to start PostgreSQL
start_postgres() {
    print_status "Starting PostgreSQL..."

    # Try to start PostgreSQL service
    if sudo systemctl start postgresql 2>/dev/null; then
        print_success "PostgreSQL started successfully"
        return 0
    fi

    # Try alternative service name
    if sudo systemctl start postgres 2>/dev/null; then
        print_success "PostgreSQL started successfully"
        return 0
    fi

    print_error "Failed to start PostgreSQL. Please ensure PostgreSQL is installed and configured."
    return 1
}

# Function to check if Redis is running
check_redis() {
    print_status "Checking Redis status..."

    # Check if Redis service is running
    if is_service_running "redis-server"; then
        print_success "Redis is already running"
        return 0
    fi

    # Try alternative service name
    if is_service_running "redis"; then
        print_success "Redis is already running"
        return 0
    fi

    # Check if Redis is running on the expected port
    if netstat -tuln 2>/dev/null | grep -q ":6379 "; then
        print_success "Redis is running on port 6379"
        return 0
    fi

    # Try to ping Redis directly
    if redis-cli ping >/dev/null 2>&1; then
        print_success "Redis is responding to ping"
        return 0
    fi

    return 1
}

# Function to start Redis
start_redis() {
    print_status "Starting Redis..."

    # Try to start Redis service
    if sudo systemctl start redis-server 2>/dev/null; then
        print_success "Redis started successfully"
        return 0
    fi

    # Try alternative service name
    if sudo systemctl start redis 2>/dev/null; then
        print_success "Redis started successfully"
        return 0
    fi

    # Try to start Redis manually in the background
    if command -v redis-server >/dev/null 2>&1; then
        print_status "Starting Redis manually..."
        redis-server --daemonize yes
        sleep 2
        if redis-cli ping >/dev/null 2>&1; then
            print_success "Redis started manually"
            return 0
        fi
    fi

    print_error "Failed to start Redis. Please ensure Redis is installed and configured."
    return 1
}

# Function to run the Go application
run_go_app() {
    print_status "Starting Go application..."

    # Change to backend directory
    cd "$(dirname "$0")/backend"

    # Check if .env file exists
    if [ ! -f ".env" ]; then
        print_warning ".env file not found in backend directory"
        print_warning "The application will use system environment variables"
        print_warning "Please refer to ENV_SETUP.md for configuration details"
    fi

    # Run the Go application
    print_status "Running: go run cmd/api/main.go"
    go run cmd/api/main.go
}

# Main execution
main() {
    print_status "Starting DesktopBuilder application..."
    echo

    # Check and start PostgreSQL
    if ! check_postgres; then
        if ! start_postgres; then
            print_error "Cannot proceed without PostgreSQL. Exiting."
            exit 1
        fi
        # Wait a moment for PostgreSQL to fully start
        sleep 3
    fi
    echo

    # Check and start Redis
    if ! check_redis; then
        if ! start_redis; then
            print_error "Cannot proceed without Redis. Exiting."
            exit 1
        fi
        # Wait a moment for Redis to fully start
        sleep 2
    fi
    echo

    # Run the Go application
    print_status "All dependencies are running. Starting the application..."
    echo
    run_go_app
}

# Handle script interruption
cleanup() {
    echo
    print_status "Shutting down..."
    exit 0
}

trap cleanup SIGINT SIGTERM

# Check if required commands are available
for cmd in systemctl netstat redis-cli go; do
    if ! command -v "$cmd" >/dev/null 2>&1; then
        print_warning "Command '$cmd' not found. Some features may not work properly."
    fi
done

# Run main function
main "$@"
