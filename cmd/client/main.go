package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/KETAKOM/grpc-sample-app/proto/ping"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()

	client := pb.NewPingClient(conn)
	message := &pb.HelloRequest{ToMessage: "hello world ikerukana"}
	res, err := client.Hello(context.TODO(), message)
	fmt.Printf("result:%#v \n", res)
	fmt.Printf("error::%#v \n", err)

}
