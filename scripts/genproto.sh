#!/bin/bash
cd `dirname $0`
cd ..
protoc --go_out=./ --go_opt=paths=source_relative --go-grpc_out=./ --go-grpc_opt=paths=source_relative proto/grpc-server/*.proto
protoc --plugin=./node_modules/.bin/protoc-gen-ts_proto --ts_proto_out=./grpc-bff proto/grpc-server/*.proto
