package main

import (
	"fmt"
)

func main() {
	tmp := 8
	switch tmp > 9 {
	case tmp%2 == 0:
		fmt.Println("right!")
	case tmp*2 == 17:
		fmt.Println("false") //這是結果
		fallthrough
	case tmp == 8:
		fmt.Println("right!") //這也是結果
	}
}
