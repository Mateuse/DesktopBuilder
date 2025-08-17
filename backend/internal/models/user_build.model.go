package models

import (
	"time"
)

// UserBuild represents the user_builds table
type UserBuild struct {
	ID          int64     `json:"id" db:"id"`
	UserID      string    `json:"user_id" db:"user_id"`
	Name        string    `json:"name" db:"name"`
	Description *string   `json:"description,omitempty" db:"description"`
	IsPublic    bool      `json:"is_public" db:"is_public"`
	IsComplete  bool      `json:"is_complete" db:"is_complete"`
	TotalPrice  *float64  `json:"total_price,omitempty" db:"total_price"`
	Currency    string    `json:"currency" db:"currency"`
	Region      string    `json:"region" db:"region"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// UserBuildWithComponents represents a user build with its associated components
type UserBuildWithComponents struct {
	UserBuild
	Components []BuildComponentWithDetails `json:"components,omitempty"`
}

// UserBuildCreate represents the data needed to create a new user build
type UserBuildCreate struct {
	UserID      string  `json:"user_id" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	IsPublic    *bool   `json:"is_public,omitempty"`
	IsComplete  *bool   `json:"is_complete,omitempty"`
	Currency    *string `json:"currency,omitempty"`
	Region      *string `json:"region,omitempty"`
}

// UserBuildUpdate represents the data that can be updated for a user build
type UserBuildUpdate struct {
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	IsPublic    *bool    `json:"is_public,omitempty"`
	IsComplete  *bool    `json:"is_complete,omitempty"`
	TotalPrice  *float64 `json:"total_price,omitempty"`
	Currency    *string  `json:"currency,omitempty"`
	Region      *string  `json:"region,omitempty"`
}

// UserBuildFilter represents filters for querying user builds
type UserBuildFilter struct {
	UserID     *string `json:"user_id,omitempty"`
	IsPublic   *bool   `json:"is_public,omitempty"`
	IsComplete *bool   `json:"is_complete,omitempty"`
	Currency   *string `json:"currency,omitempty"`
	Region     *string `json:"region,omitempty"`
}
