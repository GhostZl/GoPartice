package main

import "fmt"

type AA struct {
	value int
	key string
}

func changeAA(a *AA) {
	a.key = "123"
	a.value = 456
}
func main() {
	aa := AA{
		value: 1,
		key:   "1",
	}
	changeAA(&aa)
	fmt.Println(aa)
}
