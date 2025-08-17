package testutils

import (
	"database/sql"
	"net/http"
	"net/http/httptest"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/mock"
)

// MockDB provides utilities for mocking database connections
type MockDB struct {
	DB   *sql.DB
	Mock sqlmock.Sqlmock
}

// NewMockDB creates a new mock database connection
func NewMockDB() (*MockDB, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, err
	}

	return &MockDB{
		DB:   db,
		Mock: mock,
	}, nil
}

// Close closes the mock database connection
func (m *MockDB) Close() error {
	return m.DB.Close()
}

// MockRedisClient is a mock implementation of Redis client
type MockRedisClient struct {
	mock.Mock
}

// Ping mocks the Redis Ping method
func (m *MockRedisClient) Ping(ctx interface{}) *redis.StatusCmd {
	args := m.Called(ctx)
	return args.Get(0).(*redis.StatusCmd)
}

// Set mocks the Redis Set method
func (m *MockRedisClient) Set(ctx interface{}, key string, value interface{}, expiration interface{}) *redis.StatusCmd {
	args := m.Called(ctx, key, value, expiration)
	return args.Get(0).(*redis.StatusCmd)
}

// Get mocks the Redis Get method
func (m *MockRedisClient) Get(ctx interface{}, key string) *redis.StringCmd {
	args := m.Called(ctx, key)
	return args.Get(0).(*redis.StringCmd)
}

// Del mocks the Redis Del method
func (m *MockRedisClient) Del(ctx interface{}, keys ...string) *redis.IntCmd {
	args := m.Called(ctx, keys)
	return args.Get(0).(*redis.IntCmd)
}

// Exists mocks the Redis Exists method
func (m *MockRedisClient) Exists(ctx interface{}, keys ...string) *redis.IntCmd {
	args := m.Called(ctx, keys)
	return args.Get(0).(*redis.IntCmd)
}

// Close mocks the Redis Close method
func (m *MockRedisClient) Close() error {
	args := m.Called()
	return args.Error(0)
}

// TestServer creates a test HTTP server for integration tests
func TestServer(handler http.Handler) *httptest.Server {
	return httptest.NewServer(handler)
}

// TestRequest creates a test HTTP request
func TestRequest(method, url string, body interface{}) *http.Request {
	req := httptest.NewRequest(method, url, nil)
	req.Header.Set("Content-Type", "application/json")
	return req
}

// TestResponseRecorder creates a test HTTP response recorder
func TestResponseRecorder() *httptest.ResponseRecorder {
	return httptest.NewRecorder()
}
