// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.3
// source: inventory_cache_init.proto

package plutoservice

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

type CacheObjType int32

const (
	CacheObjType_Book CacheObjType = 0
	CacheObjType_Mac  CacheObjType = 1
)

// Enum value maps for CacheObjType.
var (
	CacheObjType_name = map[int32]string{
		0: "Book",
		1: "Mac",
	}
	CacheObjType_value = map[string]int32{
		"Book": 0,
		"Mac":  1,
	}
)

func (x CacheObjType) Enum() *CacheObjType {
	p := new(CacheObjType)
	*p = x
	return p
}

func (x CacheObjType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CacheObjType) Descriptor() protoreflect.EnumDescriptor {
	return file_inventory_cache_init_proto_enumTypes[0].Descriptor()
}

func (CacheObjType) Type() protoreflect.EnumType {
	return &file_inventory_cache_init_proto_enumTypes[0]
}

func (x CacheObjType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CacheObjType.Descriptor instead.
func (CacheObjType) EnumDescriptor() ([]byte, []int) {
	return file_inventory_cache_init_proto_rawDescGZIP(), []int{0}
}

type InventoryCacheInitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Obj CacheObjType `protobuf:"varint,1,opt,name=obj,proto3,enum=plutoservice.CacheObjType" json:"obj,omitempty"`
}

func (x *InventoryCacheInitRequest) Reset() {
	*x = InventoryCacheInitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_cache_init_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InventoryCacheInitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InventoryCacheInitRequest) ProtoMessage() {}

func (x *InventoryCacheInitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_cache_init_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InventoryCacheInitRequest.ProtoReflect.Descriptor instead.
func (*InventoryCacheInitRequest) Descriptor() ([]byte, []int) {
	return file_inventory_cache_init_proto_rawDescGZIP(), []int{0}
}

func (x *InventoryCacheInitRequest) GetObj() CacheObjType {
	if x != nil {
		return x.Obj
	}
	return CacheObjType_Book
}

type InventoryCacheInitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Succeed bool   `protobuf:"varint,1,opt,name=Succeed,proto3" json:"Succeed,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *InventoryCacheInitResponse) Reset() {
	*x = InventoryCacheInitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_cache_init_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InventoryCacheInitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InventoryCacheInitResponse) ProtoMessage() {}

func (x *InventoryCacheInitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_cache_init_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InventoryCacheInitResponse.ProtoReflect.Descriptor instead.
func (*InventoryCacheInitResponse) Descriptor() ([]byte, []int) {
	return file_inventory_cache_init_proto_rawDescGZIP(), []int{1}
}

func (x *InventoryCacheInitResponse) GetSucceed() bool {
	if x != nil {
		return x.Succeed
	}
	return false
}

func (x *InventoryCacheInitResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_inventory_cache_init_proto protoreflect.FileDescriptor

var file_inventory_cache_init_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x5f, 0x69, 0x6e, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x6c,
	0x75, 0x74, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x49, 0x0a, 0x19, 0x49, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x61, 0x63, 0x68, 0x65, 0x49, 0x6e, 0x69, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x03, 0x6f, 0x62, 0x6a, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x70, 0x6c, 0x75, 0x74, 0x6f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x4f, 0x62, 0x6a, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x03, 0x6f, 0x62, 0x6a, 0x22, 0x50, 0x0a, 0x1a, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x43, 0x61, 0x63, 0x68, 0x65, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x21, 0x0a, 0x0c, 0x43, 0x61, 0x63, 0x68, 0x65,
	0x4f, 0x62, 0x6a, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x10,
	0x00, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x61, 0x63, 0x10, 0x01, 0x42, 0x22, 0x5a, 0x20, 0x61, 0x6d,
	0x75, 0x73, 0x69, 0x6e, 0x67, 0x2d, 0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x75, 0x74,
	0x6f, 0x2f, 0x70, 0x6c, 0x75, 0x74, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inventory_cache_init_proto_rawDescOnce sync.Once
	file_inventory_cache_init_proto_rawDescData = file_inventory_cache_init_proto_rawDesc
)

func file_inventory_cache_init_proto_rawDescGZIP() []byte {
	file_inventory_cache_init_proto_rawDescOnce.Do(func() {
		file_inventory_cache_init_proto_rawDescData = protoimpl.X.CompressGZIP(file_inventory_cache_init_proto_rawDescData)
	})
	return file_inventory_cache_init_proto_rawDescData
}

var file_inventory_cache_init_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_inventory_cache_init_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_inventory_cache_init_proto_goTypes = []interface{}{
	(CacheObjType)(0),                  // 0: plutoservice.CacheObjType
	(*InventoryCacheInitRequest)(nil),  // 1: plutoservice.InventoryCacheInitRequest
	(*InventoryCacheInitResponse)(nil), // 2: plutoservice.InventoryCacheInitResponse
}
var file_inventory_cache_init_proto_depIdxs = []int32{
	0, // 0: plutoservice.InventoryCacheInitRequest.obj:type_name -> plutoservice.CacheObjType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_inventory_cache_init_proto_init() }
func file_inventory_cache_init_proto_init() {
	if File_inventory_cache_init_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_inventory_cache_init_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InventoryCacheInitRequest); i {
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
		file_inventory_cache_init_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InventoryCacheInitResponse); i {
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
			RawDescriptor: file_inventory_cache_init_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_inventory_cache_init_proto_goTypes,
		DependencyIndexes: file_inventory_cache_init_proto_depIdxs,
		EnumInfos:         file_inventory_cache_init_proto_enumTypes,
		MessageInfos:      file_inventory_cache_init_proto_msgTypes,
	}.Build()
	File_inventory_cache_init_proto = out.File
	file_inventory_cache_init_proto_rawDesc = nil
	file_inventory_cache_init_proto_goTypes = nil
	file_inventory_cache_init_proto_depIdxs = nil
}