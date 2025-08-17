import { render, screen } from '@testing-library/react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { MantineProvider } from '@mantine/core';
import { HealthComponent } from './health.component';
import { mainHealth } from '../api/health.api';
import { vi } from 'vitest';

// Mock the health API
vi.mock('../api/health.api', () => ({
  mainHealth: vi.fn(),
}));

const mockedMainHealth = vi.mocked(mainHealth);

// Test wrapper component
const TestWrapper = ({ children }: { children: React.ReactNode }) => {
  const queryClient = new QueryClient({
    defaultOptions: {
      queries: {
        retry: false,
      },
    },
  });

  return (
    <QueryClientProvider client={queryClient}>
      <MantineProvider>
        {children}
      </MantineProvider>
    </QueryClientProvider>
  );
};

describe('HealthComponent', () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  it('renders loading state initially', () => {
    mockedMainHealth.mockImplementation(() => new Promise(() => {})); // Never resolves

    render(
      <TestWrapper>
        <HealthComponent />
      </TestWrapper>
    );

    expect(screen.getByText('Loading...')).toBeInTheDocument();
  });

  it('renders health data when API call succeeds', async () => {
    const mockHealthData = { message: 'Service is healthy!' };
    mockedMainHealth.mockResolvedValue(mockHealthData);

    render(
      <TestWrapper>
        <HealthComponent />
      </TestWrapper>
    );

    expect(await screen.findByText('Health')).toBeInTheDocument();
    expect(await screen.findByText('Service is healthy!')).toBeInTheDocument();
  });

  it('renders error message when API call fails', async () => {
    const errorMessage = 'Failed to fetch health status';
    mockedMainHealth.mockRejectedValue(new Error(errorMessage));

    render(
      <TestWrapper>
        <HealthComponent />
      </TestWrapper>
    );

    expect(await screen.findByText(`Error: ${errorMessage}`)).toBeInTheDocument();
  });
});
