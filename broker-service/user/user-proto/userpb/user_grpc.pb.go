// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: user.proto

package userpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateUser(ctx context.Context, in *UserPayload, opts ...grpc.CallOption) (*User, error)
	GetMany(ctx context.Context, in *Limit, opts ...grpc.CallOption) (*Users, error)
	GetById(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*User, error)
	DeleteById(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateById(ctx context.Context, in *UserPayload, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *UserPayload, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetMany(ctx context.Context, in *Limit, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/user.UserService/GetMany", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetById(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user.UserService/GetById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteById(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.UserService/DeleteById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateById(ctx context.Context, in *UserPayload, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	CreateUser(context.Context, *UserPayload) (*User, error)
	GetMany(context.Context, *Limit) (*Users, error)
	GetById(context.Context, *UserId) (*User, error)
	DeleteById(context.Context, *UserId) (*emptypb.Empty, error)
	UpdateById(context.Context, *UserPayload) (*emptypb.Empty, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) CreateUser(context.Context, *UserPayload) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) GetMany(context.Context, *Limit) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMany not implemented")
}
func (UnimplementedUserServiceServer) GetById(context.Context, *UserId) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedUserServiceServer) DeleteById(context.Context, *UserId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteById not implemented")
}
func (UnimplementedUserServiceServer) UpdateById(context.Context, *UserPayload) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateById not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*UserPayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Limit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetMany",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetMany(ctx, req.(*Limit))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetById(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/DeleteById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteById(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateById(ctx, req.(*UserPayload))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "GetMany",
			Handler:    _UserService_GetMany_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _UserService_GetById_Handler,
		},
		{
			MethodName: "DeleteById",
			Handler:    _UserService_DeleteById_Handler,
		},
		{
			MethodName: "UpdateById",
			Handler:    _UserService_UpdateById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}