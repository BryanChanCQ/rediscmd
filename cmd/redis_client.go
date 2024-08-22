package cmd

import (
	"context"
	"errors"
	"fmt"
	"redisCmd/tools"
	"time"

	"github.com/redis/go-redis/v9"
)

func CreateRedisConnection() (redis.Cmdable, error) {
	var redisOptions redisLogin
	marshlSucess := tools.FileUnmarshl(tools.InitFilePath(), tools.JSON, &redisOptions)
	if !marshlSucess {
		return nil, errors.New("please login first!!!!!\n")
	}
	addr := redisOptions.Host + ":" + redisOptions.Port
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: redisOptions.Password,
		DB:       int(redisOptions.Db),
	})

	// 检查连接
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error connecting to Redis:%v\n", err))
	}

	// fmt.Printf("Successfully connected to Redis. Response: %s\n", pingResp)

	return client, nil

}
