package comment

import (
	"context"

	"github.com/djquan/skeleton/commentservice/internal/platform/database"

	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	db *database.Database
}

//Create provides an RPC call that creates a comment
func (s *server) Create(_ context.Context, request *CreateRequest) (*Response, error) {
	response := Response{
		Id:        uuid.New().String(),
		Comment:   request.Comment,
		Name:      request.Name,
		CreatedAt: ptypes.TimestampNow(),
	}

	_, err := s.db.Exec(
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

//Register accepts a pointer to a grpc.Server and registers the gRPC backend
func Register(s *grpc.Server, db *database.Database) {
	RegisterCommentServiceServer(s, &server{db: db})
}
