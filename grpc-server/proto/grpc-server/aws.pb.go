// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.9
// source: proto/grpc-server/aws.proto

package aws

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Result int32

const (
	Result_UNDEFINED Result = 0
	Result_OK        Result = 1
	Result_ERR       Result = 2
)

// Enum value maps for Result.
var (
	Result_name = map[int32]string{
		0: "UNDEFINED",
		1: "OK",
		2: "ERR",
	}
	Result_value = map[string]int32{
		"UNDEFINED": 0,
		"OK":        1,
		"ERR":       2,
	}
)

func (x Result) Enum() *Result {
	p := new(Result)
	*p = x
	return p
}

func (x Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Result) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_grpc_server_aws_proto_enumTypes[0].Descriptor()
}

func (Result) Type() protoreflect.EnumType {
	return &file_proto_grpc_server_aws_proto_enumTypes[0]
}

func (x Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Result.Descriptor instead.
func (Result) EnumDescriptor() ([]byte, []int) {
	return file_proto_grpc_server_aws_proto_rawDescGZIP(), []int{0}
}

type CreateBucketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateBucketRequest) Reset() {
	*x = CreateBucketRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_aws_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBucketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBucketRequest) ProtoMessage() {}

func (x *CreateBucketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_aws_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBucketRequest.ProtoReflect.Descriptor instead.
func (*CreateBucketRequest) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_aws_proto_rawDescGZIP(), []int{0}
}

type CreateBucketResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result Result `protobuf:"varint,1,opt,name=result,proto3,enum=Result" json:"result,omitempty"`
}

func (x *CreateBucketResponse) Reset() {
	*x = CreateBucketResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_aws_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBucketResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBucketResponse) ProtoMessage() {}

func (x *CreateBucketResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_aws_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBucketResponse.ProtoReflect.Descriptor instead.
func (*CreateBucketResponse) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_aws_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBucketResponse) GetResult() Result {
	if x != nil {
		return x.Result
	}
	return Result_UNDEFINED
}

type ListBucketsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListBucketsRequest) Reset() {
	*x = ListBucketsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_aws_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBucketsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBucketsRequest) ProtoMessage() {}

func (x *ListBucketsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_aws_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBucketsRequest.ProtoReflect.Descriptor instead.
func (*ListBucketsRequest) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_aws_proto_rawDescGZIP(), []int{2}
}

type ListBucketsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result  Result                       `protobuf:"varint,1,opt,name=result,proto3,enum=Result" json:"result,omitempty"`
	Buckets []*ListBucketsResponseBucket `protobuf:"bytes,2,rep,name=buckets,proto3" json:"buckets,omitempty"`
}

func (x *ListBucketsResponse) Reset() {
	*x = ListBucketsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_aws_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBucketsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBucketsResponse) ProtoMessage() {}

func (x *ListBucketsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_aws_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBucketsResponse.ProtoReflect.Descriptor instead.
func (*ListBucketsResponse) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_aws_proto_rawDescGZIP(), []int{3}
}

func (x *ListBucketsResponse) GetResult() Result {
	if x != nil {
		return x.Result
	}
	return Result_UNDEFINED
}

func (x *ListBucketsResponse) GetBuckets() []*ListBucketsResponseBucket {
	if x != nil {
		return x.Buckets
	}
	return nil
}

type PutObjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key     string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *PutObjectRequest) Reset() {
	*x = PutObjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_aws_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutObjectRequest) ProtoMessage() {}

func (x *PutObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_aws_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutObjectRequest.ProtoReflect.Descriptor instead.
func (*PutObjectRequest) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_aws_proto_rawDescGZIP(), []int{4}
}

func (x *PutObjectRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PutObjectRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type PutObjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result Result `protobuf:"varint,1,opt,name=result,proto3,enum=Result" json:"result,omitempty"`
	Key    string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *PutObjectResponse) Reset() {
	*x = PutObjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_aws_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutObjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutObjectResponse) ProtoMessage() {}

