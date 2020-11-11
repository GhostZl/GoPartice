package main

import (
	"fmt"
)

func tryDefer() {
	for i := 0; i < 30; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("FINISH")
}
func main() {
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//fmt.Println(3)
	tryDefer()
}
