package health

import (
	"context"

	"google.golang.org/grpc"
)

type server struct{}

//Check exists to provide an RPC call for clients to determine if the server is healthy.
func (*server) Check(_ context.Context, _ *CheckRequest) (*CheckResult, error) {
	return &CheckResult{
		Message: "healthy",
	}, nil
}

//RegisterWithServer accepts a pointer to a grpc.Server and registers the RPC backend.
func RegisterWithServer(s *grpc.Server) {
	RegisterHealthServiceServer(s, &server{})
}
