package entity

import (
	"context"
	"fmt"
	"share-proto/proto-gen/message"
)

type user_service struct {
}

func NewUserGRPCService() *user_service {
	return &user_service{}
}

func (*user_service) UserLogin(ctx context.Context, in *message.Credentials) (*message.LoginResult, error) {
	fmt.Println(in)
	return &message.LoginResult{
		Ok: true,
		Data: &message.AccessToken{
			AccessToken: "Login sucessfully..... here is token xxxxxx",
		},
	}, nil
}
func (*user_service) UserRegister(ctx context.Context, in *message.FormRegister) (*message.RegisterResult, error) {
	fmt.Println(in)
	return &message.RegisterResult{
		Ok: true,
		Data: &message.AccessToken{
			AccessToken: "Register sucessfully..... here is token xxxxxx",
		},
	}, nil

}

func (*user_service) GetUserDataByUserName(ctx context.Context, in *message.UserName) (*message.GetUserDataByUserNameResult, error) {
	return &message.GetUserDataByUserNameResult{
		Ok:    true,
		Error: "",
	}, nil
}

func (*user_service) UpdateUserDataByUserName(ctx context.Context, in *message.UpdateUserDataByUserName) (*message.UpdateUserDataByUserNameResult, error) {
	return &message.UpdateUserDataByUserNameResult{
		Ok:    true,
		Error: "",
	}, nil
}
