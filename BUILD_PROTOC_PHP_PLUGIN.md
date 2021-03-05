```sh
     brew install autoconf automake libtool shtool
     brew install protoc bazel
     git clone --recursive --depth 1 --shallow-submodules -b v1.35.0 https://github.com/grpc/grpc
     cd grpc
     bazel build :all
     bazel build src/compiler:grpc_php_plugin
     cd bazel-bin/src/compiler
     protoc --proto_path=. --php_out=. --grpc_out=. --plugin=protoc-gen-grpc=grpc/bazel-bin/src/compiler/grpc_php_plugin SayHello.proto
```
