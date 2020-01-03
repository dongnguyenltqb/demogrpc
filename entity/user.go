package entity

import (
	"context"
	"demogrpc/rpc"
	"fmt"
)

type user_server struct {
}

func NewUserGRPCServer() *user_server {
	return &user_server{}
}

func (*user_server) UserLogin(ctx context.Context, in *rpc.Credentials) (*rpc.LoginResult, error) {
	fmt.Println(in)
	return &rpc.LoginResult{
		Ok: true,
		Data: &rpc.AccessToken{
			AccessToken: "!23,",
		},
	}, nil
}
func (*user_server) UserRegister(ctx context.Context, in *rpc.Credentials) (*rpc.LoginResult, error) {
	fmt.Println(in)
	return &rpc.LoginResult{
		Ok: true,
		Data: &rpc.AccessToken{
			AccessToken: "regisger.... == >>>>>   !23,",
		},
	}, nil
}
