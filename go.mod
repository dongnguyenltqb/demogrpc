module demogrpc

go 1.13

replace share-proto => ../share-proto

require (
	github.com/gin-gonic/gin v1.5.0
	github.com/golang/protobuf v1.3.2
	google.golang.org/grpc v1.26.0
	share-proto v0.0.0-00010101000000-000000000000
)
