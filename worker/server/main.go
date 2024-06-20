package main

import (
	"encoding/gob"
	"spider/module"
	"spider/rpcBase"
	"spider/worker"
)

// worker server entrance
func main() {
	gob.Register(module.Weather{})
	_ = rpcBase.ServeRPC(":1234", &worker.WorkerService{})
}
