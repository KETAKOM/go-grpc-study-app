package service

import (
	"context"
	"fmt"

	"github.com/KETAKOM/grpc-sample-app/proto/ping"
)

type PingService struct{}

func (s *PingService) Hello(ctx context.Context, req *ping.HelloRequest) (*ping.HelloResponse, error) {
	toMessage := req.GetToMessage()
	fmt.Println(toMessage)
	response := ping.HelloResponse{
		ResMessage: toMessage,
	}

	return &response, nil
}
