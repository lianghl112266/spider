package engine

import (
	"encoding/gob"
	"fmt"
	"net/rpc"
	"spider/module"
)

// ConcurrentScheduler
// Concurrent scheduler object, including the scheduler,
// the number of worker processes and the pipeline for the
// saver to save data
type ConcurrentScheduler struct {
	Scheduler Scheduler
	WorkerCnt int
	ItemChan  chan interface{}
}

// Scheduler interface
type Scheduler interface {

	// ReadyNotifier Mainly used for queue schedulers.
	//Workers report to the scheduler that they are ready.
	ReadyNotifier

	// Submit The used by scheduler submitting a request
	Submit(Request)

	// WorkChan used by the worker
	WorkChan() chan Request

	//Run scheduler object
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

// Run engine
func (c *ConcurrentScheduler) Run(seeds ...Request) {

	//Worker output pipeline
	out := make(chan ParseResult, 100)

	//Create connection pool
	clients := CreateClientPoll(c.WorkerCnt)
	//Run the scheduler Run the scheduler
	c.Scheduler.Run()

	//Submit seed requests
	for _, r := range seeds {
		c.Scheduler.Submit(r)
	}

	//Create worker coroutine
	for i := 0; i < c.WorkerCnt; i++ {
		CreateWorker(c.Scheduler.WorkChan(), out, c.Scheduler, <-clients)
	}

	//Handle requests output by workers
	for result := range out {
		for _, item := range result.Items {
			go func() { c.ItemChan <- item }()
		}
		for _, r := range result.Requests {
			c.Scheduler.Submit(r)
		}
	}

}

func CreateWorker(in chan Request, out chan<- ParseResult, ready ReadyNotifier, client *rpc.Client) {
	go func() {
		//Because rpc is called later, the interface type may have module.weather,
		//so it needs to be registered in advance.
		gob.Register(module.Weather{})
		for {
			//Reports to the scheduler that it is ready and requests to get the request
			ready.WorkerReady(in)
			r := <-in
			if r.Url == "" {
				continue
			}

			//If the request is legitimate, start the coroutine to complete the work
			go func() {
				res := &ParseResult{}
				if err := client.Call("WorkerService.Worker", r, res); err == nil {
					out <- *res
				} else {
					fmt.Printf("%v\n%v", r, res)
				}
			}()
		}
	}()
}
