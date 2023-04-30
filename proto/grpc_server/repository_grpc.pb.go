// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.6
// source: proto/grpc-server/repository.proto

package grpc_server

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
	Repository_CreateFileInfo_FullMethodName = "/repository/CreateFileInfo"
	Repository_ReadFileInfo_FullMethodName   = "/repository/ReadFileInfo"
	Repository_UpdateFileInfo_FullMethodName = "/repository/UpdateFileInfo"
	Repository_DeleteFileInfo_FullMethodName = "/repository/DeleteFileInfo"
)

// RepositoryClient is the client API for Repository service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RepositoryClient interface {
	// CreateFileInfo
	CreateFileInfo(ctx context.Context, in *CreateFileInfoRequest, opts ...grpc.CallOption) (*CreateFileInfoResponse, error)
	// ReadFileInfo
	ReadFileInfo(ctx context.Context, in *ReadFileInfoRequest, opts ...grpc.CallOption) (*ReadFileInfoResponse, error)
	// UpdateFileInfo
	UpdateFileInfo(ctx context.Context, in *UpdateFileInfoRequest, opts ...grpc.CallOption) (*UpdateFileInfoResponse, error)
	// DeleteFileInfo
	DeleteFileInfo(ctx context.Context, in *DeleteFileInfoRequest, opts ...grpc.CallOption) (*DeleteFileInfoResponse, error)
}

type repositoryClient struct {
	cc grpc.ClientConnInterface
}

func NewRepositoryClient(cc grpc.ClientConnInterface) RepositoryClient {
	return &repositoryClient{cc}
}

func (c *repositoryClient) CreateFileInfo(ctx context.Context, in *CreateFileInfoRequest, opts ...grpc.CallOption) (*CreateFileInfoResponse, error) {
	out := new(CreateFileInfoResponse)
	err := c.cc.Invoke(ctx, Repository_CreateFileInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryClient) ReadFileInfo(ctx context.Context, in *ReadFileInfoRequest, opts ...grpc.CallOption) (*ReadFileInfoResponse, error) {
	out := new(ReadFileInfoResponse)
	err := c.cc.Invoke(ctx, Repository_ReadFileInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryClient) UpdateFileInfo(ctx context.Context, in *UpdateFileInfoRequest, opts ...grpc.CallOption) (*UpdateFileInfoResponse, error) {
	out := new(UpdateFileInfoResponse)
	err := c.cc.Invoke(ctx, Repository_UpdateFileInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repositoryClient) DeleteFileInfo(ctx context.Context, in *DeleteFileInfoRequest, opts ...grpc.CallOption) (*DeleteFileInfoResponse, error) {
	out := new(DeleteFileInfoResponse)
	err := c.cc.Invoke(ctx, Repository_DeleteFileInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RepositoryServer is the server API for Repository service.
// All implementations must embed UnimplementedRepositoryServer
// for forward compatibility
type RepositoryServer interface {
	// CreateFileInfo
	CreateFileInfo(context.Context, *CreateFileInfoRequest) (*CreateFileInfoResponse, error)
	// ReadFileInfo
	ReadFileInfo(context.Context, *ReadFileInfoRequest) (*ReadFileInfoResponse, error)
	// UpdateFileInfo
	UpdateFileInfo(context.Context, *UpdateFileInfoRequest) (*UpdateFileInfoResponse, error)
	// DeleteFileInfo
	DeleteFileInfo(context.Context, *DeleteFileInfoRequest) (*DeleteFileInfoResponse, error)
	mustEmbedUnimplementedRepositoryServer()
}

// UnimplementedRepositoryServer must be embedded to have forward compatible implementations.
type UnimplementedRepositoryServer struct {
}

func (UnimplementedRepositoryServer) CreateFileInfo(context.Context, *CreateFileInfoRequest) (*CreateFileInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFileInfo not implemented")
}
func (UnimplementedRepositoryServer) ReadFileInfo(context.Context, *ReadFileInfoRequest) (*ReadFileInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadFileInfo not implemented")
}
func (UnimplementedRepositoryServer) UpdateFileInfo(context.Context, *UpdateFileInfoRequest) (*UpdateFileInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFileInfo not implemented")
}
func (UnimplementedRepositoryServer) DeleteFileInfo(context.Context, *DeleteFileInfoRequest) (*DeleteFileInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFileInfo not implemented")
}
func (UnimplementedRepositoryServer) mustEmbedUnimplementedRepositoryServer() {}

// UnsafeRepositoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RepositoryServer will
// result in compilation errors.
type UnsafeRepositoryServer interface {
	mustEmbedUnimplementedRepositoryServer()
}

func RegisterRepositoryServer(s grpc.ServiceRegistrar, srv RepositoryServer) {
	s.RegisterService(&Repository_ServiceDesc, srv)
}

func _Repository_CreateFileInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFileInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServer).CreateFileInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Repository_CreateFileInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServer).CreateFileInfo(ctx, req.(*CreateFileInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repository_ReadFileInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadFileInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServer).ReadFileInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Repository_ReadFileInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServer).ReadFileInfo(ctx, req.(*ReadFileInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repository_UpdateFileInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFileInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServer).UpdateFileInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Repository_UpdateFileInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServer).UpdateFileInfo(ctx, req.(*UpdateFileInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repository_DeleteFileInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFileInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryServer).DeleteFileInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Repository_DeleteFileInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryServer).DeleteFileInfo(ctx, req.(*DeleteFileInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Repository_ServiceDesc is the grpc.ServiceDesc for Repository service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Repository_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "repository",
	HandlerType: (*RepositoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFileInfo",
			Handler:    _Repository_CreateFileInfo_Handler,
		},
		{
			MethodName: "ReadFileInfo",
			Handler:    _Repository_ReadFileInfo_Handler,
		},
		{
			MethodName: "UpdateFileInfo",
			Handler:    _Repository_UpdateFileInfo_Handler,
		},
		{
			MethodName: "DeleteFileInfo",
			Handler:    _Repository_DeleteFileInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/grpc-server/repository.proto",
}
