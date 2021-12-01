#!/bin/bash

echo "protoc code generate"
protoc --go_out=charonservice --go_opt=paths=source_relative --go-grpc_out=charonservice --go-grpc_opt=paths=source_relative service.proto *.proto

echo "succeed"