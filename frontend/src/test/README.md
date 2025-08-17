# Testing Setup

This directory contains the testing configuration and utilities for the frontend application.

## Test Structure

We're using the **co-located testing approach** where test files are placed next to the components/modules they test:

```
src/
├── components/
│   ├── health.component.tsx
│   └── health.component.test.tsx
├── api/
│   ├── health.api.ts
│   └── health.api.test.ts
└── test/
    ├── setup.ts          # Global test setup
    └── README.md         # This file
```

## Testing Stack

- **Vitest**: Fast test runner built for Vite projects
- **React Testing Library**: For testing React components
- **@testing-library/jest-dom**: Additional matchers for DOM elements
- **@testing-library/user-event**: For simulating user interactions

## Available Scripts

```bash
npm test            # Run tests in watch mode
npm run test:run    # Run tests once
npm run test:ui     # Run tests with UI (requires @vitest/ui)
npm run test:coverage # Run tests with coverage report
```

## Test Patterns

### Component Testing
```tsx
import { render, screen } from '@testing-library/react';
import { MyComponent } from './MyComponent';

describe('MyComponent', () => {
  it('renders correctly', () => {
    render(<MyComponent />);
    expect(screen.getByText('Hello')).toBeInTheDocument();
  });
});
```

### API Testing
```ts
import { myApiFunction } from './my.api';

// Mock fetch globally
global.fetch = vi.fn();
const mockFetch = vi.mocked(fetch);

describe('myApiFunction', () => {
  it('fetches data successfully', async () => {
    mockFetch.mockResolvedValue({
      ok: true,
      json: async () => ({ data: 'test' }),
    } as Response);

    const result = await myApiFunction();
    expect(result).toEqual({ data: 'test' });
  });
});
```

## Setup File

The `setup.ts` file contains:
- Jest-DOM matchers import
- Browser API mocks (matchMedia, ResizeObserver, etc.) for Mantine compatibility
- Global test configuration

## File Naming Conventions

- `*.test.tsx` - React component tests
- `*.test.ts` - TypeScript module tests  
- `*.spec.tsx` - Specification/behavior tests
- `__tests__/` - Alternative directory structure for grouping tests
