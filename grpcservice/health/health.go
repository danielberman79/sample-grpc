package health

import (
	"context"
)

type server struct{}

//Check exists to provide an RPC call for clients to determine if the server is healthy
func (*server) Check(_ context.Context, _ *CheckRequest) (*CheckResult, error) {
	return &CheckResult{
		Message: "healthy",
	}, nil
}

//NewHealthService returns a new struct that provides RPC methods for HealthService
func NewHealthService() *server {
	return &server{}
}
