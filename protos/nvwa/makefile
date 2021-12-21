.DEFAULT_GOAL := all

all: generate-proto

generate-proto:
	protoc -I ./proto \
      --go_out ./service --go_opt paths=source_relative \
      --go-grpc_out ./service --go-grpc_opt paths=source_relative \
      --grpc-gateway_out ./service --grpc-gateway_opt paths=source_relative \
      ./proto/helloword/*.proto