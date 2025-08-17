// Validation utilities using Zod
// This file contains helper functions for common validation patterns
import { z } from 'zod';

/**
 * Validates API response data against a Zod schema
 * @param schema - The Zod schema to validate against
 * @param data - The data to validate
 * @returns The validated and parsed data
 * @throws Error if validation fails
 */
export const validateApiResponse = <T>(schema: z.ZodSchema<T>, data: unknown): T => {
  try {
    return schema.parse(data);
  } catch (error) {
    if (error instanceof z.ZodError) {
      console.error('API Response Validation Error:', error.issues);
      const errorMessages = error.issues.map(e => e.message).join(', ');
      throw new Error(`Invalid API response: ${errorMessages}`);
    }
    throw error;
  }
};

/**
 * Safely validates API response data, returning null on failure
 * @param schema - The Zod schema to validate against
 * @param data - The data to validate
 * @returns The validated data or null if validation fails
 */
export const safeValidateApiResponse = <T>(schema: z.ZodSchema<T>, data: unknown): T | null => {
  const result = schema.safeParse(data);
  if (result.success) {
    return result.data;
  }
  console.error('API Response Validation Error:', result.error.issues);
  return null;
};

/**
 * Creates a fetch wrapper that automatically validates responses
 * @param schema - The Zod schema to validate the response against
 * @returns A function that fetches and validates the response
 */
export const createValidatedFetch = <T>(schema: z.ZodSchema<T>) => {
  return async (url: string, options?: RequestInit): Promise<T> => {
    const response = await fetch(url, options);
    
    if (!response.ok) {
      throw new Error(`HTTP ${response.status}: ${response.statusText}`);
    }
    
    const data: unknown = await response.json();
    return validateApiResponse(schema, data);
  };
};
