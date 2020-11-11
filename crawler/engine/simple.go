package engine

import (
	"log"
)

type SimpleEngine struct {
}

func (s *SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, s := range seeds {
		requests = append(requests, s)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		result, _ := worker(r)
		requests = append(requests, result.Requests...)
		for _, item := range result.Items {
			log.Printf("Got Item:%v", item)
		}
	}
}
