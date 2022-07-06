package entity

import (
	"context"
	"fmt"
	"share-proto/proto-gen/message"
)

type UserService struct {
}

func NewUserGRPCService() *UserService {
	return &UserService{}
}

func (*UserService) Login(ctx context.Context, in *message.LoginRequest) (*message.LoginResponse, error) {
	fmt.Println(in)
	return &message.LoginResponse{
		Success: true,
		Data: &message.AccessToken{
			AccessToken:  "access",
			RefreshToken: "refresh",
		},
	}, nil
}
