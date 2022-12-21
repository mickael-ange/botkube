// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: executor.proto

package executor

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// RawYAML contains the Executor configuration in YAML definitions.
	RawYAML []byte `protobuf:"bytes,1,opt,name=rawYAML,proto3" json:"rawYAML,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_executor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_executor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_executor_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetRawYAML() []byte {
	if x != nil {
		return x.RawYAML
	}
	return nil
}

type ExecuteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Commands represents the exact command that was specified by the user.
	Command string `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	// Configs is a list of Executor configurations specified by users.
	Configs []*Config `protobuf:"bytes,2,rep,name=configs,proto3" json:"configs,omitempty"`
}

func (x *ExecuteRequest) Reset() {
	*x = ExecuteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_executor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteRequest) ProtoMessage() {}

func (x *ExecuteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_executor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteRequest.ProtoReflect.Descriptor instead.
func (*ExecuteRequest) Descriptor() ([]byte, []int) {
	return file_executor_proto_rawDescGZIP(), []int{1}
}

func (x *ExecuteRequest) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *ExecuteRequest) GetConfigs() []*Config {
	if x != nil {
		return x.Configs
	}
	return nil
}

type ExecuteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ExecuteResponse) Reset() {
	*x = ExecuteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_executor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteResponse) ProtoMessage() {}

func (x *ExecuteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_executor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteResponse.ProtoReflect.Descriptor instead.
func (*ExecuteResponse) Descriptor() ([]byte, []int) {
	return file_executor_proto_rawDescGZIP(), []int{2}
}

func (x *ExecuteResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type MetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version     string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *MetadataResponse) Reset() {
	*x = MetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_executor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataResponse) ProtoMessage() {}

func (x *MetadataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_executor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataResponse.ProtoReflect.Descriptor instead.
func (*MetadataResponse) Descriptor() ([]byte, []int) {
	return file_executor_proto_rawDescGZIP(), []int{3}
}

func (x *MetadataResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *MetadataResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_executor_proto protoreflect.FileDescriptor

var file_executor_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x61, 0x77, 0x59, 0x41, 0x4d, 0x4c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x72, 0x61, 0x77, 0x59, 0x41, 0x4d, 0x4c, 0x22, 0x56, 0x0a, 0x0e, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x2a, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x6f, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x73, 0x22, 0x25, 0x0a, 0x0f, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4e, 0x0a, 0x10, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x8e, 0x01, 0x0a, 0x08, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x12, 0x40, 0x0a, 0x07, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x12, 0x18, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x08, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e,
	0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x70,
	0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_executor_proto_rawDescOnce sync.Once
	file_executor_proto_rawDescData = file_executor_proto_rawDesc
)

func file_executor_proto_rawDescGZIP() []byte {
	file_executor_proto_rawDescOnce.Do(func() {
		file_executor_proto_rawDescData = protoimpl.X.CompressGZIP(file_executor_proto_rawDescData)
	})
	return file_executor_proto_rawDescData
}

var file_executor_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_executor_proto_goTypes = []interface{}{
	(*Config)(nil),           // 0: executor.Config
	(*ExecuteRequest)(nil),   // 1: executor.ExecuteRequest
	(*ExecuteResponse)(nil),  // 2: executor.ExecuteResponse
	(*MetadataResponse)(nil), // 3: executor.MetadataResponse
	(*emptypb.Empty)(nil),    // 4: google.protobuf.Empty
}
var file_executor_proto_depIdxs = []int32{
	0, // 0: executor.ExecuteRequest.configs:type_name -> executor.Config
	1, // 1: executor.Executor.Execute:input_type -> executor.ExecuteRequest
	4, // 2: executor.Executor.Metadata:input_type -> google.protobuf.Empty
	2, // 3: executor.Executor.Execute:output_type -> executor.ExecuteResponse
	3, // 4: executor.Executor.Metadata:output_type -> executor.MetadataResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_executor_proto_init() }
func file_executor_proto_init() {
	if File_executor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_executor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_executor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteRequest); i {
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
		file_executor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteResponse); i {
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
		file_executor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataResponse); i {
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
			RawDescriptor: file_executor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_executor_proto_goTypes,
		DependencyIndexes: file_executor_proto_depIdxs,
		MessageInfos:      file_executor_proto_msgTypes,
	}.Build()
	File_executor_proto = out.File
	file_executor_proto_rawDesc = nil
	file_executor_proto_goTypes = nil
	file_executor_proto_depIdxs = nil
}