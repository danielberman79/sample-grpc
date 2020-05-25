SHELL := /bin/bash

proto-generate:
	protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative grpcservice/protos/*.proto

ping-test:
	grpcurl -plaintext localhost:8080 ping.PingService/Ping