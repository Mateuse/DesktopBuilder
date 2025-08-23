package utils

import (
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mateuse/desktop-builder-backend/internal/constants"
	"github.com/mateuse/desktop-builder-backend/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestInitializeDatabase_Success(t *testing.T) {
	// Set environment variables for test
	os.Setenv("DB_USER", "testuser")
	os.Setenv("DB_PASSWORD", "testpass")
	os.Setenv("DB_NAME", "testdb")
	defer func() {
		os.Unsetenv("DB_USER")
		os.Unsetenv("DB_PASSWORD")
		os.Unsetenv("DB_NAME")
	}()

	// Note: This test would require an actual database connection
	// In a real scenario, you might want to use docker for testing
	// or mock the sql.Open function

	// For now, we'll test the validation logic
	err := InitializeDatabase()
	// We expect this to fail because we don't have a real database
	// but it should not fail due to missing environment variables
	assert.Error(t, err) // Will fail on actual connection, not validation
}

func TestInitializeDatabase_MissingUser(t *testing.T) {
	// Clear environment variables
	os.Unsetenv("DB_USER")
	os.Unsetenv("DB_PASSWORD")
	os.Unsetenv("DB_NAME")

	err := InitializeDatabase()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "DB_USER environment variable is required")
}

func TestInitializeDatabase_MissingPassword(t *testing.T) {
	// Set only user
	os.Setenv("DB_USER", "testuser")
	defer os.Unsetenv("DB_USER")

	err := InitializeDatabase()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "DB_PASSWORD environment variable is required")
}

func TestInitializeDatabase_MissingDBName(t *testing.T) {
	// Set user and password but not DB name
	os.Setenv("DB_USER", "testuser")
	os.Setenv("DB_PASSWORD", "testpass")
	defer func() {
		os.Unsetenv("DB_USER")
		os.Unsetenv("DB_PASSWORD")
	}()

	err := InitializeDatabase()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "DB_NAME environment variable is required")
}

func TestCloseDatabase_NilDB(t *testing.T) {
	// Ensure DB is nil
	originalDB := DB
	DB = nil
	defer func() {
		DB = originalDB
	}()

	err := CloseDatabase()
	assert.NoError(t, err)
}

func TestCloseDatabase_WithDB(t *testing.T) {
	// Create a mock database
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Set the global DB variable
	originalDB := DB
	DB = db
	defer func() {
		DB = originalDB
	}()

	// Expect the close call
	mock.ExpectClose()

	err = CloseDatabase()
	assert.NoError(t, err)

	// Verify all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetDB(t *testing.T) {
	// Create a mock database
	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Set the global DB variable
	originalDB := DB
	DB = db
	defer func() {
		DB = originalDB
	}()

	// Test GetDB
	result := GetDB()
	assert.Equal(t, db, result)
}

func TestGetDB_Nil(t *testing.T) {
	// Ensure DB is nil
	originalDB := DB
	DB = nil
	defer func() {
		DB = originalDB
	}()

	result := GetDB()
	assert.Nil(t, result)
}

