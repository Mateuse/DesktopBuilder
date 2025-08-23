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
	CategoryCaseAccessory       Category = "case_accessory"
	CategoryCaseFan             Category = "case_fan"
	CategoryCase                Category = "case"
	CategoryCPUCooler           Category = "cpu_cooler"
	CategoryCPU                 Category = "cpu"
	CategoryExternalHDD         Category = "external_hdd"
	CategoryFanController       Category = "fan_controller"
	CategoryHeadphone           Category = "headphone"
	CategoryInternalHDD         Category = "internal_hdd"
	CategoryKeyboard            Category = "keyboard"
	CategoryMemory              Category = "memory"
	CategoryMonitor             Category = "monitor"
	CategoryMotherboard         Category = "motherboard"
	CategoryMouse               Category = "mouse"
	CategoryOpticalDrive        Category = "optical_drive"
	CategoryOS                  Category = "os"
	CategoryPowerSupply         Category = "power_supply"
	CategoryUPS                 Category = "ups"
	CategorySoundCard           Category = "sound_card"
	CategorySpeaker             Category = "speaker"
	CategoryThermalPaste        Category = "thermal_paste"
	CategoryVideoCard           Category = "video_card"
	CategoryWebcam              Category = "webcam"
	CategoryWiredNetworkCard    Category = "wired_network_card"
	CategoryWirelessNetworkCard Category = "wireless_network_card"
	CategoryWaterCooling        Category = "water_cooling"
	CategoryOther               Category = "other"
)

// Valid returns true if the category is valid
func (c Category) Valid() bool {
	switch c {
	case CategoryCaseAccessory, CategoryCaseFan, CategoryCase, CategoryCPUCooler,
		CategoryCPU, CategoryExternalHDD, CategoryFanController, CategoryHeadphone,
		CategoryInternalHDD, CategoryKeyboard, CategoryMemory, CategoryMonitor,
		CategoryMotherboard, CategoryMouse, CategoryOpticalDrive, CategoryOS,
		CategoryPowerSupply, CategoryUPS, CategorySoundCard, CategorySpeaker,
		CategoryThermalPaste, CategoryVideoCard, CategoryWebcam, CategoryWiredNetworkCard,
		CategoryWirelessNetworkCard, CategoryWaterCooling, CategoryOther:
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
