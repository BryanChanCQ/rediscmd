package tools

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	JSON = "json"
)

func InitFilePath() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		panic("can not get user home path")
	}
	filePath := dir + "/.loginInfo.json"
	return filePath
}

// 将当前文件路径下的内容，反序列化成结构体
// filepath 目录文件 format 目前只支持json
func FileUnmarshl(filepath string, format string, resp any) bool {
	//TODO:判断resp是否是指针,如果resp 非指针类型则反序列化会失败
	if !FileIsExists(filepath) {
		return false
	}
	content, err := os.ReadFile(filepath)
	if err != nil {
		return false
	}

	if format == JSON {
		err := json.Unmarshal(content, &resp)
		if err != nil {
			fmt.Printf("read login file err:%v\n", err)
			return false
		}
	}
	return true
}

func FileIsExists(filepath string) bool {
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}
