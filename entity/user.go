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

func (*UserService) Login(ctx context.Context, input *message.LoginRequest) (*message.LoginResponse, error) {
	fmt.Printf("%+v", input)
	return &message.LoginResponse{
		Success: true,
		Data: &message.AccessToken{
			AccessToken:  "access",
			RefreshToken: "refresh",
		},
	}, nil
}

func (*UserService) GetUserById(ctx context.Context, input *message.GetUserByIdRequest) (*message.GetUserByIdResponse, error) {
	fmt.Printf("get user entity for %d\n", input.Id)
	return &message.GetUserByIdResponse{
		Success: true,
		Data: &message.User{
			Age:   100,
			Email: "example@emai.com",
		},
	}, nil
}
