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

func New(port int) *Client {
	certFile := os.Getenv("CERT_PATH")
	fmt.Println("certFile: ", certFile)
	creds, err := credentials.NewClientTLSFromFile(certFile, "dongnguyen.dev")
	if err != nil {
		panic(err)
	}
	fmt.Println("connecting to :", port)
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", port), grpc.WithTransportCredentials(creds))
	if err != nil {
		panic(err)
	}

	return &Client{
		User: rpc.NewUserClient(conn),
	}
}
