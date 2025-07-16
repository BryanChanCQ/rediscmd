package main

import (
	"fmt"
	"os"
	"redisCmd/cmd"
)

func main() {
	// 执行登录和时候不走连接
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "login", "show":
			cmd.Execute(nil)
		default:
			useConnection()
		}
	} else {
		cmd.Execute(nil)
	}
}

func useConnection() {
	connection, err := cmd.CreateRedisConnection()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	cmd.Execute(connection)
}
