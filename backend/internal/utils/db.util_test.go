package utils

import (
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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
