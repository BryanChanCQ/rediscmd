package cmd

import (
	"context"
	"redisCmd/tools"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
)

type keysStruct struct {
	cmd redis.Cmdable
	tools.FormatPrint
}

func NewKeysStruct(cmd redis.Cmdable) *keysStruct {
	return &keysStruct{
		cmd,
		&tools.KeysPrint{},
	}
}

func (k keysStruct) CreateKeysCmd() {
	var keysCmd = &cobra.Command{
		Use:   "keys",
		Short: "search keys",
		Long:  "search keys by pattern",
		Args:  cobra.ExactArgs(1),
		Run:   createKeysFunc(k.cmd, k),
	}
	keysCmd.Flags().BoolP("delete", "d", false, "delete search keys")
	keysCmd.Flags().BoolP("expire", "t", false, "show keys expire")
	rootCmd.AddCommand(keysCmd)
}

func createKeysFunc(redisCmd redis.Cmdable, keysStruct keysStruct) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		result, err := redisCmd.Keys(ctx, "*"+args[0]+"*").Result()
		if err != nil {
			panic(err)
		}
		keysStruct.PrintTitle()
		for i := range result {
			eFlag, _ := cmd.Flags().GetBool("expire")
			if eFlag {
				duration := redisCmd.TTL(ctx, result[i])
				keysStruct.PrintRows(i, result[i], duration.Val())
			} else {
				keysStruct.PrintRows(i, result[i])
			}
		}
		dFlag, _ := cmd.Flags().GetBool("delete")
		if dFlag {
			redisCmd.Del(ctx, result...)
		}
	}
}
