package health

import (
	"context"
)

type server struct{}

func (*server) Check(_ context.Context, _ *CheckRequest) (*CheckResult, error) {
	return &CheckResult{
		Message: "healthy",
	}, nil
}

func NewHealthService() *server {
	return &server{}
}
