package main

import (
	"encoding/gob"
	"spider/src/spider/module"
	"spider/src/spider_distribute/config"
	"spider/src/spider_distribute/rpc_support"
	"spider/src/spider_distribute/worker"
)

// worker server entrance
func main() {
	gob.Register(module.Weather{})
	//go func() { _ = rpc_support.ServeRPC(config.WORKERHOST1, &worker.WorkerService{}) }()
	//go func() { _ = rpc_support.ServeRPC(config.WORKERHOST2, &worker.WorkerService{}) }()
	//for {
	//}
	_ = rpc_support.ServeRPC(config.WORKERHOST1, &worker.WorkerService{})
}
