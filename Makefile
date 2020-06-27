SHELL := /bin/bash

proto-generate:
	for proto in api/*.proto; do \
	  protoc --go_out=plugins=grpc:. $$proto ;\
	done

run-server:
	go run github.com/djquan/skeleton/cmd/server

build-server:
	go build github.com/djquan/skeleton/cmd/server

migrate-db:
	go run github.com/djquan/skeleton/cmd/admin --command migrate

reset-db:
	go run github.com/djquan/skeleton/cmd/admin --command reset

grpc-healthcheck:
	grpcurl -plaintext localhost:8080 grpc.health.v1.Health/Check
	grpcurl -d '{"service": "ping.PingService"}' -plaintext localhost:8080 grpc.health.v1.Health/Check
	grpcurl -d '{"service": "comment.CommentService"}' -plaintext localhost:8080 grpc.health.v1.Health/Check
	grpcurl -d '{"service": "doesnotexist.DOESNOTEXISt"}' -plaintext localhost:8080 grpc.health.v1.Health/Check || true
	grpcurl -d '{"service": "comment.CommentService"}' -plaintext localhost:8080 grpc.health.v1.Health/Watch

grpc-ping:
	grpcurl -plaintext localhost:8080 ping.PingService/Ping

grpc-ping-watch:
	grpcurl -plaintext localhost:8080 ping.PingService/Watch

grpc-comment-create:
	grpcurl -plaintext -d '{"comment": "hi", "name": "dan"}' localhost:8080 comment.CommentService/Create
