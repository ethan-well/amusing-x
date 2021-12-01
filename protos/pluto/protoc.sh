#!/bin/bash

echo "protoc code generate"
protoc --go_out=plutoservice --go_opt=paths=source_relative --go-grpc_out=plutoservice --go-grpc_opt=paths=source_relative service.proto inventory_*.proto

echo "succeed"