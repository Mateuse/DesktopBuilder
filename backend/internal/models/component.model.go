package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

// Category represents the component category enum
type Category string

const (
	CategoryCPU           Category = "cpu"
	CategoryMotherboard   Category = "motherboard"
	CategoryMemory        Category = "memory"
	CategoryStorage       Category = "storage"
	CategoryGPU           Category = "gpu"
	CategoryPowerSupply   Category = "powersupply"
	CategoryCase          Category = "case"
	CategoryCooler        Category = "cooler"
	CategoryMonitor       Category = "monitor"
	CategoryExpansionCard Category = "expansioncard"
	CategoryPeripherals   Category = "peripherals"
	CategoryOther         Category = "other"
)

// Valid returns true if the category is valid
func (c Category) Valid() bool {
	switch c {
	case CategoryCPU, CategoryMotherboard, CategoryMemory, CategoryStorage,
		CategoryGPU, CategoryPowerSupply, CategoryCase, CategoryCooler,
		CategoryMonitor, CategoryExpansionCard, CategoryPeripherals, CategoryOther:
		return true
	}
	return false
}

// Value implements the driver.Valuer interface for database storage
func (c Category) Value() (driver.Value, error) {
	if !c.Valid() {
		return nil, fmt.Errorf("invalid category: %s", c)
	}
	return string(c), nil
}

// Scan implements the sql.Scanner interface for database retrieval
func (c *Category) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	switch v := value.(type) {
	case string:
		*c = Category(v)
	case []byte:
		*c = Category(v)
	default:
		return fmt.Errorf("cannot scan %T into Category", value)
	}

	if !c.Valid() {
		return fmt.Errorf("invalid category value: %s", *c)
	}

	return nil
}

// Component represents the components table
type Component struct {
	ID        int64           `json:"id" db:"id"`
	Category  Category        `json:"category" db:"category"`
	Brand     string          `json:"brand" db:"brand"`
	Model     string          `json:"model" db:"model"`
	SKU       *string         `json:"sku,omitempty" db:"sku"`
	UPC       *string         `json:"upc,omitempty" db:"upc"`
	Specs     json.RawMessage `json:"specs" db:"specs"`
	CreatedAt time.Time       `json:"created_at" db:"created_at"`
}

// ComponentCreate represents the data needed to create a new component
type ComponentCreate struct {
	Category Category        `json:"category" validate:"required"`
	Brand    string          `json:"brand" validate:"required"`
	Model    string          `json:"model" validate:"required"`
	SKU      *string         `json:"sku,omitempty"`
	UPC      *string         `json:"upc,omitempty"`
	Specs    json.RawMessage `json:"specs" validate:"required"`
}

// ComponentUpdate represents the data that can be updated for a component
type ComponentUpdate struct {
	Brand *string          `json:"brand,omitempty"`
	Model *string          `json:"model,omitempty"`
	SKU   *string          `json:"sku,omitempty"`
	UPC   *string          `json:"upc,omitempty"`
	Specs *json.RawMessage `json:"specs,omitempty"`
}
