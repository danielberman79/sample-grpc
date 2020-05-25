package health

import (
	"context"

	"google.golang.org/grpc"
)

type server struct{}

//Watch exists to provide a streaming RPC call for clients to determine if the server is healthy.
func (s *server) Watch(_ *CheckRequest, stream HealthService_WatchServer) error {
	if err := stream.Send(&CheckResponse{Status: CheckResponse_SERVING}); err != nil {
		return err
	}

	return nil
}

//Check exists to provide a RPC call for clients to determine if the server is healthy.
func (*server) Check(_ context.Context, _ *CheckRequest) (*CheckResponse, error) {
	return &CheckResponse{
		Status: CheckResponse_SERVING,
	}, nil
}

//RegisterWithServer accepts a pointer to a grpc.Server and registers the RPC backend.
func RegisterWithServer(s *grpc.Server) {
	RegisterHealthServiceServer(s, &server{})
}
