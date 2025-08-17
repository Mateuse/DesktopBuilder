package models

import (
	"encoding/json"
	"time"
)

// Retailer represents the retailers table
type Retailer struct {
	ID           int64           `json:"id" db:"id"`
	Name         string          `json:"name" db:"name"`
	WebsiteURL   *string         `json:"website_url,omitempty" db:"website_url"`
	LogoURL      *string         `json:"logo_url,omitempty" db:"logo_url"`
	ShippingInfo json.RawMessage `json:"shipping_info,omitempty" db:"shipping_info"`
	ReturnPolicy json.RawMessage `json:"return_policy,omitempty" db:"return_policy"`
	IsActive     bool            `json:"is_active" db:"is_active"`
	CreatedAt    time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at" db:"updated_at"`
}

// RetailerCreate represents the data needed to create a new retailer
type RetailerCreate struct {
	Name         string          `json:"name" validate:"required"`
	WebsiteURL   *string         `json:"website_url,omitempty"`
	LogoURL      *string         `json:"logo_url,omitempty"`
	ShippingInfo json.RawMessage `json:"shipping_info,omitempty"`
	ReturnPolicy json.RawMessage `json:"return_policy,omitempty"`
	IsActive     *bool           `json:"is_active,omitempty"`
}

// RetailerUpdate represents the data that can be updated for a retailer
type RetailerUpdate struct {
	Name         *string          `json:"name,omitempty"`
	WebsiteURL   *string          `json:"website_url,omitempty"`
	LogoURL      *string          `json:"logo_url,omitempty"`
	ShippingInfo *json.RawMessage `json:"shipping_info,omitempty"`
	ReturnPolicy *json.RawMessage `json:"return_policy,omitempty"`
	IsActive     *bool            `json:"is_active,omitempty"`
}
