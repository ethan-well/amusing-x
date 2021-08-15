# proto to go & grpc code

cd amusing-x/protos/amusingxuserserv

protoc --go_out=userservice --go_opt=paths=source_relative --go-grpc_out=userservice --go-grpc_opt=paths=source_relative server.proto server.proto login.proto join.proto country_code.proto verification_code.proto verification_code_check.proto reset_password.proto 