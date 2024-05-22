package scheduler

import (
	"spider/src/spider/engine"
)

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (q *QueuedScheduler) Submit(r engine.Request) {
	q.requestChan <- r
}

func (q *QueuedScheduler) WorkChan() chan engine.Request {
	return make(chan engine.Request)
}

func (q *QueuedScheduler) WorkerReady(w chan engine.Request) {
	q.workerChan <- w
}

func (q *QueuedScheduler) Run() {
	q.requestChan = make(chan engine.Request)
	q.workerChan = make(chan chan engine.Request)
	go func() {
		var requests []engine.Request
		var workers []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requests) > 0 && len(workers) > 0 {
				activeRequest = requests[0]
				activeWorker = workers[0]
			}
			select {
			case r := <-q.requestChan:
				requests = append(requests, r)
			case w := <-q.workerChan:
				workers = append(workers, w)
			case activeWorker <- activeRequest:
				requests = requests[1:]
				workers = workers[1:]
			}
		}
	}()
}
