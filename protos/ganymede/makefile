.DEFAULT_GOAL := all

all: generate-proto

generate-proto:
	protoc --go_out=service --go_opt=paths=source_relative --go-grpc_out=service --go-grpc_opt=paths=source_relative server.proto server.proto login.proto join.proto country_code.proto verification_code.proto verification_code_check.proto reset_password.proto
