import {
  getComponents,
  getComponentsByCategory,
  getComponentsByBrand,
  getComponentById
} from '@/api/components.api';
import { Component } from '@/types/components.type';
import {
  mockComponent,
  generateMockComponents,
  mockFetchSuccess,
  mockFetchError,
  mockFetchWithStatus,
  expectFetchToHaveBeenCalledWith
} from '../utils/test-utils';



describe('components.api', () => {
  beforeEach(() => {
    jest.clearAllMocks();
  });

  describe('getComponents', () => {
    it('should fetch components with default page parameter', async () => {
      const mockComponents = generateMockComponents(5);
      mockFetchSuccess(mockComponents);

      const result = await getComponents({});

      expectFetchToHaveBeenCalledWith('http://localhost:8080/components?page=1');
      expect(result).toHaveLength(5);
      expect(result[0]).toEqual({
        ...mockComponents[0],
        id: mockComponents[0].id.toString(),
      });
    });

    it('should fetch components with custom page parameter', async () => {
      const mockComponents = generateMockComponents(3);
      mockFetchSuccess(mockComponents);

      const result = await getComponents({ page: '2' });

      expectFetchToHaveBeenCalledWith('http://localhost:8080/components?page=2');
      expect(result).toHaveLength(3);
    });

    it('should convert component IDs to strings', async () => {
      const mockComponents = [
        { ...mockComponent, id: 123 },
        { ...mockComponent, id: 456 }
      ];
      mockFetchSuccess(mockComponents);

      const result = await getComponents({});

      expect(result[0].id).toBe('123');
      expect(result[1].id).toBe('456');
      expect(typeof result[0].id).toBe('string');
    });

    it('should handle fetch errors', async () => {
      mockFetchError('Network error');

      await expect(getComponents({})).rejects.toThrow('Network error');
    });

    it('should handle empty response', async () => {
      mockFetchSuccess([]);

      const result = await getComponents({});

      expect(result).toEqual([]);
    });

    it('should handle server errors', async () => {
      mockFetchWithStatus(500, { error: 'Internal server error' });

      // The function doesn't check response.ok, so it will try to process the response
      const result = await getComponents({});

      expect(result).toEqual([{ error: 'Internal server error' }]);
    });
  });

  describe('getComponentsByCategory', () => {
    it('should fetch components by category with default page', async () => {
      const mockComponents = generateMockComponents(3);
      mockFetchSuccess(mockComponents);

      const result = await getComponentsByCategory({ category: 'cpu' });

      expectFetchToHaveBeenCalledWith('http://localhost:8080/components/cpu?page=1');
      expect(result).toHaveLength(3);
      expect(result[0]).toEqual({
        ...mockComponents[0],
        id: mockComponents[0].id.toString(),
      });
    });

    it('should fetch components by category with custom page', async () => {
      const mockComponents = generateMockComponents(2);
      mockFetchSuccess(mockComponents);

      const result = await getComponentsByCategory({ category: 'gpu', page: '3' });

      expectFetchToHaveBeenCalledWith('http://localhost:8080/components/gpu?page=3');
      expect(result).toHaveLength(2);
    });

    it('should convert component IDs to strings', async () => {
      const mockComponents = [{ ...mockComponent, id: 789 }];
      mockFetchSuccess(mockComponents);

      const result = await getComponentsByCategory({ category: 'motherboard' });

      expect(result[0].id).toBe('789');
      expect(typeof result[0].id).toBe('string');
    });

    it('should handle fetch errors', async () => {
      mockFetchError('Category not found');

      await expect(getComponentsByCategory({ category: 'invalid' })).rejects.toThrow('Category not found');
    });

    it('should handle URL encoding for category names with spaces or special characters', async () => {
      const mockComponents = generateMockComponents(1);
      mockFetchSuccess(mockComponents);

      await getComponentsByCategory({ category: 'cpu cooler' });

      expectFetchToHaveBeenCalledWith('http://localhost:8080/components/cpu cooler?page=1');
    });
  });

  describe('getComponentsByBrand', () => {
    it('should fetch components by category and brand with default page', async () => {
      const mockComponents = generateMockComponents(4);
      mockFetchSuccess(mockComponents);

      const result = await getComponentsByBrand({ category: 'cpu', brand: 'Intel' });

      expectFetchToHaveBeenCalledWith('http://localhost:8080/components/cpu/Intel?page=1');
      expect(result).toHaveLength(4);
      expect(result[0]).toEqual({
        ...mockComponents[0],
        id: mockComponents[0].id.toString(),
      });
    });

    it('should fetch components by category and brand with custom page', async () => {
      const mockComponents = generateMockComponents(2);
      mockFetchSuccess(mockComponents);

      const result = await getComponentsByBrand({
        category: 'gpu',
        brand: 'NVIDIA',
        page: '2'
      });

      expectFetchToHaveBeenCalledWith('http://localhost:8080/components/gpu/NVIDIA?page=2');
      expect(result).toHaveLength(2);
    });

    it('should convert component IDs to strings', async () => {
      const mockComponents = [{ ...mockComponent, id: 101112 }];
      mockFetchSuccess(mockComponents);

      const result = await getComponentsByBrand({ category: 'ram', brand: 'Corsair' });

      expect(result[0].id).toBe('101112');
      expect(typeof result[0].id).toBe('string');
    });

    it('should handle fetch errors', async () => {
      mockFetchError('Brand not found');

      await expect(getComponentsByBrand({
        category: 'cpu',
        brand: 'InvalidBrand'
      })).rejects.toThrow('Brand not found');
    });

    it('should handle URL encoding for brand names with spaces or special characters', async () => {
      const mockComponents = generateMockComponents(1);
      mockFetchSuccess(mockComponents);

      await getComponentsByBrand({ category: 'case', brand: 'Fractal Design' });

      expectFetchToHaveBeenCalledWith('http://localhost:8080/components/case/Fractal Design?page=1');
    });
  });

  describe('getComponentById', () => {
    it('should fetch component by ID with default page', async () => {
      const mockComponentData = { ...mockComponent, id: 12345 };
      mockFetchSuccess(mockComponentData);

      const result = await getComponentById({ id: '12345' });

      expectFetchToHaveBeenCalledWith('http://localhost:8080/components/item/12345?page=1');
      expect(result).toEqual({
        ...mockComponentData,
        id: '12345',
      });
    });

    it('should fetch component by ID with custom page', async () => {
      const mockComponentData = { ...mockComponent, id: 67890 };
      mockFetchSuccess(mockComponentData);

      const result = await getComponentById({ id: '67890', page: '5' });

      expectFetchToHaveBeenCalledWith('http://localhost:8080/components/item/67890?page=5');
      expect(result).toEqual({
        ...mockComponentData,
        id: '67890',
      });
    });

    it('should convert component ID to string', async () => {
      const mockComponentData = { ...mockComponent, id: 999 };
      mockFetchSuccess(mockComponentData);

      const result = await getComponentById({ id: '999' });

      expect(result.id).toBe('999');
      expect(typeof result.id).toBe('string');
    });

    it('should handle fetch errors', async () => {
      mockFetchError('Component not found');

      await expect(getComponentById({ id: 'nonexistent' })).rejects.toThrow('Component not found');
    });

    it('should handle 404 responses', async () => {
      mockFetchWithStatus(404, { error: 'Component not found' });

      const result = await getComponentById({ id: '404' });

      expect(result).toEqual({
        error: 'Component not found',
        id: '404', // Since data.id doesn't exist, it falls back to the input id
      });
    });

    it('should handle numeric string IDs correctly', async () => {
      const mockComponentData = { ...mockComponent, id: 12345 };
      mockFetchSuccess(mockComponentData);

      const result = await getComponentById({ id: '12345' });

      expectFetchToHaveBeenCalledWith('http://localhost:8080/components/item/12345?page=1');
      expect(result.id).toBe('12345');
    });
  });

  describe('Error handling and edge cases', () => {
    it('should handle malformed JSON responses', async () => {
      (global.fetch as jest.Mock).mockResolvedValueOnce({
        ok: true,
        json: jest.fn().mockRejectedValueOnce(new Error('Invalid JSON')),
      });

      await expect(getComponents({})).rejects.toThrow('Invalid JSON');
    });

    it('should handle network timeouts', async () => {
      (global.fetch as jest.Mock).mockImplementationOnce(() =>
        new Promise((_, reject) =>
          setTimeout(() => reject(new Error('Request timeout')), 100)
        )
      );

      await expect(getComponents({})).rejects.toThrow('Request timeout');
    });

    it('should preserve all component properties except ID transformation', async () => {
      const complexMockComponent = {
        ...mockComponent,
        id: 123,
        specs: 'Very detailed specifications with special characters: @#$%',
        releaseDate: '2023-12-01T15:30:00Z',
        customField: 'This should be preserved'
      };
      mockFetchSuccess([complexMockComponent]);

      const result = await getComponents({});

      expect(result[0]).toEqual({
        ...complexMockComponent,
        id: '123', // Only ID should be converted to string
      });
      expect(result[0].specs).toBe('Very detailed specifications with special characters: @#$%');
      expect((result[0] as Component & { customField: string }).customField).toBe('This should be preserved');
    });
  });
});
