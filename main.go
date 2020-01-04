package main

import (
	"demogrpc/entity"
	server "demogrpc/server"
	"fmt"
	"log"
	"net"
	"share-proto/proto-gen/rpc"

	"google.golang.org/grpc"
)

func runGRPCServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
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
	fmt.Println("gRPC server is running....")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func main() {
	go server.Start()
	runGRPCServer()
}
