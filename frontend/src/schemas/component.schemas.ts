// Component Schemas
import { z } from 'zod';
import { BaseEntitySchema, JsonObjectSchema } from './base.schemas';

// Component category enum schema
export const CategorySchema = z.enum([
  'cpu',
  'motherboard',
  'memory',
  'storage',
  'gpu',
  'powersupply',
  'case',
  'cooler',
  'monitor',
  'expansioncard',
  'peripherals',
  'other'
]);

// Core component schema
export const ComponentSchema = BaseEntitySchema.extend({
  category: CategorySchema,
  brand: z.string(),
  model: z.string(),
  sku: z.string().nullable().optional(),
  upc: z.string().nullable().optional(),
  specs: JsonObjectSchema, // JSON object for component specifications
});

// Component creation schema (for API requests)
export const ComponentCreateSchema = z.object({
  category: CategorySchema,
  brand: z.string().min(1),
  model: z.string().min(1),
  sku: z.string().nullable().optional(),
  upc: z.string().nullable().optional(),
  specs: JsonObjectSchema.optional(),
});

// Component update schema (for API requests)
export const ComponentUpdateSchema = z.object({
  category: CategorySchema.optional(),
  brand: z.string().min(1).optional(),
  model: z.string().min(1).optional(),
  sku: z.string().nullable().optional(),
  upc: z.string().nullable().optional(),
  specs: JsonObjectSchema.optional(),
});

// Component filter schema
export const ComponentFilterSchema = z.object({
  category: CategorySchema.optional(),
  brand: z.string().optional(),
  model: z.string().optional(),
  sku: z.string().optional(),
  upc: z.string().optional(),
});

// Export types
export type Category = z.infer<typeof CategorySchema>;
export type Component = z.infer<typeof ComponentSchema>;
export type ComponentCreate = z.infer<typeof ComponentCreateSchema>;
export type ComponentUpdate = z.infer<typeof ComponentUpdateSchema>;
export type ComponentFilter = z.infer<typeof ComponentFilterSchema>;

// Export all component schemas
export const ComponentSchemas = {
  Category: CategorySchema,
  Component: ComponentSchema,
  ComponentCreate: ComponentCreateSchema,
  ComponentUpdate: ComponentUpdateSchema,
  ComponentFilter: ComponentFilterSchema,
} as const;
