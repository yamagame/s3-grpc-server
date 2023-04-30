#!/bin/bash
cd `dirname $0`
cd ..
go run grpc_server/cmd/server/main.go $1
