#!/bin/bash
cd `dirname $0`
cd ..
export GRPC_HOST=localhost:50051

export MYSQL_HOST=localhost:3306
export MYSQL_ROOT_PASSWORD=pass
export MYSQL_DATABASE=mysqldb

go run grpc_client/cmd/main.go
