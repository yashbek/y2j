// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: main/v1/queue_up.proto

package mainv1

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

type QueueUpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *QueueUpRequest) Reset() {
	*x = QueueUpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_v1_queue_up_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueUpRequest) ProtoMessage() {}

func (x *QueueUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_main_v1_queue_up_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueUpRequest.ProtoReflect.Descriptor instead.
func (*QueueUpRequest) Descriptor() ([]byte, []int) {
	return file_main_v1_queue_up_proto_rawDescGZIP(), []int{0}
}

type QueueUpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *QueueUpResponse) Reset() {
	*x = QueueUpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_main_v1_queue_up_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueUpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueUpResponse) ProtoMessage() {}

func (x *QueueUpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_main_v1_queue_up_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueUpResponse.ProtoReflect.Descriptor instead.
func (*QueueUpResponse) Descriptor() ([]byte, []int) {
	return file_main_v1_queue_up_proto_rawDescGZIP(), []int{1}
}

func (x *QueueUpResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_main_v1_queue_up_proto protoreflect.FileDescriptor

var file_main_v1_queue_up_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f,
	0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x22, 0x10, 0x0a, 0x0e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x29, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x75, 0x65, 0x55, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x83,
	0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0c,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x55, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x29,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x73, 0x68, 0x62,
	0x65, 0x6b, 0x2f, 0x79, 0x32, 0x6a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f,
	0x76, 0x31, 0x3b, 0x6d, 0x61, 0x69, 0x6e, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x58, 0x58, 0xaa,
	0x02, 0x07, 0x4d, 0x61, 0x69, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x4d, 0x61, 0x69, 0x6e,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13, 0x4d, 0x61, 0x69, 0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x4d, 0x61, 0x69, 0x6e,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_main_v1_queue_up_proto_rawDescOnce sync.Once
	file_main_v1_queue_up_proto_rawDescData = file_main_v1_queue_up_proto_rawDesc
)

func file_main_v1_queue_up_proto_rawDescGZIP() []byte {
	file_main_v1_queue_up_proto_rawDescOnce.Do(func() {
		file_main_v1_queue_up_proto_rawDescData = protoimpl.X.CompressGZIP(file_main_v1_queue_up_proto_rawDescData)
	})
	return file_main_v1_queue_up_proto_rawDescData
}

var file_main_v1_queue_up_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_main_v1_queue_up_proto_goTypes = []interface{}{
	(*QueueUpRequest)(nil),  // 0: main.v1.QueueUpRequest
	(*QueueUpResponse)(nil), // 1: main.v1.QueueUpResponse
}
var file_main_v1_queue_up_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_main_v1_queue_up_proto_init() }
func file_main_v1_queue_up_proto_init() {
	if File_main_v1_queue_up_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_main_v1_queue_up_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueUpRequest); i {
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
		file_main_v1_queue_up_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueUpResponse); i {
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
			RawDescriptor: file_main_v1_queue_up_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_main_v1_queue_up_proto_goTypes,
		DependencyIndexes: file_main_v1_queue_up_proto_depIdxs,
		MessageInfos:      file_main_v1_queue_up_proto_msgTypes,
	}.Build()
	File_main_v1_queue_up_proto = out.File
	file_main_v1_queue_up_proto_rawDesc = nil
	file_main_v1_queue_up_proto_goTypes = nil
	file_main_v1_queue_up_proto_depIdxs = nil
}
