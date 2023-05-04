#!/bin/bash
cd `dirname $0`
cd ..
evans -r --host localhost --port 50051 repl
