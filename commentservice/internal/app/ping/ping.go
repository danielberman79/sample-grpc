package ping

import (
	"context"
	"time"

	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
)

type server struct{}

//NewServer creates a new ping.server
func NewServer() *server {
	return &server{}
}

//Watch exists to provide a streaming RPC call for clients to respond with "ping"
func (s *server) Watch(_ *PingRequest, stream PingService_WatchServer) error {
	if err := stream.Send(&PingResponse{Message: "ping"}); err != nil {
		return err
	}

	ticker := time.NewTicker(5 * time.Second)
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

//Check always returns serving, as it has no dependencies.
func (*server) Check() healthgrpc.HealthCheckResponse_ServingStatus {
	return healthgrpc.HealthCheckResponse_SERVING
}
