 
 # gen protobuf

 ### every message have package name  = message and option go_package = demogrpc/proto-gen/message
protoc --proto_path=protobuf  --go_out=plugins=grpc:.. protobuf/message/*.proto

 ### every service have package name  = rpc and option go_package = demogrpc/proto-gen/rpc
protoc --proto_path=protobuf  --go_out=plugins=grpc:.. protobuf/rpc/*.proto 

