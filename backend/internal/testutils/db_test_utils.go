package testutils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/mateuse/desktop-builder-backend/internal/models"
	"github.com/mateuse/desktop-builder-backend/internal/utils"
)

// TestDBConfig holds configuration for test database
type TestDBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// GetTestDBConfig returns test database configuration from environment
func GetTestDBConfig() TestDBConfig {
	return TestDBConfig{
		Host:     getEnvOrDefault("TEST_DB_HOST", "localhost"),
		Port:     getEnvOrDefault("TEST_DB_PORT", "5432"),
		User:     getEnvOrDefault("TEST_DB_USER", "test_user"),
		Password: getEnvOrDefault("TEST_DB_PASSWORD", "test_password"),
		DBName:   getEnvOrDefault("TEST_DB_NAME", "test_desktop_builder"),
		SSLMode:  getEnvOrDefault("TEST_DB_SSLMODE", "disable"),
	}
}

// SetupTestDB creates a test database connection
// Note: This requires a test database to be available
func SetupTestDB(t *testing.T) *sql.DB {
	config := GetTestDBConfig()

	// Skip if test database is not configured
	if config.User == "test_user" && config.Password == "test_password" {
		t.Skip("Test database not configured. Set TEST_DB_* environment variables to run integration tests.")
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("Failed to open test database connection: %v", err)
	}

	if err := db.Ping(); err != nil {
		t.Fatalf("Failed to ping test database: %v", err)
	}

	// Set the global DB for utils package
	utils.DB = db

	t.Logf("Connected to test database: %s", config.DBName)
	return db
}

// CleanupTestDB closes the test database connection and cleans up
func CleanupTestDB(t *testing.T, db *sql.DB) {
	if db != nil {
		// Clean up test data
		CleanupTestData(t, db)

		// Close connection
		if err := db.Close(); err != nil {
			t.Logf("Warning: Failed to close test database connection: %v", err)
		}

		// Reset global DB
		utils.DB = nil

		t.Log("Test database cleanup completed")
	}
}

// CleanupTestData removes test data from the database
func CleanupTestData(t *testing.T, db *sql.DB) {
	// Delete test components (be careful with this in production!)
	_, err := db.Exec("DELETE FROM components WHERE brand LIKE 'Test%' OR model LIKE 'Test%'")
	if err != nil {
		t.Logf("Warning: Failed to cleanup test data: %v", err)
	}
}

