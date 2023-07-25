// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.8.0
// source: europa/proto/sub_product_pictures.proto

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

type SubProductPicturesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubProductIds []int64 `protobuf:"varint,1,rep,packed,name=sub_product_ids,json=subProductIds,proto3" json:"sub_product_ids,omitempty"`
}

func (x *SubProductPicturesRequest) Reset() {
	*x = SubProductPicturesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_europa_proto_sub_product_pictures_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubProductPicturesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubProductPicturesRequest) ProtoMessage() {}

func (x *SubProductPicturesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_europa_proto_sub_product_pictures_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubProductPicturesRequest.ProtoReflect.Descriptor instead.
func (*SubProductPicturesRequest) Descriptor() ([]byte, []int) {
	return file_europa_proto_sub_product_pictures_proto_rawDescGZIP(), []int{0}
}

func (x *SubProductPicturesRequest) GetSubProductIds() []int64 {
	if x != nil {
		return x.SubProductIds
	}
	return nil
}

type SubProductPicture struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubProductId int64    `protobuf:"varint,1,opt,name=sub_product_id,json=subProductId,proto3" json:"sub_product_id,omitempty"`
	Pictures     []string `protobuf:"bytes,2,rep,name=pictures,proto3" json:"pictures,omitempty"`
}

func (x *SubProductPicture) Reset() {
	*x = SubProductPicture{}
	if protoimpl.UnsafeEnabled {
		mi := &file_europa_proto_sub_product_pictures_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubProductPicture) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubProductPicture) ProtoMessage() {}

func (x *SubProductPicture) ProtoReflect() protoreflect.Message {
	mi := &file_europa_proto_sub_product_pictures_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubProductPicture.ProtoReflect.Descriptor instead.
func (*SubProductPicture) Descriptor() ([]byte, []int) {
	return file_europa_proto_sub_product_pictures_proto_rawDescGZIP(), []int{1}
}

func (x *SubProductPicture) GetSubProductId() int64 {
	if x != nil {
		return x.SubProductId
	}
	return 0
}

func (x *SubProductPicture) GetPictures() []string {
	if x != nil {
		return x.Pictures
	}
	return nil
}

type SubProductPicturesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubProductPictures []*SubProductPicture `protobuf:"bytes,1,rep,name=sub_product_pictures,json=subProductPictures,proto3" json:"sub_product_pictures,omitempty"`
}

func (x *SubProductPicturesResponse) Reset() {
	*x = SubProductPicturesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_europa_proto_sub_product_pictures_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubProductPicturesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubProductPicturesResponse) ProtoMessage() {}

func (x *SubProductPicturesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_europa_proto_sub_product_pictures_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubProductPicturesResponse.ProtoReflect.Descriptor instead.
func (*SubProductPicturesResponse) Descriptor() ([]byte, []int) {
	return file_europa_proto_sub_product_pictures_proto_rawDescGZIP(), []int{2}
}

func (x *SubProductPicturesResponse) GetSubProductPictures() []*SubProductPicture {
	if x != nil {
		return x.SubProductPictures
	}
	return nil
}

var File_europa_proto_sub_product_pictures_proto protoreflect.FileDescriptor

var file_europa_proto_sub_product_pictures_proto_rawDesc = []byte{
	0x0a, 0x27, 0x65, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73,
	0x75, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x70, 0x69, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x19, 0x53, 0x75, 0x62,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x75, 0x62, 0x5f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x0d, 0x73, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x73, 0x22, 0x55,
	0x0a, 0x11, 0x53, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x69, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x75, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x73, 0x75, 0x62,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x69, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x70, 0x69, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x73, 0x22, 0x62, 0x0a, 0x1a, 0x53, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x14, 0x73, 0x75, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x53, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x69,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x12, 0x73, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x73, 0x42, 0x3a, 0x5a, 0x38, 0x61, 0x6d, 0x75,
	0x73, 0x69, 0x6e, 0x67, 0x78, 0x2e, 0x66, 0x69, 0x74, 0x2f, 0x61, 0x6d, 0x75, 0x73, 0x69, 0x6e,
	0x67, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x75, 0x72, 0x6f, 0x70, 0x61,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x65, 0x75, 0x72, 0x6f, 0x70, 0x61, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_europa_proto_sub_product_pictures_proto_rawDescOnce sync.Once
	file_europa_proto_sub_product_pictures_proto_rawDescData = file_europa_proto_sub_product_pictures_proto_rawDesc
)

func file_europa_proto_sub_product_pictures_proto_rawDescGZIP() []byte {
	file_europa_proto_sub_product_pictures_proto_rawDescOnce.Do(func() {
		file_europa_proto_sub_product_pictures_proto_rawDescData = protoimpl.X.CompressGZIP(file_europa_proto_sub_product_pictures_proto_rawDescData)
	})
	return file_europa_proto_sub_product_pictures_proto_rawDescData
}

var file_europa_proto_sub_product_pictures_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_europa_proto_sub_product_pictures_proto_goTypes = []interface{}{
	(*SubProductPicturesRequest)(nil),  // 0: SubProductPicturesRequest
	(*SubProductPicture)(nil),          // 1: SubProductPicture
	(*SubProductPicturesResponse)(nil), // 2: SubProductPicturesResponse
}
var file_europa_proto_sub_product_pictures_proto_depIdxs = []int32{
	1, // 0: SubProductPicturesResponse.sub_product_pictures:type_name -> SubProductPicture
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_europa_proto_sub_product_pictures_proto_init() }
func file_europa_proto_sub_product_pictures_proto_init() {
	if File_europa_proto_sub_product_pictures_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_europa_proto_sub_product_pictures_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubProductPicturesRequest); i {
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
		file_europa_proto_sub_product_pictures_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubProductPicture); i {
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
		file_europa_proto_sub_product_pictures_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubProductPicturesResponse); i {
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
			RawDescriptor: file_europa_proto_sub_product_pictures_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_europa_proto_sub_product_pictures_proto_goTypes,
		DependencyIndexes: file_europa_proto_sub_product_pictures_proto_depIdxs,
		MessageInfos:      file_europa_proto_sub_product_pictures_proto_msgTypes,
	}.Build()
	File_europa_proto_sub_product_pictures_proto = out.File
	file_europa_proto_sub_product_pictures_proto_rawDesc = nil
	file_europa_proto_sub_product_pictures_proto_goTypes = nil
	file_europa_proto_sub_product_pictures_proto_depIdxs = nil
}