package main

import (
	"spider/src/spider/engine"
	"spider/src/spider/persist"
	"spider/src/spider/scheduler"
)

// Concurrent version of crawler function entrance
func main() {
	concurrentScheduler := &engine.ConcurrentScheduler{Scheduler: &scheduler.QueuedScheduler{}, WorkerCnt: 1, ItemChan: persist.ItemSaver()}
	concurrentScheduler.Run(engine.Request{Url: "https://www.tianqi24.com/historycity/", FuncName: "City"})
}
