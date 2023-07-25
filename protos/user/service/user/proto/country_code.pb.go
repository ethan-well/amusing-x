// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.8.0
// source: user/proto/country_code.proto

package service

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

type CountryCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Cname      string `protobuf:"bytes,2,opt,name=cname,proto3" json:"cname,omitempty"`
	Code       string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	CreateTime string `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime string `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *CountryCode) Reset() {
	*x = CountryCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_country_code_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountryCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountryCode) ProtoMessage() {}

func (x *CountryCode) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_country_code_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountryCode.ProtoReflect.Descriptor instead.
func (*CountryCode) Descriptor() ([]byte, []int) {
	return file_user_proto_country_code_proto_rawDescGZIP(), []int{0}
}

func (x *CountryCode) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CountryCode) GetCname() string {
	if x != nil {
		return x.Cname
	}
	return ""
}

func (x *CountryCode) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CountryCode) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *CountryCode) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

type CountryCodeList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CountryCodeList []*CountryCode `protobuf:"bytes,1,rep,name=country_code_list,json=countryCodeList,proto3" json:"country_code_list,omitempty"`
}

func (x *CountryCodeList) Reset() {
	*x = CountryCodeList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_country_code_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountryCodeList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountryCodeList) ProtoMessage() {}

func (x *CountryCodeList) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_country_code_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountryCodeList.ProtoReflect.Descriptor instead.
func (*CountryCodeList) Descriptor() ([]byte, []int) {
	return file_user_proto_country_code_proto_rawDescGZIP(), []int{1}
}

func (x *CountryCodeList) GetCountryCodeList() []*CountryCode {
	if x != nil {
		return x.CountryCodeList
	}
	return nil
}

var File_user_proto_country_code_proto protoreflect.FileDescriptor

var file_user_proto_country_code_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x89, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x22, 0x50, 0x0a, 0x0f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x11, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x0f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x1c, 0x5a, 0x1a, 0x61, 0x6d, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x2d, 0x78,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_country_code_proto_rawDescOnce sync.Once
	file_user_proto_country_code_proto_rawDescData = file_user_proto_country_code_proto_rawDesc
)

func file_user_proto_country_code_proto_rawDescGZIP() []byte {
	file_user_proto_country_code_proto_rawDescOnce.Do(func() {
		file_user_proto_country_code_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_country_code_proto_rawDescData)
	})
	return file_user_proto_country_code_proto_rawDescData
}

var file_user_proto_country_code_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_user_proto_country_code_proto_goTypes = []interface{}{
	(*CountryCode)(nil),     // 0: user.CountryCode
	(*CountryCodeList)(nil), // 1: user.CountryCodeList
}
var file_user_proto_country_code_proto_depIdxs = []int32{
	0, // 0: user.CountryCodeList.country_code_list:type_name -> user.CountryCode
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_proto_country_code_proto_init() }
func file_user_proto_country_code_proto_init() {
	if File_user_proto_country_code_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_country_code_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountryCode); i {
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
		file_user_proto_country_code_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountryCodeList); i {
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
			RawDescriptor: file_user_proto_country_code_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_proto_country_code_proto_goTypes,
		DependencyIndexes: file_user_proto_country_code_proto_depIdxs,
		MessageInfos:      file_user_proto_country_code_proto_msgTypes,
	}.Build()
	File_user_proto_country_code_proto = out.File
	file_user_proto_country_code_proto_rawDesc = nil
	file_user_proto_country_code_proto_goTypes = nil
	file_user_proto_country_code_proto_depIdxs = nil
}