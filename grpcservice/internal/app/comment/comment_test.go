package comment

import (
	"context"
	"testing"

	"github.com/djquan/skeleton/grpcservice/internal/platform/database/test"
)

func TestCreateComment(t *testing.T) {
	db, cleanup := test.NewDatabaseForTest(t)
	defer cleanup()

	s := &server{db: &db}
	expectedComment := "hi"
	expectedName := "dan"

	request := &CreateRequest{
		Comment: expectedComment,
		Name:    expectedName,
	}
	result, err := s.Create(context.Background(), request)

	if err != nil {
		t.Fatalf("CommentService/Create returned an error when it was not expected: %v", err)
	}

	if result.Name != request.Name || result.Comment != request.Comment {
		t.Fatalf("Expected %v to include values of %v", result, request)
	}

	dbResult := db.QueryRow(context.Background(), "SELECT name, comment FROM comments where id = $1", result.Id)

	var name, comment string
	err = dbResult.Scan(&name, &comment)

	if err != nil {
		t.Fatalf("Unable to retrieve result from the database: %v", err)
	}

	if name != expectedName || comment != expectedComment {
		t.Fatalf("Expected name: %v, comment: %v, got: name: %v, comment: %v", expectedName, expectedComment, name, comment)
	}
}
