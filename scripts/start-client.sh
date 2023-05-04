#!/bin/bash
cd `dirname $0`
cd ..
export GRPC_HOST=grpc-server:50051
go run cmd/grpc_client/main.go