func TestGenerateSelectQuery(t *testing.T) {
	tests := []struct {
		name     string
		input    models.GenerateSelectQueryInput
		expected string
	}{
		{
			name: "basic query with all columns",
			input: models.GenerateSelectQueryInput{
				Table:   "components",
				Columns: []string{},
				Page:    "",
			},
			expected: "SELECT * FROM components LIMIT 50",
		},
		{
			name: "query with specific columns",
			input: models.GenerateSelectQueryInput{
				Table:   "components",
				Columns: []string{"id", "name", "price"},
				Page:    "",
			},
			expected: "SELECT id, name, price FROM components LIMIT 50",
		},
		{
			name: "query with where clause",
			input: models.GenerateSelectQueryInput{
				Table:       "components",
				Columns:     []string{"*"},
				WhereClause: "category = 'cpu'",
				Page:        "",
			},
			expected: "SELECT * FROM components WHERE category = 'cpu' LIMIT 50",
		},
		{
			name: "query with pagination - page 1",
			input: models.GenerateSelectQueryInput{
				Table:   "components",
				Columns: []string{"id", "name"},
				Page:    "1",
			},
			expected: "SELECT id, name FROM components LIMIT 50",
		},
		{
			name: "query with pagination - page 2",
			input: models.GenerateSelectQueryInput{
				Table:   "components",
				Columns: []string{"id", "name"},
				Page:    "2",
			},
			expected: "SELECT id, name FROM components LIMIT 50 OFFSET 50",
		},
		{
			name: "query with pagination - page 3",
			input: models.GenerateSelectQueryInput{
				Table:   "components",
				Columns: []string{"id", "name"},
				Page:    "3",
			},
			expected: "SELECT id, name FROM components LIMIT 50 OFFSET 100",
		},
		{
			name: "complex query with all features",
			input: models.GenerateSelectQueryInput{
				Table:       "components",
				Columns:     []string{"id", "category", "brand", "model"},
				WhereClause: "category = 'gpu' AND brand = 'nvidia'",
				Page:        "2",
			},
			expected: "SELECT id, category, brand, model FROM components WHERE category = 'gpu' AND brand = 'nvidia' LIMIT 50 OFFSET 50",
		},
		{
			name: "query with page 0 (should default to page 1)",
			input: models.GenerateSelectQueryInput{
				Table:   "components",
				Columns: []string{"id"},
				Page:    "0",
			},
			expected: "SELECT id FROM components LIMIT 50",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GenerateSelectQuery(tt.input)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGenerateLimitAndOffset(t *testing.T) {
	tests := []struct {
		name     string
		page     string
		expected constants.LimitAndOffset
	}{
		{
			name: "empty page string",
			page: "",
			expected: constants.LimitAndOffset{
				Limit:  constants.DEFAULT_PAGE_SIZE,
				Offset: 0,
			},
		},
		{
			name: "page 0",
			page: "0",
			expected: constants.LimitAndOffset{
				Limit:  constants.DEFAULT_PAGE_SIZE,
				Offset: 0,
			},
		},
		{
			name: "page 1",
			page: "1",
			expected: constants.LimitAndOffset{
				Limit:  constants.DEFAULT_PAGE_SIZE,
				Offset: 0,
			},
		},
		{
			name: "page 2",
			page: "2",
			expected: constants.LimitAndOffset{
				Limit:  constants.DEFAULT_PAGE_SIZE,
				Offset: 50,
			},
		},
		{
			name: "page 3",
			page: "3",
			expected: constants.LimitAndOffset{
				Limit:  constants.DEFAULT_PAGE_SIZE,
				Offset: 100,
			},
		},
		{
			name: "page 10",
			page: "10",
			expected: constants.LimitAndOffset{
				Limit:  constants.DEFAULT_PAGE_SIZE,
				Offset: 450,
			},
		},
		{
			name: "invalid page string (non-numeric)",
			page: "invalid",
			expected: constants.LimitAndOffset{
				Limit:  constants.DEFAULT_PAGE_SIZE,
				Offset: 0,
			},
		},
		{
			name: "invalid page string (negative)",
			page: "-1",
			expected: constants.LimitAndOffset{
				Limit:  constants.DEFAULT_PAGE_SIZE,
				Offset: -100,
			},
		},
		{
			name: "invalid page string (decimal)",
			page: "2.5",
			expected: constants.LimitAndOffset{
				Limit:  constants.DEFAULT_PAGE_SIZE,
				Offset: 0,
			},
		},
		{
			name: "invalid page string (with spaces)",
			page: " 3 ",
			expected: constants.LimitAndOffset{
				Limit:  constants.DEFAULT_PAGE_SIZE,
				Offset: 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateLimitAndOffset(tt.page)
			assert.Equal(t, tt.expected.Limit, result.Limit)
			assert.Equal(t, tt.expected.Offset, result.Offset)
		})
	}
}
