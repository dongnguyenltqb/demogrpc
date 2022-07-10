package client

import (
	"fmt"
	"os"
	"share-proto/proto-gen/rpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Client struct {
	User rpc.UserClient
}

func New(address string) *Client {
	certFile := os.Getenv("CERT_PATH")
	fmt.Println("certFile: ", certFile)
	creds, err := credentials.NewClientTLSFromFile(certFile, os.Getenv("CERT_SERVER_NAME"))
	if err != nil {
		panic(err)
	}
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
	if err != nil {
		panic(err)
	}

	return &Client{
		User: rpc.NewUserClient(conn),
	}
}
