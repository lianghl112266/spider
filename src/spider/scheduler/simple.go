package scheduler

import (
	"spider/src/spider/engine"
)

// This simple concurrent scheduler lets all workers go to a request
// pipeline to grab data
type SimpleScheduler struct {
	in chan engine.Request
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() { s.in <- r }()
}

func (s *SimpleScheduler) WorkChan() chan engine.Request {
	return s.in
}

func (s *SimpleScheduler) Run() {
	s.in = make(chan engine.Request)
}

func (s *SimpleScheduler) WorkerReady(_ chan engine.Request) {
}