func (x *PutObjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_aws_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutObjectResponse.ProtoReflect.Descriptor instead.
func (*PutObjectResponse) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_aws_proto_rawDescGZIP(), []int{5}
}

func (x *PutObjectResponse) GetResult() Result {
	if x != nil {
		return x.Result
	}
	return Result_UNDEFINED
}

func (x *PutObjectResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetObjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetObjectRequest) Reset() {
	*x = GetObjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_aws_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetObjectRequest) ProtoMessage() {}

func (x *GetObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_aws_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetObjectRequest.ProtoReflect.Descriptor instead.
func (*GetObjectRequest) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_aws_proto_rawDescGZIP(), []int{6}
}

func (x *GetObjectRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetObjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result  Result `protobuf:"varint,1,opt,name=result,proto3,enum=Result" json:"result,omitempty"`
	Key     string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *GetObjectResponse) Reset() {
	*x = GetObjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_aws_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetObjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetObjectResponse) ProtoMessage() {}

func (x *GetObjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_aws_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetObjectResponse.ProtoReflect.Descriptor instead.
func (*GetObjectResponse) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_aws_proto_rawDescGZIP(), []int{7}
}

func (x *GetObjectResponse) GetResult() Result {
	if x != nil {
		return x.Result
	}
	return Result_UNDEFINED
}

func (x *GetObjectResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *GetObjectResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type DeleteObjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DeleteObjectRequest) Reset() {
	*x = DeleteObjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_aws_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteObjectRequest) ProtoMessage() {}

func (x *DeleteObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_aws_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteObjectRequest.ProtoReflect.Descriptor instead.
func (*DeleteObjectRequest) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_aws_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteObjectRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type DeleteObjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result Result `protobuf:"varint,1,opt,name=result,proto3,enum=Result" json:"result,omitempty"`
	Key    string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DeleteObjectResponse) Reset() {
	*x = DeleteObjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_aws_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteObjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteObjectResponse) ProtoMessage() {}

func (x *DeleteObjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_aws_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteObjectResponse.ProtoReflect.Descriptor instead.
func (*DeleteObjectResponse) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_aws_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteObjectResponse) GetResult() Result {
	if x != nil {
		return x.Result
	}
	return Result_UNDEFINED
}

func (x *DeleteObjectResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type ListBucketsResponseBucket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ListBucketsResponseBucket) Reset() {
	*x = ListBucketsResponseBucket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_aws_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBucketsResponseBucket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBucketsResponseBucket) ProtoMessage() {}

func (x *ListBucketsResponseBucket) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_aws_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBucketsResponseBucket.ProtoReflect.Descriptor instead.
func (*ListBucketsResponseBucket) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_aws_proto_rawDescGZIP(), []int{3, 0}
}

func (x *ListBucketsResponseBucket) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_proto_grpc_server_aws_proto protoreflect.FileDescriptor

var file_proto_grpc_server_aws_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x61, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x15, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x37, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x14, 0x0a,
	0x12, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x8b, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x35, 0x0a, 0x07,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x07, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x73, 0x1a, 0x1c, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x3e, 0x0a, 0x10, 0x50, 0x75, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x22, 0x46, 0x0a, 0x11, 0x50, 0x75, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x24, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22,
	0x60, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x22, 0x27, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x49, 0x0a, 0x14, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x07, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x2a, 0x28, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x06,
	0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x45, 0x52, 0x52, 0x10, 0x02, 0x32,
	0xab, 0x02, 0x0a, 0x03, 0x61, 0x77, 0x73, 0x12, 0x3d, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x13, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x34, 0x0a, 0x09, 0x50, 0x75, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x11, 0x2e, 0x50, 0x75, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x50, 0x75, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d,
	0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x14,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x11, 0x5a,
	0x0f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x77, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_grpc_server_aws_proto_rawDescOnce sync.Once
	file_proto_grpc_server_aws_proto_rawDescData = file_proto_grpc_server_aws_proto_rawDesc
)

