#!/usr/bin/env sh

docker run --rm -v $PWD/protobuf:/defs --entrypoint /bin/sh namely/protoc-all -c "rm -rf /defs/demo"
# Demo
docker run --rm -v $PWD/protobuf:/defs namely/protoc-all -f demo.proto -l go -o .