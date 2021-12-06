// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.8.0
// source: oauth_login.proto

package userservice

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

type OAuthLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider string `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	Code     string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *OAuthLoginRequest) Reset() {
	*x = OAuthLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oauth_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuthLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuthLoginRequest) ProtoMessage() {}

func (x *OAuthLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oauth_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuthLoginRequest.ProtoReflect.Descriptor instead.
func (*OAuthLoginRequest) Descriptor() ([]byte, []int) {
	return file_oauth_login_proto_rawDescGZIP(), []int{0}
}

func (x *OAuthLoginRequest) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *OAuthLoginRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type OAuthLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *OAuthLoginResponse) Reset() {
	*x = OAuthLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oauth_login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuthLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuthLoginResponse) ProtoMessage() {}

func (x *OAuthLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_oauth_login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuthLoginResponse.ProtoReflect.Descriptor instead.
func (*OAuthLoginResponse) Descriptor() ([]byte, []int) {
	return file_oauth_login_proto_rawDescGZIP(), []int{1}
}

func (x *OAuthLoginResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

var File_oauth_login_proto protoreflect.FileDescriptor

var file_oauth_login_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x22, 0x43, 0x0a, 0x11, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x2c, 0x0a, 0x12, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x42, 0x27, 0x5a, 0x25, 0x61, 0x6d, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x2d, 0x78,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_oauth_login_proto_rawDescOnce sync.Once
	file_oauth_login_proto_rawDescData = file_oauth_login_proto_rawDesc
)

func file_oauth_login_proto_rawDescGZIP() []byte {
	file_oauth_login_proto_rawDescOnce.Do(func() {
		file_oauth_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_oauth_login_proto_rawDescData)
	})
	return file_oauth_login_proto_rawDescData
}

var file_oauth_login_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_oauth_login_proto_goTypes = []interface{}{
	(*OAuthLoginRequest)(nil),  // 0: userservice.OAuthLoginRequest
	(*OAuthLoginResponse)(nil), // 1: userservice.OAuthLoginResponse
}
var file_oauth_login_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_oauth_login_proto_init() }
func file_oauth_login_proto_init() {
	if File_oauth_login_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oauth_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuthLoginRequest); i {
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
		file_oauth_login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuthLoginResponse); i {
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
			RawDescriptor: file_oauth_login_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_oauth_login_proto_goTypes,
		DependencyIndexes: file_oauth_login_proto_depIdxs,
		MessageInfos:      file_oauth_login_proto_msgTypes,
	}.Build()
	File_oauth_login_proto = out.File
	file_oauth_login_proto_rawDesc = nil
	file_oauth_login_proto_goTypes = nil
	file_oauth_login_proto_depIdxs = nil
}