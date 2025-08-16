package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

// InitializeRedis initializes the Redis client connection
func InitializeRedis() error {
	// Get Redis configuration from environment variables
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDBStr := os.Getenv("REDIS_DB")

	// Set default values if not provided
	if redisHost == "" {
		redisHost = "localhost"
	}
	if redisPort == "" {
		redisPort = "6379"
	}

	// Parse Redis DB number (default to 0)
	redisDB := 0
	if redisDBStr != "" {
		var err error
		redisDB, err = strconv.Atoi(redisDBStr)
		if err != nil {
			return fmt.Errorf("invalid REDIS_DB value: %w", err)
		}
	}

	// Build Redis address
	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)

	// Create Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:         redisAddr,
		Password:     redisPassword, // empty string if no password
		DB:           redisDB,       // default DB
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("failed to connect to Redis: %w", err)
	}

	RedisClient = rdb
	log.Printf("Successfully connected to Redis at %s", redisAddr)
	return nil
}

// CloseRedis closes the Redis connection
func CloseRedis() error {
	if RedisClient != nil {
		return RedisClient.Close()
	}
	return nil
}

// GetRedisClient returns the Redis client instance
func GetRedisClient() *redis.Client {
	return RedisClient
}

// SetWithExpiration sets a key-value pair with expiration
func SetWithExpiration(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	if RedisClient == nil {
		return fmt.Errorf("redis client not initialized")
	}
	return RedisClient.Set(ctx, key, value, expiration).Err()
}

// Get retrieves a value by key
func Get(ctx context.Context, key string) (string, error) {
	if RedisClient == nil {
		return "", fmt.Errorf("redis client not initialized")
	}
	return RedisClient.Get(ctx, key).Result()
}

// Delete removes a key
func Delete(ctx context.Context, key string) error {
	if RedisClient == nil {
		return fmt.Errorf("redis client not initialized")
	}
	return RedisClient.Del(ctx, key).Err()
}

// Exists checks if a key exists
func Exists(ctx context.Context, key string) (bool, error) {
	if RedisClient == nil {
		return false, fmt.Errorf("redis client not initialized")
	}
	result, err := RedisClient.Exists(ctx, key).Result()
	return result > 0, err
}
