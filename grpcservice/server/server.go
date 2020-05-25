package main

import (
	"log"
	"net"

	"github.com/djquan/skeleton/grpcservice/ping"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
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
	healthgrpc.RegisterHealthServer(s, health.NewServer())
	ping.Register(s)
	reflection.Register(s)
	log.Println("Beginning to Serve grpc traffic on port 8080")
	return s
}
