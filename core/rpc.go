package core

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/agaUHO/aga/models"
	"github.com/agaUHO/aga/system"
)

type OrderServer string

func ServerRPC() {
	_ = rpc.Register(new(OrderServer))
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":"+system.PortServerRPC)
	if err != nil {
		log.Fatal("listen error: ", err)
	}
	_ = http.Serve(listener, nil)
}

func (t *OrderServer) CreateLogs(args *models.Logs, reply *string) error {
	system.Log <- *args
	*reply = "Create log ok"
	return nil
}
