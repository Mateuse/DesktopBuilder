// Base Schema Definitions
// Common schemas and utilities used across the application
import { z } from 'zod';

// Common field schemas that can be reused
export const IdSchema = z.number().int().positive();
export const TimestampSchema = z.iso.datetime();
export const OptionalStringSchema = z.string().nullable().optional();
export const OptionalUrlSchema = z.url().nullable().optional();
export const CurrencySchema = z.string().min(3).max(3); // ISO currency codes
export const RegionSchema = z.string().min(2); // Region/country codes
export const UserIdSchema = z.string().min(1);

// Common validation patterns
export const PositivePriceSchema = z.number().min(0);
export const QuantitySchema = z.number().int().min(1);
export const JsonObjectSchema = z.record(z.string(), z.any());

// Base entity schema with common fields
export const BaseEntitySchema = z.object({
  id: IdSchema,
  created_at: TimestampSchema,
});

export const UpdatableEntitySchema = BaseEntitySchema.extend({
  updated_at: TimestampSchema,
});

// Export types for base schemas
export type Id = z.infer<typeof IdSchema>;
export type Timestamp = z.infer<typeof TimestampSchema>;
export type Currency = z.infer<typeof CurrencySchema>;
export type Region = z.infer<typeof RegionSchema>;
export type UserId = z.infer<typeof UserIdSchema>;
export type PositivePrice = z.infer<typeof PositivePriceSchema>;
export type Quantity = z.infer<typeof QuantitySchema>;
export type JsonObject = z.infer<typeof JsonObjectSchema>;
export type BaseEntity = z.infer<typeof BaseEntitySchema>;
export type UpdatableEntity = z.infer<typeof UpdatableEntitySchema>;
