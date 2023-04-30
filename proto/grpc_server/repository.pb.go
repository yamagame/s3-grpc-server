// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: proto/grpc-server/repository.proto

package grpc_server

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

type FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_repository_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_repository_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfo.ProtoReflect.Descriptor instead.
func (*FileInfo) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_repository_proto_rawDescGZIP(), []int{0}
}

func (x *FileInfo) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

type CreateFileInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File *FileInfo `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *CreateFileInfoRequest) Reset() {
	*x = CreateFileInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_repository_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFileInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFileInfoRequest) ProtoMessage() {}

func (x *CreateFileInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_repository_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFileInfoRequest.ProtoReflect.Descriptor instead.
func (*CreateFileInfoRequest) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_repository_proto_rawDescGZIP(), []int{1}
}

func (x *CreateFileInfoRequest) GetFile() *FileInfo {
	if x != nil {
		return x.File
	}
	return nil
}

type CreateFileInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *CreateFileInfoResponse) Reset() {
	*x = CreateFileInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_repository_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFileInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFileInfoResponse) ProtoMessage() {}

func (x *CreateFileInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_repository_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFileInfoResponse.ProtoReflect.Descriptor instead.
func (*CreateFileInfoResponse) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_repository_proto_rawDescGZIP(), []int{2}
}

func (x *CreateFileInfoResponse) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type ReadFileInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *ReadFileInfoRequest) Reset() {
	*x = ReadFileInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_repository_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadFileInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadFileInfoRequest) ProtoMessage() {}

func (x *ReadFileInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_repository_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadFileInfoRequest.ProtoReflect.Descriptor instead.
func (*ReadFileInfoRequest) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_repository_proto_rawDescGZIP(), []int{3}
}

func (x *ReadFileInfoRequest) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type ReadFileInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   int64     `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	File *FileInfo `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *ReadFileInfoResponse) Reset() {
	*x = ReadFileInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_repository_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadFileInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadFileInfoResponse) ProtoMessage() {}

func (x *ReadFileInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_repository_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadFileInfoResponse.ProtoReflect.Descriptor instead.
func (*ReadFileInfoResponse) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_repository_proto_rawDescGZIP(), []int{4}
}

func (x *ReadFileInfoResponse) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ReadFileInfoResponse) GetFile() *FileInfo {
	if x != nil {
		return x.File
	}
	return nil
}

type UpdateFileInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   int64     `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	File *FileInfo `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *UpdateFileInfoRequest) Reset() {
	*x = UpdateFileInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_repository_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFileInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFileInfoRequest) ProtoMessage() {}

func (x *UpdateFileInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_repository_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFileInfoRequest.ProtoReflect.Descriptor instead.
func (*UpdateFileInfoRequest) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_repository_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateFileInfoRequest) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UpdateFileInfoRequest) GetFile() *FileInfo {
	if x != nil {
		return x.File
	}
	return nil
}

type UpdateFileInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *UpdateFileInfoResponse) Reset() {
	*x = UpdateFileInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_repository_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFileInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFileInfoResponse) ProtoMessage() {}

func (x *UpdateFileInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_repository_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFileInfoResponse.ProtoReflect.Descriptor instead.
func (*UpdateFileInfoResponse) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_repository_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateFileInfoResponse) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type DeleteFileInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DeleteFileInfoRequest) Reset() {
	*x = DeleteFileInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_repository_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFileInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFileInfoRequest) ProtoMessage() {}

func (x *DeleteFileInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_repository_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFileInfoRequest.ProtoReflect.Descriptor instead.
func (*DeleteFileInfoRequest) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_repository_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteFileInfoRequest) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type DeleteFileInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DeleteFileInfoResponse) Reset() {
	*x = DeleteFileInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_repository_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFileInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFileInfoResponse) ProtoMessage() {}

func (x *DeleteFileInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_repository_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFileInfoResponse.ProtoReflect.Descriptor instead.
func (*DeleteFileInfoResponse) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_repository_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteFileInfoResponse) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

var File_proto_grpc_server_repository_proto protoreflect.FileDescriptor

var file_proto_grpc_server_repository_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x36, 0x0a, 0x15,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x66, 0x69, 0x6c, 0x65, 0x22, 0x28, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x22, 0x25,
	0x0a, 0x13, 0x52, 0x65, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x49, 0x44, 0x22, 0x45, 0x0a, 0x14, 0x52, 0x65, 0x61, 0x64, 0x46, 0x69, 0x6c,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1d, 0x0a,
	0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x22, 0x46, 0x0a, 0x15,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1d, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x66, 0x69, 0x6c, 0x65, 0x22, 0x28, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x22, 0x27,
	0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x22, 0x28, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49,
	0x44, 0x32, 0x9a, 0x02, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x43, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x16, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0c, 0x52, 0x65, 0x61, 0x64, 0x46, 0x69, 0x6c,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x52, 0x65,
	0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x14,
	0x5a, 0x12, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_grpc_server_repository_proto_rawDescOnce sync.Once
	file_proto_grpc_server_repository_proto_rawDescData = file_proto_grpc_server_repository_proto_rawDesc
)

func file_proto_grpc_server_repository_proto_rawDescGZIP() []byte {
	file_proto_grpc_server_repository_proto_rawDescOnce.Do(func() {
		file_proto_grpc_server_repository_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_grpc_server_repository_proto_rawDescData)
	})
	return file_proto_grpc_server_repository_proto_rawDescData
}

var file_proto_grpc_server_repository_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_grpc_server_repository_proto_goTypes = []interface{}{
	(*FileInfo)(nil),               // 0: FileInfo
	(*CreateFileInfoRequest)(nil),  // 1: CreateFileInfoRequest
	(*CreateFileInfoResponse)(nil), // 2: CreateFileInfoResponse
	(*ReadFileInfoRequest)(nil),    // 3: ReadFileInfoRequest
	(*ReadFileInfoResponse)(nil),   // 4: ReadFileInfoResponse
	(*UpdateFileInfoRequest)(nil),  // 5: UpdateFileInfoRequest
	(*UpdateFileInfoResponse)(nil), // 6: UpdateFileInfoResponse
	(*DeleteFileInfoRequest)(nil),  // 7: DeleteFileInfoRequest
	(*DeleteFileInfoResponse)(nil), // 8: DeleteFileInfoResponse
}
var file_proto_grpc_server_repository_proto_depIdxs = []int32{
	0, // 0: CreateFileInfoRequest.file:type_name -> FileInfo
	0, // 1: ReadFileInfoResponse.file:type_name -> FileInfo
	0, // 2: UpdateFileInfoRequest.file:type_name -> FileInfo
	1, // 3: repository.CreateFileInfo:input_type -> CreateFileInfoRequest
	3, // 4: repository.ReadFileInfo:input_type -> ReadFileInfoRequest
	5, // 5: repository.UpdateFileInfo:input_type -> UpdateFileInfoRequest
	7, // 6: repository.DeleteFileInfo:input_type -> DeleteFileInfoRequest
	2, // 7: repository.CreateFileInfo:output_type -> CreateFileInfoResponse
	4, // 8: repository.ReadFileInfo:output_type -> ReadFileInfoResponse
	6, // 9: repository.UpdateFileInfo:output_type -> UpdateFileInfoResponse
	8, // 10: repository.DeleteFileInfo:output_type -> DeleteFileInfoResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_grpc_server_repository_proto_init() }
func file_proto_grpc_server_repository_proto_init() {
	if File_proto_grpc_server_repository_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_grpc_server_repository_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileInfo); i {
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
		file_proto_grpc_server_repository_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFileInfoRequest); i {
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
		file_proto_grpc_server_repository_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFileInfoResponse); i {
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
		file_proto_grpc_server_repository_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadFileInfoRequest); i {
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
		file_proto_grpc_server_repository_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadFileInfoResponse); i {
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
		file_proto_grpc_server_repository_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFileInfoRequest); i {
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
		file_proto_grpc_server_repository_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFileInfoResponse); i {
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
		file_proto_grpc_server_repository_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFileInfoRequest); i {
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
		file_proto_grpc_server_repository_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFileInfoResponse); i {
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
			RawDescriptor: file_proto_grpc_server_repository_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_grpc_server_repository_proto_goTypes,
		DependencyIndexes: file_proto_grpc_server_repository_proto_depIdxs,
		MessageInfos:      file_proto_grpc_server_repository_proto_msgTypes,
	}.Build()
	File_proto_grpc_server_repository_proto = out.File
	file_proto_grpc_server_repository_proto_rawDesc = nil
	file_proto_grpc_server_repository_proto_goTypes = nil
	file_proto_grpc_server_repository_proto_depIdxs = nil
}