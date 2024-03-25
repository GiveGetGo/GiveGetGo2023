package middleware

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClientInterface interface {
	Ping(ctx context.Context) *redis.StatusCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(ctx context.Context, key string) *redis.StringCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
}

// func SetupRedis(url string) *redis.Client {
func SetupRedis() *redis.Client {
	// Create a redis client.
	option, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Fatalf("Failed to create a redis client: %v", err)
		return nil
	}
	client := redis.NewClient(option)

	ctx := context.Background()

	// Ping the Redis server and check if any errors occurred.
	_, err = client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}

	return client
}
