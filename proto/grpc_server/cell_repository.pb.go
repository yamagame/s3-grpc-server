// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: proto/grpc_server/cell_repository.proto

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

type CreateCellRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cell *Cell `protobuf:"bytes,1,opt,name=cell,proto3" json:"cell,omitempty"`
}

func (x *CreateCellRequest) Reset() {
	*x = CreateCellRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_cell_repository_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCellRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCellRequest) ProtoMessage() {}

func (x *CreateCellRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_cell_repository_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCellRequest.ProtoReflect.Descriptor instead.
func (*CreateCellRequest) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_cell_repository_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCellRequest) GetCell() *Cell {
	if x != nil {
		return x.Cell
	}
	return nil
}

type CreateCellResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *CreateCellResponse) Reset() {
	*x = CreateCellResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_cell_repository_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCellResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCellResponse) ProtoMessage() {}

func (x *CreateCellResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_cell_repository_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCellResponse.ProtoReflect.Descriptor instead.
func (*CreateCellResponse) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_cell_repository_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCellResponse) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type ReadCellRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *ReadCellRequest) Reset() {
	*x = ReadCellRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_cell_repository_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadCellRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadCellRequest) ProtoMessage() {}

func (x *ReadCellRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_cell_repository_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadCellRequest.ProtoReflect.Descriptor instead.
func (*ReadCellRequest) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_cell_repository_proto_rawDescGZIP(), []int{2}
}

func (x *ReadCellRequest) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type ReadCellResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cell *Cell `protobuf:"bytes,1,opt,name=cell,proto3" json:"cell,omitempty"`
}

func (x *ReadCellResponse) Reset() {
	*x = ReadCellResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_cell_repository_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadCellResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadCellResponse) ProtoMessage() {}

func (x *ReadCellResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_cell_repository_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadCellResponse.ProtoReflect.Descriptor instead.
func (*ReadCellResponse) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_cell_repository_proto_rawDescGZIP(), []int{3}
}

func (x *ReadCellResponse) GetCell() *Cell {
	if x != nil {
		return x.Cell
	}
	return nil
}

type UpdateCellRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cell *Cell `protobuf:"bytes,1,opt,name=cell,proto3" json:"cell,omitempty"`
}

func (x *UpdateCellRequest) Reset() {
	*x = UpdateCellRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_cell_repository_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCellRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCellRequest) ProtoMessage() {}

func (x *UpdateCellRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_cell_repository_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCellRequest.ProtoReflect.Descriptor instead.
func (*UpdateCellRequest) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_cell_repository_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateCellRequest) GetCell() *Cell {
	if x != nil {
		return x.Cell
	}
	return nil
}

type UpdateCellResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *UpdateCellResponse) Reset() {
	*x = UpdateCellResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_cell_repository_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCellResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCellResponse) ProtoMessage() {}

func (x *UpdateCellResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_cell_repository_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCellResponse.ProtoReflect.Descriptor instead.
func (*UpdateCellResponse) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_cell_repository_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateCellResponse) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type DeleteCellRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DeleteCellRequest) Reset() {
	*x = DeleteCellRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_cell_repository_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCellRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCellRequest) ProtoMessage() {}

func (x *DeleteCellRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_cell_repository_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCellRequest.ProtoReflect.Descriptor instead.
func (*DeleteCellRequest) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_cell_repository_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteCellRequest) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type DeleteCellResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DeleteCellResponse) Reset() {
	*x = DeleteCellResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_cell_repository_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCellResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCellResponse) ProtoMessage() {}

func (x *DeleteCellResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_cell_repository_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCellResponse.ProtoReflect.Descriptor instead.
func (*DeleteCellResponse) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_cell_repository_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteCellResponse) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type ListCellRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListCellRequest) Reset() {
	*x = ListCellRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_cell_repository_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCellRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCellRequest) ProtoMessage() {}

func (x *ListCellRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_cell_repository_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCellRequest.ProtoReflect.Descriptor instead.
func (*ListCellRequest) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_cell_repository_proto_rawDescGZIP(), []int{8}
}

type ListCellResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cells []*Cell `protobuf:"bytes,1,rep,name=cells,proto3" json:"cells,omitempty"`
}

func (x *ListCellResponse) Reset() {
	*x = ListCellResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_server_cell_repository_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCellResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCellResponse) ProtoMessage() {}

