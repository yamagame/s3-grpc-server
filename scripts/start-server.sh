#!/bin/bash
cd `dirname $0`
cd ..
go run cmd/grpc_server/main.go $1
