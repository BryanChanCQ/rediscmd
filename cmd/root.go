package cmd

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "redisCmd",
	Short: "redis cmd root",
	Long:  "this is a hello world redis cmd",
}

func Execute(cmd redis.Cmdable) {
	NewKeysStruct(cmd).CreateKeysCmd()
	err := rootCmd.Execute()
	if err != nil {
		return
	}
}
