import { mainHealth } from './health.api';
import { ROUTES } from '../constants/urls';

// Mock fetch globally
global.fetch = vi.fn();
const mockFetch = vi.mocked(fetch);

describe('mainHealth API', () => {
  beforeEach(() => {
    vi.clearAllMocks();
  });

  it('should fetch health data successfully', async () => {
    const mockHealthData = { message: 'Service is healthy!' };
    
    mockFetch.mockResolvedValue({
      ok: true,
      json: async () => mockHealthData,
    } as Response);

    const result = await mainHealth();

    expect(mockFetch).toHaveBeenCalledWith(ROUTES.MAIN_BE_HEALTH);
    expect(result).toEqual(mockHealthData);
  });

  it('should throw error when response is not ok', async () => {
    mockFetch.mockResolvedValue({
      ok: false,
      status: 500,
    } as Response);

    await expect(mainHealth()).rejects.toThrow('Failed to fetch health');
    expect(mockFetch).toHaveBeenCalledWith(ROUTES.MAIN_BE_HEALTH);
  });

  it('should throw error when fetch fails', async () => {
    const networkError = new Error('Network error');
    mockFetch.mockRejectedValue(networkError);

    await expect(mainHealth()).rejects.toThrow('Network error');
    expect(mockFetch).toHaveBeenCalledWith(ROUTES.MAIN_BE_HEALTH);
  });
});
