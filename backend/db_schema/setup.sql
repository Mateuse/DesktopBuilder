-- DesktopBuilder Database Setup Script
-- Run this script to create all tables and indexes for first-time setup

-- Create custom types
CREATE TYPE category AS ENUM (
  'cpu','motherboard','memory','storage','gpu','powersupply','case','cooler','monitor','expansioncard','peripherals','other'
);

-- Create tables
CREATE TABLE components (
  id BIGSERIAL PRIMARY KEY,
  category category NOT NULL,
  brand TEXT NOT NULL,
  model TEXT NOT NULL,
  sku TEXT,
  upc TEXT,
  specs JSONB NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  UNIQUE(sku),
  UNIQUE(upc)
);

CREATE TABLE retailers (
  id BIGSERIAL PRIMARY KEY,
  name TEXT NOT NULL UNIQUE,
  website_url TEXT,
  logo_url TEXT,
  shipping_info JSONB,
  return_policy JSONB,
  is_active BOOLEAN NOT NULL DEFAULT true,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE prices (
  id BIGSERIAL PRIMARY KEY,
  component_id BIGINT NOT NULL REFERENCES components(id) ON DELETE CASCADE,
  retailer_id BIGINT NOT NULL REFERENCES retailers(id) ON DELETE CASCADE,
  region TEXT NOT NULL,
  currency TEXT NOT NULL,
  price DECIMAL(10,2) NOT NULL,
  in_stock BOOLEAN NOT NULL DEFAULT false,
  product_url TEXT,
  last_updated TIMESTAMPTZ NOT NULL DEFAULT now(),
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE user_builds (
  id BIGSERIAL PRIMARY KEY,
  user_id TEXT NOT NULL,
  name TEXT NOT NULL,
  description TEXT,
  is_public BOOLEAN NOT NULL DEFAULT false,
  is_complete BOOLEAN NOT NULL DEFAULT false,
  total_price DECIMAL(10,2),
  currency TEXT DEFAULT 'USD',
  region TEXT DEFAULT 'USA',
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE build_components (
  id BIGSERIAL PRIMARY KEY,
  build_id BIGINT NOT NULL REFERENCES user_builds(id) ON DELETE CASCADE,
  component_id BIGINT NOT NULL REFERENCES components(id) ON DELETE CASCADE,
  quantity INTEGER NOT NULL DEFAULT 1,
  selected_price_id BIGINT REFERENCES prices(id),
  notes TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  UNIQUE(build_id, component_id)
);

-- Create indexes
CREATE INDEX idx_components_category ON components(category);
CREATE INDEX idx_components_specs_gin ON components USING GIN (specs);

CREATE INDEX idx_prices_component_retailer_region ON prices(component_id, retailer_id, region);
CREATE INDEX idx_prices_last_updated ON prices(last_updated);
CREATE INDEX idx_prices_in_stock ON prices(in_stock);

CREATE INDEX idx_user_builds_user_id ON user_builds(user_id);
CREATE INDEX idx_user_builds_public ON user_builds(is_public);
CREATE INDEX idx_user_builds_complete ON user_builds(is_complete);

CREATE INDEX idx_build_components_build_id ON build_components(build_id);
