// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.8.0
// source: charon/proto/category_update.proto

package charonservice

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

type CategoryUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category *Category `protobuf:"bytes,1,opt,name=Category,proto3" json:"Category,omitempty"`
}

func (x *CategoryUpdateRequest) Reset() {
	*x = CategoryUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_charon_proto_category_update_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryUpdateRequest) ProtoMessage() {}

func (x *CategoryUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_charon_proto_category_update_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryUpdateRequest.ProtoReflect.Descriptor instead.
func (*CategoryUpdateRequest) Descriptor() ([]byte, []int) {
	return file_charon_proto_category_update_proto_rawDescGZIP(), []int{0}
}

func (x *CategoryUpdateRequest) GetCategory() *Category {
	if x != nil {
		return x.Category
	}
	return nil
}

type CategoryUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *CategoryUpdateResponse) Reset() {
	*x = CategoryUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_charon_proto_category_update_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryUpdateResponse) ProtoMessage() {}

func (x *CategoryUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_charon_proto_category_update_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryUpdateResponse.ProtoReflect.Descriptor instead.
func (*CategoryUpdateResponse) Descriptor() ([]byte, []int) {
	return file_charon_proto_category_update_proto_rawDescGZIP(), []int{1}
}

func (x *CategoryUpdateResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

var File_charon_proto_category_update_proto protoreflect.FileDescriptor

var file_charon_proto_category_update_proto_rawDesc = []byte{
	0x0a, 0x22, 0x63, 0x68, 0x61, 0x72, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x68, 0x61, 0x72, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x1a, 0x20, 0x63, 0x68, 0x61, 0x72, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x15, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33,
	0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x22, 0x30, 0x0a, 0x16, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x24, 0x5a, 0x22, 0x61, 0x6d, 0x75, 0x73, 0x69, 0x6e, 0x67,
	0x2d, 0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x6f, 0x6e, 0x2f, 0x63, 0x68,
	0x61, 0x72, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_charon_proto_category_update_proto_rawDescOnce sync.Once
	file_charon_proto_category_update_proto_rawDescData = file_charon_proto_category_update_proto_rawDesc
)

func file_charon_proto_category_update_proto_rawDescGZIP() []byte {
	file_charon_proto_category_update_proto_rawDescOnce.Do(func() {
		file_charon_proto_category_update_proto_rawDescData = protoimpl.X.CompressGZIP(file_charon_proto_category_update_proto_rawDescData)
	})
	return file_charon_proto_category_update_proto_rawDescData
}

var file_charon_proto_category_update_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_charon_proto_category_update_proto_goTypes = []interface{}{
	(*CategoryUpdateRequest)(nil),  // 0: charonservice.CategoryUpdateRequest
	(*CategoryUpdateResponse)(nil), // 1: charonservice.CategoryUpdateResponse
	(*Category)(nil),               // 2: charonservice.Category
}
var file_charon_proto_category_update_proto_depIdxs = []int32{
	2, // 0: charonservice.CategoryUpdateRequest.Category:type_name -> charonservice.Category
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_charon_proto_category_update_proto_init() }
func file_charon_proto_category_update_proto_init() {
	if File_charon_proto_category_update_proto != nil {
		return
	}
	file_charon_proto_category_info_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_charon_proto_category_update_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryUpdateRequest); i {
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
		file_charon_proto_category_update_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryUpdateResponse); i {
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
			RawDescriptor: file_charon_proto_category_update_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_charon_proto_category_update_proto_goTypes,
		DependencyIndexes: file_charon_proto_category_update_proto_depIdxs,
		MessageInfos:      file_charon_proto_category_update_proto_msgTypes,
	}.Build()
	File_charon_proto_category_update_proto = out.File
	file_charon_proto_category_update_proto_rawDesc = nil
	file_charon_proto_category_update_proto_goTypes = nil
	file_charon_proto_category_update_proto_depIdxs = nil
}
