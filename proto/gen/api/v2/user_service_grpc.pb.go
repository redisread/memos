// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/v2/user_service.proto

package apiv2

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	UserService_GetUser_FullMethodName               = "/memos.api.v2.UserService/GetUser"
	UserService_UpdateUser_FullMethodName            = "/memos.api.v2.UserService/UpdateUser"
	UserService_ListUserAccessTokens_FullMethodName  = "/memos.api.v2.UserService/ListUserAccessTokens"
	UserService_CreateUserAccessToken_FullMethodName = "/memos.api.v2.UserService/CreateUserAccessToken"
	UserService_DeleteUserAccessToken_FullMethodName = "/memos.api.v2.UserService/DeleteUserAccessToken"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	// ListUserAccessTokens returns a list of access tokens for a user.
	ListUserAccessTokens(ctx context.Context, in *ListUserAccessTokensRequest, opts ...grpc.CallOption) (*ListUserAccessTokensResponse, error)
	// CreateUserAccessToken creates a new access token for a user.
	CreateUserAccessToken(ctx context.Context, in *CreateUserAccessTokenRequest, opts ...grpc.CallOption) (*CreateUserAccessTokenResponse, error)
	// DeleteUserAccessToken deletes an access token for a user.
	DeleteUserAccessToken(ctx context.Context, in *DeleteUserAccessTokenRequest, opts ...grpc.CallOption) (*DeleteUserAccessTokenResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, UserService_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, UserService_UpdateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListUserAccessTokens(ctx context.Context, in *ListUserAccessTokensRequest, opts ...grpc.CallOption) (*ListUserAccessTokensResponse, error) {
	out := new(ListUserAccessTokensResponse)
	err := c.cc.Invoke(ctx, UserService_ListUserAccessTokens_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateUserAccessToken(ctx context.Context, in *CreateUserAccessTokenRequest, opts ...grpc.CallOption) (*CreateUserAccessTokenResponse, error) {
	out := new(CreateUserAccessTokenResponse)
	err := c.cc.Invoke(ctx, UserService_CreateUserAccessToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUserAccessToken(ctx context.Context, in *DeleteUserAccessTokenRequest, opts ...grpc.CallOption) (*DeleteUserAccessTokenResponse, error) {
	out := new(DeleteUserAccessTokenResponse)
	err := c.cc.Invoke(ctx, UserService_DeleteUserAccessToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	// ListUserAccessTokens returns a list of access tokens for a user.
	ListUserAccessTokens(context.Context, *ListUserAccessTokensRequest) (*ListUserAccessTokensResponse, error)
	// CreateUserAccessToken creates a new access token for a user.
	CreateUserAccessToken(context.Context, *CreateUserAccessTokenRequest) (*CreateUserAccessTokenResponse, error)
	// DeleteUserAccessToken deletes an access token for a user.
	DeleteUserAccessToken(context.Context, *DeleteUserAccessTokenRequest) (*DeleteUserAccessTokenResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServiceServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserServiceServer) ListUserAccessTokens(context.Context, *ListUserAccessTokensRequest) (*ListUserAccessTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserAccessTokens not implemented")
}
func (UnimplementedUserServiceServer) CreateUserAccessToken(context.Context, *CreateUserAccessTokenRequest) (*CreateUserAccessTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserAccessToken not implemented")
}
func (UnimplementedUserServiceServer) DeleteUserAccessToken(context.Context, *DeleteUserAccessTokenRequest) (*DeleteUserAccessTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserAccessToken not implemented")
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

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListUserAccessTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserAccessTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListUserAccessTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ListUserAccessTokens_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListUserAccessTokens(ctx, req.(*ListUserAccessTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateUserAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUserAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CreateUserAccessToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUserAccessToken(ctx, req.(*CreateUserAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUserAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUserAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_DeleteUserAccessToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUserAccessToken(ctx, req.(*DeleteUserAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "memos.api.v2.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
		{
			MethodName: "ListUserAccessTokens",
			Handler:    _UserService_ListUserAccessTokens_Handler,
		},
		{
			MethodName: "CreateUserAccessToken",
			Handler:    _UserService_CreateUserAccessToken_Handler,
		},
		{
			MethodName: "DeleteUserAccessToken",
			Handler:    _UserService_DeleteUserAccessToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v2/user_service.proto",
}
