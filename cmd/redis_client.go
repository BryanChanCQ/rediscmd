package cmd

import (
	"context"
	"fmt"
	"redisCmd/tools"
	"time"

	"github.com/redis/go-redis/v9"
)

func CreateRedisConnection() redis.Cmdable {
	var redisOptions redisLogin
	tools.FileUnmarshl(tools.InitFilePath(), tools.JSON, &redisOptions)

	addr := redisOptions.Host + ":" + redisOptions.Port
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: redisOptions.Password,
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
