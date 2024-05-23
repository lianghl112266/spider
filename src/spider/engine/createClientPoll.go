package engine

import (
	"fmt"
	"net/rpc"
	"spider/src/spider_distribute/rpc_support"
)

// Create a connection pool to continuously send out client connections for workers to use
func CreateClientPoll(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	out := make(chan *rpc.Client)
	for _, host := range hosts {
		client, err := rpc_support.NewClient(host)

		if err != nil {
			fmt.Println(err)
		}
		clients = append(clients, client)
	}

	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()

	return out
}
