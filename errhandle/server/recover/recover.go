package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("拦截到了", err)
		} else {
			panic(r)
		}
	}()
	a := []int{1, 2}
	fmt.Println(a[2])
	//panic(errors.New("报错啦"))
}
func main() {
	tryRecover()
}
