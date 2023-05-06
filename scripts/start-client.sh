#!/bin/bash
cd `dirname $0`
cd ..
export GRPC_HOST=grpc-server:50051

# go run cmd/client/main.go
go run cmd/fxclient/main.go
