# Frontend Tests

This directory contains all the tests for the frontend application.

## Structure

- `api/` - Tests for API functions
- `utils/` - Test utilities and helpers

## Running Tests

```bash
# Run all tests
npm test

# Run tests in watch mode
npm run test:watch

# Run tests with coverage
npm run test:coverage
```

## Test Utilities

The `utils/test-utils.ts` file contains helpful utilities for testing:

- `mockComponent` - Mock component data
- `generateMockComponents(count)` - Generate multiple mock components
- `mockFetchSuccess(data)` - Mock successful fetch responses
- `mockFetchError(error)` - Mock failed fetch responses
- `mockFetchWithStatus(status, data)` - Mock responses with specific status codes
- `expectFetchToHaveBeenCalledWith(url, options)` - Verify fetch calls

## API Tests

The API tests cover:

- ✅ All four API functions (`getComponents`, `getComponentsByCategory`, `getComponentsByBrand`, `getComponentById`)
- ✅ Default and custom pagination
- ✅ ID conversion to strings
- ✅ Error handling
- ✅ Network failures
- ✅ Malformed responses
- ✅ Edge cases

## Coverage

The tests aim for comprehensive coverage of:
- Happy path scenarios
- Error handling
- Edge cases
- Data transformation
- Network failures
