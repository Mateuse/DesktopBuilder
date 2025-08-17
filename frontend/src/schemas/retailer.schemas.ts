// Retailer Schemas
import { z } from 'zod';
import { UpdatableEntitySchema, OptionalUrlSchema, JsonObjectSchema } from './base.schemas';

// Core retailer schema
export const RetailerSchema = UpdatableEntitySchema.extend({
  name: z.string(),
  website_url: OptionalUrlSchema,
  logo_url: OptionalUrlSchema,
  shipping_info: JsonObjectSchema.optional(),
  return_policy: JsonObjectSchema.optional(),
  is_active: z.boolean(),
});

// Retailer creation schema (for API requests)
export const RetailerCreateSchema = z.object({
  name: z.string().min(1),
  website_url: OptionalUrlSchema,
  logo_url: OptionalUrlSchema,
  shipping_info: JsonObjectSchema.optional(),
  return_policy: JsonObjectSchema.optional(),
  is_active: z.boolean().optional(),
});

// Retailer update schema (for API requests)
export const RetailerUpdateSchema = z.object({
  name: z.string().min(1).optional(),
  website_url: OptionalUrlSchema,
  logo_url: OptionalUrlSchema,
  shipping_info: JsonObjectSchema.optional(),
  return_policy: JsonObjectSchema.optional(),
  is_active: z.boolean().optional(),
});

// Retailer filter schema
export const RetailerFilterSchema = z.object({
  name: z.string().optional(),
  is_active: z.boolean().optional(),
});

// Export types
export type Retailer = z.infer<typeof RetailerSchema>;
export type RetailerCreate = z.infer<typeof RetailerCreateSchema>;
export type RetailerUpdate = z.infer<typeof RetailerUpdateSchema>;
export type RetailerFilter = z.infer<typeof RetailerFilterSchema>;

// Export all retailer schemas
export const RetailerSchemas = {
  Retailer: RetailerSchema,
  RetailerCreate: RetailerCreateSchema,
  RetailerUpdate: RetailerUpdateSchema,
  RetailerFilter: RetailerFilterSchema,
} as const;
