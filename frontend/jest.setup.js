// Optional: configure or set up a testing framework before each test.
// If you delete this file, remove `setupFilesAfterEnv` from `jest.config.js`

// Used for __tests__/testing-library.js
// Learn more: https://github.com/testing-library/jest-dom
import '@testing-library/jest-dom'

// Mock fetch globally
global.fetch = jest.fn()

// Set environment variables for testing
process.env.NEXT_PUBLIC_BACKEND_URL = 'http://localhost:8080'

// Reset fetch mock before each test
beforeEach(() => {
  jest.resetAllMocks()
})
