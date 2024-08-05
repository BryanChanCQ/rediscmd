package main

import (
	"redisCmd/cmd"
)

func main() {
	// 执行登录的时候不走连接
	cmd.Execute(cmd.CreateRedisConnection())
}
