package store

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/dranonymousnet/eventhorizon/internal/config"
)

var rdb *redis.Client

func InitRedis() {
	if rdb != nil {
		log.Println("Redis is already initialized")
		return // Early return if Redis is already initialized
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:     config.RedisSetting.Host,     // e.g., "localhost:6379"
		Password: config.RedisSetting.Password, // e.g., "" (no password)
		DB:       config.RedisSetting.DB,       // e.g., 0 (default DB)
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Printf("Failed to connect to Redis: %v", err)
		rdb = nil // Ensure rdb is nil to avoid using an uninitialized client
		return
	}

	log.Println("Redis connected")
}

func CloseRedisConn() {
	if rdb == nil {
		log.Println("Redis connection is not established")
		return
	}
	if err := rdb.Close(); err != nil {
		log.Printf("Error closing Redis connection: %v", err)
	}
}

func GetFromCache(key string) (string, error) {
	if rdb == nil {
		return "", fmt.Errorf("redis is not initialized")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("key does not exist: %s", key)
	} else if err != nil {
		return "", fmt.Errorf("error getting value from redis: %v", err)
	}
	return val, nil
}

func SetInCache(key, value string, expiration time.Duration) error {
	if rdb == nil {
		return fmt.Errorf("redis is not initialized")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := rdb.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return fmt.Errorf("error setting value in redis: %v", err)
	}
	return nil
}
