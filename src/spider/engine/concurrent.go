package engine

import (
	"encoding/gob"
	"fmt"
	"net/rpc"
	"spider/src/spider/module"
	"spider/src/spider_distribute/config"
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

// scheduler interface
type Scheduler interface {

	//Mainly used for queue schedulers.
	//Workers report to the scheduler that they are ready.
	ReadyNotifier

	//The scheduler submits a request
	Submit(Request)

	//The pipe used by the worker
	WorkChan() chan Request

	//Run scheduler object
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

// run engine
func (c *ConcurrentScheduler) Run(seeds ...Request) {

	//Worker output pipeline
	out := make(chan ParseResult, 100)

	//Create connection pool
	clients := CreateClientPoll([]string{config.WORKERHOST1, config.WORKERHOST2})
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
					fmt.Println(err)
					fmt.Printf("%v\n%v", r, res)
				}
			}()
		}
	}()
}
