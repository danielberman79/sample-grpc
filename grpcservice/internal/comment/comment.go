package comment

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	db map[string]Response
}

//Create provides an RPC call that creates a comment
func (s *server) Create(_ context.Context, request *CreateRequest) (*Response, error) {
	response := Response{
		Id:        uuid.New().String(),
		Comment:   request.Comment,
		Name:      request.Name,
		CreatedAt: ptypes.TimestampNow(),
	}

	s.db[response.Id] = response

	return &response, nil
}

//Register accepts a pointer to a grpc.Server and registers the gRPC backend
func Register(s *grpc.Server) {
	RegisterCommentServiceServer(s, &server{db: make(map[string]Response)})
}
