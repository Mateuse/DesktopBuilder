// Price Schemas
import { z } from 'zod';
import { BaseEntitySchema, IdSchema, CurrencySchema, RegionSchema, PositivePriceSchema, TimestampSchema, OptionalUrlSchema } from './base.schemas';
import { ComponentSchema } from './component.schemas';
import { RetailerSchema } from './retailer.schemas';

// Core price schema
export const PriceSchema = BaseEntitySchema.extend({
  component_id: IdSchema,
  retailer_id: IdSchema,
  region: RegionSchema,
  currency: CurrencySchema,
  price: PositivePriceSchema,
  in_stock: z.boolean(),
  product_url: OptionalUrlSchema,
  last_updated: TimestampSchema,
});

// Price with related entity details
export const PriceWithDetailsSchema = PriceSchema.extend({
  component: ComponentSchema.optional(),
  retailer: RetailerSchema.optional(),
});

// Price creation schema (for API requests)
export const PriceCreateSchema = z.object({
  component_id: IdSchema,
  retailer_id: IdSchema,
  region: RegionSchema,
  currency: CurrencySchema,
  price: PositivePriceSchema,
  in_stock: z.boolean(),
  product_url: OptionalUrlSchema,
});

// Price update schema (for API requests)
export const PriceUpdateSchema = z.object({
  region: RegionSchema.optional(),
  currency: CurrencySchema.optional(),
  price: PositivePriceSchema.optional(),
  in_stock: z.boolean().optional(),
  product_url: OptionalUrlSchema,
});

// Price filter schema
export const PriceFilterSchema = z.object({
  component_id: IdSchema.optional(),
  retailer_id: IdSchema.optional(),
  region: RegionSchema.optional(),
  currency: CurrencySchema.optional(),
  in_stock: z.boolean().optional(),
  min_price: PositivePriceSchema.optional(),
  max_price: PositivePriceSchema.optional(),
});

// Export types
export type Price = z.infer<typeof PriceSchema>;
export type PriceWithDetails = z.infer<typeof PriceWithDetailsSchema>;
export type PriceCreate = z.infer<typeof PriceCreateSchema>;
export type PriceUpdate = z.infer<typeof PriceUpdateSchema>;
export type PriceFilter = z.infer<typeof PriceFilterSchema>;

// Export all price schemas
export const PriceSchemas = {
  Price: PriceSchema,
  PriceWithDetails: PriceWithDetailsSchema,
  PriceCreate: PriceCreateSchema,
  PriceUpdate: PriceUpdateSchema,
  PriceFilter: PriceFilterSchema,
} as const;
