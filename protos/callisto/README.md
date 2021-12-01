# proto to go & grpc code

protoc --go_out=riskservice --go_opt=paths=source_relative --go-grpc_out=riskservice --go-grpc_opt=paths=source_relative service.proto loginrisk.proto pong.proto