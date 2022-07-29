package main

import "fmt"

func main() {
	i := 1
	num(i)       //印出1
	defer num(i) //這時i等於1,所以最後執行時印出1,不會因為後面i++而改變
	i++
	num(i) //印出2
}

func num(n int) {
	fmt.Println(n)
}
