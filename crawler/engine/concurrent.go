package engine

import (
	"fmt"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan Item
}
type Scheduler interface {
	ReadyNotify
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotify interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}
	for {
		result := <-out
		fmt.Println("OUT RESULT")
		for _, item := range result.Items {
			go func() { e.ItemChan <- item }()
		}
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, r ReadyNotify) {
	go func() {
		for {
			// worker可以接受新的任务
			r.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
