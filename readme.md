 
 # gen protobuf
 protoc --proto_path=protobuf  --go_out=plugins=grpc:rpc protobuf/user.rpc.proto 