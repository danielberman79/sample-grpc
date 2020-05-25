package main

import (
	"log"
	"net"

	"github.com/djquan/skeleton/grpcservice/health"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("oh nos %v", err)
	}

	s := newServer()
	if err = s.Serve(lis); err != nil {
		log.Fatalf("oh nos %v", err)
	}
}

func newServer() *grpc.Server {
	s := grpc.NewServer()
	health.RegisterWithServer(s)
	reflection.Register(s)
	return s
}
