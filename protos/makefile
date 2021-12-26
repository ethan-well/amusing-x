.DEFAULT_GOAL := all

all: generate-proto-pangu

generate-proto-pangu:
	protoc -I . \
          --go_out ./pangu/service \
          --go_opt paths=source_relative \
          --go-grpc_out ./pangu/service --go-grpc_opt paths=source_relative \
          --grpc-gateway_out ./pangu/service --grpc-gateway_opt paths=source_relative \
          ./pangu/*.proto