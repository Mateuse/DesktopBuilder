// Build and Build Component Schemas
import { z } from 'zod';
import { UpdatableEntitySchema, IdSchema, UserIdSchema, CurrencySchema, RegionSchema, PositivePriceSchema, QuantitySchema, OptionalStringSchema } from './base.schemas';
import { ComponentSchema } from './component.schemas';
import { PriceSchema } from './price.schemas';

// User Build schemas
export const UserBuildSchema = UpdatableEntitySchema.extend({
  user_id: UserIdSchema,
  name: z.string(),
  description: OptionalStringSchema,
  is_public: z.boolean(),
  is_complete: z.boolean(),
  total_price: PositivePriceSchema.nullable().optional(),
  currency: CurrencySchema,
  region: RegionSchema,
});

export const UserBuildCreateSchema = z.object({
  user_id: UserIdSchema,
  name: z.string().min(1),
  description: OptionalStringSchema,
  is_public: z.boolean().optional(),
  is_complete: z.boolean().optional(),
  currency: CurrencySchema.optional(),
  region: RegionSchema.optional(),
});

export const UserBuildUpdateSchema = z.object({
  name: z.string().min(1).optional(),
  description: OptionalStringSchema,
  is_public: z.boolean().optional(),
  is_complete: z.boolean().optional(),
  total_price: PositivePriceSchema.nullable().optional(),
  currency: CurrencySchema.optional(),
  region: RegionSchema.optional(),
});

export const UserBuildFilterSchema = z.object({
  user_id: UserIdSchema.optional(),
  is_public: z.boolean().optional(),
  is_complete: z.boolean().optional(),
  currency: CurrencySchema.optional(),
  region: RegionSchema.optional(),
});

// Build Component schemas
export const BuildComponentSchema = z.object({
  id: IdSchema,
  build_id: IdSchema,
  component_id: IdSchema,
  quantity: QuantitySchema,
  selected_price_id: IdSchema.nullable().optional(),
  notes: OptionalStringSchema,
  created_at: z.iso.datetime(),
});

export const BuildComponentWithDetailsSchema = BuildComponentSchema.extend({
  component: ComponentSchema.optional(),
  selected_price: PriceSchema.optional(),
});

export const BuildComponentCreateSchema = z.object({
  build_id: IdSchema,
  component_id: IdSchema,
  quantity: QuantitySchema.optional(),
  selected_price_id: IdSchema.nullable().optional(),
  notes: OptionalStringSchema,
});

export const BuildComponentUpdateSchema = z.object({
  quantity: QuantitySchema.optional(),
  selected_price_id: IdSchema.nullable().optional(),
  notes: OptionalStringSchema,
});

export const BuildComponentFilterSchema = z.object({
  build_id: IdSchema.optional(),
  component_id: IdSchema.optional(),
});

// Composite schemas
export const UserBuildWithComponentsSchema = UserBuildSchema.extend({
  components: z.array(BuildComponentWithDetailsSchema).optional(),
});

// Export types
export type UserBuild = z.infer<typeof UserBuildSchema>;
export type UserBuildCreate = z.infer<typeof UserBuildCreateSchema>;
export type UserBuildUpdate = z.infer<typeof UserBuildUpdateSchema>;
export type UserBuildFilter = z.infer<typeof UserBuildFilterSchema>;
export type BuildComponent = z.infer<typeof BuildComponentSchema>;
export type BuildComponentWithDetails = z.infer<typeof BuildComponentWithDetailsSchema>;
export type BuildComponentCreate = z.infer<typeof BuildComponentCreateSchema>;
export type BuildComponentUpdate = z.infer<typeof BuildComponentUpdateSchema>;
export type BuildComponentFilter = z.infer<typeof BuildComponentFilterSchema>;
export type UserBuildWithComponents = z.infer<typeof UserBuildWithComponentsSchema>;

// Export all build schemas
export const BuildSchemas = {
  // User Builds
  UserBuild: UserBuildSchema,
  UserBuildCreate: UserBuildCreateSchema,
  UserBuildUpdate: UserBuildUpdateSchema,
  UserBuildFilter: UserBuildFilterSchema,

  // Build Components
  BuildComponent: BuildComponentSchema,
  BuildComponentWithDetails: BuildComponentWithDetailsSchema,
  BuildComponentCreate: BuildComponentCreateSchema,
  BuildComponentUpdate: BuildComponentUpdateSchema,
  BuildComponentFilter: BuildComponentFilterSchema,

  // Composite types
  UserBuildWithComponents: UserBuildWithComponentsSchema,
} as const;
