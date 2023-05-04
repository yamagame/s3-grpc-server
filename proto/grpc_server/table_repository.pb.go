// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: proto/grpc_server/table_repository.proto

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

type CreateTableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Table *Table `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
}

func (x *CreateTableRequest) Reset() {
	*x = CreateTableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_table_repository_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTableRequest) ProtoMessage() {}

func (x *CreateTableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_table_repository_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTableRequest.ProtoReflect.Descriptor instead.
func (*CreateTableRequest) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_table_repository_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTableRequest) GetTable() *Table {
	if x != nil {
		return x.Table
	}
	return nil
}

type CreateTableResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *CreateTableResponse) Reset() {
	*x = CreateTableResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_table_repository_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTableResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTableResponse) ProtoMessage() {}

func (x *CreateTableResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_table_repository_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTableResponse.ProtoReflect.Descriptor instead.
func (*CreateTableResponse) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_table_repository_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTableResponse) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type ReadTableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *ReadTableRequest) Reset() {
	*x = ReadTableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_table_repository_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadTableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadTableRequest) ProtoMessage() {}

func (x *ReadTableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_table_repository_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadTableRequest.ProtoReflect.Descriptor instead.
func (*ReadTableRequest) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_table_repository_proto_rawDescGZIP(), []int{2}
}

func (x *ReadTableRequest) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type ReadTableResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Table *Table `protobuf:"bytes,2,opt,name=table,proto3" json:"table,omitempty"`
}

func (x *ReadTableResponse) Reset() {
	*x = ReadTableResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_table_repository_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadTableResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadTableResponse) ProtoMessage() {}

func (x *ReadTableResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_table_repository_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadTableResponse.ProtoReflect.Descriptor instead.
func (*ReadTableResponse) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_table_repository_proto_rawDescGZIP(), []int{3}
}

func (x *ReadTableResponse) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ReadTableResponse) GetTable() *Table {
	if x != nil {
		return x.Table
	}
	return nil
}

type UpdateTableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Table *Table `protobuf:"bytes,2,opt,name=table,proto3" json:"table,omitempty"`
}

func (x *UpdateTableRequest) Reset() {
	*x = UpdateTableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_table_repository_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTableRequest) ProtoMessage() {}

func (x *UpdateTableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_table_repository_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTableRequest.ProtoReflect.Descriptor instead.
func (*UpdateTableRequest) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_table_repository_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateTableRequest) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UpdateTableRequest) GetTable() *Table {
	if x != nil {
		return x.Table
	}
	return nil
}

type UpdateTableResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *UpdateTableResponse) Reset() {
	*x = UpdateTableResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_table_repository_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTableResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTableResponse) ProtoMessage() {}

func (x *UpdateTableResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_table_repository_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTableResponse.ProtoReflect.Descriptor instead.
func (*UpdateTableResponse) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_table_repository_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateTableResponse) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type DeleteTableRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DeleteTableRequest) Reset() {
	*x = DeleteTableRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_table_repository_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTableRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTableRequest) ProtoMessage() {}

func (x *DeleteTableRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_table_repository_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTableRequest.ProtoReflect.Descriptor instead.
func (*DeleteTableRequest) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_table_repository_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteTableRequest) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type DeleteTableResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DeleteTableResponse) Reset() {
	*x = DeleteTableResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_table_repository_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTableResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTableResponse) ProtoMessage() {}

func (x *DeleteTableResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_table_repository_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTableResponse.ProtoReflect.Descriptor instead.
func (*DeleteTableResponse) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_table_repository_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteTableResponse) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

var File_proto_grpc_server_table_repository_proto protoreflect.FileDescriptor

var file_proto_grpc_server_table_repository_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32, 0x0a, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1c, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x06, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x25,
	0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x49, 0x44, 0x22, 0x22, 0x0a, 0x10, 0x52, 0x65, 0x61, 0x64, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x22, 0x41, 0x0a, 0x11, 0x52, 0x65, 0x61,
	0x64, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1c,
	0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x42, 0x0a, 0x12,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x1c, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x06, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x22, 0x25, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x22, 0x25, 0x0a,
	0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x49, 0x44, 0x32, 0xe7, 0x01, 0x0a, 0x0f, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x35, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x13, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x2f, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x11, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x52, 0x65, 0x61,
	0x64, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x35, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x13, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x14,
	0x5a, 0x12, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_grpc_server_table_repository_proto_rawDescOnce sync.Once
	file_proto_grpc_server_table_repository_proto_rawDescData = file_proto_grpc_server_table_repository_proto_rawDesc
)

func file_proto_grpc_server_table_repository_proto_rawDescGZIP() []byte {
	file_proto_grpc_server_table_repository_proto_rawDescOnce.Do(func() {
		file_proto_grpc_server_table_repository_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_grpc_server_table_repository_proto_rawDescData)
	})
	return file_proto_grpc_server_table_repository_proto_rawDescData
}

var file_proto_grpc_server_table_repository_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_grpc_server_table_repository_proto_goTypes = []interface{}{
	(*CreateTableRequest)(nil),  // 0: CreateTableRequest
	(*CreateTableResponse)(nil), // 1: CreateTableResponse
	(*ReadTableRequest)(nil),    // 2: ReadTableRequest
	(*ReadTableResponse)(nil),   // 3: ReadTableResponse
	(*UpdateTableRequest)(nil),  // 4: UpdateTableRequest
	(*UpdateTableResponse)(nil), // 5: UpdateTableResponse
	(*DeleteTableRequest)(nil),  // 6: DeleteTableRequest
	(*DeleteTableResponse)(nil), // 7: DeleteTableResponse
	(*Table)(nil),               // 8: Table
}
var file_proto_grpc_server_table_repository_proto_depIdxs = []int32{
	8, // 0: CreateTableRequest.table:type_name -> Table
	8, // 1: ReadTableResponse.table:type_name -> Table
	8, // 2: UpdateTableRequest.table:type_name -> Table
	0, // 3: TableRepository.Create:input_type -> CreateTableRequest
	2, // 4: TableRepository.Read:input_type -> ReadTableRequest
	4, // 5: TableRepository.Update:input_type -> UpdateTableRequest
	6, // 6: TableRepository.Delete:input_type -> DeleteTableRequest
	1, // 7: TableRepository.Create:output_type -> CreateTableResponse
	3, // 8: TableRepository.Read:output_type -> ReadTableResponse
	5, // 9: TableRepository.Update:output_type -> UpdateTableResponse
	7, // 10: TableRepository.Delete:output_type -> DeleteTableResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_grpc_server_table_repository_proto_init() }
func file_proto_grpc_server_table_repository_proto_init() {
	if File_proto_grpc_server_table_repository_proto != nil {
		return
	}
	file_proto_grpc_server_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_grpc_server_table_repository_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTableRequest); i {
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
		file_proto_grpc_server_table_repository_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTableResponse); i {
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
		file_proto_grpc_server_table_repository_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadTableRequest); i {
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
		file_proto_grpc_server_table_repository_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadTableResponse); i {
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
		file_proto_grpc_server_table_repository_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTableRequest); i {
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
		file_proto_grpc_server_table_repository_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTableResponse); i {
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
		file_proto_grpc_server_table_repository_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTableRequest); i {
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
		file_proto_grpc_server_table_repository_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTableResponse); i {
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
			RawDescriptor: file_proto_grpc_server_table_repository_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_grpc_server_table_repository_proto_goTypes,
		DependencyIndexes: file_proto_grpc_server_table_repository_proto_depIdxs,
		MessageInfos:      file_proto_grpc_server_table_repository_proto_msgTypes,
	}.Build()
	File_proto_grpc_server_table_repository_proto = out.File
	file_proto_grpc_server_table_repository_proto_rawDesc = nil
	file_proto_grpc_server_table_repository_proto_goTypes = nil
	file_proto_grpc_server_table_repository_proto_depIdxs = nil
}
