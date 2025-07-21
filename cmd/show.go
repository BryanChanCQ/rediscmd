package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/BryanChanCQ/rediscmd/tools"
	"github.com/spf13/cobra"
)

func CreateShowCmd() {
	showCmd := &cobra.Command{
		Use:   "show",
		Short: "show redis database",
		Long:  "show redis database you can use swith database index to choose database",
		Args:  cobra.NoArgs,
		Run:   createShowRunFunc(),
	}
	showCmd.Flags().BoolP("switch", "s", false, "switch database")
	rootCmd.AddCommand(showCmd)
}

func createShowRunFunc() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		var dbShowSlice []redisLogin
		tools.FileUnmarshl(tools.DbSwitchFilePath(), tools.JSONFile, &dbShowSlice)
		for index, dbShow := range dbShowSlice {
			fmt.Printf("index:%d  host  %s  dataBaseName  %s\n", index, dbShow.Host, dbShow.DataBaseName)
		}
		switchFlag, err := cmd.Flags().GetBool("switch")
		if err == nil && switchFlag {
			buffer := bufio.NewScanner(os.Stdin)
			fmt.Print("please input database index:")
			if buffer.Scan() {
				dataBaseIndex := buffer.Text()
				index, err := strconv.ParseInt(dataBaseIndex, 10, 8)
				if err != nil {
					fmt.Println("please input correct index")
				}
				if len(dbShowSlice) > 0 && (int(index) < len(dbShowSlice)) {
					dbShow := dbShowSlice[index]
					WriteLoginData2File(&dbShow)
				}
			}
			if err := buffer.Err(); err != nil {
				fmt.Println("Error reading input:", err)
			}
		}
	}
}
