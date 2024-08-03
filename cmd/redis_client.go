package cmd

import "github.com/redis/go-redis/v9"

func CreateRedisConnection() redis.Cmdable {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
