# GitHub Actions CI/CD Pipeline

This repository uses GitHub Actions for continuous integration and deployment. The pipeline automatically runs tests for both the backend (Go) and frontend (TypeScript/React) components.

## Workflow Overview

The CI/CD pipeline (`ci.yml`) consists of three main jobs:

### 1. Backend Tests (`backend-tests`)
- **Language**: Go 1.24.5
- **Services**: PostgreSQL 15, Redis 7
- **Steps**:
  - Code linting (go vet, gofmt)
  - Unit tests (`make test-unit`)
  - Integration tests (`make test-integration`)
  - Coverage report generation
  - Coverage upload to Codecov

### 2. Frontend Tests (`frontend-tests`)
- **Language**: Node.js 20 with TypeScript
- **Steps**:
  - ESLint code linting
  - TypeScript type checking
  - Unit tests with Vitest
  - Coverage report generation
  - Coverage upload to Codecov

### 3. Build Test (`build-test`)
- **Dependencies**: Runs after both backend and frontend tests pass
- **Steps**:
  - Build backend binary
  - Build frontend for production
  - Verify build artifacts

## Triggers

The workflow runs on:
- **Push** to `master`, `main`, or `develop` branches
- **Pull requests** targeting `master`, `main`, or `develop` branches

## Environment Variables

The backend tests use the following environment variables for database connections:
- `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME` (PostgreSQL)
- `REDIS_HOST`, `REDIS_PORT` (Redis)

## Coverage Reports

Coverage reports are automatically uploaded to [Codecov](https://codecov.io) for both backend and frontend components. You can view detailed coverage reports by setting up Codecov integration for your repository.

## Local Development

To run the same tests locally:

### Backend
```bash
cd backend
make lint          # Run linter
make test-unit      # Run unit tests
make test-integration  # Run integration tests (requires PostgreSQL and Redis)
make test-coverage  # Run tests with coverage
```

### Frontend
```bash
cd frontend
npm run lint        # Run ESLint
npx tsc --noEmit   # Type checking
npm run test:run    # Run tests
npm run test:coverage  # Run tests with coverage
```

## Status Badges

You can add status badges to your main README.md:

```markdown
![CI/CD Pipeline](https://github.com/yourusername/DesktopBuilder/workflows/CI/CD%20Pipeline/badge.svg)
[![codecov](https://codecov.io/gh/yourusername/DesktopBuilder/branch/master/graph/badge.svg)](https://codecov.io/gh/yourusername/DesktopBuilder)
```

Replace `yourusername` with your actual GitHub username.