// InsertTestComponent inserts a test component into the database
func InsertTestComponent(t *testing.T, db *sql.DB, component models.Component) int64 {
	query := `
		INSERT INTO components (category, brand, model, sku, upc, specs)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	var id int64
	err := db.QueryRow(query, component.Category, component.Brand, component.Model,
		component.SKU, component.UPC, component.Specs).Scan(&id)

	if err != nil {
		t.Fatalf("Failed to insert test component: %v", err)
	}

	t.Logf("Inserted test component with ID: %d", id)
	return id
}

// CreateTestComponents creates a set of test components for testing
func CreateTestComponents() []models.Component {
	return []models.Component{
		{
			Category: models.CategoryCPU,
			Brand:    "Test Intel",
			Model:    "Test Core i7-12700K",
			SKU:      stringPtr("TEST-BX8071512700K"),
			UPC:      stringPtr("TEST-735858491174"),
			Specs:    json.RawMessage(`{"cores": 12, "threads": 20, "base_clock": "3.6 GHz", "test": true}`),
		},
		{
			Category: models.CategoryCPU,
			Brand:    "Test AMD",
			Model:    "Test Ryzen 7 5800X",
			SKU:      stringPtr("TEST-100-100000063WOF"),
			UPC:      stringPtr("TEST-730143312042"),
			Specs:    json.RawMessage(`{"cores": 8, "threads": 16, "base_clock": "3.8 GHz", "test": true}`),
		},
		{
			Category: models.CategoryMemory,
			Brand:    "Test Corsair",
			Model:    "Test Vengeance LPX 32GB",
			SKU:      stringPtr("TEST-CMK32GX4M2E3200C16"),
			UPC:      stringPtr("TEST-843591098908"),
			Specs:    json.RawMessage(`{"capacity": "32GB", "speed": "3200MHz", "type": "DDR4", "test": true}`),
		},
		{
			Category: models.CategoryVideoCard,
			Brand:    "Test NVIDIA",
			Model:    "Test RTX 4080",
			SKU:      stringPtr("TEST-RTX4080"),
			UPC:      stringPtr("TEST-123456789"),
			Specs:    json.RawMessage(`{"memory": "16GB", "memory_type": "GDDR6X", "test": true}`),
		},
	}
}

// InsertTestComponents inserts multiple test components and returns their IDs
func InsertTestComponents(t *testing.T, db *sql.DB, components []models.Component) []int64 {
	var ids []int64

	for _, component := range components {
		id := InsertTestComponent(t, db, component)
		ids = append(ids, id)
	}

	t.Logf("Inserted %d test components", len(ids))
	return ids
}

// VerifyComponentExists checks if a component exists in the database
func VerifyComponentExists(t *testing.T, db *sql.DB, id int64) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM components WHERE id = $1", id).Scan(&count)
	if err != nil {
		t.Fatalf("Failed to verify component existence: %v", err)
	}
	return count > 0
}

// VerifyComponentsExist checks if multiple components exist in the database
func VerifyComponentsExist(t *testing.T, db *sql.DB, ids []int64) {
	for _, id := range ids {
		if !VerifyComponentExists(t, db, id) {
			t.Fatalf("Component with ID %d does not exist", id)
		}
	}
	t.Logf("Verified existence of %d components", len(ids))
}

// CountComponents returns the total number of components in the database
func CountComponents(t *testing.T, db *sql.DB) int {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM components").Scan(&count)
	if err != nil {
		t.Fatalf("Failed to count components: %v", err)
	}
	return count
}

// CountComponentsByCategory returns the number of components in a specific category
func CountComponentsByCategory(t *testing.T, db *sql.DB, category string) int {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM components WHERE category = $1", category).Scan(&count)
	if err != nil {
		t.Fatalf("Failed to count components by category: %v", err)
	}
	return count
}

// CountComponentsByBrand returns the number of components for a specific brand in a category
func CountComponentsByBrand(t *testing.T, db *sql.DB, category, brand string) int {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM components WHERE category = $1 AND brand = $2",
		category, brand).Scan(&count)
	if err != nil {
		t.Fatalf("Failed to count components by brand: %v", err)
	}
	return count
}

// Helper function to get environment variable with default
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// StringPtr creates a pointer to a string (exported for use in tests)
func StringPtr(s string) *string {
	return &s
}

// Helper function to create string pointers (for internal use)
func stringPtr(s string) *string {
	return &s
}

// WaitForDB waits for the database to be ready (useful for Docker containers)
func WaitForDB(t *testing.T, db *sql.DB, maxAttempts int) {
	for i := 0; i < maxAttempts; i++ {
		if err := db.Ping(); err == nil {
			return
		}
		t.Logf("Waiting for database... attempt %d/%d", i+1, maxAttempts)
		time.Sleep(time.Second)
	}
	t.Fatalf("Database not ready after %d attempts", maxAttempts)
}

/*
DATABASE TEST UTILITIES USAGE GUIDE:
=====================================

These utilities provide a foundation for integration testing with a real database.
They handle test database setup, data insertion, cleanup, and verification.

Setup Instructions:

1. **Test Database Configuration**:
   Set these environment variables for integration tests:
   ```bash
   export TEST_DB_HOST=localhost
   export TEST_DB_PORT=5432
   export TEST_DB_USER=test_user
   export TEST_DB_PASSWORD=test_password
   export TEST_DB_NAME=test_desktop_builder
   export TEST_DB_SSLMODE=disable
   ```

2. **Docker Test Database** (recommended):
   ```bash
   docker run --name test-postgres \
     -e POSTGRES_DB=test_desktop_builder \
     -e POSTGRES_USER=test_user \
     -e POSTGRES_PASSWORD=test_password \
     -p 5433:5432 \
     -d postgres:15
   ```

3. **Example Integration Test**:
   ```go
   func TestGetAllComponents_Integration(t *testing.T) {
       db := testutils.SetupTestDB(t)
       defer testutils.CleanupTestDB(t, db)

       // Insert test data
       components := testutils.CreateTestComponents()
       ids := testutils.InsertTestComponents(t, db, components)

       // Test the actual repository function
       result, err := repository.GetAllComponents()
       assert.NoError(t, err)
       assert.GreaterOrEqual(t, len(result), len(components))

       // Verify test data exists
       testutils.VerifyComponentsExist(t, db, ids)
   }
   ```

4. **Running Integration Tests**:
   ```bash
   # Run only unit tests (default)
   go test ./...

   # Run integration tests with database
   TEST_DB_HOST=localhost TEST_DB_USER=test_user TEST_DB_PASSWORD=test_password \
   TEST_DB_NAME=test_desktop_builder go test ./... -v
   ```

Safety Features:
- Tests are skipped if test database is not configured
- Test data uses "Test" prefixes for easy identification
- Cleanup functions remove only test data
- Parameterized queries prevent SQL injection

Best Practices:
- Always use defer for cleanup
- Use meaningful test data that's easy to identify
- Verify test data insertion before running tests
- Use transactions for complex test scenarios
- Keep test database separate from development database
*/
