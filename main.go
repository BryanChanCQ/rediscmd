package main

import (
	"fmt"
	"os"
	"redisCmd/cmd"
)

func main() {
	// 执行登录的时候不走连接
	if len(os.Args) > 1 && os.Args[1] == "login" {
		cmd.Execute(nil)
	} else {
		connection, err := cmd.CreateRedisConnection()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		cmd.Execute(connection)
	}

}
