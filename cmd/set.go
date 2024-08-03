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
		Args:  cobra.RangeArgs(2, 3),
		Run:   createSetFunc(k.cmd),
	}
	rootCmd.AddCommand(setCmd)
}

func createSetFunc(redisCmd redis.Cmdable) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		if len(args) == 3 {
			redisCmd.Set(ctx, args[0], args[1], time.Duration(conv2Int64(args[2]))*chooseTime(args[2]))
		} else {
			redisCmd.Set(ctx, args[0], args[1], 0)
		}
	}
}

func conv2Int64(times string) int64 {
	times = times[0 : len(times)-1]
	parseInt, err := strconv.ParseInt(times, 10, 64)
	if err != nil {
		panic("please input correct time")
	}
	return parseInt
}

func chooseTime(times string) time.Duration {
	times = times[len(times)-1:]
	return timeMap[times]
}
