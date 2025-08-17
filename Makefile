.PHONY: lint frontend-lint backend-lint test frontend-test backend-test setup-hooks install-precommit

frontend-lint:
	@cd frontend && npm run lint --silent

backend-lint:
	$(MAKE) -C backend lint

lint: frontend-lint backend-lint
	@echo "All linters passed"

frontend-test:
	@cd frontend && npm run test:run

backend-test:
	$(MAKE) -C backend test

test: frontend-test backend-test
	@echo "All tests passed"

# Install pre-commit if not already installed
install-precommit:
	@if ! command -v pre-commit >/dev/null 2>&1; then \
		echo "Installing pre-commit..."; \
		if command -v pipx >/dev/null 2>&1; then \
			pipx install pre-commit; \
		elif command -v apt >/dev/null 2>&1; then \
			echo "Trying to install via apt..."; \
			sudo apt update && sudo apt install -y pre-commit; \
		else \
			echo "Please install pre-commit manually:"; \
			echo "  Option 1: pipx install pre-commit"; \
			echo "  Option 2: sudo apt install pre-commit"; \
			echo "  Option 3: pip install --user pre-commit"; \
			exit 1; \
		fi; \
	else \
		echo "pre-commit is already installed"; \
	fi

# Setup pre-commit hooks
setup-hooks: install-precommit
	@echo "Setting up pre-commit hooks..."
	pre-commit install
	@echo "Pre-commit hooks installed successfully!"
	@echo "To test the hooks, run: make test-hooks"

# Test pre-commit hooks without committing
test-hooks:
	@echo "Testing pre-commit hooks on all files..."
	pre-commit run --all-files


