# Database Schema

## Custom Types

### Category Enum
```sql
CREATE TYPE category AS ENUM (
  'cpu','motherboard','memory','storage','gpu','powersupply','case','cooler','monitor','expansioncard','peripherals','other'
);
```

## Tables

### Components Table
```sql
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

CREATE INDEX idx_components_category ON components(category);
CREATE INDEX idx_components_specs_gin ON components USING GIN (specs);
```

**Design Notes:**
- Each component entry represents a unique product (including variants)
- Different variants of the same product are separate component entries
- `sku` and `upc` are unique across all components for product identification
- `specs` JSONB contains all specifications including variant-specific attributes
- Example: Two RAM speeds = two separate component entries with different SKUs and specs

### Retailers Table
```sql
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
```

### Prices Table
```sql
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

CREATE INDEX idx_prices_component_retailer_region ON prices(component_id, retailer_id, region);
CREATE INDEX idx_prices_last_updated ON prices(last_updated);
CREATE INDEX idx_prices_in_stock ON prices(in_stock);
```

### User Builds Table
```sql
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

CREATE INDEX idx_user_builds_user_id ON user_builds(user_id);
CREATE INDEX idx_user_builds_public ON user_builds(is_public);
CREATE INDEX idx_user_builds_complete ON user_builds(is_complete);
```

### Build Components Table
```sql
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

CREATE INDEX idx_build_components_build_id ON build_components(build_id);
```
