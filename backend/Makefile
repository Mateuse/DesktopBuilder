.PHONY: lint

lint:
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not found, running basic checks (go vet, go fmt verification)"; \
		go vet ./...; \
		out=$$(gofmt -l .); if [ -n "$$out" ]; then echo "Files need gofmt:"; echo "$$out"; exit 1; fi; \
	fi


