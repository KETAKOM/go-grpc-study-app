package main

import (
	"fmt"
	"log"
	"net"

	"github.com/KETAKOM/grpc-sample-app/proto/ping"

	"google.golang.org/grpc"
)

var port = ":19003"

func main() {
	lp, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	pingService := service.pingService()
	ping.RegisterPingServer(server, &pingService)

	fmt.Println("[server started] localhost%s", port)
	if err != nil {
		log.Fatalln(err)
	}
}
