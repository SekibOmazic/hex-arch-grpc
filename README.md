# Hexagonal architecture with Go and gRPC

Found [here](https://t.co/QaN1cAzDmu?amp=1)


## gRPC

To generate Go files from *.proto files run following:

```
protoc --go_out=internal/adapters/framework/left/grpc --proto_path=internal/adapters/framework/left/grpc/proto internal/adapters/framework/left/grpc/proto/number_msg.proto
```

```
protoc --go-grpc_out=require_unimplemented_servers=false:internal/adapters/framework/left/grpc --proto_path=internal/adapters/framework/left/grpc/proto internal/adapters/framework/left/grpc/proto/arithmetic_svc.proto
```