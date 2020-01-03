package entity

import (
	"context"
	"demogrpc/rpc"
	"fmt"
)

type user_service struct {
}

func NewUserGRPCService() *user_service {
	return &user_service{}
}

func (*user_service) UserLogin(ctx context.Context, in *rpc.Credentials) (*rpc.LoginResult, error) {
	fmt.Println(in)
	return &rpc.LoginResult{
		Ok: true,
		Data: &rpc.AccessToken{
			AccessToken: "!23,",
		},
	}, nil
}
func (*user_service) UserRegister(ctx context.Context, in *rpc.FormRegister) (*rpc.RegisterResult, error) {
	fmt.Println(in)
	return &rpc.RegisterResult{
		Ok: true,
		Data: &rpc.AccessToken{
			AccessToken: "regisger.... == >>>>>   !23@@,",
		},
	}, nil
}
