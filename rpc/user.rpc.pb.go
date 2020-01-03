// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.rpc.proto

package rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("user.rpc.proto", fileDescriptor_edfd371b7eecadb7) }

var fileDescriptor_edfd371b7eecadb7 = []byte{
	// 109 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x2a, 0x48, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x2a, 0x48, 0x96,
	0x12, 0x02, 0x0b, 0xe6, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x42, 0x24, 0x8c, 0x2c, 0xb9, 0x58,
	0x42, 0x8b, 0x53, 0x8b, 0x84, 0x0c, 0xb9, 0x38, 0x41, 0xb4, 0x4f, 0x7e, 0x7a, 0x66, 0x9e, 0x90,
	0x00, 0x58, 0xa7, 0x73, 0x51, 0x6a, 0x4a, 0x6a, 0x5e, 0x49, 0x66, 0x62, 0x4e, 0xb1, 0x14, 0x44,
	0x04, 0x2c, 0x1b, 0x94, 0x5a, 0x5c, 0x9a, 0x53, 0xa2, 0xc4, 0x90, 0xc4, 0x06, 0x36, 0xc1, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0x3a, 0xda, 0x1c, 0x66, 0x6c, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	UserLogin(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*LoginResult, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) UserLogin(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*LoginResult, error) {
	out := new(LoginResult)
	err := c.cc.Invoke(ctx, "/rpc.User/UserLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	UserLogin(context.Context, *Credentials) (*LoginResult, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) UserLogin(ctx context.Context, req *Credentials) (*LoginResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Credentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.User/UserLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserLogin(ctx, req.(*Credentials))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserLogin",
			Handler:    _User_UserLogin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.rpc.proto",
}