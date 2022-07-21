package main

import (
	"fmt"
	"time"
)

func main() {
	var count1 *int    //宣告一個指標變數(值為為nil)
	count2 := new(int) //以new在於為int形別取得記憶體/填入該型別的零值 , 然後傳回記憶體的指標
	countTemp := 5
	count3 := &countTemp //從別的變數建立指標
	t := &time.Time{}    //從結構型別建立指標

	fmt.Printf("count1: %#v\n", count1)
	fmt.Printf("count2: %#v\n", count2)
	fmt.Printf("count3: %#v\n", count3)
	fmt.Printf("time  : %#v\n", t)
}
