// Health API Schemas
import { z } from 'zod';

// Health check response schema
export const HealthResponseSchema = z.object({
  message: z.string(),
});

// Export types
export type HealthResponse = z.infer<typeof HealthResponseSchema>;

// Export all health schemas
export const HealthSchemas = {
  HealthResponse: HealthResponseSchema,
} as const;
