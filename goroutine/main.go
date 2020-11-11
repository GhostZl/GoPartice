package main

import (
	"fmt"
)

func main() {
	var b int
	for i := 0; i < 1000; i++ {
		go func(i int) {
			fmt.Println("route:", i)
		}(i)
	}
	for i := 0; i < 1000; i++ {
		b++
	}
	fmt.Println(23)
	return

}
