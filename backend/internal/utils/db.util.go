package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
	"github.com/mateuse/desktop-builder-backend/internal/constants"
	"github.com/mateuse/desktop-builder-backend/internal/models"
)

var DB *sql.DB

// InitializeDatabase initializes the PostgreSQL database connection
func InitializeDatabase() error {
	// Get database configuration from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	// Set default values if not provided
	if dbHost == "" {
		dbHost = "localhost"
	}
	if dbPort == "" {
		dbPort = "5432"
	}
	if dbSSLMode == "" {
		dbSSLMode = "disable"
	}

	// Validate required environment variables
	if dbUser == "" {
		return fmt.Errorf("DB_USER environment variable is required")
	}
	if dbPassword == "" {
		return fmt.Errorf("DB_PASSWORD environment variable is required")
	}
	if dbName == "" {
		return fmt.Errorf("DB_NAME environment variable is required")
	}

	// Build connection string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode)

	// Open database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	DB = db
	log.Println("Successfully connected to PostgreSQL database")
	return nil
}

// CloseDatabase closes the database connection
func CloseDatabase() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}

// GetDB returns the database instance
func GetDB() *sql.DB {
	if DB != nil {
		Log(constants.DB_UTIL_GET_DB_CONNECTION_SUCCESS, nil)
	} else {
		Log(constants.DB_UTIL_GET_DB_CONNECTION_ERROR, fmt.Errorf("database connection is nil"))
	}
	return DB
}

func GenerateSelectQuery(input models.GenerateSelectQueryInput) (string, error) {
	Log(constants.DB_UTIL_GENERATE_SELECT_QUERY_START, nil, input.Table)

	columns := input.Columns
	if len(columns) == 0 {
		columns = []string{"*"}
	}

	query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(columns, ", "), input.Table)

	if input.WhereClause != "" {
		query += fmt.Sprintf(" WHERE %s", input.WhereClause)
	}

	Log(constants.DB_UTIL_GENERATE_SELECT_QUERY_SUCCESS, nil, input.Table)
	return query, nil
}
