package client

import (
	"fmt"
	"spider/module"
	"spider/rpcBase"
)

// ItemSaver called by the distributed persistence layer
// receives data from out
func ItemSaver(host string) chan interface{} {
	out := make(chan interface{})

	//Create table
	client, err := rpcBase.NewClient(host)
	if err != nil {
		fmt.Println(err)
		return out
	}

	go func() {
		for {
			//Continuously read data from the out pipe. If the data needs to be saved,
			//call the remote function.
			item := <-out
			if it, ok := item.(module.Weather); ok {
				go func() { _ = client.Call("ItemSaverService.Save", it, "") }()
			}
		}
	}()

	return out
}
