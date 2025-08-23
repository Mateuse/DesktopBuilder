package repository

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/mateuse/desktop-builder-backend/internal/models"
	"github.com/mateuse/desktop-builder-backend/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestGetAllComponents_Integration tests actual data retrieval
func TestGetAllComponents_Integration(t *testing.T) {
	db := testutils.SetupTestDB(t)
	defer testutils.CleanupTestDB(t, db)

	// Insert known test data
	testComponents := testutils.CreateTestComponents()
	insertedIDs := testutils.InsertTestComponents(t, db, testComponents)

	// Test the actual repository function
	result, err := GetAllComponents(models.GetAllComponentsInput{
		Page: "1",
	})

	// Verify no errors
	require.NoError(t, err, "GetAllComponents should not return an error")

	// Verify we got results
	assert.NotEmpty(t, result, "GetAllComponents should return components")
	assert.GreaterOrEqual(t, len(result), len(testComponents), "Should return at least the test components")

	// Verify our test data is in the results
	testutils.VerifyComponentsExist(t, db, insertedIDs)

	// Verify data structure integrity
	for _, component := range result {
		assert.NotZero(t, component.ID, "Component ID should not be zero")
		assert.NotEmpty(t, component.Category, "Component category should not be empty")
		assert.NotEmpty(t, component.Brand, "Component brand should not be empty")
		assert.NotEmpty(t, component.Model, "Component model should not be empty")
		assert.NotZero(t, component.CreatedAt, "Component created_at should not be zero")
		assert.True(t, component.Category.Valid(), "Component category should be valid")
	}

	t.Logf("Successfully retrieved %d components", len(result))
}

// TestGetComponentsByCategory_Integration tests category filtering
func TestGetComponentsByCategory_Integration(t *testing.T) {
	db := testutils.SetupTestDB(t)
	defer testutils.CleanupTestDB(t, db)

	// Insert test data with different categories
	testComponents := testutils.CreateTestComponents()
	testutils.InsertTestComponents(t, db, testComponents)

	// Count how many CPU components we inserted
	expectedCPUCount := 0
	for _, comp := range testComponents {
		if comp.Category == models.CategoryCPU {
			expectedCPUCount++
		}
	}

	// Test filtering by CPU category
	cpuComponents, err := GetComponentsByCategory(models.GetComponentsByCategoryInput{
		Category: "cpu",
	})

	// Verify no errors
	require.NoError(t, err, "GetComponentsByCategory should not return an error")

	// Verify we got CPU components
	assert.NotEmpty(t, cpuComponents, "Should return CPU components")
	assert.GreaterOrEqual(t, len(cpuComponents), expectedCPUCount, "Should return at least our test CPU components")

	// Verify ALL returned components are CPUs
	for _, component := range cpuComponents {
		assert.Equal(t, models.CategoryCPU, component.Category, "All returned components should be CPUs")
		assert.NotEmpty(t, component.Brand, "CPU component should have a brand")
		assert.NotEmpty(t, component.Model, "CPU component should have a model")
	}

	// Test with a category that shouldn't exist
	emptyResults, err := GetComponentsByCategory(models.GetComponentsByCategoryInput{
		Category: "nonexistent_category",
	})
	require.NoError(t, err, "Should not error on nonexistent category")

	// Should return empty slice, not nil
	assert.NotNil(t, emptyResults, "Should return empty slice, not nil")
	assert.Empty(t, emptyResults, "Should return no components for nonexistent category")

	t.Logf("Successfully filtered %d CPU components", len(cpuComponents))
}

// TestGetComponentsByBrand_Integration tests brand and category filtering
func TestGetComponentsByBrand_Integration(t *testing.T) {
	db := testutils.SetupTestDB(t)
	defer testutils.CleanupTestDB(t, db)

	// Insert test data
	testComponents := testutils.CreateTestComponents()
	testutils.InsertTestComponents(t, db, testComponents)

	// Test filtering by CPU + Intel brand
	intelCPUs, err := GetComponentsByBrand(models.GetComponentsByBrandInput{
		Category: "cpu",
		Brand:    "Test Intel",
	})

	// Verify no errors
	require.NoError(t, err, "GetComponentsByBrand should not return an error")

	// Verify we got results
	assert.NotEmpty(t, intelCPUs, "Should return Intel CPU components")

	// Verify ALL returned components match both filters
	for _, component := range intelCPUs {
		assert.Equal(t, models.CategoryCPU, component.Category, "All components should be CPUs")
		assert.Equal(t, "Test Intel", component.Brand, "All components should be Intel brand")
		assert.NotEmpty(t, component.Model, "Component should have a model")
	}

	// Test with brand that exists but wrong category
	wrongCategoryResults, err := GetComponentsByBrand(models.GetComponentsByBrandInput{
		Category: "memory",
		Brand:    "Test Intel",
	})
	require.NoError(t, err, "Should not error on valid brand with wrong category")
	assert.Empty(t, wrongCategoryResults, "Should return no results for Intel memory")

	// Test with nonexistent brand
	nonexistentResults, err := GetComponentsByBrand(models.GetComponentsByBrandInput{
		Category: "cpu",
		Brand:    "Nonexistent Brand",
	})
	require.NoError(t, err, "Should not error on nonexistent brand")
	assert.Empty(t, nonexistentResults, "Should return no results for nonexistent brand")

	t.Logf("Successfully filtered %d Intel CPU components", len(intelCPUs))
}

