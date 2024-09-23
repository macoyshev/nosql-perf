package redis

import (
	"os"

	"github.com/redis/go-redis/v9"
)

func NewClient() *redis.Client {
	url := os.Getenv("REDIS_URL")
	opts, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}

	return redis.NewClient(opts)
}
