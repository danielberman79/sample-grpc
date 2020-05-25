package main

import (
	"context"
	"log"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	requestId := uuid.New().String()
	ctx = context.WithValue(ctx, "requestId", requestId)

	log.Printf("[%v] Received request: %v with params %v\n", requestId, info.FullMethod, req)
	m, err := handler(ctx, req)
	if err != nil {
		log.Printf("[%v] Failed with errors, %v\n", requestId, err)
	}

	log.Printf("[%v] Sending response: %v\n", requestId, m)
	return m, err
}

func streamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	requestId := uuid.New().String()
	log.Printf("[%v] Starting stream for %v\n", requestId, info.FullMethod)
	err := handler(srv, &wrappedStream{ServerStream: ss, requestId: requestId})

	if err != nil {
		log.Printf("[%v] GRPC request failed with %v\n", requestId, err)
	}
	return err
}

type wrappedStream struct {
	grpc.ServerStream
	requestId string
}

//SendMsg logs the response and calls the original ServerStream
func (w *wrappedStream) SendMsg(res interface{}) error {
	log.Printf("[%v] Sending response: %v\n", w.requestId, res)
	return w.ServerStream.SendMsg(res)
}

//RecvMsg logs the request and calls the original ServerStream
func (w *wrappedStream) RecvMsg(req interface{}) error {
	log.Printf("[%v] Received request: %v\n", w.requestId, req)
	return w.ServerStream.RecvMsg(req)
}

//Context adds request ID to the wrapped ServerStream
func (w *wrappedStream) Context() context.Context {
	return context.WithValue(w.ServerStream.Context(), "requestId", w.requestId)
}
