// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: proto/grpc-server/model.proto

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
	return file_proto_grpc_server_model_proto_enumTypes[0].Descriptor()
}

func (Result) Type() protoreflect.EnumType {
	return &file_proto_grpc_server_model_proto_enumTypes[0]
}

func (x Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Result.Descriptor instead.
func (Result) EnumDescriptor() ([]byte, []int) {
	return file_proto_grpc_server_model_proto_rawDescGZIP(), []int{0}
}

var File_proto_grpc_server_model_proto protoreflect.FileDescriptor

var file_proto_grpc_server_model_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a,
	0x28, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x44,
	0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x01,
	0x12, 0x07, 0x0a, 0x03, 0x45, 0x52, 0x52, 0x10, 0x02, 0x42, 0x14, 0x5a, 0x12, 0x73, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_grpc_server_model_proto_rawDescOnce sync.Once
	file_proto_grpc_server_model_proto_rawDescData = file_proto_grpc_server_model_proto_rawDesc
)

func file_proto_grpc_server_model_proto_rawDescGZIP() []byte {
	file_proto_grpc_server_model_proto_rawDescOnce.Do(func() {
		file_proto_grpc_server_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_grpc_server_model_proto_rawDescData)
	})
	return file_proto_grpc_server_model_proto_rawDescData
}

var file_proto_grpc_server_model_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_grpc_server_model_proto_goTypes = []interface{}{
	(Result)(0), // 0: Result
}
var file_proto_grpc_server_model_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_grpc_server_model_proto_init() }
func file_proto_grpc_server_model_proto_init() {
	if File_proto_grpc_server_model_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_grpc_server_model_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_grpc_server_model_proto_goTypes,
		DependencyIndexes: file_proto_grpc_server_model_proto_depIdxs,
		EnumInfos:         file_proto_grpc_server_model_proto_enumTypes,
	}.Build()
	File_proto_grpc_server_model_proto = out.File
	file_proto_grpc_server_model_proto_rawDesc = nil
	file_proto_grpc_server_model_proto_goTypes = nil
	file_proto_grpc_server_model_proto_depIdxs = nil
}
