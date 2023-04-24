#!/bin/bash
cd `dirname $0`
cd ..

export BUCKET_NAME=sample-bucket

export AWS_SDK_LOAD_CONFIG=false
export AWS_PROFILE=default
export AWS_REGION=ap-northeast-1
export AWS_ACCESS_KEY_ID=my-access-key
export AWS_SECRET_ACCESS_KEY=my-secret-key
export S3_ENDPOINT=http://localhost:5000

export GCS_PROJECT_ID=
export STORAGE_EMULATOR_HOST=localhost:4443

export SFTP_HOST=localhost
export SFTP_PORT=2222
export SFTP_USERNAME=username
export SFTP_PASSWORD=password
export SFTP_SHAREDIR=share

go run grpc-server/cmd/server/main.go $1
