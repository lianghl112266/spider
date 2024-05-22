package main

import (
	"spider/src/spider/engine"
	"spider/src/spider/scheduler"
	"spider/src/spider_distribute/config"
	"spider/src/spider_distribute/persist/client"
)

func main() {
	concurrentScheduler := &engine.ConcurrentScheduler{Scheduler: &scheduler.QueuedScheduler{}, WorkerCnt: 1, ItemChan: client.ItemSaver(config.HOST)}
	concurrentScheduler.Run(engine.Request{Url: "https://www.tianqi24.com/historycity/", FuncName: "City"})
}
