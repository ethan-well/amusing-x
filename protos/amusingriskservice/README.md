# proto to go & grpc code

protoc --go_out=loginrisk --go_opt=paths=source_relative --go-grpc_out=loginrisk --go-grpc_opt=paths=source_relative loginrisk/loginrisk.proto
