package cache

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

func NewRedis() *redis.Client {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379"
	}

	return redis.NewClient(&redis.Options{
		Addr: addr,
	})
}

func Ping(client *redis.Client) error {
	return client.Ping(context.Background()).Err()
}
