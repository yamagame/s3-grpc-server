#!/bin/bash
cd `dirname $0`
cd ..

export AWS_SDK_LOAD_CONFIG=false
export AWS_PROFILE=default
export AWS_REGION=ap-northeast-1
export AWS_ACCESS_KEY_ID=my-access-key
export AWS_SECRET_ACCESS_KEY=my-secret-key
export S3_BUCKET_NAME=sample-bucket
export S3_ENDPOINT=http://localhost:5000
export COGNITO_ENDPOINT=http://localhost:5000

go run grpc-server/cmd/server/main.go