func file_proto_grpc_server_aws_proto_rawDescGZIP() []byte {
	file_proto_grpc_server_aws_proto_rawDescOnce.Do(func() {
		file_proto_grpc_server_aws_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_grpc_server_aws_proto_rawDescData)
	})
	return file_proto_grpc_server_aws_proto_rawDescData
}

var file_proto_grpc_server_aws_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_grpc_server_aws_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_grpc_server_aws_proto_goTypes = []interface{}{
	(Result)(0),                       // 0: result
	(*CreateBucketRequest)(nil),       // 1: CreateBucketRequest
	(*CreateBucketResponse)(nil),      // 2: CreateBucketResponse
	(*ListBucketsRequest)(nil),        // 3: ListBucketsRequest
	(*ListBucketsResponse)(nil),       // 4: ListBucketsResponse
	(*PutObjectRequest)(nil),          // 5: PutObjectRequest
	(*PutObjectResponse)(nil),         // 6: PutObjectResponse
	(*GetObjectRequest)(nil),          // 7: GetObjectRequest
	(*GetObjectResponse)(nil),         // 8: GetObjectResponse
	(*DeleteObjectRequest)(nil),       // 9: DeleteObjectRequest
	(*DeleteObjectResponse)(nil),      // 10: DeleteObjectResponse
	(*ListBucketsResponseBucket)(nil), // 11: ListBucketsResponse.bucket
}
var file_proto_grpc_server_aws_proto_depIdxs = []int32{
	0,  // 0: CreateBucketResponse.result:type_name -> result
	0,  // 1: ListBucketsResponse.result:type_name -> result
	11, // 2: ListBucketsResponse.buckets:type_name -> ListBucketsResponse.bucket
	0,  // 3: PutObjectResponse.result:type_name -> result
	0,  // 4: GetObjectResponse.result:type_name -> result
	0,  // 5: DeleteObjectResponse.result:type_name -> result
	1,  // 6: aws.CreateBucket:input_type -> CreateBucketRequest
	3,  // 7: aws.ListBuckets:input_type -> ListBucketsRequest
	5,  // 8: aws.PutObject:input_type -> PutObjectRequest
	7,  // 9: aws.GetObject:input_type -> GetObjectRequest
	9,  // 10: aws.DeleteObject:input_type -> DeleteObjectRequest
	2,  // 11: aws.CreateBucket:output_type -> CreateBucketResponse
	4,  // 12: aws.ListBuckets:output_type -> ListBucketsResponse
	6,  // 13: aws.PutObject:output_type -> PutObjectResponse
	8,  // 14: aws.GetObject:output_type -> GetObjectResponse
	10, // 15: aws.DeleteObject:output_type -> DeleteObjectResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_grpc_server_aws_proto_init() }
func file_proto_grpc_server_aws_proto_init() {
	if File_proto_grpc_server_aws_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_grpc_server_aws_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBucketRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_grpc_server_aws_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBucketResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_grpc_server_aws_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBucketsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_grpc_server_aws_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBucketsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_grpc_server_aws_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutObjectRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_grpc_server_aws_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutObjectResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_grpc_server_aws_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetObjectRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_grpc_server_aws_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetObjectResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_grpc_server_aws_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteObjectRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_grpc_server_aws_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteObjectResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_grpc_server_aws_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBucketsResponseBucket); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_grpc_server_aws_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_grpc_server_aws_proto_goTypes,
		DependencyIndexes: file_proto_grpc_server_aws_proto_depIdxs,
		EnumInfos:         file_proto_grpc_server_aws_proto_enumTypes,
		MessageInfos:      file_proto_grpc_server_aws_proto_msgTypes,
	}.Build()
	File_proto_grpc_server_aws_proto = out.File
	file_proto_grpc_server_aws_proto_rawDesc = nil
	file_proto_grpc_server_aws_proto_goTypes = nil
	file_proto_grpc_server_aws_proto_depIdxs = nil
}