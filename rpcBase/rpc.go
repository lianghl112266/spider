package rpcBase

import (
	"net"
	"net/rpc"
)

// To facilitate rpc calls, the server here gives the service port
// and service entity
func ServeRPC(host string, service interface{}) error {
	_ = rpc.Register(service)
	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}

	//Constantly listen to the port to obtain the connection. If the
	//connection is obtained correctly, use the coroutine to service
	//the connection.
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}

// It is convenient for the client to monitor the port and directly
// obtain the connection of the calling function.
func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	return rpc.NewClient(conn), nil
}
