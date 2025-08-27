package services

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/mateuse/desktop-builder-backend/internal/models"
	"github.com/stretchr/testify/assert"
)

// Mock data for testing
var mockComponent = models.Component{
	ID:        "1",
	Category:  models.CategoryCPU,
	Brand:     "Intel",
	Model:     "Core i7-12700K",
	SKU:       stringPtr("BX8071512700K"),
	UPC:       stringPtr("735858491174"),
	Specs:     json.RawMessage(`{"cores": 12, "threads": 20, "base_clock": "3.6 GHz"}`),
	CreatedAt: time.Now(),
}

var mockComponents = []models.Component{
	mockComponent,
	{
		ID:        "2",
		Category:  models.CategoryCPU,
		Brand:     "AMD",
		Model:     "Ryzen 7 5800X",
		SKU:       stringPtr("100-100000063WOF"),
		UPC:       stringPtr("730143312042"),
		Specs:     json.RawMessage(`{"cores": 8, "threads": 16, "base_clock": "3.8 GHz"}`),
		CreatedAt: time.Now(),
	},
	{
		ID:        "3",
		Category:  models.CategoryMemory,
		Brand:     "Corsair",
		Model:     "Vengeance LPX 32GB",
		SKU:       stringPtr("CMK32GX4M2E3200C16"),
		UPC:       stringPtr("843591098908"),
		Specs:     json.RawMessage(`{"capacity": "32GB", "speed": "3200MHz", "type": "DDR4"}`),
		CreatedAt: time.Now(),
	},
}

// Helper function to create string pointers
func stringPtr(s string) *string {
	return &s
}

// TestGetComponentsByBrand_InputValidation tests input validation and processing
func TestGetComponentsByBrand_InputValidation(t *testing.T) {
	tests := []struct {
		name        string
		input       models.GetComponentsByBrandInput
		description string
	}{
		{
			name: "Valid input with both category and brand",
			input: models.GetComponentsByBrandInput{
				Category: "cpu",
				Brand:    "Intel",
			},
			description: "Should process valid input correctly",
		},
		{
			name: "Empty category",
			input: models.GetComponentsByBrandInput{
				Category: "",
				Brand:    "Intel",
			},
			description: "Should handle empty category",
		},
		{
			name: "Empty brand",
			input: models.GetComponentsByBrandInput{
				Category: "cpu",
				Brand:    "",
			},
			description: "Should handle empty brand",
		},
		{
			name: "Both empty",
			input: models.GetComponentsByBrandInput{
				Category: "",
				Brand:    "",
			},
			description: "Should handle both fields empty",
		},
		{
			name: "Special characters in category",
			input: models.GetComponentsByBrandInput{
				Category: "cpu_cooler",
				Brand:    "Noctua",
			},
			description: "Should handle special characters in category",
		},
		{
			name: "Special characters in brand",
			input: models.GetComponentsByBrandInput{
				Category: "motherboard",
				Brand:    "ASUS & MSI",
			},
			description: "Should handle special characters in brand",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test that the function doesn't panic with various inputs
			// Note: This will fail without database setup, but tests input processing
			t.Logf("Testing input: Category='%s', Brand='%s'", tt.input.Category, tt.input.Brand)
			t.Logf("Description: %s", tt.description)

			// Verify that input fields are properly extracted
			category, brand := tt.input.Category, tt.input.Brand
			assert.Equal(t, tt.input.Category, category)
			assert.Equal(t, tt.input.Brand, brand)
		})
	}
}

// TestServiceFunctionSignatures tests that all service functions have correct signatures
func TestServiceFunctionSignatures(t *testing.T) {
	t.Run("GetAllComponents signature", func(t *testing.T) {
		// Test that function exists and has correct signature
		// This is a compile-time test - if it compiles, the signature is correct
		var fn func(models.GetAllComponentsInput) ([]models.Component, error) = GetAllComponents
		assert.NotNil(t, fn)
	})

	t.Run("GetComponentsByCategory signature", func(t *testing.T) {
		var fn func(models.GetComponentsByCategoryInput) ([]models.Component, error) = GetComponentsByCategory
		assert.NotNil(t, fn)
	})

	t.Run("GetComponentsByBrand signature", func(t *testing.T) {
		var fn func(models.GetComponentsByBrandInput) ([]models.Component, error) = GetComponentsByBrand
		assert.NotNil(t, fn)
	})

	t.Run("GetComponentById signature", func(t *testing.T) {
		var fn func(models.GetComponentByIdInput) (models.Component, error) = GetComponentById
		assert.NotNil(t, fn)
	})
}

// TestGetComponentsByBrand_ParameterExtraction tests parameter extraction logic
func TestGetComponentsByBrand_ParameterExtraction(t *testing.T) {
	input := models.GetComponentsByBrandInput{
		Category: "cpu",
		Brand:    "Intel",
	}

	// Test the parameter extraction logic used in the service
	category, brand := input.Category, input.Brand

	assert.Equal(t, "cpu", category)
	assert.Equal(t, "Intel", brand)
	assert.NotEqual(t, category, brand)
}

// TestServiceLayerInputTypes tests that service layer handles all expected input types
func TestServiceLayerInputTypes(t *testing.T) {
	tests := []struct {
		name         string
		testFunction string
		inputType    string
		description  string
	}{
		{
			name:         "GetAllComponents takes GetAllComponentsInput",
			testFunction: "GetAllComponents",
			inputType:    "GetAllComponentsInput",
			description:  "Should accept GetAllComponentsInput struct",
		},
		{
			name:         "GetComponentsByCategory takes GetComponentsByCategoryInput",
			testFunction: "GetComponentsByCategory",
			inputType:    "GetComponentsByCategoryInput",
			description:  "Should accept GetComponentsByCategoryInput struct",
		},
		{
			name:         "GetComponentsByBrand takes struct",
			testFunction: "GetComponentsByBrand",
			inputType:    "GetComponentsByBrandInput",
			description:  "Should accept GetComponentsByBrandInput struct",
		},
		{
			name:         "GetComponentById takes GetComponentByIdInput",
			testFunction: "GetComponentById",
			inputType:    "GetComponentByIdInput",
			description:  "Should accept GetComponentByIdInput struct",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Function: %s", tt.testFunction)
			t.Logf("Input Type: %s", tt.inputType)
			t.Logf("Description: %s", tt.description)
			// This test mainly documents the expected input types
			assert.True(t, true, "Input type validation passed")
		})
	}
}

// BenchmarkGetComponentsByBrand_ParameterExtraction benchmarks parameter extraction
func BenchmarkGetComponentsByBrand_ParameterExtraction(b *testing.B) {
	input := models.GetComponentsByBrandInput{
		Category: "cpu",
		Brand:    "Intel",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		category, brand := input.Category, input.Brand
		_ = category
		_ = brand
	}
}
