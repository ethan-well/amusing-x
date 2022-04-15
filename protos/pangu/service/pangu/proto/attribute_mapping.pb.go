// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.3
// source: pangu/proto/attribute_mapping.proto

package proto

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

type AttributeMapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AttrId       string `protobuf:"bytes,2,opt,name=attr_id,json=attrId,proto3" json:"attr_id,omitempty"`
	SubProductId int64  `protobuf:"varint,3,opt,name=sub_product_id,json=subProductId,proto3" json:"sub_product_id,omitempty"`
	AttrValue    string `protobuf:"bytes,4,opt,name=attr_value,json=attrValue,proto3" json:"attr_value,omitempty"`
}

func (x *AttributeMapping) Reset() {
	*x = AttributeMapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pangu_proto_attribute_mapping_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttributeMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributeMapping) ProtoMessage() {}

func (x *AttributeMapping) ProtoReflect() protoreflect.Message {
	mi := &file_pangu_proto_attribute_mapping_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributeMapping.ProtoReflect.Descriptor instead.
func (*AttributeMapping) Descriptor() ([]byte, []int) {
	return file_pangu_proto_attribute_mapping_proto_rawDescGZIP(), []int{0}
}

func (x *AttributeMapping) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AttributeMapping) GetAttrId() string {
	if x != nil {
		return x.AttrId
	}
	return ""
}

func (x *AttributeMapping) GetSubProductId() int64 {
	if x != nil {
		return x.SubProductId
	}
	return 0
}

func (x *AttributeMapping) GetAttrValue() string {
	if x != nil {
		return x.AttrValue
	}
	return ""
}

var File_pangu_proto_attribute_mapping_proto protoreflect.FileDescriptor

var file_pangu_proto_attribute_mapping_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x61, 0x6e, 0x67, 0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x61, 0x6e, 0x67, 0x75, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x22, 0x80, 0x01, 0x0a, 0x10, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x74, 0x74, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x74, 0x74, 0x72, 0x49,
	0x64, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x75, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x5f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x74, 0x74,
	0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x38, 0x5a, 0x36, 0x61, 0x6d, 0x75, 0x73, 0x69, 0x6e,
	0x67, 0x78, 0x2e, 0x66, 0x69, 0x74, 0x2f, 0x61, 0x6d, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x78, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x61, 0x6e, 0x67, 0x75, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x61, 0x6e, 0x67, 0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pangu_proto_attribute_mapping_proto_rawDescOnce sync.Once
	file_pangu_proto_attribute_mapping_proto_rawDescData = file_pangu_proto_attribute_mapping_proto_rawDesc
)

func file_pangu_proto_attribute_mapping_proto_rawDescGZIP() []byte {
	file_pangu_proto_attribute_mapping_proto_rawDescOnce.Do(func() {
		file_pangu_proto_attribute_mapping_proto_rawDescData = protoimpl.X.CompressGZIP(file_pangu_proto_attribute_mapping_proto_rawDescData)
	})
	return file_pangu_proto_attribute_mapping_proto_rawDescData
}

var file_pangu_proto_attribute_mapping_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pangu_proto_attribute_mapping_proto_goTypes = []interface{}{
	(*AttributeMapping)(nil), // 0: panguservice.AttributeMapping
}
var file_pangu_proto_attribute_mapping_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pangu_proto_attribute_mapping_proto_init() }
func file_pangu_proto_attribute_mapping_proto_init() {
	if File_pangu_proto_attribute_mapping_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pangu_proto_attribute_mapping_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttributeMapping); i {
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
			RawDescriptor: file_pangu_proto_attribute_mapping_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pangu_proto_attribute_mapping_proto_goTypes,
		DependencyIndexes: file_pangu_proto_attribute_mapping_proto_depIdxs,
		MessageInfos:      file_pangu_proto_attribute_mapping_proto_msgTypes,
	}.Build()
	File_pangu_proto_attribute_mapping_proto = out.File
	file_pangu_proto_attribute_mapping_proto_rawDesc = nil
	file_pangu_proto_attribute_mapping_proto_goTypes = nil
	file_pangu_proto_attribute_mapping_proto_depIdxs = nil
}
