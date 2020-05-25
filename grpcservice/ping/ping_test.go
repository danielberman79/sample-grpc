package ping

import (
	"context"
	"testing"
)

func TestUnaryPing(t *testing.T) {
	s := &server{}
	got, err := s.Ping(context.Background(), &PingRequest{})

	if err != nil {
		t.Fatalf("PingService/Ping returned an error when it was not expected: %v", err)
	}

	if got.Message != "ping" {
		t.Fatalf("PingService/Ping returned %v for it's message, but expected ping", got.Message)
	}
}
