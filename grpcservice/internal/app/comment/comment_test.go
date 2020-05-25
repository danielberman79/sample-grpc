package comment

import (
	"context"
	"github.com/djquan/skeleton/grpcservice/internal/platform/database/test"
	"testing"
)

func TestCreateComment(t *testing.T) {
	_, cleanup := test.NewDatabaseForTest(t)
	s := &server{db: make(map[string]Response)}
	request := &CreateRequest{
		Comment: "hi",
		Name:    "dan",
	}
	result, err := s.Create(context.Background(), request)

	if err != nil {
		t.Fatalf("CommentService/Create returned an error when it was not expected: %v", err)
	}

	if result.Name != request.Name || result.Comment != request.Comment {
		t.Fatalf("Expected %v to include values of %v", result, request)
	}

	cleanup()
}
