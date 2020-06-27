package comment

import (
	"context"

	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/djquan/skeleton/internal/platform/database"

	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
)

//Server provides an implementation for the Comment GRPC api.
type server struct {
	Db *database.Database
}

//NewServer is a constructor for a server object
func NewServer(db *database.Database) *server {
	return &server{Db: db}
}

//Create provides an RPC call that creates a comment
func (s *server) Create(_ context.Context, request *CreateRequest) (*Response, error) {
	response := Response{
		Id:        uuid.New().String(),
		Comment:   request.Comment,
		Name:      request.Name,
		CreatedAt: ptypes.TimestampNow(),
	}

	_, err := s.Db.Exec(
		context.Background(),
		"INSERT INTO comments (id, comment, name, created_at) VALUES ($1, $2, $3, $4)",
		response.Id,
		response.Comment,
		response.Name,
		ptypes.TimestampString(response.CreatedAt),
	)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

//Check will return SERVING unless there is a problem connecting to the database.
func (s *server) Check() healthgrpc.HealthCheckResponse_ServingStatus {
	if _, err := s.Db.Exec(context.Background(), "SELECT true"); err != nil {
		return healthgrpc.HealthCheckResponse_NOT_SERVING
	}

	return healthgrpc.HealthCheckResponse_SERVING
}
