package cmd

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
	"strconv"
	"time"
)

type SetStruct struct {
	cmd redis.Cmdable
}

var timeMap = map[string]time.Duration{
	"s": time.Second,
	"m": time.Minute,
	"h": time.Hour,
}

func NewSetStruct(cmd redis.Cmdable) *SetStruct {
	return &SetStruct{
		cmd: cmd,
	}
}

func (k SetStruct) CreateSetCmd() {
	var setCmd = &cobra.Command{
		Use:   "set",
		Short: "set key value",
		Long:  "set key value",
		Args:  cobra.ExactArgs(2),
		Run:   createSetFunc(k.cmd),
	}
	setCmd.Flags().StringP("expire", "e", "", "set key expire time example: 20s, 30m, 20h")
	rootCmd.AddCommand(setCmd)
}

func createSetFunc(redisCmd redis.Cmdable) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		expireFlagValue := cmd.Flags().Lookup("expire").Value.String()
		redisCmd.Set(ctx, args[0], args[1], time.Duration(conv2Int64(expireFlagValue))*chooseTime(expireFlagValue))
	}
}

func conv2Int64(times string) int64 {
	if len(times) == 0 {
		return 0
	}
	times = times[0 : len(times)-1]
	parseInt, err := strconv.ParseInt(times, 10, 64)
	if err != nil {
		return 0
	}
	return parseInt
}

func chooseTime(times string) time.Duration {
	if len(times) == 0 {
		return 0
	}
	times = times[len(times)-1:]
	return timeMap[times]
}
