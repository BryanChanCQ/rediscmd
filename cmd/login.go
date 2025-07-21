package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"redisCmd/tools"

	"github.com/spf13/cobra"
)

type redisLogin struct {
	DataBaseName string
	User         string
	Password     string
	Host         string
	Port         string
	Db           uint8
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
	loginCmd.Flags().StringP("dbname", "n", "", "input your redis dbname")
	rootCmd.AddCommand(loginCmd)
}

func createLoginRunFunc() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		user, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")
		port, _ := cmd.Flags().GetString("port")
		host, _ := cmd.Flags().GetString("host")
		dbname, _ := cmd.Flags().GetString("dbname")
		login := redisLogin{
			User:         user,
			Password:     password,
			Port:         port,
			Host:         host,
			DataBaseName: dbname,
		}
		WriteLoginData2File(&login)
	}
}

func WriteLoginData2File(login *redisLogin) {
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
	writeLoginData2SwitchDB(login)
}

func writeLoginData2SwitchDB(login *redisLogin) {
	var dbSwitch []redisLogin
	filepath := tools.DbSwitchFilePath()
	tools.FileUnmarshl(filepath, tools.JSONFile, &dbSwitch)
	hasHost := false
	for index := range dbSwitch {
		loginDb := dbSwitch[index]
		if loginDb.Host == login.Host {
			hasHost = true
			break
		}
	}
	if !hasHost {
		dbSwitch = append(dbSwitch, *login)
		file, err := os.Create(filepath)
		if err != nil {
			panic(fmt.Sprintf("switch file create err:%v", err))
		}
		defer file.Close()
		marshal, err := json.Marshal(dbSwitch)
		if err != nil {
			panic("json format error")
		}
		_, err = file.Write(marshal)
		if err != nil {
			panic("write file error")
		}
		fmt.Println("add to switch database success!")
	}
}
