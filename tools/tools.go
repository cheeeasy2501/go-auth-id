package tools

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
    _ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
    _ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
    _ "google.golang.org/protobuf/cmd/protoc-gen-go"
)


// protoc -I . --grpc-gateway_out ./gen/go  --grpc-gateway_opt logtostderr=true  --grpc-gateway_opt paths=source_relative  --grpc-gateway_opt generate_unbound_methods=true your/service/v1/your_service.proto

// protoc --go_out=. --go-grpc_out=. --grpc-gateway_out=. --grpc-gateway_opt generate_unbound_methods=true --openapiv2_out . internal/transport/grpc/v1/authorization.proto
