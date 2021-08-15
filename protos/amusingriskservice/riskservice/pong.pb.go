// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.3
// source: pong.proto

package riskservice

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

type PongReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *PongReply) Reset() {
	*x = PongReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pong_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PongReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PongReply) ProtoMessage() {}

func (x *PongReply) ProtoReflect() protoreflect.Message {
	mi := &file_pong_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PongReply.ProtoReflect.Descriptor instead.
func (*PongReply) Descriptor() ([]byte, []int) {
	return file_pong_proto_rawDescGZIP(), []int{0}
}

func (x *PongReply) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_pong_proto protoreflect.FileDescriptor

var file_pong_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x6f, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x72, 0x69,
	0x73, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x23, 0x0a, 0x09, 0x50, 0x6f, 0x6e,
	0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x27,
	0x5a, 0x25, 0x61, 0x6d, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x2d, 0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x72, 0x69, 0x73, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x69, 0x73, 0x6b,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pong_proto_rawDescOnce sync.Once
	file_pong_proto_rawDescData = file_pong_proto_rawDesc
)

func file_pong_proto_rawDescGZIP() []byte {
	file_pong_proto_rawDescOnce.Do(func() {
		file_pong_proto_rawDescData = protoimpl.X.CompressGZIP(file_pong_proto_rawDescData)
	})
	return file_pong_proto_rawDescData
}

var file_pong_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pong_proto_goTypes = []interface{}{
	(*PongReply)(nil), // 0: riskservice.PongReply
}
var file_pong_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pong_proto_init() }
func file_pong_proto_init() {
	if File_pong_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pong_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PongReply); i {
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
			RawDescriptor: file_pong_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pong_proto_goTypes,
		DependencyIndexes: file_pong_proto_depIdxs,
		MessageInfos:      file_pong_proto_msgTypes,
	}.Build()
	File_pong_proto = out.File
	file_pong_proto_rawDesc = nil
	file_pong_proto_goTypes = nil
	file_pong_proto_depIdxs = nil
}
