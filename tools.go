//go:build tools

package tools

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
