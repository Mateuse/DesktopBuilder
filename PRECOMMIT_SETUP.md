# Pre-commit Hooks Setup

This project uses pre-commit hooks to ensure code quality by running linters and tests before each commit.

## Quick Setup

Run the following command to install and configure pre-commit hooks:

```bash
make setup-hooks
```

This will:
1. Install `pre-commit` if not already installed
2. Install the pre-commit hooks defined in `.pre-commit-config.yaml`
3. Configure git to run these hooks before each commit

## What Gets Checked

### Frontend (when frontend files are changed):
- **ESLint**: Code linting with `npm run lint`
- **Tests**: All tests with `npm run test:run`

### Backend (when backend files are changed):
- **Go Linting**: Code linting with `make lint` (golangci-lint or go vet/gofmt)
- **Go Tests**: Unit and integration tests with `make test`

### General (all commits):
- Trailing whitespace removal
- End-of-file fixer
- YAML/JSON syntax validation
- Merge conflict detection
- Large file detection

## Manual Commands

```bash
# Test all hooks without committing
make test-hooks

# Run only linters
make lint

# Run only tests
make test

# Frontend only
make frontend-lint
make frontend-test

# Backend only
make backend-lint
make backend-test
```

## How It Works

1. **Before each commit**: Pre-commit hooks automatically run
2. **If any check fails**: The commit is blocked
3. **Fix the issues**: Address linting errors or failing tests
4. **Commit again**: Once all checks pass, the commit proceeds

## Bypassing Hooks (Not Recommended)

In emergency situations, you can bypass hooks with:
```bash
git commit --no-verify -m "your message"
```

**Note**: This should only be used in exceptional circumstances as it defeats the purpose of having quality gates.

## Troubleshooting

### Pre-commit not found
```bash
pip install pre-commit
```

### Hooks not running
```bash
pre-commit install
```

### Update hooks
```bash
pre-commit autoupdate
```

### Clear hook cache
```bash
pre-commit clean
```
