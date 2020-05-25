package main

import (
	"context"
	"fmt"
	"github.com/djquan/skeleton/grpcservice/comment"
	"github.com/jackc/pgx/v4"
	"log"
	"net"

	"github.com/djquan/skeleton/grpcservice/ping"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func main() {
	config := readConfig()
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", config.Server.Host, config.Server.Port))
	if err != nil {
		log.Fatalf("oh nos %v", err)
	}

	_, err = pgx.Connect(context.Background(), config.Database.Url())

	if err != nil {
		log.Fatalf("unable to talk to databaseName %v\n", err)
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
