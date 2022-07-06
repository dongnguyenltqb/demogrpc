package client

import (
	"context"
	"fmt"
	"share-proto/proto-gen/message"
	"testing"
)

func TestLogin(t *testing.T) {
	client := New(7777)
	response, err := client.User.Login(context.Background(), &message.LoginRequest{
		Username: "dong",
		Password: "nguyen huu",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v", response)
}