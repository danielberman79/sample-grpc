SHELL := /bin/bash

proto-generate:
	protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative grpcservice/*/*.proto

grpc-healthcheck:
	grpcurl -plaintext localhost:8080 health.HealthService/Check