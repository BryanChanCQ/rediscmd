package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func CreateRedisConnection() redis.Cmdable {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// 检查连接
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error connecting to Redis:%v\n", err))
	}

	// fmt.Printf("Successfully connected to Redis. Response: %s\n", pingResp)

	return client

}
