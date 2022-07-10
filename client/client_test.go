package client

import (
	"context"
	"fmt"
	"os"
	"share-proto/proto-gen/message"
	"testing"
)

func TestLogin(t *testing.T) {
	client := New(os.Getenv("SERVER_ADDR"))
	response, err := client.User.Login(context.Background(), &message.LoginRequest{
		Username: "dong",
		Password: "nguyen huu",
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v", response)
}
