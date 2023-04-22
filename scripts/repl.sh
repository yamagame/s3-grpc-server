#!/bin/bash
cd `dirname $0`
evans --proto ../proto/grpc-server/storage.proto --host localhost --port 50051 repl
