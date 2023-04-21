// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.6
// source: proto/grpc-server/aws.proto

package aws

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
	Aws_CreateBucket_FullMethodName = "/aws/CreateBucket"
	Aws_ListBuckets_FullMethodName  = "/aws/ListBuckets"
	Aws_PutObject_FullMethodName    = "/aws/PutObject"
	Aws_GetObject_FullMethodName    = "/aws/GetObject"
	Aws_DeleteObject_FullMethodName = "/aws/DeleteObject"
)

// AwsClient is the client API for Aws service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AwsClient interface {
	// CreateBucket
	CreateBucket(ctx context.Context, in *CreateBucketRequest, opts ...grpc.CallOption) (*CreateBucketResponse, error)
	// ListBuckets
	ListBuckets(ctx context.Context, in *ListBucketsRequest, opts ...grpc.CallOption) (*ListBucketsResponse, error)
	// PutObject
	PutObject(ctx context.Context, in *PutObjectRequest, opts ...grpc.CallOption) (*PutObjectResponse, error)
	// GetObject
	GetObject(ctx context.Context, in *GetObjectRequest, opts ...grpc.CallOption) (*GetObjectResponse, error)
	// DeleteObject
	DeleteObject(ctx context.Context, in *DeleteObjectRequest, opts ...grpc.CallOption) (*DeleteObjectResponse, error)
}

type awsClient struct {
	cc grpc.ClientConnInterface
}

func NewAwsClient(cc grpc.ClientConnInterface) AwsClient {
	return &awsClient{cc}
}

func (c *awsClient) CreateBucket(ctx context.Context, in *CreateBucketRequest, opts ...grpc.CallOption) (*CreateBucketResponse, error) {
	out := new(CreateBucketResponse)
	err := c.cc.Invoke(ctx, Aws_CreateBucket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *awsClient) ListBuckets(ctx context.Context, in *ListBucketsRequest, opts ...grpc.CallOption) (*ListBucketsResponse, error) {
	out := new(ListBucketsResponse)
	err := c.cc.Invoke(ctx, Aws_ListBuckets_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *awsClient) PutObject(ctx context.Context, in *PutObjectRequest, opts ...grpc.CallOption) (*PutObjectResponse, error) {
	out := new(PutObjectResponse)
	err := c.cc.Invoke(ctx, Aws_PutObject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *awsClient) GetObject(ctx context.Context, in *GetObjectRequest, opts ...grpc.CallOption) (*GetObjectResponse, error) {
	out := new(GetObjectResponse)
	err := c.cc.Invoke(ctx, Aws_GetObject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *awsClient) DeleteObject(ctx context.Context, in *DeleteObjectRequest, opts ...grpc.CallOption) (*DeleteObjectResponse, error) {
	out := new(DeleteObjectResponse)
	err := c.cc.Invoke(ctx, Aws_DeleteObject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AwsServer is the server API for Aws service.
// All implementations must embed UnimplementedAwsServer
// for forward compatibility
type AwsServer interface {
	// CreateBucket
	CreateBucket(context.Context, *CreateBucketRequest) (*CreateBucketResponse, error)
	// ListBuckets
	ListBuckets(context.Context, *ListBucketsRequest) (*ListBucketsResponse, error)
	// PutObject
	PutObject(context.Context, *PutObjectRequest) (*PutObjectResponse, error)
	// GetObject
	GetObject(context.Context, *GetObjectRequest) (*GetObjectResponse, error)
	// DeleteObject
	DeleteObject(context.Context, *DeleteObjectRequest) (*DeleteObjectResponse, error)
	mustEmbedUnimplementedAwsServer()
}

// UnimplementedAwsServer must be embedded to have forward compatible implementations.
type UnimplementedAwsServer struct {
}

func (UnimplementedAwsServer) CreateBucket(context.Context, *CreateBucketRequest) (*CreateBucketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBucket not implemented")
}
func (UnimplementedAwsServer) ListBuckets(context.Context, *ListBucketsRequest) (*ListBucketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBuckets not implemented")
}
func (UnimplementedAwsServer) PutObject(context.Context, *PutObjectRequest) (*PutObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutObject not implemented")
}
func (UnimplementedAwsServer) GetObject(context.Context, *GetObjectRequest) (*GetObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObject not implemented")
}
func (UnimplementedAwsServer) DeleteObject(context.Context, *DeleteObjectRequest) (*DeleteObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteObject not implemented")
}
func (UnimplementedAwsServer) mustEmbedUnimplementedAwsServer() {}

// UnsafeAwsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AwsServer will
// result in compilation errors.
type UnsafeAwsServer interface {
	mustEmbedUnimplementedAwsServer()
}

func RegisterAwsServer(s grpc.ServiceRegistrar, srv AwsServer) {
	s.RegisterService(&Aws_ServiceDesc, srv)
}

func _Aws_CreateBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AwsServer).CreateBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Aws_CreateBucket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AwsServer).CreateBucket(ctx, req.(*CreateBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Aws_ListBuckets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBucketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AwsServer).ListBuckets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Aws_ListBuckets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AwsServer).ListBuckets(ctx, req.(*ListBucketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Aws_PutObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AwsServer).PutObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Aws_PutObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AwsServer).PutObject(ctx, req.(*PutObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Aws_GetObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AwsServer).GetObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Aws_GetObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AwsServer).GetObject(ctx, req.(*GetObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Aws_DeleteObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AwsServer).DeleteObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Aws_DeleteObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AwsServer).DeleteObject(ctx, req.(*DeleteObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Aws_ServiceDesc is the grpc.ServiceDesc for Aws service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Aws_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aws",
	HandlerType: (*AwsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBucket",
			Handler:    _Aws_CreateBucket_Handler,
		},
		{
			MethodName: "ListBuckets",
			Handler:    _Aws_ListBuckets_Handler,
		},
		{
			MethodName: "PutObject",
			Handler:    _Aws_PutObject_Handler,
		},
		{
			MethodName: "GetObject",
			Handler:    _Aws_GetObject_Handler,
		},
		{
			MethodName: "DeleteObject",
			Handler:    _Aws_DeleteObject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/grpc-server/aws.proto",
}
