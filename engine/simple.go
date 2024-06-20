package engine

import (
	"fmt"
)

type SimpleScheduler struct {
}

// Run Simple serial execution engine
func (_ *SimpleScheduler) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	//s := &SimpleScheduler{}
	for len(requests) > 0 {
		//for i := 0; i < 10; i++ {
		r := requests[0]
		requests = requests[1:]
		if parseResult, err := Worker(r); err == nil {
			requests = append(requests, parseResult.Requests...)
			for _, item := range parseResult.Items {
				fmt.Println(item)
			}
		} else {
			fmt.Println("Fetch err", err)
		}
	}
}
