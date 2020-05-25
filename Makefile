SHELL := /bin/bash

proto-generate:
	protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative grpcservice/*/*.proto

grpc-healthcheck:
	grpcurl -plaintext localhost:8080 grpc.health.v1.Health/Check

grpc-ping:
	grpcurl -plaintext localhost:8080 ping.PingService/Ping

grpc-ping-watch:
	grpcurl -plaintext localhost:8080 ping.PingService/Watch
