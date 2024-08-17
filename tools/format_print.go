package tools

import (
	"fmt"
	"time"
)

type FormatPrint interface {
	PrintTitle()
	PrintRows(index int, value string, time ...time.Duration)
}

type KeysPrint struct {
}

func (keys *KeysPrint) PrintTitle() {
	fmt.Println("index \t\t keyName \t\t expire")
}

func (keys *KeysPrint) PrintRows(index int, value string, time ...time.Duration) {
	if len(time) == 1 {
		fmt.Printf("%v \t\t %v \t\t %v\n", index, value, time[0])
	} else {
		fmt.Printf("%v \t\t %v \t\t %v\n", index, value, "")
	}
}
