package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"redisCmd/tools"

	"github.com/spf13/cobra"
)

type redisLogin struct {
	User     string
	Password string
	Host     string
	Port     string
	Db       int8
}

var filepath string

func CreateLoginCmd() {
	filepath = tools.InitFilePath()
	loginCmd := &cobra.Command{
		Use:   "login",
		Short: "login redis",
		Long:  "login redis with user or password",
		Args:  cobra.NoArgs,
		Run:   createLoginRunFunc(),
	}
	loginCmd.Flags().StringP("user", "u", "default", "input your redis user")
	loginCmd.Flags().StringP("password", "s", "", "input your redis password")
	loginCmd.Flags().StringP("port", "p", "6379", "input your redis port")
	// shorthand "h" conflict with help -h
	loginCmd.Flags().StringP("host", "o", "127.0.0.1", "input your redis host")
	rootCmd.AddCommand(loginCmd)
}

func createLoginRunFunc() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		user, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")
		port, _ := cmd.Flags().GetString("port")
		host, _ := cmd.Flags().GetString("host")
		login := redisLogin{
			User:     user,
			Password: password,
			Port:     port,
			Host:     host,
		}
		marshal, err := json.Marshal(login)
		if err != nil {
			panic("json format error")
		}

		file, err := os.Create(filepath)

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
