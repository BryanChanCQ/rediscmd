package cmd

import (
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
	"os"
	"redisCmd/tools"
)

type SelectStruct struct {
	cmd redis.Cmdable
}

func NewSelectStruct(cmd redis.Cmdable) *SelectStruct {
	return &SelectStruct{
		cmd: cmd,
	}
}

func (k *SelectStruct) CreateSelectCmd() {
	command := cobra.Command{
		Use:   "select",
		Short: "select db index",
		Long:  "select db index",
		Run:   createSelectFunc(),
	}
	command.Flags().Uint8P("db", "b", 0, "db index")
	rootCmd.AddCommand(&command)
}

func createSelectFunc() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		var redisOptions redisLogin
		marshlSucess := tools.FileUnmarshl(tools.InitFilePath(), tools.JSON, &redisOptions)
		if !marshlSucess {
			panic("file unmarshl fail")
		}
		redisOptions.Db, _ = cmd.Flags().GetUint8("db")
		file, _ := os.Create(tools.InitFilePath())
		defer file.Close()
		marshal, _ := json.Marshal(redisOptions)
		_, err := file.Write(marshal)
		if err != nil {
			panic("write file fail")
		}
	}
}
