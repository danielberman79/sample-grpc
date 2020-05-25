package ping

import (
	"context"
	"time"

	"google.golang.org/grpc"
)

type server struct{}

//Watch exists to provide a streaming RPC call for clients to respond with "ping"
func (s *server) Watch(_ *PingRequest, stream PingService_WatchServer) error {
	if err := stream.Send(&PingResponse{Message: "ping"}); err != nil {
		return err
	}

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-stream.Context().Done():
			return nil
		case <-ticker.C:
			if err := stream.Send(&PingResponse{Message: "ping"}); err != nil {
				return err
			}
		}
	}
}

//Ping exists to provide a RPC call for clients to respond with "ping"
func (*server) Ping(_ context.Context, _ *PingRequest) (*PingResponse, error) {
	return &PingResponse{
		Message: "ping",
	}, nil
}

//Register accepts a pointer to a grpc.Server and registers the RPC backend.
func Register(s *grpc.Server) {
	RegisterPingServiceServer(s, &server{})
}
