package main

import (
	"fmt"
	"log"
	"net"

	"github.com/KETAKOM/grpc-sample-app/proto/ping"
	"github.com/KETAKOM/grpc-sample-app/service"
	"google.golang.org/grpc"
)

var port = ":19003"

func main() {
	lp, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	pingService := service.PingService{}
	ping.RegisterPingServer(server, &pingService)

	fmt.Printf("[server started] localhost%s", port)
	err = server.Serve(lp)
	if err != nil {
		log.Fatalln(err)
	}
}
