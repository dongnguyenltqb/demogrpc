package server

import (
	"demogrpc/entity"
	"fmt"
	"log"
	"net"
	"share-proto/proto-gen/rpc"

	"google.golang.org/grpc"
)

func New(port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// list all rpc service to register
	var (
		user = entity.NewUserGRPCService()
	)

	// create a server instance and register service to server
	grpcServer := grpc.NewServer()
	rpc.RegisterUserServer(grpcServer, user)

	// start the server
	fmt.Printf("gRPC server is running on port %d\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		panic(fmt.Errorf("failed to serve: %s", err))
	}
}
