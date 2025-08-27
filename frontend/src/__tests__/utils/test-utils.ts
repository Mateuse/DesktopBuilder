import { Component } from '@/types/components.type';

/**
 * Mock component data for testing
 */
export const mockComponent: Component = {
  id: 1,
  category: 'cpu',
  brand: 'Intel',
  model: 'Core i7-12700K',
  sku: 'BX8071512700K',
  upc: '735858491921',
  specs: 'Socket: LGA1700, Cores: 12, Threads: 20, Base Clock: 3.6GHz',
  releaseDate: '2021-11-04T00:00:00Z',
  createdAt: '2023-01-01T00:00:00Z',
};

/**
 * Generate array of mock components
 */
export const generateMockComponents = (count: number): Component[] => {
  return Array.from({ length: count }, (_, index) => ({
    ...mockComponent,
    id: index + 1,
    model: `${mockComponent.model}-${index + 1}`,
    sku: `${mockComponent.sku}-${index + 1}`,
  }));
};

/**
 * Mock successful fetch response
 */
export const mockFetchSuccess = (data: unknown) => {
  (global.fetch as jest.Mock).mockResolvedValueOnce({
    ok: true,
    json: jest.fn().mockResolvedValueOnce(data),
  });
};

/**
 * Mock failed fetch response
 */
export const mockFetchError = (error: string = 'Network error') => {
  (global.fetch as jest.Mock).mockRejectedValueOnce(new Error(error));
};

/**
 * Mock fetch response with specific status
 */
export const mockFetchWithStatus = (status: number, data?: unknown) => {
  (global.fetch as jest.Mock).mockResolvedValueOnce({
    ok: status >= 200 && status < 300,
    status,
    json: jest.fn().mockResolvedValueOnce(data || {}),
  });
};

/**
 * Verify fetch was called with correct URL and options
 */
export const expectFetchToHaveBeenCalledWith = (url: string, options?: RequestInit) => {
  if (options === undefined) {
    expect(global.fetch).toHaveBeenCalledWith(url);
  } else {
    expect(global.fetch).toHaveBeenCalledWith(url, options);
  }
};
