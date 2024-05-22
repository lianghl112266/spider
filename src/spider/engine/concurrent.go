package engine

type ConcurrentScheduler struct {
	Scheduler Scheduler
	WorkerCnt int
	ItemChan  chan interface{}
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (c *ConcurrentScheduler) Run(seeds ...Request) {
	out := make(chan ParseResult, 100)
	c.Scheduler.Run()
	for _, r := range seeds {
		c.Scheduler.Submit(r)
	}

	for i := 0; i < c.WorkerCnt; i++ {
		CreateWorker(c.Scheduler.WorkChan(), out, c.Scheduler)
	}

	for result := range out {
		for _, item := range result.Items {
			go func() { c.ItemChan <- item }()
		}
		for _, r := range result.Requests {
			c.Scheduler.Submit(r)
		}
	}

}

func CreateWorker(in chan Request, out chan<- ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			r := <-in
			if p, err := Worker(r); err == nil {
				out <- p
			}
		}
	}()
}
