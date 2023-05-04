#!/bin/bash
cd `dirname $0`
cd ..

export MYSQL_HOST=localhost:3306
export MYSQL_ROOT_PASSWORD=pass
export MYSQL_DATABASE=mysqldb

go run cmd/droptable/main.go $1
