.DEFAULT_GOAL := all

all: generate-proto

generate-proto:
	protoc -I . --grpc-gateway_out service \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    *.proto