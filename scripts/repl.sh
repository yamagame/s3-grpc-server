#!/bin/bash
cd `dirname $0`
cd ..
evans --proto ./proto/grpc-server/repository.proto --host localhost --port 50051 repl
# evans --proto ../proto/grpc-server/storage.proto --host localhost --port 50051 repl
