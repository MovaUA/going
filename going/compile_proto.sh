#!/bin/sh

set -ex

DIR=$(dirname $0)

echo $DIR

protoc \
  --proto_path="${DIR}" \
  --go_out="${DIR}" \
  --go_opt=paths=source_relative \
  --go-grpc_out="${DIR}" \
  --go-grpc_opt=paths=source_relative \
  "${DIR}/going.proto"
