#!/bin/bash
# 生成 Go 和 JavaScript 语言的 Protobuf 代码。

protoc --go_out=gen --go-grpc_out=gen  proto/*.proto  # 生成 代码

