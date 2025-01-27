// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: proto/user/user.proto

package user

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UserService_GetUserLastLogin_FullMethodName = "/user.UserService/GetUserLastLogin"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUserLastLogin(ctx context.Context, in *GetUserLastLoginRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetUserLastLoginResponse], error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUserLastLogin(ctx context.Context, in *GetUserLastLoginRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetUserLastLoginResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &UserService_ServiceDesc.Streams[0], UserService_GetUserLastLogin_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetUserLastLoginRequest, GetUserLastLoginResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type UserService_GetUserLastLoginClient = grpc.ServerStreamingClient[GetUserLastLoginResponse]

// UserServiceServer is the server API for UserService service.
// All implementations should embed UnimplementedUserServiceServer
// for forward compatibility.
type UserServiceServer interface {
	GetUserLastLogin(*GetUserLastLoginRequest, grpc.ServerStreamingServer[GetUserLastLoginResponse]) error
}

// UnimplementedUserServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServiceServer struct{}

func (UnimplementedUserServiceServer) GetUserLastLogin(*GetUserLastLoginRequest, grpc.ServerStreamingServer[GetUserLastLoginResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetUserLastLogin not implemented")
}
func (UnimplementedUserServiceServer) testEmbeddedByValue() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetUserLastLogin_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetUserLastLoginRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserServiceServer).GetUserLastLogin(m, &grpc.GenericServerStream[GetUserLastLoginRequest, GetUserLastLoginResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type UserService_GetUserLastLoginServer = grpc.ServerStreamingServer[GetUserLastLoginResponse]

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUserLastLogin",
			Handler:       _UserService_GetUserLastLogin_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/user/user.proto",
}
