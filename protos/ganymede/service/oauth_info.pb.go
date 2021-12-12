// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.8.0
// source: oauth_info.proto

package ganymedeservice

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

type OAuthInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider    string `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	OauthUrl    string `protobuf:"bytes,2,opt,name=oauth_url,json=oauthUrl,proto3" json:"oauth_url,omitempty"`
	ClientId    string `protobuf:"bytes,3,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Scope       string `protobuf:"bytes,4,opt,name=scope,proto3" json:"scope,omitempty"`
	State       string `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`
	GrantType   string `protobuf:"bytes,6,opt,name=grant_type,json=grantType,proto3" json:"grant_type,omitempty"`
	RedirectUrl string `protobuf:"bytes,7,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
}

func (x *OAuthInfoResponse) Reset() {
	*x = OAuthInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oauth_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuthInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuthInfoResponse) ProtoMessage() {}

func (x *OAuthInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_oauth_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuthInfoResponse.ProtoReflect.Descriptor instead.
func (*OAuthInfoResponse) Descriptor() ([]byte, []int) {
	return file_oauth_info_proto_rawDescGZIP(), []int{0}
}

func (x *OAuthInfoResponse) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *OAuthInfoResponse) GetOauthUrl() string {
	if x != nil {
		return x.OauthUrl
	}
	return ""
}

func (x *OAuthInfoResponse) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *OAuthInfoResponse) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *OAuthInfoResponse) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *OAuthInfoResponse) GetGrantType() string {
	if x != nil {
		return x.GrantType
	}
	return ""
}

func (x *OAuthInfoResponse) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

type OAuthInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider string `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
}

func (x *OAuthInfoRequest) Reset() {
	*x = OAuthInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oauth_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuthInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuthInfoRequest) ProtoMessage() {}

func (x *OAuthInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oauth_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuthInfoRequest.ProtoReflect.Descriptor instead.
func (*OAuthInfoRequest) Descriptor() ([]byte, []int) {
	return file_oauth_info_proto_rawDescGZIP(), []int{1}
}

func (x *OAuthInfoRequest) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

var File_oauth_info_proto protoreflect.FileDescriptor

var file_oauth_info_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x67, 0x61, 0x6e, 0x79, 0x6d, 0x64, 0x65, 0x22, 0xd7, 0x01, 0x0a, 0x11,
	0x4f, 0x41, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1b, 0x0a,
	0x09, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x2e, 0x0a, 0x10, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x42, 0x28, 0x5a, 0x26, 0x61, 0x6d, 0x75, 0x73, 0x69, 0x6e, 0x67,
	0x2d, 0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x6e, 0x79, 0x6d, 0x65, 0x64, 0x65, 0x2f,
	0x67, 0x61, 0x6e, 0x79, 0x6d, 0x65, 0x64, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_oauth_info_proto_rawDescOnce sync.Once
	file_oauth_info_proto_rawDescData = file_oauth_info_proto_rawDesc
)

func file_oauth_info_proto_rawDescGZIP() []byte {
	file_oauth_info_proto_rawDescOnce.Do(func() {
		file_oauth_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_oauth_info_proto_rawDescData)
	})
	return file_oauth_info_proto_rawDescData
}

var file_oauth_info_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_oauth_info_proto_goTypes = []interface{}{
	(*OAuthInfoResponse)(nil), // 0: ganymde.OAuthInfoResponse
	(*OAuthInfoRequest)(nil),  // 1: ganymde.OAuthInfoRequest
}
var file_oauth_info_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_oauth_info_proto_init() }
func file_oauth_info_proto_init() {
	if File_oauth_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oauth_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuthInfoResponse); i {
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
		file_oauth_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuthInfoRequest); i {
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
			RawDescriptor: file_oauth_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_oauth_info_proto_goTypes,
		DependencyIndexes: file_oauth_info_proto_depIdxs,
		MessageInfos:      file_oauth_info_proto_msgTypes,
	}.Build()
	File_oauth_info_proto = out.File
	file_oauth_info_proto_rawDesc = nil
	file_oauth_info_proto_goTypes = nil
	file_oauth_info_proto_depIdxs = nil
}
