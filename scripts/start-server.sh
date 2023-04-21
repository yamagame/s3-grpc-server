#!/bin/bash
cd `dirname $0`
cd ..
go run grpc-server/cmd/server/main.go
