package server

import (
	"context"
	"demogrpc/entity"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type AuthContextKey string

func Authenticate(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	path := info.FullMethod
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		err = fmt.Errorf("REQ: %s : could not receive metadata", path)
		return
	}
	fmt.Printf("REQ: %s \nMetadata: %+v\n", path, md)
	tokens := md.Get("Authenticate")
	if len(tokens) == 1 {
		authCtx := context.WithValue(ctx, AuthContextKey("claims"), entity.Claims{
			UserId: 1000,
		})
		return handler(authCtx, req)
	}
	return handler(ctx, req)
}
