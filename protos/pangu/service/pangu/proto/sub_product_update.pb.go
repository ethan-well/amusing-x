// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.3
// source: pangu/proto/sub_product_update.proto

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

type SubProductUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Desc        string     `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	ProductId   int64      `protobuf:"varint,4,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Currency    string     `protobuf:"bytes,5,opt,name=currency,proto3" json:"currency,omitempty"`
	Price       int64      `protobuf:"varint,6,opt,name=price,proto3" json:"price,omitempty"`
	Stock       int64      `protobuf:"varint,7,opt,name=stock,proto3" json:"stock,omitempty"`
	AttributeId []int64    `protobuf:"varint,8,rep,packed,name=attribute_id,json=attributeId,proto3" json:"attribute_id,omitempty"`
	Pictures    []*Picture `protobuf:"bytes,9,rep,name=pictures,proto3" json:"pictures,omitempty"`
}

func (x *SubProductUpdateRequest) Reset() {
	*x = SubProductUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pangu_proto_sub_product_update_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubProductUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubProductUpdateRequest) ProtoMessage() {}

func (x *SubProductUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pangu_proto_sub_product_update_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubProductUpdateRequest.ProtoReflect.Descriptor instead.
func (*SubProductUpdateRequest) Descriptor() ([]byte, []int) {
	return file_pangu_proto_sub_product_update_proto_rawDescGZIP(), []int{0}
}

func (x *SubProductUpdateRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SubProductUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SubProductUpdateRequest) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *SubProductUpdateRequest) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *SubProductUpdateRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *SubProductUpdateRequest) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *SubProductUpdateRequest) GetStock() int64 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *SubProductUpdateRequest) GetAttributeId() []int64 {
	if x != nil {
		return x.AttributeId
	}
	return nil
}

func (x *SubProductUpdateRequest) GetPictures() []*Picture {
	if x != nil {
		return x.Pictures
	}
	return nil
}

var File_pangu_proto_sub_product_update_proto protoreflect.FileDescriptor

var file_pangu_proto_sub_product_update_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x61, 0x6e, 0x67, 0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x75,
	0x62, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x61, 0x6e, 0x67, 0x75, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x1a, 0x1d, 0x70, 0x61, 0x6e, 0x67, 0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x73, 0x75, 0x62, 0x5f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x8e, 0x02, 0x0a, 0x17, 0x53, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x21,
	0x0a, 0x0c, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x03, 0x52, 0x0b, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x49,
	0x64, 0x12, 0x31, 0x0a, 0x08, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x09, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x61, 0x6e, 0x67, 0x75, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x08, 0x70, 0x69, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x42, 0x38, 0x5a, 0x36, 0x61, 0x6d, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x78,
	0x2e, 0x66, 0x69, 0x74, 0x2f, 0x61, 0x6d, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x78, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x61, 0x6e, 0x67, 0x75, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x70, 0x61, 0x6e, 0x67, 0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pangu_proto_sub_product_update_proto_rawDescOnce sync.Once
	file_pangu_proto_sub_product_update_proto_rawDescData = file_pangu_proto_sub_product_update_proto_rawDesc
)

func file_pangu_proto_sub_product_update_proto_rawDescGZIP() []byte {
	file_pangu_proto_sub_product_update_proto_rawDescOnce.Do(func() {
		file_pangu_proto_sub_product_update_proto_rawDescData = protoimpl.X.CompressGZIP(file_pangu_proto_sub_product_update_proto_rawDescData)
	})
	return file_pangu_proto_sub_product_update_proto_rawDescData
}

var file_pangu_proto_sub_product_update_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pangu_proto_sub_product_update_proto_goTypes = []interface{}{
	(*SubProductUpdateRequest)(nil), // 0: panguservice.SubProductUpdateRequest
	(*Picture)(nil),                 // 1: panguservice.Picture
}
var file_pangu_proto_sub_product_update_proto_depIdxs = []int32{
	1, // 0: panguservice.SubProductUpdateRequest.pictures:type_name -> panguservice.Picture
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pangu_proto_sub_product_update_proto_init() }
func file_pangu_proto_sub_product_update_proto_init() {
	if File_pangu_proto_sub_product_update_proto != nil {
		return
	}
	file_pangu_proto_sub_product_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pangu_proto_sub_product_update_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubProductUpdateRequest); i {
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
			RawDescriptor: file_pangu_proto_sub_product_update_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pangu_proto_sub_product_update_proto_goTypes,
		DependencyIndexes: file_pangu_proto_sub_product_update_proto_depIdxs,
		MessageInfos:      file_pangu_proto_sub_product_update_proto_msgTypes,
	}.Build()
	File_pangu_proto_sub_product_update_proto = out.File
	file_pangu_proto_sub_product_update_proto_rawDesc = nil
	file_pangu_proto_sub_product_update_proto_goTypes = nil
	file_pangu_proto_sub_product_update_proto_depIdxs = nil
}
