package client

import (
	"fmt"
	"log"
	"share-proto/proto-gen/rpc"

	"google.golang.org/grpc"
)

type Client struct {
	User rpc.UserClient
}

func New(port int) *Client {
	conn, err := grpc.Dial(fmt.Sprintf(":%d", port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error connect: %s", err)
	}

	return &Client{
		User: rpc.NewUserClient(conn),
	}
}
