.PHONY: lint test test-unit test-integration test-coverage test-race test-bench test-all clean-test

lint:
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not found, running basic checks (go vet, go fmt verification)"; \
		go vet ./...; \
		out=$$(gofmt -l .); if [ -n "$$out" ]; then echo "Files need gofmt:"; echo "$$out"; exit 1; fi; \
	fi

# Run all tests
test: test-unit test-integration

# Run unit tests only
test-unit:
	@echo "Running unit tests..."
	go test -v -short ./internal/...

# Run integration tests only
test-integration:
	@echo "Running integration tests..."
	go test -v ./test/integration/...

# Run tests with coverage report
test-coverage:
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run tests with race detection
test-race:
	@echo "Running tests with race detection..."
	go test -race -v ./...

# Run benchmark tests
test-bench:
	@echo "Running benchmark tests..."
	go test -bench=. -benchmem ./...

# Run all test types
test-all: lint test-unit test-integration test-race test-bench

# Clean test artifacts
clean-test:
	@echo "Cleaning test artifacts..."
	rm -f coverage.out coverage.html
	go clean -testcache

# Run tests in watch mode (requires entr: brew install entr or apt-get install entr)
test-watch:
	@echo "Running tests in watch mode (requires 'entr' to be installed)..."
	find . -name "*.go" | entr -c make test-unit

# Generate test coverage and open in browser (macOS/Linux)
test-coverage-open: test-coverage
	@if command -v open >/dev/null 2>&1; then \
		open coverage.html; \
	elif command -v xdg-open >/dev/null 2>&1; then \
		xdg-open coverage.html; \
	else \
		echo "Coverage report generated: coverage.html"; \
	fi


