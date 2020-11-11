package main

import (
	"fmt"
	"time"
)

func chanDemo() {
	c := make(chan int)
	go func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()
	c <- 1
	c <- 2
}
func worker(id int, c chan int) {
	fmt.Println(<-c)
	for {
		fmt.Printf("Worker：%d, Receive：%d\n", id, <-c)
	}
}
func chanDemo1() {
	c := make(chan int, 3)
	c <- 4
	c <- 2
	c <- 1
	c <- 5
	go worker(0, c)
	time.Sleep(time.Millisecond * 10)
}
func main() {
	//chanDemo()
	chanDemo1()
}
