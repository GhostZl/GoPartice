package scheduler

import (
	"GoPartice/crawler/engine"
)

type QueueScheduler struct {
	requestChan chan engine.Request
	workChan    chan chan engine.Request
}

func (q *QueueScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (q *QueueScheduler) Submit(r engine.Request) {
	q.requestChan <- r
}

func (q *QueueScheduler) WorkerReady(w chan engine.Request) {
	q.workChan <- w
}

func (q *QueueScheduler) Run() {
	q.workChan = make(chan chan engine.Request)
	q.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWork chan engine.Request
			if len(requestQ) > 0 && len(workQ) > 0 {
				activeWork = workQ[0]
				activeRequest = requestQ[0]
			}
			select {
			case r := <-q.requestChan:
				requestQ = append(requestQ, r)
			case w := <-q.workChan:
				workQ = append(workQ, w)
			case activeWork <- activeRequest:
				requestQ = requestQ[1:]
				workQ = workQ[1:]
			}
		}
	}()
}
