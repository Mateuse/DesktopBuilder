package models

import (
	"time"
)

// BuildComponent represents the build_components table
type BuildComponent struct {
	ID              int64     `json:"id" db:"id"`
	BuildID         int64     `json:"build_id" db:"build_id"`
	ComponentID     int64     `json:"component_id" db:"component_id"`
	Quantity        int       `json:"quantity" db:"quantity"`
	SelectedPriceID *int64    `json:"selected_price_id,omitempty" db:"selected_price_id"`
	Notes           *string   `json:"notes,omitempty" db:"notes"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
}

// BuildComponentWithDetails represents a build component with related information
type BuildComponentWithDetails struct {
	BuildComponent
	Component     *Component `json:"component,omitempty"`
	SelectedPrice *Price     `json:"selected_price,omitempty"`
}

// BuildComponentCreate represents the data needed to create a new build component
type BuildComponentCreate struct {
	BuildID         int64   `json:"build_id" validate:"required"`
	ComponentID     int64   `json:"component_id" validate:"required"`
	Quantity        *int    `json:"quantity,omitempty" validate:"omitempty,min=1"`
	SelectedPriceID *int64  `json:"selected_price_id,omitempty"`
	Notes           *string `json:"notes,omitempty"`
}

// BuildComponentUpdate represents the data that can be updated for a build component
type BuildComponentUpdate struct {
	Quantity        *int    `json:"quantity,omitempty" validate:"omitempty,min=1"`
	SelectedPriceID *int64  `json:"selected_price_id,omitempty"`
	Notes           *string `json:"notes,omitempty"`
}

// BuildComponentFilter represents filters for querying build components
type BuildComponentFilter struct {
	BuildID     *int64 `json:"build_id,omitempty"`
	ComponentID *int64 `json:"component_id,omitempty"`
}
