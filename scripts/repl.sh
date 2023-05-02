#!/bin/bash
cd `dirname $0`
cd ..
evans --proto ./proto/grpc_server/repository.proto --host localhost --port 50051 repl
# evans --proto ../proto/grpc_server/storage.proto --host localhost --port 50051 repl
