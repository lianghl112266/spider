package main

import (
	"github.com/spf13/viper"
	"spider/config"
	"spider/engine"
	"spider/persist/client"
	"spider/scheduler"
)

func init() {
	if err := config.InitConfig(); err != nil {
		panic(err.Error())
	}
}

// Entry point
func main() {
	concurrentScheduler := &engine.ConcurrentScheduler{Scheduler: &scheduler.QueuedScheduler{}, WorkerCnt: 1, ItemChan: client.ItemSaver(viper.GetString("persist.port"))}
	concurrentScheduler.Run(engine.Request{Url: "https://www.tianqi24.com/historycity/", FuncName: "City"})
}