func (x *ListCellResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_server_cell_repository_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCellResponse.ProtoReflect.Descriptor instead.
func (*ListCellResponse) Descriptor() ([]byte, []int) {
	return file_proto_grpc_server_cell_repository_proto_rawDescGZIP(), []int{9}
}

func (x *ListCellResponse) GetCells() []*Cell {
	if x != nil {
		return x.Cells
	}
	return nil
}

var File_proto_grpc_server_cell_repository_proto protoreflect.FileDescriptor

var file_proto_grpc_server_cell_repository_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x63, 0x65, 0x6c, 0x6c, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x04, 0x63, 0x65, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x43,
	0x65, 0x6c, 0x6c, 0x52, 0x04, 0x63, 0x65, 0x6c, 0x6c, 0x22, 0x24, 0x0a, 0x12, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x22,
	0x21, 0x0a, 0x0f, 0x52, 0x65, 0x61, 0x64, 0x43, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x49, 0x44, 0x22, 0x2d, 0x0a, 0x10, 0x52, 0x65, 0x61, 0x64, 0x43, 0x65, 0x6c, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x63, 0x65, 0x6c, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x43, 0x65, 0x6c, 0x6c, 0x52, 0x04, 0x63, 0x65, 0x6c,
	0x6c, 0x22, 0x2e, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x65, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x04, 0x63, 0x65, 0x6c, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x43, 0x65, 0x6c, 0x6c, 0x52, 0x04, 0x63, 0x65, 0x6c,
	0x6c, 0x22, 0x24, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x65, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x22, 0x24, 0x0a, 0x12,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x49, 0x44, 0x22, 0x11, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x65, 0x6c, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2f, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x65, 0x6c,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x05, 0x63, 0x65, 0x6c,
	0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x43, 0x65, 0x6c, 0x6c, 0x52,
	0x05, 0x63, 0x65, 0x6c, 0x6c, 0x73, 0x32, 0x8d, 0x02, 0x0a, 0x0e, 0x43, 0x65, 0x6c, 0x6c, 0x52,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x33, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x65, 0x6c, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2d,
	0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x10, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x43, 0x65, 0x6c,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x43,
	0x65, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x33, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x10, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x14, 0x5a, 0x12, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_grpc_server_cell_repository_proto_rawDescOnce sync.Once
	file_proto_grpc_server_cell_repository_proto_rawDescData = file_proto_grpc_server_cell_repository_proto_rawDesc
)

func file_proto_grpc_server_cell_repository_proto_rawDescGZIP() []byte {
	file_proto_grpc_server_cell_repository_proto_rawDescOnce.Do(func() {
		file_proto_grpc_server_cell_repository_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_grpc_server_cell_repository_proto_rawDescData)
	})
	return file_proto_grpc_server_cell_repository_proto_rawDescData
}

var file_proto_grpc_server_cell_repository_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_grpc_server_cell_repository_proto_goTypes = []interface{}{
	(*CreateCellRequest)(nil),  // 0: CreateCellRequest
	(*CreateCellResponse)(nil), // 1: CreateCellResponse
	(*ReadCellRequest)(nil),    // 2: ReadCellRequest
	(*ReadCellResponse)(nil),   // 3: ReadCellResponse
	(*UpdateCellRequest)(nil),  // 4: UpdateCellRequest
	(*UpdateCellResponse)(nil), // 5: UpdateCellResponse
	(*DeleteCellRequest)(nil),  // 6: DeleteCellRequest
	(*DeleteCellResponse)(nil), // 7: DeleteCellResponse
	(*ListCellRequest)(nil),    // 8: ListCellRequest
	(*ListCellResponse)(nil),   // 9: ListCellResponse
	(*Cell)(nil),               // 10: Cell
}
var file_proto_grpc_server_cell_repository_proto_depIdxs = []int32{
	10, // 0: CreateCellRequest.cell:type_name -> Cell
	10, // 1: ReadCellResponse.cell:type_name -> Cell
	10, // 2: UpdateCellRequest.cell:type_name -> Cell
	10, // 3: ListCellResponse.cells:type_name -> Cell
	0,  // 4: CellRepository.Create:input_type -> CreateCellRequest
	2,  // 5: CellRepository.Read:input_type -> ReadCellRequest
	4,  // 6: CellRepository.Update:input_type -> UpdateCellRequest
	6,  // 7: CellRepository.Delete:input_type -> DeleteCellRequest
	8,  // 8: CellRepository.List:input_type -> ListCellRequest
	1,  // 9: CellRepository.Create:output_type -> CreateCellResponse
	3,  // 10: CellRepository.Read:output_type -> ReadCellResponse
	5,  // 11: CellRepository.Update:output_type -> UpdateCellResponse
	7,  // 12: CellRepository.Delete:output_type -> DeleteCellResponse
	9,  // 13: CellRepository.List:output_type -> ListCellResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_grpc_server_cell_repository_proto_init() }
func file_proto_grpc_server_cell_repository_proto_init() {
	if File_proto_grpc_server_cell_repository_proto != nil {
		return
	}
	file_proto_grpc_server_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_grpc_server_cell_repository_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCellRequest); i {
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
		file_proto_grpc_server_cell_repository_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCellResponse); i {
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
		file_proto_grpc_server_cell_repository_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadCellRequest); i {
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
		file_proto_grpc_server_cell_repository_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadCellResponse); i {
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
		file_proto_grpc_server_cell_repository_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCellRequest); i {
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
		file_proto_grpc_server_cell_repository_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCellResponse); i {
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
		file_proto_grpc_server_cell_repository_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCellRequest); i {
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
		file_proto_grpc_server_cell_repository_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCellResponse); i {
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
		file_proto_grpc_server_cell_repository_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCellRequest); i {
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
		file_proto_grpc_server_cell_repository_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCellResponse); i {
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
			RawDescriptor: file_proto_grpc_server_cell_repository_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_grpc_server_cell_repository_proto_goTypes,
		DependencyIndexes: file_proto_grpc_server_cell_repository_proto_depIdxs,
		MessageInfos:      file_proto_grpc_server_cell_repository_proto_msgTypes,
	}.Build()
	File_proto_grpc_server_cell_repository_proto = out.File
	file_proto_grpc_server_cell_repository_proto_rawDesc = nil
	file_proto_grpc_server_cell_repository_proto_goTypes = nil
	file_proto_grpc_server_cell_repository_proto_depIdxs = nil
}
