package cmd

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
)

type keysStruct struct {
	cmd redis.Cmdable
}

func NewKeysStruct(cmd redis.Cmdable) *keysStruct {
	return &keysStruct{
		cmd: cmd,
	}
}

func (k keysStruct) CreateKeysCmd() {
	var keysCmd = &cobra.Command{
		Use:   "keys",
		Short: "search keys",
		Long:  "search keys by pattern",
		Args:  cobra.ExactArgs(1),
		Run:   createKeysFunc(k.cmd),
	}
	keysCmd.Flags().BoolP("delete", "d", false, "delete search keys")
	rootCmd.AddCommand(keysCmd)
}

func createKeysFunc(redisCmd redis.Cmdable) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		result, err := redisCmd.Keys(ctx, "*"+args[0]+"*").Result()
		if err != nil {
			panic(err)
		}
		for i := range result {
			fmt.Printf("key:%s\n", result[i])
		}
		value := cmd.Flags().Lookup("delete").Value.String()
		fmt.Println(value)
		if value == "true" {
			redisCmd.Del(ctx, result...)
		}
	}
}