// TestGetComponentById_Integration tests single component retrieval
func TestGetComponentById_Integration(t *testing.T) {
	db := testutils.SetupTestDB(t)
	defer testutils.CleanupTestDB(t, db)

	// Insert test data
	testComponents := testutils.CreateTestComponents()
	insertedIDs := testutils.InsertTestComponents(t, db, testComponents)

	// Test retrieving a specific component
	targetID := insertedIDs[0]
	component, err := GetComponentById(models.GetComponentByIdInput{
		ID:   fmt.Sprintf("%d", targetID),
		Page: "1",
	})

	// Verify no errors
	require.NoError(t, err, "GetComponentById should not return an error")

	// Verify we got the correct component
	assert.Equal(t, targetID, component.ID, "Should return component with correct ID")
	assert.NotEmpty(t, component.Category, "Component should have a category")
	assert.NotEmpty(t, component.Brand, "Component should have a brand")
	assert.NotEmpty(t, component.Model, "Component should have a model")
	assert.NotZero(t, component.CreatedAt, "Component should have a created_at timestamp")
	assert.True(t, component.Category.Valid(), "Component category should be valid")

	// Verify the component data matches what we inserted
	originalComponent := testComponents[0]
	assert.Equal(t, originalComponent.Category, component.Category, "Category should match")
	assert.Equal(t, originalComponent.Brand, component.Brand, "Brand should match")
	assert.Equal(t, originalComponent.Model, component.Model, "Model should match")

	// Test retrieving nonexistent component
	_, err = GetComponentById(models.GetComponentByIdInput{
		ID:   "99999",
		Page: "1",
	})
	assert.Error(t, err, "Should return error for nonexistent component")

	// Test with invalid ID format
	_, err = GetComponentById(models.GetComponentByIdInput{
		ID:   "invalid_id",
		Page: "1",
	})
	assert.Error(t, err, "Should return error for invalid ID format")

	t.Logf("Successfully retrieved component with ID: %d", targetID)
}

// TestComponentDataIntegrity_Integration tests data integrity and JSON handling
func TestComponentDataIntegrity_Integration(t *testing.T) {
	db := testutils.SetupTestDB(t)
	defer testutils.CleanupTestDB(t, db)

	// Create a component with complex JSON specs
	complexSpecs := json.RawMessage(`{
		"cores": 12,
		"threads": 20,
		"base_clock": "3.6 GHz",
		"boost_clock": "5.0 GHz",
		"cache": {
			"l1": "768 KB",
			"l2": "12 MB",
			"l3": "25 MB"
		},
		"features": ["AVX-512", "Intel Turbo Boost", "Hyper-Threading"],
		"tdp": 125,
		"socket": "LGA1700"
	}`)

	testComponent := models.Component{
		Category: models.CategoryCPU,
		Brand:    "Test Intel Complex",
		Model:    "Test Core i7-12700K Complex",
		SKU:      testutils.StringPtr("TEST-COMPLEX-BX8071512700K"),
		UPC:      testutils.StringPtr("TEST-COMPLEX-735858491174"),
		Specs:    complexSpecs,
	}

	// Insert the component
	componentID := testutils.InsertTestComponent(t, db, testComponent)

	// Retrieve it back
	retrievedComponent, err := GetComponentById(models.GetComponentByIdInput{
		ID:   fmt.Sprintf("%d", componentID),
		Page: "1",
	})
	require.NoError(t, err, "Should retrieve complex component without error")

	// Verify JSON specs integrity
	assert.JSONEq(t, string(complexSpecs), string(retrievedComponent.Specs), "JSON specs should be preserved exactly")

	// Parse and verify JSON structure
	var specs map[string]interface{}
	err = json.Unmarshal(retrievedComponent.Specs, &specs)
	require.NoError(t, err, "Should be able to parse retrieved JSON specs")

	assert.Equal(t, float64(12), specs["cores"], "Cores should be preserved")
	assert.Equal(t, "3.6 GHz", specs["base_clock"], "Base clock should be preserved")
	assert.Contains(t, specs, "cache", "Cache object should be preserved")
	assert.Contains(t, specs, "features", "Features array should be preserved")

	// Verify optional fields (SKU, UPC)
	require.NotNil(t, retrievedComponent.SKU, "SKU should not be nil")
	require.NotNil(t, retrievedComponent.UPC, "UPC should not be nil")
	assert.Equal(t, *testComponent.SKU, *retrievedComponent.SKU, "SKU should match")
	assert.Equal(t, *testComponent.UPC, *retrievedComponent.UPC, "UPC should match")

	t.Logf("Successfully verified data integrity for component with ID: %d", componentID)
}

