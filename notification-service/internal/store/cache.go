package store

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"

	"github.com/dranonymousnet/eventhorizon/internal/config"
)


var rdb *redis.Client

func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:    config.RedisSetting.Host,
		Password: config.RedisSetting.Password,
		DB:       0,               // Use the default DB
	})


    // Check the connection
    ctx := context.Background()
    _, err := rdb.Ping(ctx).Result()
    if err != nil {
        log.Fatalf("Failed to connect to Redis: %v", err)
    }
}


var ctx = context.Background()

func main() {
	// Create a new Redis client.

	// Example of setting a value with expiration
	err := rdb.Set(ctx, "key", "value", 0).Err() // Use 0 for no expiration, or a duration like time.Second*10 for 10 seconds
	if err != nil {
		log.Fatalf("Error setting value in Redis: %v", err)
	}

	// Example of getting a value
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		log.Fatalf("Error getting value from Redis: %v", err)
	}
	fmt.Printf("The value of 'key' is: %s\n", val)

	// Optionally, close the connection when your application exits
	defer rdb.Close()
}
