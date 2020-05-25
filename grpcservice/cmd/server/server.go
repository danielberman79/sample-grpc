package main

import (
	"fmt"
	"github.com/djquan/skeleton/grpcservice/internal"
	"github.com/djquan/skeleton/grpcservice/internal/platform/database"
	"log"
	"net"

	"github.com/djquan/skeleton/grpcservice/internal/app/comment"
	"github.com/djquan/skeleton/grpcservice/internal/app/ping"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func main() {
	config := internal.ReadConfig()
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", config.Server.Host, config.Server.Port))
	if err != nil {
		log.Fatalf("oh nos %v", err)
	}

	_, err = database.FromConfig(config.Database)

	if err != nil {
		log.Fatalf("unable to talk to database %v\n", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(unaryInterceptor), grpc.StreamInterceptor(streamInterceptor))
	setupServer(s)
	log.Println("Beginning to Serve grpc traffic on port 8080")
	if err = s.Serve(lis); err != nil {
		log.Fatalf("oh nos %v", err)
	}
}

func setupServer(s *grpc.Server) {
	healthgrpc.RegisterHealthServer(s, health.NewServer())
	ping.Register(s)
	comment.Register(s)
	reflection.Register(s)
}
