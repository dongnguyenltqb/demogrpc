package rpc

import (
	"context"
	"fmt"
)

func (*Server) UserLogin(ctx context.Context, in *Credentials) (*LoginResult, error) {
	fmt.Println(in)
	return &LoginResult{
		Ok: true,
		Data: &AccessToken{
			AccessToken: "!23,",
		},
	}, nil
}
