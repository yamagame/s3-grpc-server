#!/bin/bash
cd `dirname $0`
cd ..
go run grpc-client/cmd/main.go
