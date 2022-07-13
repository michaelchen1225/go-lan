package main

import (
	"fmt"
)

func main() {
	sl1 := []int{1, 3}
	sl2 := []int{555, 666}
	sl1 = append(sl1, sl2...)
	fmt.Println(sl1)

}
