package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
)

const (
	timeout = 10
)

type GetStruct struct {
	cmd redis.Cmdable
}

func NewGetStruct(cmd redis.Cmdable) *GetStruct {
	return &GetStruct{
		cmd: cmd,
	}
}

func (k *GetStruct) CreateGetCmd() {
	var setCmd = &cobra.Command{
		Use:   "get",
		Short: "get value by key",
		Long:  "get value by key",
		Args:  cobra.ExactArgs(1),
		Run:   createGetFunc(k.cmd),
	}
	rootCmd.AddCommand(setCmd)
}

func createGetFunc(redisCmd redis.Cmdable) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), timeout*time.Second)
		defer cancel()
		stringCmd := redisCmd.Get(ctx, args[0])
		val := stringCmd.Val()
		if len(val) <= 0 {
			fmt.Printf("key:%s not found\n", args[0])
			return
		}
		fmt.Println(stringCmd.Val())
	}
}
