package server

import (
	"demogrpc/entity"
	"fmt"
	"log"
	"net"
	"path"
	"runtime"
	"share-proto/proto-gen/rpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("unable to get the current filename")
	}
	certFile := path.Join(filename, "../../cert/server.crt")
	keyFile := path.Join(filename, "../../cert/server.key")
	fmt.Println("cert file path")
	fmt.Println(certFile)
	fmt.Println(keyFile)
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	rpc.RegisterUserServer(grpcServer, user)

	// start the server
	fmt.Printf("gRPC server is running on port %d\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		panic(fmt.Errorf("failed to serve: %s", err))
	}
}
