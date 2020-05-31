package health

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"

	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
)

//Checker is an interface that looks for a method Check, which returns a Serving Status
type Checker interface {
	Check() healthgrpc.HealthCheckResponse_ServingStatus
}

type server struct {
	serviceMap map[string]Checker
}

//NewServer creates a health.server object
func NewServer(serviceMap map[string]Checker) *server {
	return &server{serviceMap: serviceMap}
}

//Check will provide the status of a requested service. If no service is requested, it will return serving.
func (s *server) Check(_ context.Context, request *healthgrpc.HealthCheckRequest) (*healthgrpc.HealthCheckResponse, error) {
	if request.Service == "" {
		return &healthgrpc.HealthCheckResponse{
			Status: healthgrpc.HealthCheckResponse_SERVING,
		}, nil
	}

	if checker, ok := s.serviceMap[request.Service]; ok {
		return &healthgrpc.HealthCheckResponse{
			Status: checker.Check(),
		}, nil
	}

	return nil, status.Errorf(codes.Unknown, "Unknown service: %v", request.Service)
}

//Watch will monitor a requested service and respond with the updated status when it changes
func (s *server) Watch(request *healthgrpc.HealthCheckRequest, stream healthgrpc.Health_WatchServer) error {
	response, err := s.Check(context.Background(), request)
	if err != nil {
		return err
	}

	if err := stream.Send(response); err != nil {
		return err
	}

	c := make(chan healthgrpc.HealthCheckResponse_ServingStatus)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	go func() {
		responseStatus := response.Status
		for {
			select {
			case <-ticker.C:
				updatedResponse := s.serviceMap[request.Service].Check()
				if updatedResponse != responseStatus {
					c <- updatedResponse
					responseStatus = updatedResponse
				}
			}
		}
	}()

	for {
		select {
		case <-stream.Context().Done():
			return nil
		case s := <-c:
			if err := stream.SendMsg(&healthgrpc.HealthCheckResponse{Status: s}); err != nil {
				return err
			}
		}
	}
}
