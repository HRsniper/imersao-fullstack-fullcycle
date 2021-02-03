#!/bin/sh

PROTOC=path/to/protoc

rm -rf application/grpc/pb
mkdir -p application/grpc/pb

$PROTOC \
  --go_out=application/grpc/pb \
  --go_opt=paths=source_relative \
  --go_grpc_out=application/grpc/pb \
  --go_grpc_opt=paths=source_relative \
  --proto_path=application/grpc/protofiles application/grpc/protofiles/*.proto
