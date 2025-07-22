package cmd

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/BryanChanCQ/rediscmd/tools"

	"github.com/redis/go-redis/v9"
)

var (
	ErrNotLogin = errors.New("please login first")
)
func CreateRedisConnection() (redis.Cmdable, error) {
	var redisOptions redisLogin
	marshlSucess := tools.FileUnmarshl(tools.InitFilePath(), tools.JSONFile, &redisOptions)
	if !marshlSucess {
		return nil, ErrNotLogin
	}
	addr := redisOptions.Host + ":" + redisOptions.Port
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: redisOptions.Password,
		DB:       int(redisOptions.Db),
	})

	// 检查连接
	ctx, cancel := context.WithTimeout(context.Background(), tools.Timeout*time.Second)
	defer cancel()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error connecting to Redis:%v\n", err))
	}

	// fmt.Printf("Successfully connected to Redis. Response: %s\n", pingResp)

	return client, nil

}