// TestRepositoryErrorHandling_Integration tests error scenarios
func TestRepositoryErrorHandling_Integration(t *testing.T) {
	db := testutils.SetupTestDB(t)
	defer testutils.CleanupTestDB(t, db)

	// Test with malformed ID that might cause database errors
	malformedIDs := []string{
		"",                             // empty string
		"abc",                          // non-numeric
		"-1",                           // negative
		"999999999",                    // very large number
		"1.5",                          // decimal
		"1; DROP TABLE components; --", // SQL injection attempt
	}

	for _, malformedID := range malformedIDs {
		t.Run("Malformed ID: "+malformedID, func(t *testing.T) {
			_, err := GetComponentById(models.GetComponentByIdInput{
				ID:   malformedID,
				Page: "1",
			})
			// Should either return an error or handle gracefully
			// The exact behavior depends on your error handling strategy
			t.Logf("GetComponentById('%s') returned error: %v", malformedID, err)
		})
	}

	// Test with special characters in category/brand filters
	specialInputs := []struct {
		category string
		brand    string
	}{
		{"", ""},              // empty strings
		{"cpu'test", "Intel"}, // single quote in category
		{"cpu", "Intel'; DROP TABLE components; --"}, // SQL injection in brand
		{"cpu", "Brand with spaces"},                 // spaces
		{"cpu", "Brand/with/slashes"},                // special characters
	}

	for _, input := range specialInputs {
		t.Run("Special input", func(t *testing.T) {
			// These should not panic or cause database corruption
			_, err := GetComponentsByBrand(models.GetComponentsByBrandInput{
				Category: input.category,
				Brand:    input.brand,
			})
			t.Logf("GetComponentsByBrand('%s', '%s') returned error: %v",
				input.category, input.brand, err)
		})
	}
}

// TestRepositoryPerformance_Integration tests performance with larger datasets
func TestRepositoryPerformance_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping performance test in short mode")
	}

	db := testutils.SetupTestDB(t)
	defer testutils.CleanupTestDB(t, db)

	// Insert a reasonable amount of test data
	baseComponents := testutils.CreateTestComponents()
	var allComponents []models.Component

	// Create variations of the base components
	for i := 0; i < 25; i++ { // 25 * 4 = 100 components
		for _, comp := range baseComponents {
			variant := comp
			variant.Model = comp.Model + " Variant " + string(rune('A'+i))
			variant.SKU = testutils.StringPtr(*comp.SKU + "-V" + string(rune('A'+i)))
			allComponents = append(allComponents, variant)
		}
	}

	// Insert all components
	start := time.Now()
	insertedIDs := testutils.InsertTestComponents(t, db, allComponents)
	insertDuration := time.Since(start)

	t.Logf("Inserted %d components in %v", len(insertedIDs), insertDuration)

	// Test GetAllComponents performance
	start = time.Now()
	allResults, err := GetAllComponents(models.GetAllComponentsInput{
		Page: "1",
	})
	getAllDuration := time.Since(start)

	require.NoError(t, err)
	assert.GreaterOrEqual(t, len(allResults), len(allComponents))
	t.Logf("Retrieved %d components in %v", len(allResults), getAllDuration)

	// Test category filtering performance
	start = time.Now()
	cpuResults, err := GetComponentsByCategory(models.GetComponentsByCategoryInput{
		Category: "cpu",
	})
	categoryDuration := time.Since(start)

	require.NoError(t, err)
	t.Logf("Filtered %d CPU components in %v", len(cpuResults), categoryDuration)

	// Test brand filtering performance
	start = time.Now()
	brandResults, err := GetComponentsByBrand(models.GetComponentsByBrandInput{
		Category: "cpu",
		Brand:    "Test Intel",
	})
	brandDuration := time.Since(start)

	require.NoError(t, err)
	t.Logf("Filtered %d Intel CPU components in %v", len(brandResults), brandDuration)

	// Performance assertions (adjust thresholds as needed)
	assert.Less(t, getAllDuration.Milliseconds(), int64(1000), "GetAllComponents should complete within 1 second")
	assert.Less(t, categoryDuration.Milliseconds(), int64(500), "Category filtering should complete within 500ms")
	assert.Less(t, brandDuration.Milliseconds(), int64(500), "Brand filtering should complete within 500ms")
}
