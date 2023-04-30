#!/bin/bash
cd `dirname $0`
cd ..
export GRPC_HOST=localhost:50051
# export GRPC_HOST=grpc-server:50051
go run grpc_client/cmd/storage/main.go
