#!/bin/bash
cd `dirname $0`
cd ..
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpc-server/*.proto
