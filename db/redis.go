package db

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var (
	Rdb *redis.Client
	Ctx context.Context
)

// InitRedis initializes the Redis client
func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	Ctx = context.Background()
}
