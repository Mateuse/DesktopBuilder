package utils

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestInitializeRedis_DefaultValues(t *testing.T) {
	// Clear environment variables to test defaults
	os.Unsetenv("REDIS_HOST")
	os.Unsetenv("REDIS_PORT")
	os.Unsetenv("REDIS_PASSWORD")
	os.Unsetenv("REDIS_DB")

	// Store original client and restore after test
	originalClient := RedisClient
	defer func() {
		if RedisClient != nil {
			RedisClient.Close()
		}
		RedisClient = originalClient
	}()

	// This will fail because we don't have a real Redis server
	// but we can test that it doesn't panic and handles defaults
	err := InitializeRedis()
	if err == nil {
		// If Redis is actually running locally, that's fine too
		// Just make sure we can close it properly
		assert.NoError(t, CloseRedis())
	} else {
		// Expected to fail on connection if Redis is not running
		assert.Contains(t, err.Error(), "failed to connect to Redis")
	}
}

func TestInitializeRedis_CustomValues(t *testing.T) {
	// Set custom environment variables
	os.Setenv("REDIS_HOST", "custom-host")
	os.Setenv("REDIS_PORT", "6380")
	os.Setenv("REDIS_PASSWORD", "testpass")
	os.Setenv("REDIS_DB", "1")
	defer func() {
		os.Unsetenv("REDIS_HOST")
		os.Unsetenv("REDIS_PORT")
		os.Unsetenv("REDIS_PASSWORD")
		os.Unsetenv("REDIS_DB")
	}()

	// Store original client and restore after test
	originalClient := RedisClient
	defer func() {
		if RedisClient != nil {
			RedisClient.Close()
		}
		RedisClient = originalClient
	}()

	// This will fail because we don't have a real Redis server
	// but we can test that it processes the environment variables
	err := InitializeRedis()
	assert.Error(t, err) // Expected to fail on connection to custom-host
	assert.Contains(t, err.Error(), "failed to connect to Redis")
}

func TestInitializeRedis_InvalidDB(t *testing.T) {
	// Set invalid DB value
	os.Setenv("REDIS_DB", "invalid")
	defer os.Unsetenv("REDIS_DB")

	err := InitializeRedis()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid REDIS_DB value")
}

func TestCloseRedis_NilClient(t *testing.T) {
	// Ensure RedisClient is nil
	originalClient := RedisClient
	RedisClient = nil
	defer func() {
		RedisClient = originalClient
	}()

	err := CloseRedis()
	assert.NoError(t, err)
}

func TestGetRedisClient(t *testing.T) {
	// Create a mock client
	mockClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Set the global RedisClient variable
	originalClient := RedisClient
	RedisClient = mockClient
	defer func() {
		RedisClient = originalClient
		mockClient.Close()
	}()

	result := GetRedisClient()
	assert.Equal(t, mockClient, result)
}

func TestGetRedisClient_Nil(t *testing.T) {
	// Ensure RedisClient is nil
	originalClient := RedisClient
	RedisClient = nil
	defer func() {
		RedisClient = originalClient
	}()

	result := GetRedisClient()
	assert.Nil(t, result)
}

func TestSetWithExpiration_NilClient(t *testing.T) {
	// Ensure RedisClient is nil
	originalClient := RedisClient
	RedisClient = nil
	defer func() {
		RedisClient = originalClient
	}()

	ctx := context.Background()
	err := SetWithExpiration(ctx, "test-key", "test-value", time.Minute)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "redis client not initialized")
}

func TestGet_NilClient(t *testing.T) {
	// Ensure RedisClient is nil
	originalClient := RedisClient
	RedisClient = nil
	defer func() {
		RedisClient = originalClient
	}()

	ctx := context.Background()
	_, err := Get(ctx, "test-key")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "redis client not initialized")
}

func TestDelete_NilClient(t *testing.T) {
	// Ensure RedisClient is nil
	originalClient := RedisClient
	RedisClient = nil
	defer func() {
		RedisClient = originalClient
	}()

	ctx := context.Background()
	err := Delete(ctx, "test-key")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "redis client not initialized")
}

func TestExists_NilClient(t *testing.T) {
	// Ensure RedisClient is nil
	originalClient := RedisClient
	RedisClient = nil
	defer func() {
		RedisClient = originalClient
	}()

	ctx := context.Background()
	_, err := Exists(ctx, "test-key")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "redis client not initialized")
}

// Integration tests would require a real Redis instance
// These could be added using docker-compose or testcontainers
// for more comprehensive testing
