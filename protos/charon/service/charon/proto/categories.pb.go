// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.8.0
// source: charon/proto/categories.proto

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

type CategoryListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Page  int64  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Limit int64  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *CategoryListRequest) Reset() {
	*x = CategoryListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_charon_proto_categories_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryListRequest) ProtoMessage() {}

func (x *CategoryListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_charon_proto_categories_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryListRequest.ProtoReflect.Descriptor instead.
func (*CategoryListRequest) Descriptor() ([]byte, []int) {
	return file_charon_proto_categories_proto_rawDescGZIP(), []int{0}
}

func (x *CategoryListRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *CategoryListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *CategoryListRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type CategoryListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page       int64       `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit      int64       `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Total      int64       `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	HasNext    bool        `protobuf:"varint,4,opt,name=hasNext,proto3" json:"hasNext,omitempty"`
	Categories []*Category `protobuf:"bytes,5,rep,name=categories,proto3" json:"categories,omitempty"`
}

func (x *CategoryListResponse) Reset() {
	*x = CategoryListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_charon_proto_categories_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryListResponse) ProtoMessage() {}

func (x *CategoryListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_charon_proto_categories_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryListResponse.ProtoReflect.Descriptor instead.
func (*CategoryListResponse) Descriptor() ([]byte, []int) {
	return file_charon_proto_categories_proto_rawDescGZIP(), []int{1}
}

func (x *CategoryListResponse) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *CategoryListResponse) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *CategoryListResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *CategoryListResponse) GetHasNext() bool {
	if x != nil {
		return x.HasNext
	}
	return false
}

func (x *CategoryListResponse) GetCategories() []*Category {
	if x != nil {
		return x.Categories
	}
	return nil
}

var File_charon_proto_categories_proto protoreflect.FileDescriptor

var file_charon_proto_categories_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x68, 0x61, 0x72, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x63, 0x68, 0x61, 0x72, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x20,
	0x63, 0x68, 0x61, 0x72, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x55, 0x0a, 0x13, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0xa9, 0x01, 0x0a, 0x14, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x18, 0x0a, 0x07, 0x68, 0x61, 0x73, 0x4e, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x68, 0x61, 0x73, 0x4e, 0x65, 0x78, 0x74, 0x12, 0x37, 0x0a, 0x0a, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x63, 0x68, 0x61, 0x72, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x42, 0x24, 0x5a, 0x22, 0x61, 0x6d, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x2d, 0x78,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x6f, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x72,
	0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_charon_proto_categories_proto_rawDescOnce sync.Once
	file_charon_proto_categories_proto_rawDescData = file_charon_proto_categories_proto_rawDesc
)

func file_charon_proto_categories_proto_rawDescGZIP() []byte {
	file_charon_proto_categories_proto_rawDescOnce.Do(func() {
		file_charon_proto_categories_proto_rawDescData = protoimpl.X.CompressGZIP(file_charon_proto_categories_proto_rawDescData)
	})
	return file_charon_proto_categories_proto_rawDescData
}

var file_charon_proto_categories_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_charon_proto_categories_proto_goTypes = []interface{}{
	(*CategoryListRequest)(nil),  // 0: charonservice.CategoryListRequest
	(*CategoryListResponse)(nil), // 1: charonservice.CategoryListResponse
	(*Category)(nil),             // 2: charonservice.Category
}
var file_charon_proto_categories_proto_depIdxs = []int32{
	2, // 0: charonservice.CategoryListResponse.categories:type_name -> charonservice.Category
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_charon_proto_categories_proto_init() }
func file_charon_proto_categories_proto_init() {
	if File_charon_proto_categories_proto != nil {
		return
	}
	file_charon_proto_category_info_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_charon_proto_categories_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryListRequest); i {
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
		file_charon_proto_categories_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryListResponse); i {
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
			RawDescriptor: file_charon_proto_categories_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_charon_proto_categories_proto_goTypes,
		DependencyIndexes: file_charon_proto_categories_proto_depIdxs,
		MessageInfos:      file_charon_proto_categories_proto_msgTypes,
	}.Build()
	File_charon_proto_categories_proto = out.File
	file_charon_proto_categories_proto_rawDesc = nil
	file_charon_proto_categories_proto_goTypes = nil
	file_charon_proto_categories_proto_depIdxs = nil
}