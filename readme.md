- Generate grpc's - protoc --go_out=gen --go-grpc_out=gen --grpc-gateway_out=gen --grpc-gateway_opt generate_unbound_methods=true --openapiv2_out ./docs proto/*.proto