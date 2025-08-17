package models

import (
	"time"
)

// Price represents the prices table
type Price struct {
	ID          int64     `json:"id" db:"id"`
	ComponentID int64     `json:"component_id" db:"component_id"`
	RetailerID  int64     `json:"retailer_id" db:"retailer_id"`
	Region      string    `json:"region" db:"region"`
	Currency    string    `json:"currency" db:"currency"`
	Price       float64   `json:"price" db:"price"`
	InStock     bool      `json:"in_stock" db:"in_stock"`
	ProductURL  *string   `json:"product_url,omitempty" db:"product_url"`
	LastUpdated time.Time `json:"last_updated" db:"last_updated"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// PriceWithDetails represents a price with related component and retailer information
type PriceWithDetails struct {
	Price
	Component *Component `json:"component,omitempty"`
	Retailer  *Retailer  `json:"retailer,omitempty"`
}

// PriceCreate represents the data needed to create a new price entry
type PriceCreate struct {
	ComponentID int64   `json:"component_id" validate:"required"`
	RetailerID  int64   `json:"retailer_id" validate:"required"`
	Region      string  `json:"region" validate:"required"`
	Currency    string  `json:"currency" validate:"required"`
	Price       float64 `json:"price" validate:"required,min=0"`
	InStock     *bool   `json:"in_stock,omitempty"`
	ProductURL  *string `json:"product_url,omitempty"`
}

// PriceUpdate represents the data that can be updated for a price entry
type PriceUpdate struct {
	Price      *float64 `json:"price,omitempty" validate:"omitempty,min=0"`
	InStock    *bool    `json:"in_stock,omitempty"`
	ProductURL *string  `json:"product_url,omitempty"`
}

// PriceFilter represents filters for querying prices
type PriceFilter struct {
	ComponentID *int64   `json:"component_id,omitempty"`
	RetailerID  *int64   `json:"retailer_id,omitempty"`
	Region      *string  `json:"region,omitempty"`
	Currency    *string  `json:"currency,omitempty"`
	InStock     *bool    `json:"in_stock,omitempty"`
	MinPrice    *float64 `json:"min_price,omitempty"`
	MaxPrice    *float64 `json:"max_price,omitempty"`
}
