package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str interface{} = 42
	if s, ok := str.(string); ok { //做型別斷言 , 並檢查是否轉換成功
		fmt.Println(strconv.Atoi(s))
	} else {
		fmt.Println("Type assertion failed")
	}
}
