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

func Execute(cmdAble redis.Cmdable) {
	AddCmdUnderRoot(cmdAble)
	err := rootCmd.Execute()
	if err != nil {
		return
	}
}

func AddCmdUnderRoot(cmdAble redis.Cmdable) {
	NewKeysStruct(cmdAble).CreateKeysCmd()
	NewSetStruct(cmdAble).CreateSetCmd()
	NewGetStruct(cmdAble).CreateGetCmd()
	CreateLoginCmd()
}
