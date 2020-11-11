package main

import (
	"GoPartice/crawler/config"
	"GoPartice/crawler/engine"
	"GoPartice/crawler/persist"
	"GoPartice/crawler/scheduler"
	"GoPartice/crawler/xcar/parser"
)

func main() {
	itemChan, err := persist.ItemSaver(
		config.ElasticIndex)
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}
	e.Run(engine.Request{
		Url:       "http://newcar.xcar.com.cn",
		ParseFunc: parser.ParseCarList,
	})
}
