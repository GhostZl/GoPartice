package main

import "fmt"

type S struct {
	A uint64
	B uint32
	//C uint32
	//D complex128
	E []int64
}

//func main() {
//	e := S{
//		E: []int64{1, 2, 3},
//	}
//	fmt.Println(unsafe.Sizeof(e.E))
//	fmt.Println(unsafe.Sizeof(e))
//}
func main() {
	defer_call()
}
func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
	fmt.Println("123")
}
