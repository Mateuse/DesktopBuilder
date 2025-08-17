import { mainHealth } from './health.api';
import { ROUTES } from '../constants/urls';
import type { HealthResponse } from '../schemas/health.schemas';

// Mock fetch globally
global.fetch = vi.fn();
const mockFetch = vi.mocked(fetch);

describe('mainHealth API', () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  it('should fetch health data successfully', async () => {
    const mockHealthData: HealthResponse = { message: 'Service is healthy!' };

    mockFetch.mockResolvedValue({
      ok: true,
      json: async () => await Promise.resolve(mockHealthData),
    } as Response);

    const result = await mainHealth();

    expect(mockFetch).toHaveBeenCalledWith(ROUTES.MAIN_BE_HEALTH, undefined);
    expect(result).toEqual(mockHealthData);
  });

  it('should throw error when response is not ok', async () => {
    mockFetch.mockResolvedValue({
      ok: false,
      status: 500,
      statusText: 'Internal Server Error',
    } as Response);

    await expect(mainHealth()).rejects.toThrow('HTTP 500: Internal Server Error');
    expect(mockFetch).toHaveBeenCalledWith(ROUTES.MAIN_BE_HEALTH, undefined);
  });

  it('should throw error when fetch fails', async () => {
    const networkError = new Error('Network error');
    mockFetch.mockRejectedValue(networkError);

    await expect(mainHealth()).rejects.toThrow('Network error');
    expect(mockFetch).toHaveBeenCalledWith(ROUTES.MAIN_BE_HEALTH, undefined);
  });

  it('should validate response data with Zod schema', async () => {
    const validData = { message: 'Service is healthy!' };

    mockFetch.mockResolvedValue({
      ok: true,
      json: async () => await Promise.resolve(validData),
    } as Response);

    const result = await mainHealth();
    expect(result).toEqual(validData);
  });

  it('should throw error for invalid response data', async () => {
    const invalidData = { wrongField: 'invalid' }; // Missing required 'message' field

    mockFetch.mockResolvedValue({
      ok: true,
      json: async () => await Promise.resolve(invalidData),
    } as Response);

    await expect(mainHealth()).rejects.toThrow('Invalid API response:');
    expect(mockFetch).toHaveBeenCalledWith(ROUTES.MAIN_BE_HEALTH, undefined);
  });
});
