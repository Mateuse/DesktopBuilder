.PHONY: lint frontend-lint backend-lint

frontend-lint:
	@cd frontend && npm run lint --silent

backend-lint:
	$(MAKE) -C backend lint

lint: frontend-lint backend-lint
	@echo "All linters passed"


