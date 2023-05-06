#!/bin/bash
cd `dirname $0`
cd ..
# go run cmd/server/main.go $1
go run cmd/fxserver/main.go $1
