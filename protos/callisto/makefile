.DEFAULT_GOAL := all

all: generate-proto

generate-proto:
	protoc --go_out=service --go_opt=paths=source_relative --go-grpc_out=service --go-grpc_opt=paths=source_relative service.proto loginrisk.proto pong.proto
