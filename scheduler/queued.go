package scheduler

import (
	"spider/engine"
)

// QueuedScheduler with its own worker queue and request queue
type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (q *QueuedScheduler) Submit(r engine.Request) {
	q.requestChan <- r
}

// WorkChan give each worker a new pipeline to facilitate subsequent management
func (q *QueuedScheduler) WorkChan() chan engine.Request {
	return make(chan engine.Request)
}

func (q *QueuedScheduler) WorkerReady(w chan engine.Request) {
	q.workerChan <- w
}

func (q *QueuedScheduler) Run() {
	//Init pipe
	q.requestChan = make(chan engine.Request)
	q.workerChan = make(chan chan engine.Request)

	//Run scheduler
	go func() {
		var requests []engine.Request
		var workers []chan engine.Request
		for {

			//Get requested and ready workers
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requests) > 0 && len(workers) > 0 {
				activeRequest = requests[0]
				activeWorker = workers[0]
			}

			//Because we don’t know which worker or request comes first,
			//we use select
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
