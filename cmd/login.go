package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type redisLogin struct {
	User     string
	Password string
	Db       int8
}

var filePath string

func CreateLoginCmd() {
	initFilePath()
	loginCmd := &cobra.Command{
		Use:   "login",
		Short: "login redis",
		Long:  "login redis with user or password",
		Args:  cobra.NoArgs,
		Run:   createLoginRunFunc(),
	}
	loginCmd.Flags().StringP("user", "u", "default", "input your redis user")
	loginCmd.Flags().StringP("password", "p", "", "input your redis password")
	rootCmd.AddCommand(loginCmd)
}

func initFilePath() {
	dir, err := os.UserHomeDir()
	if err != nil {
		panic("can not get user home path")
	}
	filePath = dir + "/.loginInfo.json"
}

func createLoginRunFunc() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		user, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")
		login := redisLogin{
			User:     user,
			Password: password,
		}
		marshal, err := json.Marshal(login)
		if err != nil {
			panic("json format error")
		}
		file, err := os.Create(filePath)

		if err != nil {
			panic(fmt.Sprintf("login file create err:%v", err))
		}
		defer file.Close()

		_, err = file.Write(marshal)
		if err != nil {
			panic("write file error")
		}
		fmt.Println("login success!")
	}
}
