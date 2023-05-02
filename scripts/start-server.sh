#!/bin/bash
cd `dirname $0`
cd ..
go run grpc_server/cmd/main.go $1
