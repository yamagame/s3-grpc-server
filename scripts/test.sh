#!/bin/bash
cd `dirname $0`

function callEvans() {
  evans --proto ../proto/grpc-server/storage.proto --host localhost --port 50051  cli call storage.$1
}

function runTest() {
  echo CreateBucket
  echo '{}' | callEvans CreateBucket | jq
  echo ListBuckets
  echo '{}' | callEvans ListBuckets | jq
  echo PutObject
  echo '{ "key": "test/a", "content": "hello test-a" }' | callEvans PutObject | jq
  echo '{ "key": "test/b", "content": "hello test-b" }' | callEvans PutObject | jq
  echo GetObject
  echo '{ "key": "test/a" }' | callEvans GetObject | jq
  echo '{ "key": "test/b" }' | callEvans GetObject | jq
  echo ListObjects
  echo '{ "prefix": "test" }' | callEvans ListObjects | jq
  echo DeleteObject
  echo '{ "key": "test/a" }' | callEvans DeleteObject | jq
  echo '{ "key": "test/b" }' | callEvans DeleteObject | jq
  echo ListObjects
  echo '{ "prefix": "test" }' | callEvans ListObjects | jq
}

runTest > ../work/test.log
diff ../work/test.log ./test/test.log

cmdstatus=$?
if [ $cmdstatus -ne 0 ]; then
  echo NG
  exit $cmdstatus
fi

echo OK
