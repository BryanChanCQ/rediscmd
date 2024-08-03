package main

import "redisCmd/cmd"

func main() {
	cmd.Execute(cmd.CreateRedisConnection())
}
