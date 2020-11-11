package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generate() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}
func worker(id int, c chan chanValue) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d from %v \n", id, n.value, n.chanName)
	}
}
func createWorker(id int) chan<- chanValue {
	c := make(chan chanValue)
	go worker(id, c)
	return c
}

type chanValue struct {
	value    int
	chanName string
}

func main() {
	var c1, c2 = generate(), generate()
	w := createWorker(0)
	var values []chanValue
	n := 0
	after := time.After(time.Second * 10)
	ticker := time.Tick(time.Second)
	for {
		var activeWorker chan<- chanValue
		var activeValue chanValue
		if len(values) > 0 {
			activeValue = values[0]
			activeWorker = w
		}
		select {
		case n = <-c1:
			values = append(values, chanValue{
				value:    n,
				chanName: "c1",
			})
		case n = <-c2:
			values = append(values, chanValue{
				value:    n,
				chanName: "c2",
			})
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(time.Millisecond * 800):
			fmt.Println("timeout")
		case <-ticker:
			fmt.Println("Queue Length:", len(values))
		case <-after:
			fmt.Println("bye")
			return
		}
	}
}
