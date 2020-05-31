package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"runtime"

	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/djquan/skeleton/commentservice/internal/app/health"

	"github.com/djquan/skeleton/commentservice/internal"
	"github.com/djquan/skeleton/commentservice/internal/app/comment"
	"github.com/djquan/skeleton/commentservice/internal/app/ping"
	"github.com/djquan/skeleton/commentservice/internal/platform/database"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	runtime.SetBlockProfileRate(1)
	config := internal.ReadConfig()
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", config.Server.Host, config.Server.Port))
	if err != nil {
		log.Fatalf("oh nos %v", err)
	}

	db, err := database.FromConfig(config.Database)

	if err != nil {
		log.Fatalf("unable to talk to database %v\n", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(unaryInterceptor), grpc.StreamInterceptor(streamInterceptor))
	setupServer(s, db)

	go func() {
		err := http.ListenAndServe("localhost:9080", nil)

		if err != nil {
			log.Printf("pprof server returned an err: %v", err)
		}
	}()

	log.Println("Beginning to Serve grpc traffic on port 8080")
	if err = s.Serve(lis); err != nil {
		log.Fatalf("oh nos %v", err)
	}
}

func setupServer(s *grpc.Server, db *database.Database) {
	commentServer := comment.NewServer(db)
	comment.RegisterCommentServiceServer(s, commentServer)

	pingServer := ping.NewServer()
	ping.RegisterPingServiceServer(s, pingServer)

	healthServer := health.NewServer(map[string]health.Checker{
		"ping.PingService":       pingServer,
		"comment.CommentService": commentServer,
	})

	healthgrpc.RegisterHealthServer(s, healthServer)
	reflection.Register(s)
}
