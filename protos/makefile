.DEFAULT_GOAL := all

all: generate-proto-pangu generate-charon

generate-proto-pangu:
	protoc -I . \
          --go_out ./pangu/service \
          --go_opt paths=source_relative \
          --go-grpc_out ./pangu/service --go-grpc_opt paths=source_relative \
          --grpc-gateway_out ./pangu/service --grpc-gateway_opt paths=source_relative \
          ./pangu/*.proto

generate-charon:
	protoc -I . \
              --go_out ./charon/service \
              --go_opt paths=source_relative \
              --go-grpc_out ./charon/service --go-grpc_opt paths=source_relative \
              --grpc-gateway_out ./charon/service --grpc-gateway_opt paths=source_relative \
              ./charon/proto/*.proto

genreate-pluto:
	protoc -I . \
              --go_out ./pluto/service \
              --go_opt paths=source_relative \
              --go-grpc_out ./pluto/service --go-grpc_opt paths=source_relative \
              --grpc-gateway_out ./pluto/service --grpc-gateway_opt paths=source_relative \
              ./pluto/proto/*.proto

genreate-ganymede:
	protoc -I . \
              --go_out ./ganymede/service \
              --go_opt paths=source_relative \
              --go-grpc_out ./ganymede/service --go-grpc_opt paths=source_relative \
              --grpc-gateway_out ./ganymede/service --grpc-gateway_opt paths=source_relative \
              ./ganymede/proto/*.proto