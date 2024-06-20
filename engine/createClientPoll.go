package engine

import (
	"github.com/spf13/viper"
	"net/rpc"
	"spider/rpcBase"
)

// CreateClientPoll create a connection pool to continuously send out client connections for workers to use
func CreateClientPoll(workerNum int) chan *rpc.Client {
	var clients []*rpc.Client
	out := make(chan *rpc.Client)
	hosts := viper.GetStringSlice("worker.ports")
	for _, host := range hosts {
		client, _ := rpcBase.NewClient(host)
		clients = append(clients, client)
	}

	go func() {
		n := len(clients)
		for i := 0; i < workerNum; i++ {
			out <- clients[i%n]
		}
	}()

	return out
}
