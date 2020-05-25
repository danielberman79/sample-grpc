package main

import (
	"context"
	"github.com/djquan/skeleton/grpcservice/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type server struct{}

func (*server) Ping(_ context.Context, _ *ping.PingRequest) (*ping.PingResult, error) {
	return &ping.PingResult{
		Message:   "ping",
		CreatedAt: nil,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("oh nos %v", err)
	}

	s := grpc.NewServer()
	ping.RegisterPingServiceServer(s, &server{})
	reflection.Register(s)

	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("oh nos %v", err)
	}
}
